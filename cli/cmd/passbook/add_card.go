package passbook

import (
	"github.com/22Fariz22/passbook/cli/pkg"
	pb "github.com/22Fariz22/passbook/server/proto"
	"github.com/spf13/cobra"
)

var addCardRequest pb.AddCardRequest

var addCardCmd = &cobra.Command{
	Use:   "card",
	Short: "add the card to save",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		c := pkg.ConnGRPCServer()

		err := pkg.AddCard(c, &addCardRequest)
		if err != nil {
			return
		}
	}}
