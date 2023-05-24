package passbook

import (
	"github.com/22Fariz22/passbook/cli/pkg"
	pb "github.com/22Fariz22/passbook/server/proto"
	"github.com/spf13/cobra"
	"log"
)

var logoutCmd = &cobra.Command{
	Use:     "logout",
	Aliases: []string{"out"},
	Short:   "out",
	Long:    "выходим из системы",
	Run: func(cmd *cobra.Command, args []string) {
		c := pkg.ConnGRPCServer()

		err := pkg.Logout(c, &pb.LogoutRequest{})
		if err != nil {
			log.Fatal(err)
			return
		}
	},
}
