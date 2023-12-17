package domain

import (
	"reflect"
)

type ValueObject[T any] struct {
	value T
}

func NewValueObject[T any](v T) ValueObject[T] {
	return ValueObject[T]{value: v}
}

func (v ValueObject[T]) Value() T {
	return v.value
}

func (v ValueObject[T]) Equals(other ValueObject[T]) bool {
	return reflect.DeepEqual(v, other)
}
