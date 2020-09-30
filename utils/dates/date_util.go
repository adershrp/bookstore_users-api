package dates

import "time"

// GetNow returns current UTC time
func GetNow() time.Time {
	return time.Now().UTC()
}

// GetNowString returns current UTC time in String
func GetNowString() string {
	return GetNow().Format(time.RFC3339)
}
