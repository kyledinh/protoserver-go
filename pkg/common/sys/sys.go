package sys

// Logging keys
const (
	CALLREF     = "callRef"
	INTERNALREF = "internalRef"
	LOG         = "log"
)

// Status Strings
const (
	FAILURE = "failure"
	SUCCESS = "success"
)

// Actions for model.route.Action, tells the endpoint what to do
const (
	FORWARD = "forward"
	GET     = "get"
	GETPOST = "getpost"
	POST    = "post"
	RESPOND = "respond"
)

const (
	MACRO_STD_RESPONSE          = "macro_std_response"
	MACRO_TRACEID               = "macro_traceid"
	MACRO_MAKE_TOKEN            = "macro_make_token"
	MACRO_SEND_SCHEDULER_DEPLOY = "macro_send_scheduler_deploy"
	MACRO_SQS_ADD               = "macro_sqs_add"
)
