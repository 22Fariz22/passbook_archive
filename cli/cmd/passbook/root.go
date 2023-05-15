package passbook

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "passbook",
	Short: "passbook help you to keep your secrets",
	Long:  "a longer description for passbook",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("passbook is ready")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
