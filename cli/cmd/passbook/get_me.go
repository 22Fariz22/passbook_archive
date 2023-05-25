package passbook

import (
	"fmt"
	"github.com/22Fariz22/passbook/cli/pkg"
	pb "github.com/22Fariz22/passbook/server/proto"
	"github.com/spf13/cobra"
	"log"
)

var getMeCmd = &cobra.Command{
	Use:     "getme",
	Aliases: []string{"me"},
	Short:   "Get your currently account session",
	Run: func(cmd *cobra.Command, args []string) {
		c := pkg.ConnGRPCServer()

		user, err := pkg.GetMe(c, &pb.GetMeRequest{})
		if err != nil {
			log.Println("you don't have a session")
			return
		}
		fmt.Println("ID:", user.User.Uuid)
		fmt.Println("Login:", user.User.Login)
	},
}
