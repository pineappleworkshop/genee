package services

import (
	"os"
	"path/filepath"
)

func ParseTemplateDirectory(path string) ([]string, []string, error) {
	var dirs, files []string
	err := filepath.Walk(path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				dirs = append(dirs, path)
			} else {
				files = append(files, path)
			}
			return nil
		},
	)
	if err != nil {
		return nil, nil, err
	}

	return dirs[1:], files, nil
}
