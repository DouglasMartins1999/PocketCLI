package helpers

import (
	"io"
	"os"

	"github.com/samber/lo"
)

func copy(source string, destination string) error {
	src := lo.Must(os.Open(source))
	defer src.Close()

	dest := lo.Must(os.Create(destination))
	defer dest.Close()

	lo.Must(io.Copy(dest, src))
	return dest.Sync()
}

func Backup(source string, dest string, backup bool) string {
	if dest != "" {
		return dest
	}

	if backup {
		copy(source, source+".bak")
	}

	return source
}
