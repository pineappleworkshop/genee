package models

import (
	"bufio"
	"errors"
	"fmt"
	"strings"
)

type Config struct {
	Vars     []Var    `yaml:"vars"`
	Commands []string `yaml:"commands"`
}

type Var struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
}

func (c *Config) SearchReplaceVars() error {
	for i, command := range c.Commands {
		count := c.countVars(command)
		commandParsed, err := c.parseAndReplaceVar(command)
		if err != nil {
			return err
		}

		for i := 0; i < count; i++ {
			commandParsed, err = c.parseAndReplaceVar(commandParsed)
			if err != nil {
				return err
			}
		}
		c.Commands[i] = commandParsed
	}

	return nil
}

func (c *Config) parseAndReplaceVar(command string) (string, error) {
	data := bufio.NewScanner(strings.NewReader(command))
	data.Split(bufio.ScanRunes)

	begin1 := -1
	begin2 := -1
	begin3 := -1
	end1 := -1
	end2 := -1
	end3 := -1

	index := 0
	for data.Scan() {
		if data.Text() == "{" {
			if begin1 == -1 {
				begin1 = index
				index++
				continue
			}
		}
		if begin1 != -1 {
			if data.Text() == "{" {
				if begin2 == -1 {
					begin2 = index
					index++
					continue
				}
			}
			if begin2 != -1 {
				if data.Text() == "<" {
					begin3 = index
					index++
					continue
				}
			}
		}
		if data.Text() == ">" {
			if end1 == -1 {
				end1 = index
				index++
				continue
			}
		}
		if end1 != -1 {
			if data.Text() == "}" {
				if end2 == -1 {
					end2 = index
					index++
					continue
				}
			}
			if end2 != -1 {
				if data.Text() == "}" {
					end3 = index
					index++
					continue
				}
			}
		}

		if end3 != -1 {
			break
		}
		index++
	}

	data = bufio.NewScanner(strings.NewReader(command))
	data.Split(bufio.ScanRunes)
	varName := ""
	index = 0
	for data.Scan() {
		if index > begin3 {
			if index < end1 {
				varName = fmt.Sprintf("%s%s", varName, data.Text())
			}
		}
		index++
	}

	data = bufio.NewScanner(strings.NewReader(command))
	data.Split(bufio.ScanRunes)
	commandParsed := ""
	index = 0
	for data.Scan() {
		if index == begin1 {
			varValue, err := findVarValue(varName, c)
			if err != nil {
				return "", err
			}
			commandParsed = fmt.Sprintf("%s%s", commandParsed, varValue)
		}

		if index >= begin1 {
			if index <= end3 {
				index++
				continue
			}
		}

		commandParsed = fmt.Sprintf("%s%s", commandParsed, data.Text())
		index++
	}
	return commandParsed, nil
}

func findVarValue(varName string, c *Config) (string, error) {
	var varValue string
	for _, xvar := range c.Vars {
		if xvar.Name == varName {
			varValue = xvar.Value
		}
	}

	if varValue == "" {
		return "", errors.New(fmt.Sprintf("`%s` was not found in config", varName))
	}
	return varValue, nil
}

func (c *Config) countVars(command string) int {
	count := 0

	data := bufio.NewScanner(strings.NewReader(command))
	data.Split(bufio.ScanRunes)

	begin1 := -1
	begin2 := -1
	begin3 := -1
	end1 := -1
	end2 := -1
	end3 := -1

	index := 0
	for data.Scan() {
		if data.Text() == "{" {
			if begin1 == -1 {
				begin1 = index
				index++
				continue
			}
		}
		if begin1 != -1 {
			if data.Text() == "{" {
				if begin2 == -1 {
					begin2 = index
					index++
					continue
				}
			}
			if begin2 != -1 {
				if data.Text() == "<" {
					begin3 = index
					index++
					continue
				}
			}
		}
		if data.Text() == ">" {
			if end1 == -1 {
				end1 = index
				index++
				continue
			}
		}
		if end1 != -1 {
			if data.Text() == "}" {
				if end2 == -1 {
					end2 = index
					index++
					continue
				}
			}
			if end2 != -1 {
				if data.Text() == "}" {
					end3 = index
					index++
					continue
				}
			}
		}

		if end3 != -1 {
			begin1 = -1
			begin2 = -1
			begin3 = -1
			end1 = -1
			end2 = -1
			end3 = -1
			count++
		}
		index++
	}

	if begin3 == -1 {
		return count
	}

	return count
}
