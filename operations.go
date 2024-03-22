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
