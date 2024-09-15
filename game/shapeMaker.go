package game

import "math/rand/v2"

var OPTIONS [7]tetromino

type shapeMaker interface {
	Make() *tetromino
}

var _ shapeMaker = (*random)(nil)

type random struct{}

func (r *random) Make() *tetromino {
	return &OPTIONS[rand.IntN(len(OPTIONS))]
}

var _ shapeMaker = (*sevenBag)(nil)

type sevenBag struct {
	bag []tetromino
}

func (s *sevenBag) Make() *tetromino {
	if len(s.bag) == 0 {
		for _, tetromino := range OPTIONS {
			s.bag = append(s.bag, tetromino)
		}
	}

	selection := rand.IntN(len(s.bag))
	tet := s.bag[selection]
	s.bag = append(s.bag[:selection], s.bag[selection+1:]...)
	return &tet
}
