#!/bin/bash

## ENVs and VARS
TARGET=docker-desktop
NAMESPACE=txferry
# NAMESPACE=default
THIS_FILE=$0

## CHECKS
if kubectl config current-context | grep -q $TARGET
then 
	echo; echo "## Current context $TARGET"
else
	echo; echo "## Wrong current context: $(kubectl config current-context)!"
	echo "Aborting script. No actions taken."
	exit 1
fi

function fn_chrome {
	open -a "Google Chrome" http://api.txferry.local/v0/heartbeat
	open -a "Google Chrome" http://api.txferry.local/vx/contact
	open -a "Google Chrome" http://api.txferry.local/vx/agent
}

function fn_down {
	kubectl delete deployment txferry-api -n $NAMESPACE
	kubectl delete service/svc-txferry -n $NAMESPACE
	kubectl delete ingress ingress-api-txferry-local -n $NAMESPACE
	kubectl delete cm cm-txferry-api-protoserver-routes-json -n $NAMESPACE

	kubectl delete deployment agent-api -n $NAMESPACE
	kubectl delete service/svc-agent -n $NAMESPACE
	kubectl delete cm cm-agent-api-protoserver-routes-json -n $NAMESPACE

	kubectl delete deployment bank-api -n $NAMESPACE
	kubectl delete service/svc-bank -n $NAMESPACE
	kubectl delete cm cm-bank-api-protoserver-routes-json -n $NAMESPACE

}

function fn_ex {
	POD=$(kubectl get pods -n txferry | grep txferry | awk '{print $1}')
    echo "kubectl exec -it -n txferry $POD -- ash"
    kubectl exec -it -n txferry $POD -- ash
}

function fn_help {
	echo; echo "USAGE: $0 <arguments>. Try '$0 up | down | info | log | ... '"
	tail -50 $THIS_FILE | grep ")"
}

function fn_info {
	echo; kubectl get ingress -n $NAMESPACE
	echo; echo "Deployments"
	kubectl get deployments -n $NAMESPACE
	echo; echo "Services"
	kubectl get svc  -n $NAMESPACE
	echo; echo "Pods"
	kubectl get pods -n $NAMESPACE
	echo; echo "Config Maps"
	kubectl get cm -n $NAMESPACE
}

function fn_log {
	POD=$(kubectl get pods -n $NAMESPACE | grep $1 | awk '{print $1}')
	echo; echo "kubectl logs $POD -n $NAMESPACE"
    kubectl logs $POD -n $NAMESPACE
}

function fn_previous {
	echo; echo "kubectl logs txferry-api --previous -n $NAMESPACE"
    kubectl logs txferry-api --previous -n $NAMESPACE
}

function fn_up {
	kubectl apply -f docker-desktop/deployment-txferry.yaml -n $NAMESPACE
	kubectl apply -f docker-desktop/deployment-postgres.yaml -n $NAMESPACE
}

## Creates the namespace for "txferry", ran once 
function fn_init {
	kubectl apply -f common/namespace-txferry.yaml
	kubectl apply -f common/deploy-ingress-nginx-controller.yaml
}

## MAIN/SWITCH
echo ".... in the MAIN/SWITCH"

if [ "$#" -eq 0 ]; then
	COMMAND=up
	echo; echo "Defaulting to <deploy>. Try '$0 help' for more options."; echo
else
	COMMAND=$1
fi

case "$COMMAND" in
	chrome)
		fn_chrome
		exit 0;;
	down)
		fn_down
		exit 0;;
	ex)
		fn_ex
		exit 0;;
	help)
		fn_help
		exit 0;;
	info) 
		fn_info
		exit 0;;
	init) 
		fn_init
		exit 0;;
	log) 
		fn_log ${@:2}
		exit 0;;
	logs) 
		fn_logs ${@:2}
		exit 0;;
	up)
		fn_up
		exit 0;;
	*) 
		fn_info
		exit 0;;
esac