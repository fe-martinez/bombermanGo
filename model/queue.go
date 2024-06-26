package model

type Queue[T any] struct {
	items []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

// Encolar agrega un elemento al final de la cola
func (q *Queue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
}

// Desencolar elimina y devuelve el elemento al inicio de la cola
func (q *Queue[T]) Dequeue() (T, bool) {
	if len(q.items) == 0 {
		var zero T
		return zero, false
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, true
}
