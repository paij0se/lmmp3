**lmmp3**

lmmp3 is a function that download a video from youtube, and convert it to a mp3 file using ffmpeg

**You need to have installed ffmpeg in your system!**

download lmmp3 with ```go get github.com/paij0se/lmmp3 ```

example

```go
package main

import(lmmp3 "github.com/paij0se/lmmp3")

func main(){
    // download the video and convert it to a mp3 file
    // output: 'Vine Boom Sound Effect.mp3'
    lmmp3.DownloadAndConvert("https://www.youtube.com/watch?v=829pvBHyG6I")
}
```
