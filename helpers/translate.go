package helpers

import (
	"io/ioutil"
	"strings"

	gt "github.com/bas24/googletranslatefree"
	"github.com/samber/lo"
)

func Translate(source string, dest string, lang string, chunk int, ext string, backup bool) error {
	sFile := string(lo.Must(ioutil.ReadFile(source)))
	lines := strings.Split(sFile, "\n")
	grups := lo.Chunk(lines, chunk)
	result := lo.Reduce(grups, func(agg string, x []string, _ int) string {
		txt := strings.Join(x, "\n")
		return agg + lo.Must(gt.Translate(txt, "en", lang)) + "\n"
	}, "")

	dest = lo.Ternary(dest == "", source+ext, dest)
	dest = Backup(source, dest, backup)
	return ioutil.WriteFile(dest, []byte(result), 0644)
}
