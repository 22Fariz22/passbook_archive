package passbook

import (
	"github.com/22Fariz22/passbook/cli/pkg"
	pb "github.com/22Fariz22/passbook/server/proto"
	"github.com/spf13/cobra"
)

var addAccountRequest pb.AddAccountRequest

var addAccountCmd = &cobra.Command{
	Use:     "account",
	Aliases: []string{"acc"},
	Short:   "add the account to save",
	Run: func(cmd *cobra.Command, args []string) {
		c := pkg.ConnGRPCServer()

		err := pkg.AddAccount(c, &addAccountRequest)
		if err != nil {
			return
		}
	}}
