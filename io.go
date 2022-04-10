package main

import (
	"errors"
	"fmt"
	"io"
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

func copyFile(src string, dst string) {
	sourceFileStat, err := os.Stat(src)
	errorLog(err, 4, "Unable to stat file %s", src)

	if !sourceFileStat.Mode().IsRegular() {
		log(4, "File %s is not a regular file, can't copy.", src)
	}

	source, err := os.Open(src)
	errorLog(err, 4, "An error occurred while opening file %s", src)
	defer source.Close()

	destination, err := os.Create(dst)
	errorLog(err, 4, "An error occurred while creating file %s", dst)
	defer destination.Close()
	_, err = io.Copy(destination, source)
	errorLog(err, 4, "An error occurred while copying file %s to %s", dst, src)
}

func newDirSilent(name string) {
	os.MkdirAll(name, 0770) //nolint:errcheck
}

func readFile(file string, errMsg string, params ...interface{}) string {
	data, err := ioutil.ReadFile(file)
	errorLog(err, 4, fmt.Sprintf(errMsg, params...))
	return string(data)
}

func delPath(logLevel int, path string, errMsg string, params ...interface{}) {
	err := os.RemoveAll(path)
	debugLog("Deleting %s", path)
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
