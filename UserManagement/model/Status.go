package model

type Status int

const (
	Draft Status = iota + 1 // EnumIndex = 1
	Published               // EnumIndex = 2
	Closed                  // EnumIndex = 3
)

func (s Status) String() string {
	return [...]string{"Draft", "Published", "Closed"}[s-1]
}

func (s Status) EnumIndex() int {
	return int(s)
}

