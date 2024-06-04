package vo

import "time"

type Timestamp struct {
	Value uint64
}

func NewTimestamp() Timestamp {
	return Timestamp{Value: uint64(time.Now().Unix())}
}

func NewTimestampOf(timestampValue uint64) Timestamp {
	var value uint64
	if timestampValue == 0 {
		value = NewTimestamp().Value
	} else {
		value = timestampValue
	}
	return Timestamp{Value: value}
}
