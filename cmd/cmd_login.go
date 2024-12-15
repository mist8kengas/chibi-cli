package cmd

import (
	"fmt"
	"os"

	"github.com/CosmicPredator/chibi/internal"
	"github.com/spf13/cobra"
)

func handleLoginCommand() {
	var code string

	authRequest := internal.NewAuthRequest()
	fmt.Printf(
		"Open the below link in browser to login with anilist: \n%s \n",
		authRequest.GetAuthURL(),
	)
	fmt.Println("Copy the code from the browser and enter it below:")
	fmt.Print("Enter code: ")
	fmt.Scanln(&code)

	if code == "" {
		fmt.Println("Please provide a valid code")
		os.Exit(0)
	}

	authRequest.Login(code)
	fmt.Println("Login Successful")
}

var loginCmd = &cobra.Command{
	Use: "login",
	Short: "Login with anilist",
	Run: func(cmd *cobra.Command, args []string) {
		handleLoginCommand()
	},
}