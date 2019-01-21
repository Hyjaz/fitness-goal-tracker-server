package parse

import (
	"strconv"
	"time"
)

// ConvertUnixTimestampToTime get a timestamp as a string and return a time of type Time
func ConvertUnixTimestampToTime(unixTimestamp string) time.Time {
	i, err := strconv.ParseInt(unixTimestamp, 10, 64)
	if err != nil {
		panic(err)
	}
	return time.Unix(i, 0)
}
