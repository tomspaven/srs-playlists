package srs

import (
	"fmt"
	"time"
)

const ALGORITHM_LEITNER string = "LEITNER"

type algorithm interface {
	calculateDue(card Card, ease EaseFactor) (due time.Time, err error)
	calculateInterval(card Card, ease EaseFactor) (interval time.Duration, err error)
}

func getSRSInstance(algorithm string) algorithm {
	if algorithm == ALGORITHM_LEITNER {
		return &normalisedLeitner{}
	}
	return &unknownAlgorithm{}
}

type unknownAlgorithm struct{}

func (u *unknownAlgorithm) calculateDue(card Card, ease EaseFactor) (due time.Time, err error) {
	return time.Time{}, fmt.Errorf("unknownAlgorithm: Not Implemented")
}
func (u *unknownAlgorithm) calculateInterval(card Card, ease EaseFactor) (interval time.Duration, err error) {
	return 0, fmt.Errorf("unknownAlgorithm: Not Implmented")
}
