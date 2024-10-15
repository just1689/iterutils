package iterutils

import (
	"iter"
)

type Iterator[V any] struct {
	iter iter.Seq[V]
}

func (i Iterator[V]) Collect() []V {
	collect := make([]V, 0)
	for e := range i.iter {
		collect = append(collect, e)
	}
	return collect
}

func From[V any](slice []V) *Iterator[V] {
	return &Iterator[V]{
		iter: func(yield func(V) bool) {
			for _, v := range slice {
				if !yield(v) {
					return
				}
			}
		},
	}
}

func (i *Iterator[V]) Each(f func(V)) {
	for i := range i.iter {
		f(i)
	}
}

func (i *Iterator[V]) Reverse() *Iterator[V] {
	collect := i.Collect()
	counter := len(collect) - 1
	for e := range i.iter {
		collect[counter] = e
		counter--
	}
	return From(collect)
}

func (i *Iterator[V]) Map(f func(V) V) *Iterator[V] {
	cpy := i.iter
	i.iter = func(yield func(V) bool) {
		for v := range cpy {
			v = f(v)
			if !yield(v) {
				return
			}
		}
	}
	return i
}

func (i *Iterator[V]) Filter(f func(V) bool) *Iterator[V] {
	cpy := i.iter
	i.iter = func(yield func(V) bool) {
		for v := range cpy {
			if f(v) {
				if !yield(v) {
					return
				}
			}
		}
	}
	return i
}
