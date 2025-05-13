package solver

import (
	"github.com/DaveVodrazka/concentricringspuzzle/puzzle"
)

type Solver interface {
	Solve(p puzzle.Puzzle) [][]int
}

type SolverImpl struct {
	solutions [][]int
}

func NewSolver() Solver {
	return &SolverImpl{}
}

func (s *SolverImpl) findWinningSolutions(puzzle puzzle.Puzzle, availableMoves []int, history []int) {
	if len(availableMoves) == 0 {
		if puzzle.IsVictory() {
			s.solutions = append(s.solutions, append([]int{}, history...))
		}
		return
	}

	for _, move := range availableMoves {
		nextPuzzle := puzzle.Clone()
		nextAvailableMoves := nextPuzzle.Move(move)
		nextHistory := append(append([]int{}, history...), move)
		s.findWinningSolutions(nextPuzzle, nextAvailableMoves, nextHistory)
	}
}

func (s *SolverImpl) Solve(p puzzle.Puzzle) [][]int {
	s.solutions = [][]int{}
	initialAvailableMoves := p.Move(0)
	s.findWinningSolutions(p, initialAvailableMoves, []int{0})
	return s.solutions
}
