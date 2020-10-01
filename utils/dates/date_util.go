package dates

import "time"

const (
	apiDateFormat   = "2006-01-02T15:04:05Z"
	apiDateDBFormat = "2006-01-02 15:04:05"
)

// GetNow returns current UTC time
func GetNow() time.Time {
	return time.Now().UTC()
}

// GetNowString returns current UTC time in String
func GetNowString() string {
	return GetNow().Format(apiDateFormat)
}

// GetNowString returns current UTC time in String
func GetNowDBFormat() string {
	return GetNow().Format(apiDateDBFormat)
}
