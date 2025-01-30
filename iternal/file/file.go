package file

import (
	"fmt"
	"io"
	"log"
	"os"
)

func IsDir(namePath string) bool {
	fileInfo, err := os.Stat(namePath)
	if err != nil {
		return false
	}
	return fileInfo.IsDir()
}

func Exists(namePath string) bool {
	if _, err := os.Stat(namePath); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func CreateFilePath(path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		log.Println("function.CreateFilePath Error: ", err)
		log.Println(path)
	}
	return err
}

func GetTextInFile(path string) string {
	var result string

	data, err := os.ReadFile(path)
	if err != nil {
		log.Println(err)
	}
	result = string(data)
	return result
}

func CopyFile(from, to string) (int64, error) {
	sourceFileStat, err := os.Stat(from)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("not a regular file")
	}

	source, err := os.Open(from)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(to)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

func SaveFile(path string, text string) error {
	file, err := os.Create(path)

	if err != nil {
		fmt.Println("Unable to create file:", err)
		return err
	}
	defer file.Close()
	file.WriteString(text)
	return err
}
