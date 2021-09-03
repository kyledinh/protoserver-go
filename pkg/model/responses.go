package model

import "encoding/json"

type RespHeartbeat struct {
	Status string `json:"status"`
}

type ApiResponse struct {
	TraceUUID string          `json:"trace_uuid"`
	Metadata  string          `json:"metadata"`
	Timestamp string          `json:"timestamp"`
	Payload   json.RawMessage `json:"payload"`
}

func (ar *ApiResponse) StdErr(err error) []byte {
	newApiResp := ApiResponse{
		TraceUUID: ar.TraceUUID,
		Timestamp: ar.Timestamp,
		Payload:   json.RawMessage(`{"error":"` + err.Error() + `"}`),
	}
	ba, _ := json.Marshal(newApiResp)
	return ba
}
