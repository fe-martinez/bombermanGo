package model

type Queue struct {
	items []string
}

func NewQueue() *Queue {
	return &Queue{}
}

// Encolar agrega un elemento al final de la cola
func (q *Queue) Enqueue(item string) {
	q.items = append(q.items, item)
}

// Desencolar elimina y devuelve el elemento al inicio de la cola
func (q *Queue) Dequeue() (string, bool) {
	if len(q.items) == 0 {
		return "", false
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, true
}
