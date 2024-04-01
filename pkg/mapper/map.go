package mapper

import (
	"github.com/dolfolife/railwaygo/pkg/iterator"
)

type mapIterator[T any, V any] struct {
	source iterator.Iterator[T]
	mapper func(T) V
}

func (iter *mapIterator[T, V]) Next() bool {
	return iter.source.Next()
}

func (iter *mapIterator[T, V]) Value() V {
	value := iter.source.Value()
	return iter.mapper(value)
}

func Map[T any, V any](iter iterator.Iterator[T], f func(T) V) iterator.Iterator[V] {
	return &mapIterator[T, V]{
		iter, f,
	}
}
