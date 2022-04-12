package main

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func viewFile(url string, errMsg string, params ...interface{}) (string, error) {
	resp, err := http.Get(url)
	errMsgAdded := fmt.Sprintf(errMsg, params...) + "\n    URL: " + url
	errorLog(err, 4, errMsgAdded)

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("HTTP Error. Code: " + fmt.Sprint(resp.StatusCode))
	}

	final, err := ioutil.ReadAll(resp.Body)

	return string(final), err
}

func downloadFile(filepath string, url string, errMsg string, params ...interface{}) error {
	resp, err := http.Get(url)
	errMsgAdded := fmt.Sprintf(errMsg, params...) + "\n    URL: " + url
	errorLog(err, 4, errMsgAdded)

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return (errors.New("HTTP Error. Code: " + fmt.Sprint(resp.StatusCode)))
	}

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}

	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return err
}
