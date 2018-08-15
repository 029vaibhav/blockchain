package time

import (
	"strconv"
	"time"
)

func Now() time.Time {
	return time.Now()
}

func Millis() int64 {
	return time.Now().UnixNano() / 1000000
}

func MillisInString() string {

	return strconv.FormatInt(Millis(), 10)

}
