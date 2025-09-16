package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/kamildemocko/renamefiles/internal/backuper"
)

var (
	pattern string
	dir     string
)

func RenameDirPattern(dir string, pattern string) error {
	files, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	bckpr, err := backuper.NewBackuper(dir, "backup.zip")
	if err != nil {
		return err
	}
	defer bckpr.Close()

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
		fmt.Printf("rename '%s' to '%s'\n", filename, newFilename)

		path := filepath.Join(dir, filename)
		newPath := filepath.Join(dir, string(newFilename))

		err = bckpr.AddFile(path)
		if err != nil {
			return err
		}

		err = os.Rename(path, newPath)
		if err != nil {
			return err
		}
	}

	return nil
}

func parseArgs() error {
	var err error

	flag.StringVar(&pattern, "p", "", "example: _(NN) | _XXX")
	flag.Parse()

	args := flag.Args()
	dir = strings.Join(args, " ")
	if dir == "" {
		dir, err = os.Getwd()
		if err != nil {
			return err
		}
	}

	if pattern == "" {
		flag.Usage()
		os.Exit(1)
	}

	return nil
}

func main() {
	err := parseArgs()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		os.Exit(1)
	}

	err = RenameDirPattern(dir, pattern)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error renaming files: %s", err)
		os.Exit(1)
	}
}
