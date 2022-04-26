package main

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/schollz/progressbar/v3"
)

func viewFile(url string, errMsg string, params ...interface{}) (string, error) {
	resp, err := http.Get(url)
	errMsgAdded := fmt.Sprintf(errMsg, params...) + ". URL: " + bolden(url)
	errorLog(err, errMsgAdded)

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("HTTP Error. Code: " + fmt.Sprint(resp.StatusCode))
	}

	final, err := ioutil.ReadAll(resp.Body)

	return string(final), err
}

func downloadFileWithProg(filepath string, url string, errMsg string, params ...interface{}) {
	req, err := http.NewRequest("GET", url, nil)
	errorLog(err, "An error occurred making a new GET request")
	resp, err := http.DefaultClient.Do(req)
	errorLog(err, "An error occurred sending GET request")
	defer resp.Body.Close()

	file, _ := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY, 0o644)
	defer file.Close()

	newBar := func(maxBytes int64, description ...string) *progressbar.ProgressBar {
		desc := ""
		if len(description) > 0 {
			desc = logType[2] + description[0]
		}

		bar := progressbar.NewOptions64(
			maxBytes,
			progressbar.OptionSetDescription(desc),
			progressbar.OptionSetWriter(os.Stderr),
			progressbar.OptionShowBytes(true),
			progressbar.OptionSetWidth(10),
			progressbar.OptionThrottle(65*time.Millisecond),
			progressbar.OptionOnCompletion(func() {
				fmt.Printf("\n")
			}),
			progressbar.OptionSpinnerType(14),
			progressbar.OptionFullWidth(),

			progressbar.OptionEnableColorCodes(true),
			progressbar.OptionSetWidth(15),
			progressbar.OptionSetTheme(progressbar.Theme{
				Saucer:        config.Progressbar.Saucer,
				SaucerHead:    config.Progressbar.SaucerHead,
				AltSaucerHead: config.Progressbar.AltSaucerHead,
				SaucerPadding: config.Progressbar.SaucerPadding,
				BarStart:      config.Progressbar.BarStart,
				BarEnd:        config.Progressbar.BarEnd,
			}),
		)
		err := bar.RenderBlank()
		errorLog(err, "An error occurred while rendering loading bar.")

		return bar
	}
	bar := newBar(
		resp.ContentLength,
		" Progress",
	)

	_, err = io.Copy(io.MultiWriter(file, bar), resp.Body)
	errorLog(err, "An error occurred while running %s.", bolden("io.Copy()"))
}
