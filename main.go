package main

import (
	"log/slog"

	"fyne.io/systray"
	"github.com/gopxl/beep/speaker"
	"github.com/gopxl/beep/vorbis"

	"github.com/dwisiswant0/noizy/internal/runner"
)

func main() {
	slog.Debug("Initialize runner")
	r, err := runner.New(sounds)
	if err != nil {
		panic(err)
	}

	slog.Debug("Decode single sound file to get sample rate")
	streamer, format, err := vorbis.Decode(r.Files["brook"])
	if err != nil {
		panic(err)
	}
	defer streamer.Close()

	slog.Debug("Initialize speaker")
	err = speaker.Init(format.SampleRate, bufSize)
	if err != nil {
		panic(err)
	}

	systray.Run(r.Ready, nil)
}
