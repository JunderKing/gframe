package logger

import (
	"fmt"
	"gframe/pkg/file"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/spf13/viper"
)

type Level int

var (
	F                  *os.File
	DefaultPrefix      = ""
	DefaultCallerDepth = 2
	writer             *log.Logger
	logPrefix          = ""
	levelFlags         = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

// Setup initialize the logs instance
func Setup() {
	var err error
	filePath := getLogFilePath()
	fileName := getLogFileName()
	F, err = file.MustOpen(fileName, filePath)
	if err != nil {
		log.Fatalf("logger.Setup err: %v", err)
	}

	writer = log.New(F, DefaultPrefix, log.LstdFlags)
}

// Debug output logs at debug level
func Debug(v ...interface{}) {
	setPrefix(DEBUG)
	writer.Println(v)
}

// Info output logs at info level
func Info(v ...interface{}) {
	setPrefix(INFO)
	writer.Println(v)
}

// Warn output logs at warn level
func Warn(v ...interface{}) {
	setPrefix(WARNING)
	writer.Println(v)
}

// Error output logs at error level
func Error(v ...interface{}) {
	setPrefix(ERROR)
	writer.Println(v)
}

// Fatal output logs at fatal level
func Fatal(v ...interface{}) {
	setPrefix(FATAL)
	writer.Fatalln(v)
}

// setPrefix set the prefix of the logs output
func setPrefix(level Level) {
	_, files, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(files), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}

	writer.SetPrefix(logPrefix)
}

// getLogFilePath get the logs file save path
func getLogFilePath() string {
	return viper.GetString("log.path")
	//return fmt.Sprintf("%s%s", viper.GetString("logs.path"), setting.AppSetting.LogSavePath)
}

// getLogFileName get the save name of the logs file
func getLogFileName() string {
	return fmt.Sprintf("%s%s.%s",
		viper.GetString("log.name"),
		time.Now().Format(viper.GetString("log.timeFormat")),
		viper.GetString("log.ext"),
	)
}
