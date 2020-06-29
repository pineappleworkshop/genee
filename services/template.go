package services

import (
	"os"
	"path/filepath"
	"strings"
)

func ParseTemplateDirectory(path string) ([]string, []string, error) {
	var dirs, files []string
	err := filepath.Walk(path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				if info.Name() != ".git" {
					dirs = append(dirs, path)
				}
			} else {
				if !strings.Contains(path, ".git") {
					files = append(files, path)
				}
			}
			return nil
		},
	)
	if err != nil {
		return nil, nil, err
	}

	return dirs[1:], files, nil
}
