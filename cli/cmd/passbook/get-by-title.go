package passbook

import (
	"fmt"
	"github.com/22Fariz22/passbook/cli/pkg"
	pb "github.com/22Fariz22/passbook/server/proto"
	"github.com/spf13/cobra"
	"log"
)

var getByTitleRequest pb.GetByTitleRequest

var getByTitleCmd = &cobra.Command{
	Use:   "get_title",
	Short: "get secret by title",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		c := ConnGRPCServer()

		res, err := pkg.GetByTitle(c, &getByTitleRequest)
		if err != nil {
			log.Println("can not get by title")
			return
		}

		//выводим список секретов
		for _, v := range res.Data {
			fmt.Println(v)
		}
	}}

func init() {
	RootCmd.AddCommand(getByTitleCmd)
	getByTitleCmd.Flags().StringVarP(&getByTitleRequest.Title, "title", "t", "", "get your secret by title")
	getByTitleCmd.MarkFlagRequired("title")
}
