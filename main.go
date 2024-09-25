package main

import "fmt"



type stack struct {
	stackA []int
	stackB []int
}

func NewStack() *stack {
	return &stack {
		stackA: []int{},
		stackB: []int{},
	}
}

func (s *stack) PushA(v int) {
    s.stackA = append(s.stackA, v)
}

func (s *stack) PushB(v int) {
    s.stackB = append(s.stackB, v)
}

func (s *stack) PopA() {
    s.stackA=(s.stackA)[:len(s.stackA)-1]
}

func (s *stack) PopB()  {
    s.stackB=(s.stackB)[:len(s.stackB)-1]
}

func (s *stack) Pa() {
	s.PushA(s.stackB[len(s.stackB)-1])
	s.PopB()
}

func (s *stack) Rra() {
		val := s.stackA[0]
		s.stackA = s.stackA[1:]
		s.PushA(val)
}

func (s *stack) Rrb() {
	val := s.stackB[0]
	s.stackB = s.stackB[1:]
	s.PushB(val)
}

func (s *stack) Rrr() {
	s.Rra()
	s.Rrb()
}

func (s *stack) Pb() {
	s.PushB(s.stackA[len(s.stackA)-1])
	s.PopA()
}

func main() {
	s := NewStack()
	
	s.PushA(1)
	s.PushA(2)
	s.PushA(3)
	s.PushB(4)
	s.PushB(5)
	s.PushB(6)

	fmt.Println("stack A :",s.stackA)
	fmt.Println("stack B :",s.stackB)

	s.Rrr()

	fmt.Println("stack A :",s.stackA)
	fmt.Println("stack B :",s.stackB)

}