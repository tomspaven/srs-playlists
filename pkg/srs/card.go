package srs

import (
	"time"

	"github.com/jinzhu/gorm"
)

type cardMeta struct {
	dateAdded     time.Time
	firstReview   time.Time
	latestReview  time.Time
	due           time.Time
	interval      time.Duration
	numberReviews int
}

type Card struct {
	gorm.Model
	FrontData string
	metaData  cardMeta
	alg       algorithm
}

func (c *Card) DueDate() time.Time {
	return c.metaData.due
}

func (c *Card) Interval() time.Duration {
	return c.metaData.interval
}

func NewCard(data string, algorithm string) (*Card, error) {
	retCard := &Card{
		FrontData: data,
		metaData: cardMeta{
			dateAdded: time.Now(),
		},
		alg: getSRSInstance(algorithm),
	}
	if err := retCard.updateNextData(initialEaseFactor); err != nil {
		return nil, err
	}
	return retCard, nil
}

func (c *Card) UpdateCard(ease EaseFactor) error {
	c.metaData.latestReview = time.Now()
	c.metaData.numberReviews++
	if err := c.updateNextData(ease); err != nil {
		return err
	}
	return nil
}

func (c *Card) updateNextData(ease EaseFactor) error {
	var err error
	var next nextVals
	if next, err = c.alg.calculateNext(*c, ease); err != nil {
		return err
	}

	c.metaData.interval = next.interval
	c.metaData.due = next.due
	return nil
}
