//nolint:bodyclose
package main

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/schollz/progressbar/v3"
)

// Make sure to close the response after calling.
// Not closing automatically because the response body needs to be read.
func makeReq(url string) (http.Response, error) {
	debugLog("Making GET request to %s", bolden(url))

	req, err := http.NewRequestWithContext(context.Background(), "GET", url, nil)
	if err != nil {
		return http.Response{}, fmt.Errorf("error while making http request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return http.Response{}, fmt.Errorf("error while doing http request: %w", err)
	}

	return *resp, nil
}

func makeGithubReq(url string) (http.Response, error) {
	req, err := http.NewRequestWithContext(context.Background(), "GET", url, nil)
	req.SetBasicAuth(config.Github.Username, config.Github.Token)

	if err != nil {
		return http.Response{}, fmt.Errorf("error while making http request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return http.Response{}, fmt.Errorf("error while doing http request: %w", err)
	}

	return *resp, nil
}

func viewFile(url string) (string, int, error) {
	resp, err := makeReq(url)
	if err != nil {
		return "", resp.StatusCode, err
	}

	defer resp.Body.Close()

	final, err := ioutil.ReadAll(resp.Body)

	return string(final), resp.StatusCode, err
}

func downloadFileWithProg(filepath string, url string, errMsg string, params ...interface{}) {
	resp, err := makeReq(url)

	errorLog(err, "An error occurred while making GET request to %s", url)

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
				rawLog("\n")
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
