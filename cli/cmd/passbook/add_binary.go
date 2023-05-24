package passbook

import (
	pb "github.com/22Fariz22/passbook/server/proto"
	"github.com/spf13/cobra"
)

var addBinaryRequest pb.AddBinaryRequest
var pathToFile string

var addBinaryCmd = &cobra.Command{
	Use:     "Binary",
	Aliases: []string{"bin"},
	Short:   "binary to save",
	Run:     addBinaryCmdRun,
}

func addBinaryCmdRun(cmd *cobra.Command, args []string) {

}
