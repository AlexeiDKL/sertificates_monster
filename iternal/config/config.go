package config

import (
	"fmt"

	"github.com/spf13/viper"
)

const (
	configName = "config"
	configType = "yaml"
	configPath = "./config/"
)

type ConfigStruct struct {
	Env         string
	Logger      LogStruct
	Storages    StorageStruct
	Http_server HttpStruct
	Keys        KeysStruct
}

type LogStruct struct {
	Level string
	Path  string
	Type  string
}

type StorageStruct struct {
	Server               string
	User                 string
	Port                 int
	Password             string
	NumberOfSparePhrases int
}

type HttpStruct struct {
	Adress      string
	Port        int
	ReadTimeout int
	IdleTimeout int
}

type KeysStruct struct {
	Path string
}

// инициализация конфигурации
var Config ConfigStruct

func Init() error {
	// получаем данные из конфигурационного файла
	// и загружаем их в конфигурацию
	// используем package viper

	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
	viper.AddConfigPath(configPath)
	viper.AddConfigPath("../config/")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
		return err
	}

	err := viper.Unmarshal(&Config)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println(Config.Storages.Password)
	return nil
}
