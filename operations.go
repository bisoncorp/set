package set

func Union[T comparable](a, b Set[T]) Set[T] {
	r := New[T]()
	a.ForEach(func(value T) {
		r.Insert(value)
	})
	b.ForEach(func(value T) {
		r.Insert(value)
	})
	return r
}

func Intersection[T comparable](a, b Set[T]) Set[T] {
	return a.Filter(func(value T) bool {
		return b.Has(value)
	})
}

func Subtraction[T comparable](a, b Set[T]) Set[T] {
	return a.Filter(func(value T) bool {
		return !b.Has(value)
	})
}

func ExclusiveOr[T comparable](a, b Set[T]) Set[T] {
	return Subtraction(Union(a, b), Intersection(a, b))
}

type ClosureFunction[T comparable, S any] func(set.Set[T], S) set.Set[T]

func Closure[T comparable, S any](fn ClosureFunction[T,S]) ClosureFunction[T,S] {
	return func (a set.Set[T], arg S) set.Set[T] {
		for {
			b := fn(a.Clone(), arg)
			if b.IsSubset(a) && b.IsSuperset(a) {
				return b
			}
			a = b
		}
	}
}
