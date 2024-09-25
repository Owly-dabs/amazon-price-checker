package util

func Truncate(s string, maxWidth int) string {
	if len(s) > maxWidth {
		return s[:maxWidth-3] + "..."
	}
	return s
}
