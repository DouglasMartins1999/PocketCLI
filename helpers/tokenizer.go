package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"
	"regexp"

	"github.com/samber/lo"
)

func Tokenize(source string, dest string, glossary string, backup bool) error {
	var anime [][]string
	var cities [][]string
	var ignores [][]string
	var locations [][]string
	var pokemon []string
	var moves [][]string

	sFile := string(lo.Must(ioutil.ReadFile(source)))

	json.Unmarshal(lo.Must(ioutil.ReadFile(path.Join(glossary, "anime.json"))), &anime)
	json.Unmarshal(lo.Must(ioutil.ReadFile(path.Join(glossary, "cities.json"))), &cities)
	json.Unmarshal(lo.Must(ioutil.ReadFile(path.Join(glossary, "ignores.json"))), &ignores)
	json.Unmarshal(lo.Must(ioutil.ReadFile(path.Join(glossary, "locations.json"))), &locations)
	json.Unmarshal(lo.Must(ioutil.ReadFile(path.Join(glossary, "pokemon.json"))), &pokemon)
	json.Unmarshal(lo.Must(ioutil.ReadFile(path.Join(glossary, "moves.json"))), &moves)

	for i, x := range ignores {
		cd := "_I" + fmt.Sprintf("%02d", i) + "_"
		re := regexp.MustCompile("\\b" + x[0] + "\\b")
		sFile = re.ReplaceAllString(sFile, cd)
	}

	for i, x := range pokemon {
		cd := "_P" + fmt.Sprintf("%02d", i) + "_"
		re := regexp.MustCompile("\\b" + x + "\\b")
		sFile = re.ReplaceAllString(sFile, cd)
	}

	for i, x := range moves {
		cd := "_M" + fmt.Sprintf("%02d", i) + "_"
		re := regexp.MustCompile("\\b" + x[0] + "\\b")
		sFile = re.ReplaceAllString(sFile, cd)
	}

	for i, x := range cities {
		cd := "_C" + fmt.Sprintf("%02d", i) + "_"
		re := regexp.MustCompile("\\b" + x[0] + "\\b")
		sFile = re.ReplaceAllString(sFile, cd)
	}

	for i, x := range locations {
		cd := "_L" + fmt.Sprintf("%02d", i) + "_"
		re := regexp.MustCompile("\\b" + x[0] + "\\b")
		sFile = re.ReplaceAllString(sFile, cd)
	}

	for i, x := range anime {
		cd := "_A" + fmt.Sprintf("%02d", i) + "_"
		re := regexp.MustCompile("\\b" + x[0] + "\\b")
		sFile = re.ReplaceAllString(sFile, cd)
	}

	dest = Backup(source, dest, backup)
	return ioutil.WriteFile(dest, []byte(sFile), 0644)
}

func Untokenize(source string, dest string, glossary string, backup bool) error {
	var anime [][]string
	var cities [][]string
	var ignores [][]string
	var locations [][]string
	var pokemon []string
	var moves [][]string

	sFile := string(lo.Must(ioutil.ReadFile(source)))

	json.Unmarshal(lo.Must(ioutil.ReadFile(path.Join(glossary, "anime.json"))), &anime)
	json.Unmarshal(lo.Must(ioutil.ReadFile(path.Join(glossary, "cities.json"))), &cities)
	json.Unmarshal(lo.Must(ioutil.ReadFile(path.Join(glossary, "ignores.json"))), &ignores)
	json.Unmarshal(lo.Must(ioutil.ReadFile(path.Join(glossary, "locations.json"))), &locations)
	json.Unmarshal(lo.Must(ioutil.ReadFile(path.Join(glossary, "pokemon.json"))), &pokemon)
	json.Unmarshal(lo.Must(ioutil.ReadFile(path.Join(glossary, "moves.json"))), &moves)

	for i, x := range ignores {
		cd := "_I" + fmt.Sprintf("%02d", i) + "_"
		re := regexp.MustCompile("\\b" + cd + "\\b")
		sFile = re.ReplaceAllString(sFile, x[1])
	}

	for i, x := range pokemon {
		cd := "_P" + fmt.Sprintf("%02d", i) + "_"
		re := regexp.MustCompile("\\b" + cd + "\\b")
		sFile = re.ReplaceAllString(sFile, x)
	}

	for i, x := range moves {
		cd := "_M" + fmt.Sprintf("%02d", i) + "_"
		re := regexp.MustCompile("\\b" + cd + "\\b")
		sFile = re.ReplaceAllString(sFile, x[0])
	}

	for i, x := range cities {
		cd := "_C" + fmt.Sprintf("%02d", i) + "_"
		re := regexp.MustCompile("\\b" + cd + "\\b")
		sFile = re.ReplaceAllString(sFile, x[1])
	}

	for i, x := range locations {
		cd := "_L" + fmt.Sprintf("%02d", i) + "_"
		re := regexp.MustCompile("\\b" + cd + "\\b")
		sFile = re.ReplaceAllString(sFile, x[1])
	}

	for i, x := range anime {
		cd := "_A" + fmt.Sprintf("%02d", i) + "_"
		re := regexp.MustCompile("\\b" + cd + "\\b")
		sFile = re.ReplaceAllString(sFile, x[1])
	}

	dest = Backup(source, dest, backup)
	return ioutil.WriteFile(dest, []byte(sFile), 0644)
}
