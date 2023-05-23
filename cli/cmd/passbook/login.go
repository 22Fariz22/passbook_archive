package passbook

import (
	"github.com/22Fariz22/passbook/cli/pkg"
	pb "github.com/22Fariz22/passbook/server/proto"
	"github.com/spf13/cobra"
)

var login string
var password string

var loginCmd = &cobra.Command{
	Use:     "login",
	Aliases: []string{"log"},
	Short:   "it's like a sign-in",
	Run: func(cmd *cobra.Command, args []string) {
		c := pkg.ConnGRPCServer()

		pkg.Login(c, &pb.LoginRequest{
			Login:    login,
			Password: password,
		})
	},
}
