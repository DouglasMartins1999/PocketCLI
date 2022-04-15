package helpers

import (
	"errors"
	"io/ioutil"
	"strings"

	"github.com/samber/lo"
)

const delimitter = "Format: Layer, Start, End, Style, Name, MarginL, MarginR, MarginV, Effect, Text"

func Extract(source string, dest string, ext string, backup bool) error {
	file, err := ioutil.ReadFile(source)

	if err != nil {
		return err
	}

	if dest == "" {
		dest = source + ext
	}

	dialogs := strings.Split(string(file), delimitter)[1]
	lines := strings.Split(dialogs, "\n")

	lines = lo.Filter(lines, func(x string, _ int) bool {
		return len(x) > 0
	})

	lines = lo.Map(lines[1:], func(x string, _ int) string {
		return strings.Split(x, ",,0,0,0,,")[1]
	})
	dialogs = strings.Join(lines, "\n")

	dest = Backup(source, dest, backup)
	return ioutil.WriteFile(dest, []byte(dialogs), 0644)
}

func Restore(source string, dest string, backup bool) error {
	if source == "" {
		return errors.New("você precisa especificar o .txt de falas")
	}

	if dest == "" {
		return errors.New("você precisa especificar o arquivo original de legendas")
	}

	sFile, err := ioutil.ReadFile(source)

	if err != nil {
		return err
	}

	dFile, err := ioutil.ReadFile(dest)

	if err != nil {
		return err
	}

	slines := strings.Split(string(sFile), "\n")
	dialogs := strings.Split(string(dFile), delimitter)

	lines := strings.Split(dialogs[1], "\n")

	lines = lo.Filter(lines, func(x string, _ int) bool {
		return len(x) > 0
	})

	lines = lo.Map(lines, func(x string, i int) string {
		y := strings.Split(x, ",,0,0,0,,")
		y[1] = slines[i]
		return strings.Join(y, ",,0,0,0,,")
	})

	dialogs[1] = strings.Join(lines, "\n")
	dFile = []byte(strings.Join(dialogs, delimitter))
	dest = Backup(source, dest, backup)
	return ioutil.WriteFile(dest, dFile, 0644)
}
