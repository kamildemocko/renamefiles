package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	pattern2 = "_(NN)"
	dir      = `.\test`
)

func RenameDirPattern(dir string, pattern string) error {
	files, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.Type().IsDir() {
			continue
		}

		filename := file.Name()

		replacer := strings.NewReplacer(
			"X", "[A-z]",
			"N", "[0-9]",
			"(", `\(`,
			")", `\)`,
		)
		pattern = replacer.Replace(pattern)
		re, err := regexp.Compile(pattern)
		if err != nil {
			return err
		}

		if !re.Match([]byte(filename)) {
			continue
		}

		newFilename := re.ReplaceAll([]byte(filename), []byte(""))
		fmt.Printf("rename '%s' to '%s'", filename, newFilename)

		path := filepath.Join(dir, filename)
		newPath := filepath.Join(dir, string(newFilename))

		err = os.Rename(path, newPath)
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	err := RenameDirPattern(dir, pattern2)
	if err != nil {
		panic(err)
	}
}
