package lmmp3

import (
	"os/exec"
	"runtime"
	"testing"
)

func Test_lmmp3(t *testing.T) {
	DownloadFFmpeg()
	DownloadAndConvert("https://www.youtube.com/watch?v=829pvBHyG6I")
	DownloadAndConvert("https://www.youtube.com/watch?v=0b-qUiKGy3o")
	DownloadAndConvert("https://www.youtube.com/watch?v=AFujPyUFepc")
	if runtime.GOOS == "windows" {
		del := exec.Command("cmd", "/C", "del", "*.mpeg")
		if del.Run() != nil {
			panic("failed to delete files")
		}
	}
}
