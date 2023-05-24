package passbook

import (
	"fmt"
	"github.com/spf13/cobra"
)

var version = "v0.0.1  20.05.2023"

// RootCmd главная команда
var RootCmd = &cobra.Command{
	Short:   "Passbook keep your private information",
	Version: version,
	RunE:    RootCmdRunE,
}

// RootCmdRunE состав главной команды
func RootCmdRunE(cmd *cobra.Command, args []string) error {
	fmt.Println("passbook is ready")
	return nil
}

// Execute исполнитель остальных подкоманд
func Execute(cmd *cobra.Command) error {
	//addAccountCmd сохраняет данные аккаунта
	RootCmd.AddCommand(addAccountCmd)
	addAccountCmd.Flags().StringVarP(&addAccountRequest.Title, "title", "t", "", "title to save")
	addAccountCmd.Flags().StringVarP(&addAccountRequest.Login, "login", "l", "", "login to save")
	addAccountCmd.Flags().StringVarP(&addAccountRequest.Password, "password", "p", "", "password to save")
	addAccountCmd.MarkFlagRequired("login")
	addAccountCmd.MarkFlagRequired("password")

	//addTextCmd сохряняет произвольные текстовые данные
	RootCmd.AddCommand(addTextCmd)
	addTextCmd.Flags().StringVarP(&addTextRequest.Title, "title", "t", "", "add title")
	addTextCmd.Flags().StringVarP(&addTextRequest.Data, "data", "d", "", "add text")
	addTextCmd.MarkFlagRequired("data")

	//addBinaryCmd сохряняет произвольные бинарные данные
	RootCmd.AddCommand(addBinaryCmd)
	addBinaryCmd.Flags().StringVarP(&titleBinary, "title", "t", "", "add title")
	addBinaryCmd.Flags().StringVarP(&pathToFile, "path", "p", "", "path to file")
	addBinaryCmd.MarkFlagRequired("data")

	//addCardCmd сохраняет данные бансковской карты
	RootCmd.AddCommand(addCardCmd)
	addCardCmd.Flags().StringVarP(&addCardRequest.Title, "title", "t", "", "title to save")
	addCardCmd.Flags().StringVarP(&addCardRequest.Name, "name", "n", "", "name to save")
	addCardCmd.Flags().StringVarP(&addCardRequest.CardNumber, "card", "", "", "card number to save")
	addCardCmd.Flags().StringVarP(&addCardRequest.DateExp, "date", "d", "", "date expiration to save")
	addCardCmd.Flags().StringVarP(&addCardRequest.CvcCode, "cvc", "", "", "cvc code to save")
	addCardCmd.MarkFlagRequired("card")
	addCardCmd.MarkFlagRequired("date")
	addCardCmd.MarkFlagRequired("cvc")

	//getByTitleCmd получение данных по мета-информации
	RootCmd.AddCommand(getByTitleCmd)
	getByTitleCmd.Flags().StringVarP(&getByTitleRequest.Title, "title", "t", "", "get your secret by title")
	getByTitleCmd.MarkFlagRequired("title")

	//getFullListCmd получение всех данных
	RootCmd.AddCommand(getFullListCmd)

	//getMeCmd получение данных о текущей сессии
	RootCmd.AddCommand(getMeCmd)

	//loginCmd зайти в систему
	RootCmd.AddCommand(loginCmd)
	loginCmd.Flags().StringVarP(&login, "login", "l", "", "it is string to reverse")
	loginCmd.Flags().StringVarP(&password, "password", "p", "", "it is string to reverse")
	loginCmd.MarkFlagRequired("login")
	loginCmd.MarkFlagRequired("password")

	//logoutCmd выйти из системы
	RootCmd.AddCommand(logoutCmd)

	//registerCmd регистрация нового пользователя в системе
	RootCmd.AddCommand(registerCmd)
	registerCmd.Flags().StringVarP(&registerReq.Login, "login", "l", "", "it is string to reverse")
	registerCmd.Flags().StringVarP(&registerReq.Password, "password", "p", "", "it is string to reverse")
	registerCmd.MarkFlagRequired("login")
	registerCmd.MarkFlagRequired("password")

	cmd.Execute()

	return nil
}
