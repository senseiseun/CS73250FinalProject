package Iterator

type Iterable[T any] interface {
	Start()
	Next()
	Done() bool
	GetIterVal() T
}

type Iterator[T any] struct {
	current   any
	reference Iterable[T]
	started   bool
}

func (it *Iterator[T]) Start() {
	it.reference.Start()
}

func (it *Iterator[T]) Next() {
	it.reference.Next()
}

func (it *Iterator[T]) Get() T {
	return it.reference.GetIterVal()
}

func (it *Iterator[T]) SetCurrent(curr any) {
	it.current = curr
}

func (it *Iterator[T]) GetCurrent() any {
	return it.current
}

func (it *Iterator[T]) SetReference(ref Iterable[T]) {
	it.reference = ref
}

func (it *Iterator[T]) Done() bool {
	return it.reference.Done()
}
