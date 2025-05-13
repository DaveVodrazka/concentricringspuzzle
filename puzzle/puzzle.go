package puzzle

import (
	"fmt"
	"log"
)

type Puzzle interface {
	Show()
	Move(s int) []int
	IsVictory() bool
	Clone() Puzzle
}

type PuzzleImpl struct {
	Size  int
	Outer []int
	Inner []int
}

func NewPuzzle(n int) Puzzle {
	if n < 3 {
		log.Fatal("minimum puzzle size is 3")
	}
	p := &PuzzleImpl{
		Size:  n,
		Outer: make([]int, n),
		Inner: make([]int, n),
	}

	for i := 0; i < n; i++ {
		p.Inner[i] = i + 1
	}

	return p
}

func (p *PuzzleImpl) Show() {
	fmt.Printf("Out: %#v\nInn: %#v\n\n", p.Outer, p.Inner)
}

func (p *PuzzleImpl) rotate() {
	last := p.Inner[p.Size-1]
	for i := p.Size - 2; i >= 0; i-- {
		p.Inner[i+1] = p.Inner[i]
	}
	p.Inner[0] = last
}

func (p *PuzzleImpl) availablePositions() (pos []int) {
	for i := 0; i < p.Size; i++ {
		if p.Inner[i] != 0 && p.Outer[i] == 0 {
			pos = append(pos, i)
		}
	}
	return
}

func (p *PuzzleImpl) Move(s int) []int {
	p.Outer[s] = p.Inner[s]
	p.Inner[s] = 0
	p.rotate()

	return p.availablePositions()
}

func (p *PuzzleImpl) IsVictory() bool {
	for _, v := range p.Inner {
		if v != 0 {
			// is not victory
			return false
		}
	}
	// did not find non zero value - is victory
	return true
}

func (p *PuzzleImpl) Clone() Puzzle {
	outerClone := make([]int, len(p.Outer))
	copy(outerClone, p.Outer)
	innerClone := make([]int, len(p.Inner))
	copy(innerClone, p.Inner)
	return &PuzzleImpl{
		Size:  p.Size,
		Outer: outerClone,
		Inner: innerClone,
	}
}
