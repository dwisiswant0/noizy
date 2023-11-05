package noizy

import (
	"io/fs"
	"log/slog"

	"fyne.io/systray"
	"github.com/gopxl/beep"
	"github.com/gopxl/beep/speaker"
	"github.com/gopxl/beep/vorbis"
)

type BackgroundNoiseItem struct {
	*beep.Ctrl
	*systray.MenuItem
	fs.File
	SoundTitle string
}

func (b *BackgroundNoiseItem) OnClick() {
	item := b.MenuItem
	for range item.ClickedCh {
		b.Toggle()
	}
}

func (b *BackgroundNoiseItem) Toggle() {
	item := b.MenuItem
	if item.Checked() {
		b.ToggleStop()
	} else {
		b.TogglePlay()
	}
}

func (b *BackgroundNoiseItem) TogglePlay() {
	item := b.MenuItem
	if item.Checked() {
		return
	}

	slog.Debug("Toggle play for sound", "sound", b.SoundTitle)
	b.Play()
	item.Check()
}

func (b *BackgroundNoiseItem) ToggleStop() {
	item := b.MenuItem
	if !item.Checked() {
		return
	}

	slog.Debug("Toggle stop for sound", "sound", b.SoundTitle)
	b.Stop()
	item.Uncheck()
}

func (b *BackgroundNoiseItem) Play() {
	streamer, _, err := vorbis.Decode(b.File)
	if err != nil {
		panic(err)
	}

	ctrl := &beep.Ctrl{
		Streamer: beep.Loop(-1, streamer),
		Paused:   false,
	}

	b.Ctrl = ctrl
	speaker.Play(ctrl)
}

func (b *BackgroundNoiseItem) Stop() {
	if b.Ctrl != nil {
		speaker.Lock()
		b.Ctrl.Paused = true
		speaker.Unlock()
	}
}
