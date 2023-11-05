package runner

import (
	"os"

	"log/slog"
	"os/signal"

	"fyne.io/systray"

	"github.com/dwisiswant0/noizy/internal/pkg/noizy"
)

func (r *Runner) Ready() {
	slog.Debug("Initialize application system tray")
	addAppSystray()

	mBgNoise := systray.AddMenuItem(topMenu.Title, topMenu.Tooltip)
	mBgNoiseItems := []*noizy.BackgroundNoiseItem{}

	slog.Debug("Getting noise list")
	for _, list := range noizy.GetNoiseList() {
		slog.Info("Add sound submenu", "category", list.Category)
		mSubBgNoise := mBgNoise.AddSubMenuItem(list.Category, list.Category)
		for _, s := range list.Sounds {
			slog.Info("Add sound to submenu", "category", list.Category, "sound", s)
			mBgNoiseItem := r.getBackgroundNoiseItem(mSubBgNoise, s)
			mBgNoiseItems = append(mBgNoiseItems, mBgNoiseItem)
		}
	}

	r.Items = mBgNoiseItems
	slog.Info("Background sounds menu item ready", "count", len(r.Items))

	for _, mBgNoiseItem := range mBgNoiseItems {
		slog.Debug("Initialize on-click trigger for sound menu item", "sound", mBgNoiseItem.SoundTitle)
		go mBgNoiseItem.OnClick()
	}

	slog.Debug("Add application information menu item")
	addAppInfoMenu()

	slog.Debug("Add application reset menu item")
	r.addResetMenu()

	slog.Debug("Add application quit menu item")
	r.addQuitMenu()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, signals...)
	go r.signalHandler(sig)

	slog.Info("All set!", "pid", os.Getpid())
}
