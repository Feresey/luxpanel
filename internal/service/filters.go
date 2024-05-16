package service

import (
	"fmt"
	"regexp"
)

type Aggregator[A, V any] func(agg A, val V) A

type AggCount[T any] struct {
	Value T
	Count int
}

type Filter[T, V any] func(val T) (res V, ok bool)

func NoopFilter[T any](val T) (res T, ok bool) {
	return val, true
}

func Sum(summ float32, val float32) float32 {
	return summ + val
}

func Append[T any, E ~[]T](arr E, val T) E {
	return append(arr, val)
}

// func Average

func ProcessArray[T, A, E any](
	data []T,
	filter Filter[T, E],
	aggregator Aggregator[A, E],
) (res A) {
	for _, line := range data {
		val, ok := filter(line)
		if ok {
			res = aggregator(res, val)
		}
	}
	return res
}

type ValueFilter[T any] interface {
	Filter(T) bool
	fmt.Stringer
}

type RegexpFilter struct {
	*regexp.Regexp
}

func (r *RegexpFilter) UnmarshalText(raw []byte) error {
	rw, err := MakeRegexpFilter(string(raw))
	if err != nil {
		return err
	}

	*r = rw
	return nil
}

func MakeRegexpFilter(re string) (RegexpFilter, error) {
	r, err := regexp.Compile(re)
	return RegexpFilter{Regexp: r}, err
}

func MustRegexpFilter(re string) RegexpFilter {
	r, err := MakeRegexpFilter(re)
	if err != nil {
		panic(err)
	}
	return r
}

func (r RegexpFilter) Filter(s string) bool {
	return r.Regexp.MatchString(s)
}

type StringFilter string

func (ss StringFilter) String() string { return string(ss) }

func (ss *StringFilter) UnmarshalText(raw []byte) error {
	*ss = StringFilter(raw)
	return nil
}

func (ss StringFilter) Filter(s string) bool {
	return s == string(ss)
}
