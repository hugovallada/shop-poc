package vo

type Department struct {
	Value string
}

func NewDepartment(value string) Department {
	return Department{Value: value}
}
