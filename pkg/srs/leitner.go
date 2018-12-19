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
		return nl.initialNextVals(), nil
	}

	nextInterval := nl.calculateNextInterval(card, ease)
	return nextVals{
		due:      time.Now().Add(nextInterval),
		interval: nextInterval,
	}, nil
}

func (nl *normalisedLeitner) initialNextVals() nextVals {
	return nextVals{
		due:      time.Now().Add(nl.initInterval),
		interval: nl.initInterval,
	}
}

func (nl *normalisedLeitner) calculateNextInterval(card Card, ease EaseFactor) time.Duration {
	return time.Duration(float64(nl.initInterval) * math.Pow(float64(ease), float64(card.metaData.numberReviews)))
}
