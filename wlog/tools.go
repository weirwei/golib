package wlog

import (
	"strconv"
	"time"
)

func getRequestID() string {
	usec := uint64(time.Now().UnixNano())
	requestID := strconv.FormatUint(usec&0x7FFFFFFF|0x80000000, 10)
	return requestID
}
