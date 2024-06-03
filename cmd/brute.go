package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt"
)

func BruteForceTest(token string) {
	file, err := os.Open("wordlist.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var secrets []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		secrets = append(secrets, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println("\033[31mDictionary Attack TEST\033[0m")

	for _, secret := range secrets {
		token, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secret), nil
		})

		if err != nil {
			fmt.Println("Failed to parse with secret:", secret, "Error:", err)
			continue
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			fmt.Println("Found secret:", secret)
			fmt.Println("Claims:", claims)
			break
		} else {
			fmt.Println("Invalid token or claims do not match for secret:", secret)
		}
	}
}
