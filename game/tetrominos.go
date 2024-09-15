package game

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	B, T, L, J, I, S, Z tetromino
)

type tetromino [4][4]bool

func (t tetromino) Pretty() string {
	s := ""
	for x := 0; x < 4; x++ {
		for y := 0; y < 4; y++ {
			if t[x][y] {
				s = fmt.Sprintf("%v#", s)
			} else {
				s = fmt.Sprintf("%v ", s)
			}
		}
		s = fmt.Sprintf("%v\n", s)
	}

	return s
}

func load() error {
	f, err := os.Open("game/tetrominos.txt")
	if err != nil {
		return fmt.Errorf("error reading tetronimos data file: %w", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var lines []string
	var i uint8
	for scanner.Scan() {
		text := scanner.Text()
		if strings.HasSuffix(text, ":") {
			if len(lines) != 0 {
				loadShape(lines)
			}
			lines = make([]string, 6)
			i = 0
		}
		lines[i] = text
		i++
	}
	loadShape(lines)
	OPTIONS = [7]tetromino{B, I, S, Z, J, L, T}
	return nil
}

func loadShape(shape []string) {
	cells := [4][4]bool{}
	for y := 0; y < 4; y++ {
		slice := shape[y+1]
		for x := 0; x < len(slice); x++ {
			char := slice[x : x+1]

			if char == "\n" {
				break
			}

			switch char {
			case " ":
				cells[x][y] = false
			case "#":
				cells[x][y] = true
			default:
				panic(fmt.Sprintf("unrecognised character in tetronimos data file: %v", char))
			}
		}
	}

	thisShape := cells

	shapeIdentifier := shape[0][:1]
	switch shapeIdentifier {
	case "b":
		B = thisShape
	case "t":
		T = thisShape
	case "l":
		L = thisShape
	case "j":
		J = thisShape
	case "i":
		I = thisShape
	case "s":
		S = thisShape
	case "z":
		Z = thisShape
	default:
		panic(fmt.Sprintf("unrecognised shape identifier in tetronimos data file: %v", shapeIdentifier))
	}
}
