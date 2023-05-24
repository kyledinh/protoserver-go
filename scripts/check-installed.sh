#!/bin/bash
# Checks for installed software packages 

function fn_check_docker {
  if command -v docker >/dev/null 2>&1; then
    docker --version
  else
    echo "!!! docker required, but it's not installed."; 
  fi
}

function fn_check_go {
  if command -v go >/dev/null 2>&1; then
    go version
  else
    echo "!!! make required, but it's not installed."; 
  fi
}

function fn_check_make {
  if command -v make >/dev/null 2>&1; then
    make --version
  else
    echo "!!! make required, but it's not installed."; 
  fi
}

function fn_check_node {
  if command -v npm >/dev/null 2>&1; then
    npm version
  else
    echo "!!! node/npm/yarn required, but it's not installed."; 
  fi
}

## MAIN
echo "Checking installed software packages:"
echo "docker, go..."
echo
fn_check_docker
fn_check_go
# fn_check_make
