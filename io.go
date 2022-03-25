package main

import (
	"io/ioutil"
)

func newFile(file string, text string) error {
	err := ioutil.WriteFile(file, []byte(text), 0664)
	return err
}

func readFile(file string) (string, error) {
	data, err := ioutil.ReadFile(file)
	return string(data), err
}
