package passbook

import (
	"github.com/22Fariz22/passbook/cli/pkg"
	pb "github.com/22Fariz22/passbook/server/proto"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var loginReq pb.LoginRequest
var loginResp pb.LoginResponse

var loginCmd = &cobra.Command{
	Use:     "login",
	Aliases: []string{"log"},
	Short:   "it's like a sign-in",
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := grpc.Dial(":5001", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		// получаем переменную интерфейсного типа UsersClient,
		// через которую будем отправлять сообщения
		c := pb.NewUserServiceClient(conn)

		pkg.Login(c, &pb.LoginRequest{
			Login:    login,
			Password: password,
		})
	},
}

var login string
var password string

func init() {
	rootCmd.AddCommand(loginCmd)
	loginCmd.Flags().StringVarP(&login, "login", "l", "", "it is string to reverse")
	loginCmd.Flags().StringVarP(&password, "password", "p", "", "it is string to reverse")
}
