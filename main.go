package main

import "fmt"

type stack struct {
	stackA []int
	stackB []int
}

func NewStack() *stack {
	return &stack{
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
	s.stackA = (s.stackA)[:len(s.stackA)-1]
}

func (s *stack) PopB() {
	s.stackB = (s.stackB)[:len(s.stackB)-1]
}

func (s *stack) Pa() {
	s.PushA(s.stackB[len(s.stackB)-1])
	s.PopB()
}

func (s *stack) Ra() {
	val := s.stackA[len(s.stackA)-1]
	s.PopA()
	s.stackA = append([]int{val}, s.stackA...)
}

func (s *stack) Rb() {
	val := s.stackB[len(s.stackB)-1]
	s.PopB()
	s.stackB = append([]int{val}, s.stackB...)
}

func (s *stack) Rr() {
	s.Ra()
	s.Rb()
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

func (s *stack) Sb() {
	s.stackB[len(s.stackB)-1], s.stackB[len(s.stackB)-2] = s.stackB[len(s.stackB)-2], s.stackB[len(s.stackB)-1]
}

func (s *stack) Sa() {
	s.stackA[len(s.stackA)-1], s.stackA[len(s.stackA)-2] = s.stackA[len(s.stackA)-2], s.stackA[len(s.stackA)-1]
}

func (s *stack) Ss() {
	s.Sa()
	s.Sb()
}


func main() {

	s := NewStack()
	s.PushA(1)
	s.PushA(2)
	s.PushA(3)
	s.PushB(4)
	s.PushB(5)
	s.PushB(6)

	fmt.Println("stack A :", s.stackA)
	fmt.Println("stack B :", s.stackB)

	s.Ss()

	fmt.Println("stack A :", s.stackA)
	fmt.Println("stack B :", s.stackB)

}
