package result

import (
	"github.com/dolfolife/railwaygo/pkg/iterator"
)

type Result[V any] struct {
	Val   V
	Error error
}

type SliceResult[T any] struct {
	Results []Result[T]
	value   Result[T]
	index   int
}

func (s *SliceResult[T]) Next() bool {

	if s.index < len(s.Results) {
		s.value = s.Results[s.index]
		s.index += 1
		return true
	}

	return false
}

func (s *SliceResult[T]) Value() Result[T] {
	return s.value
}

func NewSliceResult[T any](elems []T) iterator.Iterator[Result[T]] {
	result := SliceResult[T]{}
	for _, e := range elems {
		result.Results = append(result.Results, Result[T]{Val: e, Error: nil})
	}
	return &result
}
