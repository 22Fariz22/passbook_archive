package passbook

import (
	"fmt"
	"github.com/22Fariz22/passbook/cli/pkg"
	pb "github.com/22Fariz22/passbook/server/proto"
	"github.com/spf13/cobra"
	"os"
)

var pathToFile string
var titleBinary string

var addBinaryCmd = &cobra.Command{
	Use:     "Binary",
	Aliases: []string{"bin"},
	Short:   "binary to save",
	Run:     addBinaryCmdRun,
}

func addBinaryCmdRun(cmd *cobra.Command, args []string) {
	b, err := os.ReadFile(pathToFile)
	if err != nil {
		fmt.Println("не удалось загрузить контент")
		return
	}

	c := pkg.ConnGRPCServer()

	pkg.AddBinary(c, &pb.AddBinaryRequest{
		Title: titleBinary,
		Data:  b,
	})
}
