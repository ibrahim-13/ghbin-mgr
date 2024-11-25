package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type InstallState struct {
	filePath string `json:"-"`
	Version  int    `json:"version,omitempty"`
}

const (
	__state_filename string = "ghbin-mgr-state.json"
	__state_version  int    = 1
)

func newInstallState(installDir string) *InstallState {
	filePath := filepath.Join(installDir, __state_filename)
	stat, err := os.Stat(filePath)
	if err != nil && os.IsNotExist(err) {
		return &InstallState{filePath: filePath, Version: __state_version}
	} else if err != nil {
		panic(fmt.Errorf("unable to read stat of state file: %s: %w", filePath, err))
	}
	if stat.IsDir() {
		panic("state file path is a directory: " + filePath)
	}
	stateFile, err := os.Open(filePath)
	if err != nil {
		panic(fmt.Errorf("unable to open state file: %s: %w", filePath, err))
	}
	defer stateFile.Close()
	bytes, err := io.ReadAll(stateFile)
	if err != nil {
		panic(fmt.Errorf("unable to read bytes of state file: %s: %w", filePath, err))
	}
	var state InstallState
	err = json.Unmarshal(bytes, &state)
	if err != nil {
		panic(fmt.Errorf("unable to unmarshal state file: %s: %w", filePath, err))
	}
	state.filePath = filePath
	return &state
}

func (s *InstallState) Save() error {
	stat, err := os.Stat(s.filePath)
	if err != nil {
		if !os.IsNotExist(err) {
			return fmt.Errorf("unable to read stat of state file: %s: %w", s.filePath, err)
		}
	} else if stat.IsDir() {
		return errors.New("state file path is a directory: " + s.filePath)
	}
	bytes, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return fmt.Errorf("unable to marshal state file: %s: %w", s.filePath, err)
	}
	err = os.WriteFile(s.filePath, bytes, 0666)
	if err != nil {
		return fmt.Errorf("unable to write state file: %s: %w", s.filePath, err)
	}
	return nil
}
