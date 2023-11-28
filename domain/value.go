package domain

// type primitive interface {
// 	~int | ~int8 | ~int16 | ~int32 | ~int64 |
// 		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
// 		~float32 | ~float64 |
// 		~bool |
// 		~string
// }

// type ValueObject[T any] struct {
// 	value T
// }

// func NewValueObject[T any](v T) ValueObject[T] {
// 	return ValueObject[T]{value: v}
// }

// func (v ValueObject[T]) Value() T {
// 	return v.value
// }

// func (v ValueObject[T]) String() string {
// 	return fmt.Sprintf("%v", v.value)
// }

// func (v ValueObject[T]) MarshalJSON() ([]byte, error) {
// 	return json.Marshal(v.value)
// }

// type ValueObjFunc struct{}

// func (v ValueObjFunc) Equals(obj, other any) bool {
// 	return reflect.DeepEqual(obj, other)
// }
