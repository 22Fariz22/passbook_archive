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

var addCardRequest pb.AddCardRequest

var addCardCmd = &cobra.Command{
	Use:   "card",
	Short: "add the card to save",
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

		err = pkg.AddCard(c, &addCardRequest)
		if err != nil {
			return
		}
		fmt.Println("card added")
	}}

func init() {
	rootCmd.AddCommand(addCardCmd)
	addCardCmd.Flags().StringVarP(&addCardRequest.Title, "title", "t", "", "title to save")
	addCardCmd.Flags().StringVarP(&addCardRequest.Name, "name", "n", "", "name to save")
	addCardCmd.Flags().StringVarP(&addCardRequest.CardNumber, "card", "c", "", "card number to save")
	addCardCmd.Flags().StringVarP(&addCardRequest.DateExp, "date", "d", "", "date expiration to save")
	addCardCmd.Flags().StringVarP(&addCardRequest.CvcCode, "cvc", "", "", "cvc code to save")

	addCardCmd.MarkFlagRequired("title")
	addCardCmd.MarkFlagRequired("name")
	addCardCmd.MarkFlagRequired("card")

}
