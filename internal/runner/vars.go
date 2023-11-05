package runner

import (
	"os"
	"syscall"

	"github.com/dwisiswant0/noizy/common"
)

var (
	appMenu = common.Item{
		Title:   common.AppName,
		Tooltip: common.AppDescription,
	}
	appInfoMenu = common.Item{
		Title:   common.AppInfo,
		Tooltip: "Application version",
	}
	quitMenu = common.Item{
		Title:   "Quit",
		Tooltip: "Quit application",
	}
	topMenu = common.Item{
		Title:   "Background Sounds",
		Tooltip: "Available background sounds",
	}
	resetMenu = common.Item{
		Title:   "Reset",
		Tooltip: "Reset sound settings",
	}
	iconData = common.Icon

	signals = []os.Signal{
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGKILL,
		syscall.SIGTERM,
	}
)
