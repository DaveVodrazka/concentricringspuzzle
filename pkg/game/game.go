package game

import (
	"fmt"
	"strconv"

	"github.com/DaveVodrazka/concentricringspuzzle/pkg/puzzle"
)

func contains(slice []int, target int) bool {
	for _, num := range slice {
		if num == target {
			return true
		}
	}
	return false
}

func getMove(msg string, options []int) (res int) {
	fmt.Println(msg)
	var input string
	var err error

	for err != nil || !contains(options, res) {
		_, err = fmt.Scanln(&input)
		if err != nil {
			continue
		}
		res, err = strconv.Atoi(input)
	}
	return
}

func getSize(msg string, min, max int) (res int) {
	fmt.Println(msg)
	var input string
	var err error

	for err != nil || min > res || max < res {
		_, err = fmt.Scanln(&input)
		if err != nil {
			continue
		}
		res, err = strconv.Atoi(input)
	}
	return
}

func StartGame() {
	s := getSize("Choose puzzle size (must be 3 < n < 20):", 3, 20)
	p := puzzle.NewPuzzle(s)

	moves := p.Move(0)

	for len(moves) > 0 {
		p.Show()
		m := getMove(fmt.Sprintf("Choose your move: %v", moves), moves)
		moves = p.Move(m)
	}

	if p.IsVictory() {
		fmt.Println("You won!")
	} else {
		fmt.Println("You lost!")
	}
}
