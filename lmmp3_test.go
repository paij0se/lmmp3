package lmmp3

import (
	"os/exec"
	"runtime"
	"testing"
)

func Test_lmmp3(t *testing.T) {
	DownloadFile("skull.gif", "https://c.tenor.com/XmEgf6XjPRQAAAAd/skull.gif")
	// This is only for windows
	DownloadFFmpeg()
	DownloadAndConvert("https://www.youtube.com/watch?v=ZKjIHQxG_3Q")
	DownloadAndConvert("https://www.youtube.com/watch?v=HNQY7afHRvo")
	if runtime.GOOS == "windows" {
		del := exec.Command("cmd", "/C", "del", "*.mpeg")
		if del.Run() != nil {
			panic("failed to delete files")
		}
	}
}
