package passbook

import (
	"github.com/22Fariz22/passbook/cli/pkg"
	pb "github.com/22Fariz22/passbook/server/proto"
	"github.com/spf13/cobra"
)

var loginReq pb.LoginRequest
var loginResp pb.LoginResponse

var loginCmd = &cobra.Command{
	Use:     "login",
	Aliases: []string{"log"},
	Short:   "it's like a sign-in",
	Run: func(cmd *cobra.Command, args []string) {
		c := ConnGRPCServer()

		pkg.Login(c, &pb.LoginRequest{
			Login:    login,
			Password: password,
		})
	},
}

var login string
var password string

func init() {
	RootCmd.AddCommand(loginCmd)
	loginCmd.Flags().StringVarP(&login, "login", "l", "", "it is string to reverse")
	loginCmd.Flags().StringVarP(&password, "password", "p", "", "it is string to reverse")
}
