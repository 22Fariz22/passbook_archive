package passbook

import (
	"fmt"
	"github.com/spf13/cobra"
)

var version = "v0.0.1  20.05.2023"

var RootCmd = &cobra.Command{
	RunE: RootCmdRunE,
}

func RootCmdRunE(cmd *cobra.Command, args []string) error {
	fmt.Println("passbook is ready")
	return nil
}

func Execute() {
	cobra.CheckErr(RootCmd.Execute())
}
