package cmd

import (
	"encoding/base64"
	"encoding/json"
)

type JWT struct {
	Header    Header
	Payload   map[string]interface{}
	Signature string
}

type Header struct {
	Typ string `json:"typ"`
	Alg string `json:"alg"`
}

func (jwt JWT) EncodeToString() string {
	header, _ := json.Marshal(jwt.Header)
	payload, _ := json.Marshal(jwt.Payload)

	return base64.RawURLEncoding.EncodeToString(header) + "." + base64.RawURLEncoding.EncodeToString(payload) + "." + jwt.Signature
}
