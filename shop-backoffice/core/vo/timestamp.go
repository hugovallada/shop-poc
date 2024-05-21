package vo

import "time"

type Timestamp struct {
	Value uint64
}

func NewTimestamp() Timestamp {
	return Timestamp{Value: uint64(time.Now().Unix())}
}
