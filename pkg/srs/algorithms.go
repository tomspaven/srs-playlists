package srs

import (
	"fmt"
	"time"
)

const AlgorithmLeitner string = "LEITNER"
const initialEaseFactor EaseFactor = 1.0

type nextVals struct {
	due      time.Time
	interval time.Duration
}

type algorithm interface {
	calculateNext(card Card, ease EaseFactor) (nextVals, error)
}

func getSRSInstance(algorithm string) algorithm {
	if algorithm == AlgorithmLeitner {
		return &normalisedLeitner{}
	}
	return &unknownAlgorithm{}
}

type unknownAlgorithm struct{}

func (u *unknownAlgorithm) calculateNext(card Card, ease EaseFactor) (nextVals, error) {
	return nextVals{
		due:      time.Time{},
		interval: 0,
	}, fmt.Errorf("unknownAlgorithm: Not Implemented")
}
