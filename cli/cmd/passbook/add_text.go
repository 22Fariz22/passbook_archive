package passbook

import (
	"fmt"
	"github.com/22Fariz22/passbook/cli/pkg"
	pb "github.com/22Fariz22/passbook/server/proto"
	"github.com/spf13/cobra"
)

var addTextRequest pb.AddTextRequest

var addTextCmd = &cobra.Command{
	Use:     "text",
	Aliases: []string{"tex"},
	Short:   "text to save",
	Long:    "",
	Run: func(cmd *cobra.Command, args []string) {
		c := ConnGRPCServer()

		err := pkg.AddText(c, &addTextRequest)
		if err != nil {
			return
		}
		fmt.Println("text added")
	}}

func init() {
	RootCmd.AddCommand(addTextCmd)
	addTextCmd.Flags().StringVarP(&addTextRequest.Title, "title", "t", "", "add title")
	addTextCmd.Flags().StringVarP(&addTextRequest.Data, "data", "d", "", "add text")

	addAccountCmd.MarkFlagRequired("title")
	addAccountCmd.MarkFlagRequired("data")

}
