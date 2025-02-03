package logger

import (
	"fmt"
	"log/slog"
	"os"
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
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		slog.Error(err.Error())
	}
	slog.NewJSONHandler(f, nil)
	// todo: log
}
