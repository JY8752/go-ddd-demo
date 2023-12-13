package value

import (
	"reflect"
)

type Object[T any] struct {
	value T
}

func NewObject[T any](v T) Object[T] {
	return Object[T]{value: v}
}

func (v Object[T]) Value() T {
	return v.value
}

func (v Object[T]) Equals(other Object[T]) bool {
	return reflect.DeepEqual(v, other)
}
