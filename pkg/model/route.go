package model

import "encoding/json"

type Route struct {
	Action   string          `json:"action"`
	Egress   string          `json:"egress"`
	Egresses []string        `json:"egresses"`
	File     string          `json:"file"`
	Ingress  string          `json:"ingress"`
	Logic    string          `json:"logic"`
	Macros   []string        `json:"macros"`
	Payload  json.RawMessage `json:"payload"`
}

// Actions: forward | get | respond | error
// Macros type of server prototyped
