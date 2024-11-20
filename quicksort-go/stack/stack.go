package stack

type Stack[T any] []T

func (s *Stack[T]) Push(item T) {
	*s = append(*s, item)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(*s) == 0 {
		var zero T
		return zero, false
	}
	index := len(*s) - 1
	item := (*s)[index]
	*s = (*s)[:index]
	return item, true
}