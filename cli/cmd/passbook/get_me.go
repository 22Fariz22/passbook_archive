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

var getMeCmd = &cobra.Command{
	Use:     "getme",
	Aliases: []string{"me"},
	Short:   "Get your currently account session",
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := grpc.Dial(":5001", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		// получаем переменную интерфейсного типа UsersClient,
		// через которую будем отправлять сообщения
		c := pb.NewUserServiceClient(conn)

		user, err := pkg.GetMe(c, &pb.GetMeRequest{})
		if err != nil {
			log.Println("you don't have a session")
		}
		fmt.Println("ID:", user.User.Uuid)
		fmt.Println("Login:", user.User.Login)
	},
}

func init() {
	rootCmd.AddCommand(getMeCmd)
}
