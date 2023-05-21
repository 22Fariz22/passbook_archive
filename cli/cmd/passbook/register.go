package passbook

import (
	"github.com/22Fariz22/passbook/cli/pkg"
	pb "github.com/22Fariz22/passbook/server/proto"
	"github.com/spf13/cobra"
)

var registerReq pb.LoginRequest

var registerCmd = &cobra.Command{
	Use:     "register",
	Aliases: []string{"reg"},
	Short:   "it's like a sign-up",
	Run: func(cmd *cobra.Command, args []string) {
		c := ConnGRPCServer()

		pkg.Register(c, &pb.RegisterRequest{
			Login:    registerReq.Login,
			Password: registerReq.Password,
		})

	},
}

func init() {
	RootCmd.AddCommand(registerCmd)
	registerCmd.Flags().StringVarP(&registerReq.Login, "login", "l", "", "it is string to reverse")
	registerCmd.Flags().StringVarP(&registerReq.Password, "password", "p", "", "it is string to reverse")
}
