package tests

import (
	"fmt"
	"testing"

	"dkl.dklsa.certificates_monster/iternal/config"
)

func TestConfigInit(t *testing.T) {
	config.Init()
	fmt.Println(config.Config.Debug)
	if config.Config.Env == "" || config.Config.Env != "local" {
		t.Error("Expected DBConfig.Server to be set")
	}
	if config.Config.Storages.Server == "" || config.Config.Storages.Server != "DESKTOP-6LQKT1U" {
		t.Error("Expected DBConfig.Server to be set")
	}
	if config.Config.Storages.User == "" || config.Config.Storages.User != "user" {
		t.Error("Expected DBConfig.User to be set")
	}
	if config.Config.Storages.Port == 0 || config.Config.Storages.Port != 1433 {
		t.Error("Expected DBConfig.Port to be set")
	}
	if config.Config.Http_server.Adress == "" || config.Config.Http_server.Adress != "localhost" {
		t.Error("Expected DBConfig.Adress to be set")
	}
	if config.Config.Http_server.Port == 0 || config.Config.Http_server.Port != 8080 {
		t.Error("Expected DBConfig.Port to be set")
	}

}
