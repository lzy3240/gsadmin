package f

import "time"

func DateFormat(t time.Time) string {
	if !t.IsZero() {
		return t.Format("2006-01-02 15:04:05")
	}
	return ""
}
