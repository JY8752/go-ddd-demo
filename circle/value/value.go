package value

import (
	"ddd-demo/value"
	"errors"
)

// サークルの識別子　値オブジェクト
type CircleId struct {
	value.Object[string]
}

func NewCircleId(v string) *CircleId {
	return &CircleId{value.NewObject[string](v)}
}

// サークル名 値オブジェクト
type CircleName struct {
	value.Object[string]
}

func NewCircleName(v string) (*CircleName, error) {
	nameLength := len([]rune(v))

	if nameLength < 3 {
		return nil, errors.New("circle name must be at least 3 charcters")
	}

	if nameLength > 20 {
		return nil, errors.New("circle name must be less than 20 charcters")
	}

	return &CircleName{value.NewObject[string](v)}, nil
}
