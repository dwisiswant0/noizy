package runner

import (
	"os"
	"syscall"

	"log/slog"

	"fyne.io/systray"
	"github.com/gosimple/slug"

	"github.com/dwisiswant0/noizy/internal/pkg/noizy"
)

func (r *Runner) getBackgroundNoiseItem(parent *systray.MenuItem, s string) *noizy.BackgroundNoiseItem {
	soundFile := r.Files[slug.Make(s)]

	return &noizy.BackgroundNoiseItem{
		MenuItem:   parent.AddSubMenuItemCheckbox(s, s, false),
		File:       soundFile,
		SoundTitle: s,
	}
}

func (r *Runner) stopAllSounds() {
	for _, mBgNoiseItem := range r.Items {
		mBgNoiseItem.ToggleStop()
	}
}

func (r *Runner) doReset() {
	slog.Debug("Stopping all played sounds")
	r.stopAllSounds()
}

func (r *Runner) doQuit() {
	r.doReset()
	slog.Info("Bye!")
	systray.Quit()
}

func (r *Runner) signalHandler(sigCh chan os.Signal) {
	for {
		sig := <-sigCh

		switch sig {
		case syscall.SIGHUP:
			slog.Warn("Hang up. Resetting...")
			r.doReset()
		default:
			slog.Warn("Quitting...", "signal", sig)
			r.doQuit()
		}
	}
}
