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

var addAccountRequest pb.AddAccountRequest

var addAccountCmd = &cobra.Command{
	Use:     "account",
	Aliases: []string{"acc"},
	Short:   "add the account to save",
	Long:    "",
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := grpc.Dial(":5001", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		// получаем переменную интерфейсного типа UsersClient,
		// через которую будем отправлять сообщения
		c := pb.NewUserServiceClient(conn)

		err = pkg.AddAccount(c, &addAccountRequest)
		if err != nil {
			return
		}
		fmt.Println("account added")
	}}

func init() {
	rootCmd.AddCommand(addAccountCmd)
	addAccountCmd.Flags().StringVarP(&addAccountRequest.Title, "title", "t", "", "title to save")
	addAccountCmd.Flags().StringVarP(&addAccountRequest.Login, "login", "l", "", "login to save")
	addAccountCmd.Flags().StringVarP(&addAccountRequest.Password, "password", "p", "", "password to save")

	addAccountCmd.MarkFlagRequired("title")
	addAccountCmd.MarkFlagRequired("login")
	addAccountCmd.MarkFlagRequired("password")
}
