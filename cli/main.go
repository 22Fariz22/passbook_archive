package main

import (
	"github.com/22Fariz22/passbook/cli/cmd/passbook"
)

func main() {
	passbook.Execute()
}

//examples
// go run main.go register  --login leo --password qwerty
//go run main.go login  --login leo --password qwerty
// go run main.go account --title vk.ru --login leo --password qwerty
// go run main.go text --title mybook --data "lorem iposum dolor"
