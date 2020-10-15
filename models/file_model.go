package models

import (
	"bytes"
	"text/template"
)

type File struct {
	Template    string
	Destination string
	FileStr     string
}

func (f *File) SearchReplaceVars(c *Config) error {
	t, err := template.New("").Parse(f.FileStr)
	if err != nil {
		return err
	}
	var fileParsed bytes.Buffer
	err = t.Execute(&fileParsed, c.Vars)
	if err != nil {
		return err
	}

	f.FileStr = fileParsed.String()

	return nil
}
