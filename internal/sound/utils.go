package sound

func getFilename(s string) string {
	return s[:len(s)-len(Ext)]
}
