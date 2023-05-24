package passbook

import (
	"github.com/matryer/is"
	"testing"
)

func TestRootCmd(t *testing.T) {
	is := is.New(t)
	root := RootCmd
	err := Execute(root)

	is.NoErr(err)
}
