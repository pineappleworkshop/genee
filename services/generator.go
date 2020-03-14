package services

import (
	"fmt"
	"genee/models"
	"io/ioutil"
	"os"
	"strings"

	//"genee/models"
)

func GenerateRoot(path string) error {
	//fmt.Println("-=-=-=-=-")
	//fmt.Println("GenerateRoot")
	//fmt.Println(path)
	//fmt.Println("-=-=-=-=-")
	if err := os.MkdirAll(path, 0777); err != nil {
		return err
	}

	return nil
}

func GenerateDirs(path string, dirs []string) error {
	//fmt.Println("-=-=-=-=-")
	//fmt.Println("GenerateDirs")
	//fmt.Println(path)
	//for _, dir := range dirs {
	//	fmt.Println(dir)
	//}
	//fmt.Println("-=-=-=-=-")
	var dirsParsed []string
	for _, dir := range dirs {
		dirArray := strings.Split(dir, "/")
		var dirParsed string
		for i, dirItem := range dirArray {
			if i == 0 {
				continue
			}
			dirParsed = fmt.Sprintf("%s/%s", dirParsed, dirItem)
		}

		dirsParsed = append(dirsParsed, dirParsed)
	}

	for _, dirParsed := range dirsParsed {
		fullPath := fmt.Sprintf("%s/%s", path, dirParsed)
		if err := os.MkdirAll(fullPath, 0777); err != nil {
			return err
		}
	}

	return nil
}

func GenerateFiles(c *models.Config, template, destination string, filePaths []string) error {
	//fmt.Println("-=-=-=-=-")
	//fmt.Println("GenerateDirs")
	//fmt.Println(template)
	//fmt.Println(destination)
	//for _, filePath := range filePaths {
	//	fmt.Println(filePath)
	//}
	//fmt.Println("-=-=-=-=-")
	var files []models.File
	for _, filesPath := range filePaths {
		fileArray := strings.Split(filesPath, "/")
		var fileParsed string
		for i, fileItem := range fileArray {
			if i == 0 {
				continue
			}
			fileParsed = fmt.Sprintf("%s/%s", fileParsed, fileItem)
		}

		file := models.File{}
		file.Template = fmt.Sprintf("%s%s", template, fileParsed)
		file.Destination = fmt.Sprintf("%s%s", destination, fileParsed)

		files = append(files, file)
	}

	for _, file := range files {
		data, err := ioutil.ReadFile(file.Template)
		if err != nil {
			return err
		}

		file.FileStr = string(data)
		if err := file.SearchReplaceVars(c); err != nil {
			return err
		}

		if err := ioutil.WriteFile(file.Destination, []byte(file.FileStr), 0777); err != nil {
			return err
		}

		fmt.Println(fmt.Sprintf("File: %s generated", file.Destination))
	}

	return nil
}
