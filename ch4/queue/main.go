package main

import "fmt"

type Queue struct {
	List []string
}

func main() {
	var queue Queue
	fmt.Println(queue.IsEmpty())
	fmt.Println(queue.Pop())
	queue.Append("James")
	queue.Append("Confidence")
	queue.Append("Victory")
	queue.Append("Mesoma")
	queue.Append("Marvelous")
	fmt.Println(queue.Pop())
	queue.Print()

	var queue1 Queue
	fmt.Println(queue1.IsEmpty())
	fmt.Println(queue1.Pop())
	queue1.Append("Jam")
	queue1.Append("Confid")
	queue1.Append("Vic")
	queue1.Append("Meso")
	queue1.Append("Marv")
	fmt.Println(queue1.Pop())
	queue1.Print()
}

func (q *Queue) Append(value string) {
	q.List = append(q.List, value)
}

func (q *Queue) Pop() string {
	if q.IsEmpty() {
		return "nil"
	}
	poppedValue := q.List[0]
	q.List = q.List[1:]
	return poppedValue
}

func (q *Queue) Print() {
	fmt.Println(q.List)
}

func (q *Queue) IsEmpty() bool {
	return len(q.List) == 0
}
