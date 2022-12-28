package helpers

type Result[T any] struct {
	Err   error
	Value T
}

func FromTuple[T any](value T, err error) *Result[T] {
	return &Result[T]{
		Err:   err,
		Value: value,
	}
}

func (t *Result[T]) IsFailure() bool {
	return !t.IsSuccess()
}

func (t *Result[T]) IsSuccess() bool {
	return t.Err == nil
}
