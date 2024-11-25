package util

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type AppCtx struct {
	logFile *os.File
	Log     *log.Logger
	Conf    *Configuration
	State   *InstallState
}

func NewAppCtx(logToFile bool) *AppCtx {
	conf := newConfiguration()
	var logFile *os.File
	var logger *log.Logger
	if logToFile {
		logFilePath := filepath.Join(conf.InstallDir, __config_log_filename)
		logOutput, err := os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			panic(fmt.Errorf("unable to open log file: %s: %w", logFilePath, err))
		}
		logFile = logOutput
		logger = log.New(logOutput, "", log.LstdFlags|log.Lshortfile)
	} else {
		logger = log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
	}

	state := newInstallState(conf.InstallDir)
	ctx := &AppCtx{
		logFile: logFile,
		Conf:    conf,
		Log:     logger,
		State:   state,
	}
	return ctx
}

func (ctx *AppCtx) Cleanup() {
	if ctx.logFile != nil {
		ctx.logFile.Close()
	}
}
