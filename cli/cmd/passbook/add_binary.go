package passbook

import (
	"github.com/22Fariz22/passbook/cli/pkg"
	pb "github.com/22Fariz22/passbook/server/proto"
	"github.com/spf13/cobra"
)

var AddBinaryRequest pb.AddBinaryRequest

var addBinaryCmd = &cobra.Command{
	Use:     "Binary",
	Aliases: []string{"bin"},
	Short:   "binary to save",
	Run:     addBinaryCmdRun,
}

func addBinaryCmdRun(cmd *cobra.Command, args []string) {
	c := pkg.ConnGRPCServer()

	err := pkg.AddBinary(c, &AddBinaryRequest)
	if err != nil {
		return
	}
	return
}
