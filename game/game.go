package game

import "fmt"

const MATRIX_HEIGHT uint8 = 20
const MATRIX_WIDTH uint8 = 10

type activeTetromino struct {
	tetromino tetromino
	x         uint8
	y         uint8
}

type state struct {
	matrix     [][]bool // indexable as matrix[x][y], where (0, 0) is the top left corner.
	shapeMaker shapeMaker
	active     *activeTetromino
	next       *tetromino
}

type coordSet []coord

type coord struct {
	x uint8
	y uint8
}

func NewGame() (*state, error) {
	err := load()
	if err != nil {
		return nil, fmt.Errorf("error loading game: %w", err)
	}

	matrix := make([][]bool, MATRIX_WIDTH)
	for i := range matrix {
		matrix[i] = make([]bool, MATRIX_HEIGHT)
	}

	return &state{
		matrix:     matrix,
		shapeMaker: &sevenBag{}, // todo: parameterise me at the user level
	}, nil
}

func (s *state) Step() {
	if s.next == nil {
		s.next = s.shapeMaker.Make()
	}

	if s.active == nil {
		s.active = &activeTetromino{
			tetromino: *s.next,
			x:         3,
			y:         0,
		}
		s.next = nil
	}

	// can the active tetronimo move down?
	// all cells below tetronimo coords must be false
	// -> AND(cells) == false
	cSet := s.getActiveCoordSet()
	blocked := false
	for _, c := range cSet {
		if c.y+1 >= MATRIX_HEIGHT {
			blocked = true
			break
		}

		if s.matrix[c.x][c.y+1] {
			blocked = true
			break
		}
	}

	if blocked {
		for _, c := range cSet {
			s.matrix[c.x][c.y] = true
		}
		s.active = nil
	} else {
		s.active.y++
	}
}

// getActiveCoordSet returns a list of coordinates of each cell in the active tetronimo relative to the matrix
func (s state) getActiveCoordSet() coordSet {
	cSet := make(coordSet, 0, 4)
	for y := 0; y < len(s.active.tetromino); y++ {
		for x := 0; x < len(s.active.tetromino[0]); x++ {
			if s.active.tetromino[x][y] {
				cSet = append(cSet, coord{x: uint8(x) + s.active.x, y: uint8(y) + s.active.y})
			}
		}
	}

	return cSet
}

func (s state) Draw() string {
	matrixCopy := make([][]bool, len(s.matrix))

	for i := range s.matrix {
		matrixCopy[i] = make([]bool, len(s.matrix[i]))
		copy(matrixCopy[i], s.matrix[i])
	}

	if s.active != nil {
		cSet := s.getActiveCoordSet()
		for _, c := range cSet {
			matrixCopy[c.x][c.y] = true
		}
	}

	output := ""
	for y := 0; y < int(MATRIX_HEIGHT); y++ {
		output = fmt.Sprintf("%v|", output)
		for x := 0; x < int(MATRIX_WIDTH); x++ {
			if matrixCopy[x][y] {
				output = fmt.Sprintf("%v#", output)
			} else {
				output = fmt.Sprintf("%v ", output)
			}
		}
		output = fmt.Sprintf("%v|\n", output)
	}
	output = fmt.Sprintf("%v ", output)
	for range MATRIX_WIDTH {
		output = fmt.Sprintf("%v-", output)
	}
	return output
}
