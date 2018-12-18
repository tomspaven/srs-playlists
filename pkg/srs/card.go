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
	Data     string
	metaData cardMeta
	alg      algorithm
}

func (c *Card) DueDate() time.Time {
	return c.metaData.due
}

func (c *Card) Interval() time.Duration {
	return c.metaData.interval
}

func NewCard(algorithm string) (*Card, error) {
	retCard := &Card{
		metaData: cardMeta{
			dateAdded: time.Now(),
		},
		alg: getSRSInstance(algorithm),
	}
	if err := retCard.updateDue(1); err != nil {
		return nil, err
	}
	if err := retCard.updateInterval(1); err != nil {
		return nil, err
	}
	return retCard, nil
}

func (c *Card) UpdateCard(ease EaseFactor) error {
	c.metaData.latestReview = time.Now()
	c.metaData.numberReviews++
	if err := c.updateDue(ease); err != nil {
		return err
	}

	if err := c.updateInterval(ease); err != nil {
		return err
	}
	return nil
}

func (c *Card) updateDue(ease EaseFactor) error {
	var err error
	c.metaData.due, err = c.alg.calculateDue(*c, ease)
	if err != nil {
		return err
	}
	return nil
}

func (c *Card) updateInterval(ease EaseFactor) error {
	var err error
	c.metaData.interval, err = c.alg.calculateInterval(*c, ease)
	if err != nil {
		return err
	}
	return nil
}
