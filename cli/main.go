package main

import (
	"github.com/22Fariz22/passbook/cli/cmd/passbook"
)

func main() {
	root := passbook.RootCmd

	passbook.Execute(root)
}
