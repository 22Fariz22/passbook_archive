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

var addTextRequest pb.AddTextRequest

var addTextCmd = &cobra.Command{
	Use:     "text",
	Aliases: []string{"tex"},
	Short:   "text to save",
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

		err = pkg.AddText(c, &addTextRequest)
		if err != nil {
			return
		}
		fmt.Println("text added")
	}}

func init() {
	rootCmd.AddCommand(addTextCmd)
	addTextCmd.Flags().StringVarP(&addTextRequest.Title, "title", "t", "", "add title")
	addTextCmd.Flags().StringVarP(&addTextRequest.Data, "data", "d", "", "add text")

	addAccountCmd.MarkFlagRequired("title")
	addAccountCmd.MarkFlagRequired("data")

}
