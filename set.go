package set

import (
	"fmt"
	"strings"
)

type MapFunc[T comparable] func(value T) T
type FilterFunc[T comparable] func(value T) bool
type OperationFunc[T comparable] func (value T)

type Set[T comparable] interface {
	Insert(T)
	Remove(T)
	Has(T) bool
	Len() int
	IsSubset(Set[T]) bool
	IsProperSubset(Set[T]) bool
	IsSuperset(Set[T]) bool
	IsProperSuperset(Set[T]) bool
	Map(MapFunc[T]) Set[T]
	Filter(FilterFunc[T]) Set[T]
	ForEach(OperationFunc[T])
}

type empty struct{}
type mapset[T comparable] map[T]empty

func New[T comparable]() Set[T] {
	return mapset[T](make(map[T]empty))
}

func (s mapset[T]) Insert(value T) {
	s[value] = empty{}
}

func (s mapset[T]) Remove(value T) {
	delete(s, value)
}

func (s mapset[T]) Has(value T) bool {
	_, ok := s[value]
	return ok
}

func (s mapset[T]) Len() int {
	return len(s)
}

func (s mapset[T]) IsSubset(other Set[T]) bool {
	return Subtraction[T](s, other).Len() == 0
}

func (s mapset[T]) IsProperSubset(other Set[T]) bool {
	return s.IsSubset(other) && Subtraction[T](other, s).Len() != 0
}

func (s mapset[T]) IsSuperset(other Set[T]) bool {
	return other.IsSubset(s)
}

func (s mapset[T]) IsProperSuperset(other Set[T]) bool {
	return other.IsProperSuperset(s)
}

func (s mapset[T]) Map(op MapFunc[T]) Set[T] {
	r := New[T]()
	for k := range s {
		r.Insert(op(k))
	}
	return r
}

func (s mapset[T]) Filter(op FilterFunc[T]) Set[T] {
	r := New[T]()
	for k := range s {
		if op(k) {
			r.Insert(k)
		}
	}
	return r
}

func (s mapset[T]) ForEach(op OperationFunc[T]) {
	for k := range s {
		op(k)
	}
}

func (s mapset[T]) String() string {
	els := make([]string, 0, s.Len())
	s.ForEach(func(value T) {
		els = append(els, fmt.Sprintf("%s", value))
	})	
	return fmt.Sprintf("{ %s }", strings.Join(els, ", "))
}
