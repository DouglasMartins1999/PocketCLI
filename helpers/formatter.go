package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"

	"github.com/samber/lo"
)

func Stylish(source string, dest string, ssheet string, delim string, backup bool) error {
	rgxp := regexp.MustCompile("Format: .{1,}\nStyle: .{1,}")

	ssheet = string(lo.Must(ioutil.ReadFile(ssheet)))
	sFile := string(lo.Must(ioutil.ReadFile(source)))

	blocks := strings.Split(sFile, delim)
	blocks[1] = rgxp.ReplaceAllString(blocks[1], ssheet)
	result := strings.Join(blocks, delim)

	dest = Backup(source, dest, backup)
	return ioutil.WriteFile(dest, []byte(result), 0644)
}

func Format(source string, dest string, moves string, styles string, backup bool) error {
	mJson := lo.Must(ioutil.ReadFile(moves))
	sJson := lo.Must(ioutil.ReadFile(styles))
	source = string(lo.Must(ioutil.ReadFile(source)))

	var mList [][]string
	var sDict map[string]string

	json.Unmarshal(mJson, &mList)
	json.Unmarshal(sJson, &sDict)

	fmt.Println(mList)

	for _, m := range mList {
		text := sDict[m[1]] + m[0] + "{\\r}"
		rgxp := regexp.MustCompile("\\b" + m[0] + "\\b")

		fmt.Println(text, rgxp)
		source = rgxp.ReplaceAllString(source, text)
	}

	dest = Backup(source, dest, backup)
	return ioutil.WriteFile(dest, []byte(source), 0644)
}

func Clean(source string, dest string, tags bool, lines bool, backup bool) error {
	source = string(lo.Must(ioutil.ReadFile(source)))

	if tags {
		rgxp := regexp.MustCompile(`{\\(.*?)}`)
		source = rgxp.ReplaceAllString(source, "")
	}

	if lines {
		source = strings.ReplaceAll(source, "\\N", " ")
	}

	dest = Backup(source, dest, backup)
	return ioutil.WriteFile(dest, []byte(source), 0644)
}
