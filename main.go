package main

import (
	"fmt"

	storage "dkl.dklsa.certificates_monster/iternal/storage/mssql"
)

func initConfig() {}

func initLogger() {}

func main() {
	initConfig()
	initLogger()

	phrase, err := storage.GetPhrase()
	if err != nil {
		fmt.Println("phrase error: ", err)
		return
	}
	fmt.Println(phrase)
}
