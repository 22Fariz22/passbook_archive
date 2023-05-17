package main

import (
	"github.com/22Fariz22/passbook/cli/cmd/passbook"
)

func main() {
	passbook.Execute()
}

//examples commands line
// go run main.go register  --login leo --password qwerty
// go run main.go login  --login leo --password qwerty
// go run main.go logout
// go run main.go account --title vk.ru --login leo --password qwerty
// go run main.go text --title mybook --data "lorem iposum dolor"
// go run main.go card --title sber --name "leo de catrio" --card 24234436456457 --date 21/23 --cvc sd
// go run main.go get_title --title "mybook"
// go run main.go full
