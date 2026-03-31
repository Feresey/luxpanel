package service

import (
	"iter"
)

type Aggregator[T any] interface {
	Aggregate(it iter.Seq[T])
}

type AggCount[T any] struct {
	Value T
	Count int
}

type Filter[T, V any] interface {
	Filter(val T) (res V, ok bool)
}

func NoopFilter[T any](val T) (res T, ok bool) {
	return val, true
}

func Sum(summ float32, val float32) float32 {
	return summ + val
}

func Append[T any, E ~[]T](arr E, val T) E {
	return append(arr, val)
}

func ApplyFilter[T, E any](
	data []T,
	filter Filter[T, E],
) iter.Seq[E] {
	return func(yield func(E) bool) {
		for _, line := range data {
			val, ok := filter.Filter(line)
			if ok {
				if !yield(val) {
					return
				}
			}
		}
	}
}
