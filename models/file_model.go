package models

import (
	"bufio"
	"fmt"
	"strings"
)

type File struct {
	Template    string
	Destination string
	FileStr string
}

func (f *File) SearchReplaceVars(c *Config) error {
	count := f.countVars()
	fileParsed, err := f.parseAndReplaceVar(c, f.FileStr)
	if err != nil {
		return err
	}

	for i := 0; i < count; i++ {
		fileParsed, err = f.parseAndReplaceVar(c, fileParsed)
		if err != nil {
			return err
		}
	}
	f.FileStr = fileParsed

	return nil
}

func (f *File) parseAndReplaceVar(c *Config, fileStr string) (string, error) {
	data := bufio.NewScanner(strings.NewReader(fileStr))
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

	data = bufio.NewScanner(strings.NewReader(fileStr))
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

	data = bufio.NewScanner(strings.NewReader(fileStr))
	data.Split(bufio.ScanRunes)
	commandParsed := ""
	index = 0
	for data.Scan() {
		if index == begin1 {
			//varValue, err := findVarValue(varName, c)
			//if err != nil {
			//	return "", err
			//}
			//commandParsed = fmt.Sprintf("%s%s", commandParsed, varValue)
			if varName != "" {
				varValue, err := findVarValue(varName, c)
				if err != nil {
					return "", err
				}
				commandParsed = fmt.Sprintf("%s%s", commandParsed, varValue)
			}
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

func (f *File) countVars() int {
	count := 0

	data := bufio.NewScanner(strings.NewReader(f.FileStr))
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
