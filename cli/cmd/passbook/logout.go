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

var logoutCmd = &cobra.Command{
	Use:     "logout",
	Aliases: []string{"out"},
	Short:   "out",
	Long:    "выходим из системы",
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := grpc.Dial(":5001", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		// получаем переменную интерфейсного типа UsersClient,
		// через которую будем отправлять сообщения
		c := pb.NewUserServiceClient(conn)

		err = pkg.Logout(c, &pb.LogoutRequest{})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("you got out")
	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)

}
