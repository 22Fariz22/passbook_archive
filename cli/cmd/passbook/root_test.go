package passbook

import (
	"github.com/matryer/is"
	"github.com/spf13/cobra"
	"testing"
)

func TestRootCmd(t *testing.T) {
	is := is.New(t)
	root := &cobra.Command{Use: "root", RunE: RootCmdRunE}
	err := root.Execute()
	is.NoErr(err)
}
