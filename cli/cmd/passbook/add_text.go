package passbook

import (
	"fmt"
	"github.com/22Fariz22/passbook/cli/pkg"
	pb "github.com/22Fariz22/passbook/server/proto"
	"github.com/spf13/cobra"
)

var AddTextRequest pb.AddTextRequest

var addTextCmd = &cobra.Command{
	Use:     "text",
	Aliases: []string{"tex"},
	Short:   "text to save",
	Run:     addTextCmdRun,
}

func addTextCmdRun(cmd *cobra.Command, args []string) {
	c := pkg.ConnGRPCServer()

	err := pkg.AddText(c, &AddTextRequest)
	if err != nil {
		return
	}
	fmt.Println("text added")
}
