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
	errMsgAdded := fmt.Sprintf(errMsg, params...)
	errorLog(err, 4, errMsgAdded)

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("HTTP Error. Code: " + fmt.Sprint(resp.StatusCode))
	}

	final, err := ioutil.ReadAll(resp.Body)

	return string(final), err
}

func downloadFile(filepath string, url string, errMsg string, params ...interface{}) {
	resp, err := http.Get(url)
	errMsgAdded := fmt.Sprintf(errMsg, params...) + "\n    URL: " + url + "\n   "
	errorLog(err, 4, errMsgAdded)

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errorLog(errors.New("HTTP Error. Code: "+fmt.Sprint(resp.StatusCode)), 4, errMsgAdded)
	}

	out, err := os.Create(filepath)
	errorLog(err, 4, errMsgAdded)
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	errorLog(err, 4, errMsgAdded)
}
