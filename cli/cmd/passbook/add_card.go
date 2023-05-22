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

//func init() {
//	RootCmd.AddCommand(addCardCmd)
//	addCardCmd.Flags().StringVarP(&addCardRequest.Title, "title", "t", "", "title to save")
//	addCardCmd.Flags().StringVarP(&addCardRequest.Name, "name", "n", "", "name to save")
//	addCardCmd.Flags().StringVarP(&addCardRequest.CardNumber, "card", "", "", "card number to save")
//	addCardCmd.Flags().StringVarP(&addCardRequest.DateExp, "date", "d", "", "date expiration to save")
//	addCardCmd.Flags().StringVarP(&addCardRequest.CvcCode, "cvc", "", "", "cvc code to save")
//
//	addCardCmd.MarkFlagRequired("title")
//	addCardCmd.MarkFlagRequired("name")
//	addCardCmd.MarkFlagRequired("card")
//}
