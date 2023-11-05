package sound

import "embed"

func ToMapFiles(fs embed.FS) (MapFiles, error) {
	var mf = make(MapFiles)

	entries, err := fs.ReadDir("sounds")
	if err != nil {
		return mf, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		f, err := fs.Open("sounds/" + entry.Name())
		if err != nil {
			return mf, err
		}
		defer f.Close()

		fname := getFilename(entry.Name())
		mf[fname] = f
	}

	return mf, nil
}
