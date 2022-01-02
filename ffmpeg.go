package lmmp3

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
)

func DownloadFile(filepath string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	return err
}

// This func is going to download ffmpeg if is not installed (only in windows)
func FFmpeg() error {
	// https://valledupar.tk/monda/ffmpeg/ffmpeg.exe
	if runtime.GOOS == "windows" {
		// check if ffmpeg is installed
		_, err := os.Stat("ffmpeg.exe")
		if err != nil {
			fmt.Println("downloading ffmpeg...")
			DownloadFile("ffmpeg.exe", "https://valledupar.tk/monda/ffmpeg/ffmpeg.exe")
		}
		fmt.Println("ffmpeg installed")

	}
	return nil
}
