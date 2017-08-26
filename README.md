# Earthin24

Source code for twitter [@earthin24](https://twitter.com/earthin24) and [instagram](https://www.instagram.com/earthintwentyfour) bots. 

## Run it yourself

### Installing go
We're using Go 1.8.
```bash
brew install go@1.8
```

You'll need to set your `PATH` to point at the correct version of Go, so add this to your `.bashrc` or `.zshrc`:

```bash
export PATH="/usr/local/opt/go@1.8/bin:$PATH"
```

### Installing ffmpeg
In order to create the video from the pngs you'll need `ffmpeg`

```bash
brew install ffmpeg
```

### Compile
Run `go build frames.go`

This will create the `frames` executable on your machine.

If you now run `create_video.sh` it will generate a video for you.