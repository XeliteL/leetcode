package main

import "fmt"

type MyQueue struct {
	in  []int
	out []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	this.in = append(this.in, x)
}

func (this *MyQueue) Pop() int {
	num := this.Peek()
	this.out = this.out[:len(this.out)-1]
	return num
}

func (this *MyQueue) Peek() int {
	this.move()
	return this.out[len(this.out)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.in) == 0 && len(this.out) == 0
}

func (this *MyQueue) move() {
	if len(this.out) == 0 {
		for len(this.in) > 0 {
			top := this.in[len(this.in)-1]
			this.in = this.in[:len(this.in)-1]
			this.out = append(this.out, top)
		}
	}
}

func main() {
	q := Constructor()
	q.Push(1)
	q.Push(2)
	fmt.Println(q.Peek())  // 1
	fmt.Println(q.Pop())   // 1
	fmt.Println(q.Empty()) // false

	q2 := Constructor()
	q2.Push(1)
	q2.Push(2)
	q2.Push(3)
	fmt.Println(q2.Pop()) // 1
	q2.Push(4)
	fmt.Println(q2.Pop())   // 2
	fmt.Println(q2.Peek())  // 3
	fmt.Println(q2.Pop())   // 3
	fmt.Println(q2.Pop())   // 4
	fmt.Println(q2.Empty()) // true
}
