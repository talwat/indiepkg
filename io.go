package main

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
)

func newFile(path string, text string, errMsg string, params ...interface{}) {
	debugLog("Creating file %s...", bolden(path))
	debugLog("Text to write:\n%s", text)
	err := ioutil.WriteFile(path, []byte(text), 0o770)
	errorLog(err, fmt.Sprintf(errMsg, params...))
}

func newDir(path string, errMsg string, params ...interface{}) {
	debugLog("Creating directory %s...", bolden(path))
	err := os.MkdirAll(path, 0o770)
	errorLog(err, fmt.Sprintf(errMsg, params...))
}

func safeNewDir(path string, name string, ignore bool, text string) {
	if !pathExists(path, name) || ignore {
		log(1, "Creating %s file...", bolden(name))
		newFile(path, text, "An error occurred while creating %s file", bolden(name))
	} else {
		debugLog("Skipping creation of %s.", bolden(path))
	}
}

func copyFile(src string, dst string) {
	debugLog("Copying %s to %s...", bolden(src), bolden(dst))
	debugLog("Statting %s...", bolden(src))
	sourceFileStat, err := os.Stat(src)
	errorLog(err, "Unable to stat file %s", src)

	if !sourceFileStat.Mode().IsRegular() {
		errorLogRaw("File %s is not a regular file, can't copy", src)
	}

	debugLog("Opening files %s...", bolden(src))
	source, err := os.Open(src)
	errorLog(err, "An error occurred while opening file %s", src)

	defer source.Close()

	destination, err := os.Create(dst)
	errorLog(err, "An error occurred while creating file %s", dst)

	defer destination.Close()

	debugLog("Copying...")

	_, err = io.Copy(destination, source)
	errorLog(err, "An error occurred while copying file %s to %s", dst, src)
}

func readFile(file string, errMsg string, params ...interface{}) string {
	debugLog("Reading %s...", bolden(file))
	data, err := ioutil.ReadFile(file)
	errorLog(err, fmt.Sprintf(errMsg, params...))
	debugLog("File contents of %s: %s", bolden(file), bolden(string(data)))

	return string(data)
}

func delPath(silent bool, path string, errMsg string, params ...interface{}) {
	debugLog("Deleting %s...", bolden(path))

	if err := os.RemoveAll(path); !silent {
		errorLog(err, fmt.Sprintf(errMsg, params...))
	}
}

func mvPath(path string, destPath string) {
	debugLog("Moving %s to %s", bolden(path), bolden(destPath))
	err := os.Rename(path, destPath)
	errorLog(err, "An error occurred while moving %s to %s", bolden(path), bolden(destPath))
}

func pathExists(path string, fileName string, params ...interface{}) bool {
	debugLog("Checking if %s exists...", bolden(path))
	_, err := os.Stat(path)

	if err == nil {
		return true
	}

	if errors.Is(err, os.ErrNotExist) {
		return false
	}

	errorLog(err, "An error occurred while checking if %s file exists", fileName)

	return false
}

func dirContents(dir string, errMsg string) []fs.FileInfo {
	debugLog("Getting contents of %s...", bolden(dir))
	files, err := ioutil.ReadDir(dir)
	errorLog(err, errMsg)

	return files
}

func changePerms(file string, perms fs.FileMode) {
	debugLog("Changing permissions of %s to %s...", bolden(file), bolden(perms))
	err := os.Chmod(file, perms)
	errorLog(err, "An error occurred while changing permissions for the file %s", bolden(file))
}

func appendToFile(file string, text string, params ...interface{}) {
	toAppend := fmt.Sprintf(text, params...)

	debugLog("Appending %s to %s", bolden(toAppend), bolden(file))
	fileObj, err := os.OpenFile(file,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	errorLog(err, "An error occurred while opening file %s", bolden(file))

	defer fileObj.Close()

	_, err = fileObj.WriteString(toAppend + "\n")
	errorLog(err, "An error occurred while writing to %s", bolden(file))
}
