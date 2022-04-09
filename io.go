package main

import (
	"errors"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
)

func newFile(file string, text string, errMsg string, params ...interface{}) {
	err := ioutil.WriteFile(file, []byte(text), 0770)
	errorLog(err, 4, fmt.Sprintf(errMsg, params...))
}

func newDir(name string, errMsg string, params ...interface{}) {
	err := os.MkdirAll(name, 0770)
	errorLog(err, 4, fmt.Sprintf(errMsg, params...))
}

func newDirSilent(name string) {
	os.MkdirAll(name, 0770) //nolint:errcheck
}

func readFile(file string, errMsg string, params ...interface{}) string {
	data, err := ioutil.ReadFile(file)
	errorLog(err, 4, fmt.Sprintf(errMsg, params...))
	return string(data)
}

func delFile(file string, errMsg string, params ...interface{}) {
	err := os.Remove(file)
	errorLog(err, 4, fmt.Sprintf(errMsg, params...))
}

func delDir(dir string, errMsg string, params ...interface{}) {
	err := os.RemoveAll(dir)
	errorLog(err, 4, fmt.Sprintf(errMsg, params...))
}

func pathExists(path string, errMsg string, params ...interface{}) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if errors.Is(err, os.ErrNotExist) {
		return false
	}
	errorLog(err, 4, fmt.Sprintf(errMsg, params...))
	return false
}

func dirContents(dir string, errMsg string, params ...interface{}) []fs.FileInfo {
	files, err := ioutil.ReadDir(dir)
	errorLog(err, 4, fmt.Sprintf(errMsg, params...))
	return files
}
