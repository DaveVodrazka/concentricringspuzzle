package main

import (
	"fmt"

	"github.com/DaveVodrazka/concentricringspuzzle/pkg/puzzle"
	"github.com/DaveVodrazka/concentricringspuzzle/pkg/solver"
)

func main() {
	s := solver.NewSolver()

	for n := 3; n < 16; n++ {
		p := puzzle.NewPuzzle(n)
		solutions := s.Solve(p)
		solutionsCount := len(solutions)
		if solutionsCount > 0 {
			fmt.Printf("Size %d has %d solutions\nSample solution: %#v\n\n", n, solutionsCount, solutions[0])
		}
	}
}
