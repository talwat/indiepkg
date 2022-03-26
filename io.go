package main

import (
	"errors"
	"io/fs"
	"io/ioutil"
	"os"
)

func newFile(file string, text string) error {
	return ioutil.WriteFile(file, []byte(text), 0770)
}

func newDir(name string) error {
	return os.MkdirAll(name, 0770)
}

func readFile(file string) (string, error) {
	data, err := ioutil.ReadFile(file)
	return string(data), err
}

func delFile(file string) error {
	return os.Remove(file)
}

func delDir(dir string) error {
	return os.RemoveAll(dir)
}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return false, err
}

func dirContents(dir string) ([]fs.FileInfo, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	return files, nil
}
