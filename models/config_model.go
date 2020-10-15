package models

import (
	"bytes"
	"text/template"
)

type Config struct {
	Vars     map[string]interface{} `yaml:"vars"`
	Commands []string               `yaml:"commands"`
}

func (c *Config) SearchReplaceVars() error {
	for i, command := range c.Commands {
		t, err := template.New("").Parse(command)
		if err != nil {
			return err
		}
		var commandParsed bytes.Buffer
		err = t.Execute(&commandParsed, c.Vars)
		if err != nil {
			return err
		}

		c.Commands[i] = commandParsed.String()
	}

	return nil
}
