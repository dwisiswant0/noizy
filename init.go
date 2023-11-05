package main

import (
	"fmt"
	"time"

	"log/slog"

	"github.com/dwisiswant0/noizy/common"
	"github.com/lmittmann/tint"
	"github.com/mattn/go-colorable"
)

func init() {
	slog.SetDefault(slog.New(
		tint.NewHandler(colorable.NewColorableStderr(), &tint.Options{
			Level:      slog.LevelDebug,
			TimeFormat: time.DateTime,
		}),
	))

	fmt.Fprintf(colorable.NewColorableStderr(), banner+"\n", common.AppVersion)
}
