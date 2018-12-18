package srs

import (
	"math"
	"time"
)

type normalisedLeitner struct {
	initInterval time.Duration
	eFactor      int
}

func (nl *normalisedLeitner) calculateDue(card Card, ease EaseFactor) (due time.Time, err error) {
	return time.Now().Add(card.metaData.interval), nil
}

func (nl *normalisedLeitner) calculateInterval(card Card, ease EaseFactor) (interval time.Duration, err error) {
	if card.metaData.interval == 0 {
		return nl.initInterval, nil
	}
	next := time.Duration(float64(nl.initInterval) * math.Pow(float64(ease), float64(card.metaData.numberReviews)))
	return next, nil
}
