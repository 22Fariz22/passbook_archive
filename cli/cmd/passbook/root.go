package passbook

import (
	"fmt"
	"github.com/spf13/cobra"
)

var version = "v0.0.1  20.05.2023"

var RootCmd = &cobra.Command{
	RunE: RootCmdRunE,
}

func RootCmdRunE(cmd *cobra.Command, args []string) error {
	fmt.Println("passbook is ready")
	return nil
}

func Execute(cmd *cobra.Command) error {
	RootCmd.AddCommand(addTextCmd)
	addTextCmd.Flags().StringVarP(&AddTextRequest.Title, "title", "t", "", "add title")
	addTextCmd.Flags().StringVarP(&AddTextRequest.Data, "data", "d", "", "add text")

	RootCmd.AddCommand(addCardCmd)
	addCardCmd.Flags().StringVarP(&addCardRequest.Title, "title", "t", "", "title to save")
	addCardCmd.Flags().StringVarP(&addCardRequest.Name, "name", "n", "", "name to save")
	addCardCmd.Flags().StringVarP(&addCardRequest.CardNumber, "card", "", "", "card number to save")
	addCardCmd.Flags().StringVarP(&addCardRequest.DateExp, "date", "d", "", "date expiration to save")
	addCardCmd.Flags().StringVarP(&addCardRequest.CvcCode, "cvc", "", "", "cvc code to save")

	RootCmd.AddCommand(getByTitleCmd)
	getByTitleCmd.Flags().StringVarP(&getByTitleRequest.Title, "title", "t", "", "get your secret by title")

	RootCmd.AddCommand(getFullListCmd)

	RootCmd.AddCommand(getMeCmd)

	RootCmd.AddCommand(loginCmd)
	loginCmd.Flags().StringVarP(&login, "login", "l", "", "it is string to reverse")
	loginCmd.Flags().StringVarP(&password, "password", "p", "", "it is string to reverse")

	RootCmd.AddCommand(logoutCmd)

	RootCmd.AddCommand(registerCmd)
	registerCmd.Flags().StringVarP(&registerReq.Login, "login", "l", "", "it is string to reverse")
	registerCmd.Flags().StringVarP(&registerReq.Password, "password", "p", "", "it is string to reverse")

	//cobra.CheckErr(RootCmd.Execute())
	cmd.Execute()

	return nil
}
