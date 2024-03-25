package util

import (
	"os"
	"path/filepath"
)

func GetTemplatesFilePath(filename string) string {
	wd, _ := os.Getwd()
	return filepath.Join(wd, "src", "templates", filename)
}