package model

import "errors"

type Status int

const (
	Draft     Status = iota + 1 // EnumIndex = 1
	Published                   // EnumIndex = 2
	Closed                      // EnumIndex = 3
)

func (s Status) String() string {
	return [...]string{"Draft", "Published", "Closed"}[s-1]
}

func (s Status) EnumIndex() int {
	return int(s)
}

// Get status from string
func GetStatus(statusStr string) (Status, error) {
	switch statusStr {
	case "draft":
		return Draft, nil
	case "published":
		return Published, nil
	case "closed":
		return Closed, nil
	default:
		return 0, errors.New("Invalid status")
	}
}
