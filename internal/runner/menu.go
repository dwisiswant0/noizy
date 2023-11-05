package runner

import (
	"log/slog"

	"fyne.io/systray"
)

func addAppSystray() {
	systray.SetTemplateIcon(iconData, iconData)
	systray.SetTitle(appMenu.Title)
	systray.SetTooltip(appMenu.Tooltip)
}

func addAppInfoMenu() {
	systray.AddSeparator()
	mAppInfo := systray.AddMenuItem(appInfoMenu.Title, appInfoMenu.Tooltip)
	mAppInfo.SetTemplateIcon(iconData, iconData)
	mAppInfo.Disable()
}

func (r *Runner) addResetMenu() {
	systray.AddSeparator()
	mReset := systray.AddMenuItem(resetMenu.Title, resetMenu.Tooltip)
	mReset.Enable()

	go func() {
		for range mReset.ClickedCh {
			slog.Info("Reset menu item clicked")
			r.doReset()
		}
	}()
}

func (r *Runner) addQuitMenu() {
	mQuit := systray.AddMenuItem(quitMenu.Title, quitMenu.Tooltip)
	mQuit.Enable()

	go func() {
		<-mQuit.ClickedCh
		slog.Warn("Quit menu item clicked")
		r.doQuit()
	}()
}
