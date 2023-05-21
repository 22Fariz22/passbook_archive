package passbook

import (
	"fmt"
	"github.com/22Fariz22/passbook/cli/pkg"
	pb "github.com/22Fariz22/passbook/server/proto"
	"github.com/spf13/cobra"
)

var addAccountRequest pb.AddAccountRequest

var addAccountCmd = &cobra.Command{
	Use:     "account",
	Aliases: []string{"acc"},
	Short:   "add the account to save",
	Long:    "",
	Run: func(cmd *cobra.Command, args []string) {
		c := ConnGRPCServer()

		err := pkg.AddAccount(c, &addAccountRequest)
		if err != nil {
			return
		}
		fmt.Println("account added")
	}}

func init() {
	RootCmd.AddCommand(addAccountCmd)
	addAccountCmd.Flags().StringVarP(&addAccountRequest.Title, "title", "t", "", "title to save")
	addAccountCmd.Flags().StringVarP(&addAccountRequest.Login, "login", "l", "", "login to save")
	addAccountCmd.Flags().StringVarP(&addAccountRequest.Password, "password", "p", "", "password to save")

	addAccountCmd.MarkFlagRequired("title")
	addAccountCmd.MarkFlagRequired("login")
	addAccountCmd.MarkFlagRequired("password")
}
