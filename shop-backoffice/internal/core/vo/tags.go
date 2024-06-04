package vo

type Tags struct {
	Value []string
}

func NewTags(value []string) Tags {
	return Tags{Value: value}
}
