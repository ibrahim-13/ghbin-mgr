package util

import (
	"fmt"
	"os"
	"path/filepath"
)

type Configuration struct {
	InstallDir string
}

const (
	__config_install_dir  string = "ghbin"
	__config_log_filename string = "ghbin-mgr.log"
)

func newConfiguration() *Configuration {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(fmt.Errorf("unable to get home directory: %w", err))
	}
	installDirPath := filepath.Join(home, __config_install_dir)
	info, err := os.Stat(installDirPath)
	if err != nil {
		if os.IsNotExist(err) {
			if os.Mkdir(installDirPath, 0777) != nil {
				panic("could not create installation directory: " + installDirPath)
			}
		} else {
			panic(fmt.Errorf("unable to stat installation directory: %s: %w", installDirPath, err))
		}
	}
	if !info.IsDir() {
		panic("directory is a file: " + installDirPath)
	}
	return &Configuration{
		InstallDir: installDirPath,
	}
}
