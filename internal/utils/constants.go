package utils

import "time"

const timeout = 5000

func Timeout() time.Duration {
	return time.Millisecond * timeout
}
