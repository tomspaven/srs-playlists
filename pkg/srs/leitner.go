package srs

import (
	"math"
	"time"
)

type normalisedLeitner struct {
	initInterval time.Duration
}

func (nl *normalisedLeitner) calculateNext(card Card, ease EaseFactor) (nextVals, error) {
	if card.metaData.interval == 0 {
		return nextVals{
			time.Now().Add(nl.initInterval),
			nl.initInterval,
		}, nil
	}

	nextInterval := time.Duration(float64(nl.initInterval) * math.Pow(float64(ease), float64(card.metaData.numberReviews)))
	return nextVals{
		due:      time.Now().Add(nextInterval),
		interval: nextInterval,
	}, nil
}
