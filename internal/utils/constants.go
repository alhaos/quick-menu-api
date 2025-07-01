package utils

import "time"

const timeout = 5000000

func Timeout() time.Duration {
	return time.Millisecond * timeout
}
