package cmd

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
)

func Parse(token string) (jwt JWT, err error) {
	splitToken := strings.Split(token, ".")
	if len(splitToken) != 3 {
		return JWT{}, errors.New("JWTの形式が不正です")
	}

	header, err := base64.RawURLEncoding.DecodeString(splitToken[0])
	if err != nil {
		return JWT{}, err
	}
	json.Unmarshal([]byte(header), &jwt.Header)

	payload, err := base64.RawURLEncoding.DecodeString(splitToken[1])
	if err != nil {
		return JWT{}, err
	}

	json.Unmarshal([]byte(payload), &jwt.Payload)
	jwt.Signature = splitToken[2]

	return jwt, err
}
