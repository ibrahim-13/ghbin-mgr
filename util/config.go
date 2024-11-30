package util

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

type Os rune

type Configuration struct {
	InstallDir string
	TmpDir     string
	OsType     Os
}

const (
	OsWindows Os = 1
	OsLinux   Os = 2
	OsMac     Os = 3
)

func ensureDirExistsOrPanic(dir string) {
	info, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			if os.Mkdir(dir, 0777) != nil {
				panic("could not create directory: " + dir)
			}
		} else {
			panic(fmt.Errorf("unable to stat directory: %s: %w", dir, err))
		}
	}
	if info != nil && !info.IsDir() {
		panic("directory is a file: " + dir)
	}
}

func NewConfiguration(installDir string) *Configuration {
	tempDirPath := filepath.Join(installDir, "tmp")
	ensureDirExistsOrPanic(installDir)
	ensureDirExistsOrPanic(tempDirPath)
	osTypeStr := runtime.GOOS
	var osType Os
	switch osTypeStr {
	case "darwin":
		osType = OsMac
	case "linux":
		osType = OsLinux
	case "windows":
		osType = OsWindows
	default:
		panic("unknown os type: " + osTypeStr)
	}
	return &Configuration{
		InstallDir: installDir,
		TmpDir:     tempDirPath,
		OsType:     osType,
	}
}
