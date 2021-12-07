package stackandqueues

//STACK is last in first out (push is to top, pop is from top)

//QUEUE first un first out	(Push is to the end, pop is from top)

type Stack struct {
	items []int
}

//PUSH
func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

//POP
func (s *Stack) Pop() int {
	stackLen := len(s.items) - 1
	removedItem := s.items[stackLen]
	s.items = s.items[:stackLen]
	return removedItem
}

type Queue struct {
	items []int
}

//PUSH
func (q *Queue) Push(i int) {
	q.items = append(q.items, i)
}

//POP

func (q *Queue) Pop() int {
	removedElem := q.items[0]
	q.items = q.items[1:]
	return removedElem
}
