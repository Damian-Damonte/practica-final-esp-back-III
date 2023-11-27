package timeformatter

import "time"

const (
	dateformat = "02/01/2006"
)

func dateToString(time time.Time) string {
	return time.Format(dateformat)
}
