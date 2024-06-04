package vo

type Money struct {
	Value uint64
}

func NewMoney(newValue uint64) Money {
	return Money{Value: newValue}
}
