package main

type Element interface{}

type Queue struct {
	sli []Element
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Size() int {
	return len(q.sli)
}

func (q *Queue) Push(e Element) {
	q.sli = append(q.sli, e)
}

func (q *Queue) Pop() Element {
	if q.Size() == 0 {
		return nil
	}
	e := q.sli[0]
	q.sli = q.sli[1:]
	return e
}
