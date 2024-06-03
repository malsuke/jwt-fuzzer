/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/malsuke/jwt-fuzzer/internal/request"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "jwt-fuzzer",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		prompt := promptui.Select{
			Label: "which test?",
			Items: []string{
				"All Test",
				"Algo None Test",
				"Null Signature Test",
				"Black Password Test",
				"Dictionary Attack Test",
			},
			CursorPos: 0,
		}
		idx, _, err := prompt.Run()
		if err != nil {
			fmt.Println(err)
			return
		}

		templates := &promptui.PromptTemplates{
			Prompt:  "{{ . | bold }} ",
			Valid:   "{{ . | bold | green }} ",
			Invalid: "{{ . | bold | red }} ",
			Success: "{{ . | bold | green }} ",
		}

		validate := func(input string) error {
			if len(input) < 1 {
				return errors.New("token is required")
			}
			return nil
		}

		token := promptui.Prompt{
			Label:     "Input JWT Token",
			Templates: templates,
			Validate:  validate,
		}

		var input string

		if os.Getenv("ENV_JWT_TOKEN") == "" {
			input, err = token.Run()
			if err != nil {
				fmt.Println(err)
				return
			}
		} else {
			input = os.Getenv("ENV_JWT_TOKEN")
			fmt.Printf("\033[32mInput JWT Token %s\033[0m\n", input)
		}
		client, err := request.NewClient(os.Getenv("ENV_ENDPOINT"))
		if err != nil {
			fmt.Printf("Error creating client: %v", err)
			return
		}

		jwt, err := Parse(input)
		if err != nil {
			fmt.Printf("Error parsing token: %v", err)
			return
		}

		switch idx {
		case 0:
			All(*client, jwt)
		case 1:
			algoTest(*client, jwt)
		case 2:
			nullSigTest(*client, jwt)
		case 3:
			BlankPasswordTest(*client, input)
		case 4:
			BruteForceTest(input)
		}
	},
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.jwt-fuzzer.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
