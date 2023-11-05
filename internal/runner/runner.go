package runner

import (
	"embed"

	"github.com/dwisiswant0/noizy/internal/pkg/noizy"
	"github.com/dwisiswant0/noizy/internal/sound"
)

type Runner struct {
	Files sound.MapFiles
	Items []*noizy.BackgroundNoiseItem
}

func New(fs embed.FS) (*Runner, error) {
	mf, err := sound.ToMapFiles(fs)
	if err != nil {
		return nil, err
	}

	r := &Runner{Files: mf}

	return r, nil
}
