package passbook

import (
	"fmt"
	"github.com/22Fariz22/passbook/cli/pkg"
	pb "github.com/22Fariz22/passbook/server/proto"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var getFullListCmd = &cobra.Command{
	Use:   "full",
	Short: "get all your secrets",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := grpc.Dial(":5001", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		// получаем переменную интерфейсного типа UsersClient,
		// через которую будем отправлять сообщения
		c := pb.NewUserServiceClient(conn)

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

func init() {
	rootCmd.AddCommand(getFullListCmd)
}
