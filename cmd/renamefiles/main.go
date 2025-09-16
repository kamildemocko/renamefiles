package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/kamildemocko/renamefiles/internal/renamer"
)

var (
	pattern string
	dir     string
)

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

	r := renamer.NewRenamer(dir, pattern)

	err = r.RenameDirPattern()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error renaming files: %s", err)
		os.Exit(1)
	}
}
