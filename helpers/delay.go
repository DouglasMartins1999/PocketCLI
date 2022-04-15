package helpers

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/samber/lo"
)

func parse(d string) time.Duration {
	dots := strings.Split(d, ".")
	sep := strings.Split(dots[0], ":")

	h := time.Duration(lo.Must(strconv.ParseInt(sep[0], 10, 0))) * time.Hour
	m := time.Duration(lo.Must(strconv.ParseInt(sep[1], 10, 0))) * time.Minute
	s := time.Duration(lo.Must(strconv.ParseInt(sep[2], 10, 0))) * time.Second
	ms := time.Duration(lo.Must(strconv.ParseInt(dots[1], 10, 0))) * time.Millisecond * 10

	return h + m + s + ms
}

func stringify(d time.Duration) string {
	z := time.Unix(0, 0).UTC()
	h := d.Truncate(time.Hour).Hours()
	m := z.Add(time.Duration(d)).Format(":04:05.00")

	return fmt.Sprintf("%0.f%v", h, m)
}

func Delay(source string, dest string, delay string, backup bool) error {
	dMap := make(map[string]string)
	duration := parse(delay)

	rgxp := regexp.MustCompile(`\d:\d{2}:\d{2}\.\d{2}`)
	source = string(lo.Must(ioutil.ReadFile(source)))

	times := rgxp.FindAllString(source, -1)

	for _, t := range times {
		d := parse(t) + duration
		dMap[t] = stringify(d)
	}

	for old, new := range dMap {
		source = strings.ReplaceAll(source, old, new)
	}

	dest = Backup(source, dest, backup)
	return ioutil.WriteFile(dest, []byte(source), 0644)
}
