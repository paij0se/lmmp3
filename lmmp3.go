package lmmp3

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strings"

	"github.com/kkdai/youtube/v2"
)

var (
	ytregex = regexp.MustCompile(`(http:|https:)?\/\/(www\.)?(youtube.com|youtu.be)\/(watch)?(\?v=)?(\S+)?`)
	Version = "1.0.0"
)

func searchffmpeg() {
	path, err := exec.LookPath("ffmpeg")
	if err != nil {
		fmt.Println("ffmpeg not found", path)
	}
}

// This function is going to download and convert the video to a mp3 file
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

	fileVideo := strings.ReplaceAll(video.Title+".mpeg", "/", "|")
	//So.. the character / is and space, so i need to replace it.
	mp3file := strings.ReplaceAll(video.Title+".mp3", "/", "|")
	fmt.Println(fileVideo + " -> " + mp3file)
	fmt.Println("video downloaded:", video.Title)
	file, err := os.Create(fileVideo)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = io.Copy(file, stream)
	if err != nil {
		panic(err)
	}
	switch runtime.GOOS {
	case "linux", "darwin":
		cmd := exec.Command("ffmpeg", "-i", fileVideo, mp3file)
		if cmd.Run() != nil {
			fmt.Println(err)
		}
	case "windows":
		cmd := exec.Command("ffmpeg.exe", "-i", fileVideo, mp3file)
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
		del := exec.Command("cmd", "/C", "del", "*.mpeg")
		del.Run()
	default:
		fmt.Println("Unknown OS")
	}
}
