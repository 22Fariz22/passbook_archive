package passbook

import (
	"fmt"
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
		c := ConnGRPCServer()

		err := pkg.Logout(c, &pb.LogoutRequest{})
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Println("you got out")
	},
}

func init() {
	RootCmd.AddCommand(logoutCmd)

}
