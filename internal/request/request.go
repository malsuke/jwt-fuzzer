package request

import (
	"context"
	"fmt"
	"io"
	"strings"
	"time"
)

func RequestWithJWT(client Client, token string) (err error) {
	params := &CheckLoginParams{
		Jwt: token,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := client.CheckLogin(ctx, params)
	if err != nil {
		fmt.Printf("Error calling CheckLogin: %v\n", err)
		return
	}
	defer resp.Body.Close()
	temp, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return
	}
	cleanMessage := strings.ReplaceAll(string(temp), "\n", "")

	fmt.Println("┌────────────────┬────────────────────────────────────┐")
	fmt.Printf("│     %-6s     │               %-8s             │\n", "Status", "Message")
	fmt.Println("├────────────────┼────────────────────────────────────┤")
	fmt.Printf("│      %-3.3s       │   %-30.30s   │\n", resp.Status, cleanMessage)
	fmt.Println("└────────────────┴────────────────────────────────────┘\n")

	return nil
}
