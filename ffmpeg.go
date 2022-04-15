package lmmp3

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"runtime"

	"github.com/schollz/progressbar/v3"
)

// This function is going to work like curl making a GET http request
// To the server and create a file with the output
func DownloadFile(filepath string, url string) error {
	req, _ := http.NewRequest("GET", url, nil)
	resp, _ := http.DefaultClient.Do(req)
	f, _ := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()

	bar := progressbar.DefaultBytes(
		resp.ContentLength,
		"Downloading "+filepath,
	)
	io.Copy(io.MultiWriter(f, bar), resp.Body)
	return nil

}

// This func is going to download ffmpeg if is not installed (only in windows)
func DownloadFFmpeg() error {
	// https://valledupar.tk/monda/ffmpeg/ffmpeg.exe
	if runtime.GOOS == "windows" {
		_, err := exec.LookPath("ffmpeg.exe")
		if err != nil {
			fmt.Println("Downloading ffmpeg")
			DownloadFile("ffmpeg.exe", "https://valledupar.tk/monda/ffmpeg/ffmpeg.exe")
		}
	}
	return nil
}
