package cmd

import "github.com/malsuke/jwt-fuzzer/internal/request"

func All(client request.Client, token JWT) error {
	algoTest(client, token)
	BlankPasswordTest(client, token.EncodeToString())
	nullSigTest(client, token)
	BruteForceTest(token.EncodeToString())
	return nil
}
