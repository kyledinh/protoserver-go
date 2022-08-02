#!/bin/bash




JWT_TOKEN=$(
    curl -X POST http://localhost:8000/v1/login  | jq -r '.token'
)

echo "JWT_TOKEN: $JWT_TOKEN"
echo "using now to fetch a heartbeat"

HEARTBEAT=$(curl http://localhost:8000/v1/heartbeat \
    -H  "X-Authentication-Token: $JWT_TOKEN"
)

echo "HEARTBEAT: $HEARTBEAT"