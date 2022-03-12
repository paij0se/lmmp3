**lmmp3**

lmmp3 is a function that download a video from youtube, and convert it to a mp3 file using ffmpeg

**You need to have installed ffmpeg in your system!**

download lmmp3 with ```go get github.com/paij0se/lmmp3 ```

example (windows)

https://user-images.githubusercontent.com/69026987/158023658-82a7ee33-5b83-4259-ab35-1b1dc70f5e69.mp4

```go
package main

import (
	"os/exec"

	lmmp3 "github.com/paij0se/lmmp3"
)

func main() {
	lmmp3.DownloadFFmpeg() //  download ffmpeg if is not installed (only in windows)
	// output: 'Vine Boom Sound Effect.mp3'
	lmmp3.DownloadAndConvert("https://www.youtube.com/watch?v=829pvBHyG6I")
	// delete the original mpeg file (only for windows)

	del := exec.Command("cmd", "/C", "del", "*.mpeg")
	if del.Run() != nil {
		panic("failed to delete files")
	}

}

```

example linux and mac

```go
package main

import (
	lmmp3 "github.com/paij0se/lmmp3"
)

func main() {
	// download the video and convert it to a mp3 file
	// output: 'Vine Boom Sound Effect.mp3'
	lmmp3.DownloadAndConvert("https://www.youtube.com/watch?v=829pvBHyG6I")
}
```
