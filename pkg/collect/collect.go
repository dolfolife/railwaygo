package collect

import (
	"github.com/dolfolife/railwaygo/pkg/iterator"
	"github.com/dolfolife/railwaygo/pkg/result"
)

func Collect[T any](iter iterator.Iterator[T]) []T {
	var xs []T

	for iter.Next() {
		xs = append(xs, iter.Value())
	}

	return xs
}

func Fold[T any](r result.Result[T]) T {
	return r.Val
}
