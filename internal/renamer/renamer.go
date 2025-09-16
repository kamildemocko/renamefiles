package renamer

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/kamildemocko/renamefiles/internal/backuper"
)

type Renamer struct {
	Dir     string
	Pattern string
}

func NewRenamer(dir, pattern string) Renamer {
	replacer := strings.NewReplacer(
		"X", "[A-z]",
		"N", "[0-9]",
		"(", `\(`,
		")", `\)`,
	)
	pattern = replacer.Replace(pattern)

	return Renamer{dir, pattern}
}

func (r *Renamer) RenameDirPattern() error {
	files, err := os.ReadDir(r.Dir)
	if err != nil {
		return err
	}

	bckpr, err := backuper.NewBackuper(r.Dir, "backup.zip")
	if err != nil {
		return err
	}
	defer bckpr.Close()

	for _, file := range files {
		if file.Type().IsDir() {
			continue
		}

		filename := file.Name()

		re, err := regexp.Compile(r.Pattern)
		if err != nil {
			return err
		}

		if !re.Match([]byte(filename)) {
			continue
		}

		newFilename := re.ReplaceAll([]byte(filename), []byte(""))
		fmt.Printf("rename '%s' to '%s'\n", filename, newFilename)

		path := filepath.Join(r.Dir, filename)
		newPath := filepath.Join(r.Dir, string(newFilename))

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
