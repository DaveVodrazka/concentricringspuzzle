package main

import (
	"fmt"

	"github.com/DaveVodrazka/concentricringspuzzle/puzzle"
	"github.com/DaveVodrazka/concentricringspuzzle/solver"
)

func main() {
	s := solver.NewSolver()

	for n := 3; n < 16; n++ {
		p := puzzle.NewPuzzle(n)
		solutions := s.Solve(p)
		solutionsCount := len(solutions)
		if solutionsCount > 0 {
			fmt.Printf("Size %d had %d solutions\n\n", n, solutionsCount)
		}
	}
}
