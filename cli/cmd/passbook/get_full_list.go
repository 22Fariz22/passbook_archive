package passbook

import (
	"fmt"
	"github.com/22Fariz22/passbook/cli/pkg"
	pb "github.com/22Fariz22/passbook/server/proto"
	"github.com/spf13/cobra"
	"log"
)

var getFullListCmd = &cobra.Command{
	Use:   "full",
	Short: "get all your secrets",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		c := pkg.ConnGRPCServer()

		res, err := pkg.GetFullList(c, &pb.GetFullListRequest{})

		if err != nil {
			log.Println("can not get full list.")
			return
		}

		//выводим список секретов
		for _, v := range res.Data {
			fmt.Println(v)
		}
	}}
