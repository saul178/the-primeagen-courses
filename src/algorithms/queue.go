// NOTE: you can also use slices and linked lists to implement a Queue

package algorithms

import (
	"errors"
	"fmt"
)

type Queue struct {
	Elements []int
	Size     int
}

func (q *Queue) Enqueue(elem int) {
	if q.GetLength() == q.Size {
		fmt.Println("out of bounds")
		return
	}
	q.Elements = append(q.Elements, elem)
}

func (q *Queue) Dequeue() int {
	if q.IsEmpty() {
		fmt.Print("queue is empty")
		return 0
	}
	element := q.Elements[0]
	if q.GetLength() == 1 {
		q.Elements = nil
		return element
	}
	q.Elements = q.Elements[1:]
	return element
}

func (q *Queue) GetLength() int {
	return len(q.Elements)
}

func (q *Queue) IsEmpty() bool {
	return len(q.Elements) == 0
}

func (q *Queue) Peek() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("empty queue")
	}
	return q.Elements[0], nil
}
