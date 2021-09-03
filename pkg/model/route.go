package model

import "encoding/json"

type Route struct {
	Ingress  string          `json:"ingress"`
	Action   string          `json:"action"`
	Egress   string          `json:"egress"`
	Egresses []string        `json:"egresses"`
	Macros   []string        `json:"macros"`
	Logic    string          `json:"logic"`
	Payload  json.RawMessage `json:"payload"`
	File     string          `json:"file"`
}

// Actions: forward | get | respond | error
// Macros type of server prototyped
