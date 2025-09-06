package project

import (
	"os"
	"path/filepath"
)

var projectRoot string

func findProjectRoot() {
	dir, _ := os.Getwd()

	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			projectRoot = dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}
}

func GetProjectRoot() string {
	return projectRoot
}
