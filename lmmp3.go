package lmmp3

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"regexp"
	"runtime"

	"github.com/kkdai/youtube/v2"
)

var (
	ytregex = regexp.MustCompile(`(http:|https:)?\/\/(www\.)?(youtube.com|youtu.be)\/(watch)?(\?v=)?(\S+)?`)
	stdout  bytes.Buffer
	stderr  bytes.Buffer
	version = "0.0.1"
)

func searchffmpeg() {
	path, err := exec.LookPath("ffmpeg")
	if err != nil {
		fmt.Println("ffmpeg not found", path)
	}
}
func Version() {
	fmt.Println(version)
}

func DownloadAndConvert(url string) {
	searchffmpeg()
	if !ytregex.MatchString(url) {
		fmt.Println("not a youtube url")
	}
	videoID := url
	client := youtube.Client{}

	video, err := client.GetVideo(videoID)
	if err != nil {
		panic(err)
	}

	formats := video.Formats.WithAudioChannels()
	stream, _, err := client.GetStream(video, &formats[2])
	if err != nil {
		panic(err)
	}
	fileVideo := video.Title + ".mpeg"
	mp3file := video.Title + ".mp3"
	file, err := os.Create(fileVideo)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = io.Copy(file, stream)
	if err != nil {
		panic(err)
	}
	// convert the file to mp3
	switch runtime.GOOS {
	case "linux", "darwin":
		cmd := exec.Command("ffmpeg", "-i", fileVideo, mp3file)
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr
		if cmd.Run() != nil {
			fmt.Println(err)

		}
	case "windows":
		cmd := exec.Command("cmd", "/c", "ffmpeg -i", fileVideo, mp3file)
		if cmd.Run() != nil {
			fmt.Println(err)

		}
	default:
		fmt.Println("Unknown OS")
	}
	switch runtime.GOOS {
	case "linux", "darwin":
		del := exec.Command("sh", "-c", "rm *.mpeg").Run()
		if del != nil {
			fmt.Println(del)
		}
	case "windows":
		del := exec.Command("cmd", "/c", "del *.mpeg").Run()
		if del != nil {
			fmt.Println(del)
		}
	default:
		fmt.Println("Unknown OS")
	}
}
