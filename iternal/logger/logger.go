package logger

import (
	"fmt"
	"time"

	"dkl.dklsa.certificates_monster/iternal/config"
)

const (
	LevelDebug = "DEBUG"
	LevelInfo  = "INFO"
	LevelWarn  = "WARNING"
	LevelError = "ERROR"
)

func Init() {
	pathLog := config.Config.Logger.Path
	name := time.Now().Format("20060102")
	logType := config.Config.Logger.Type
	path := fmt.Sprintf("%s%s%s", pathLog, name, logType)
	fmt.Println("Initializing logger...")
	fmt.Println(path)
	// todo: log
}
