package srs_test

import (
	"srs-playlists/pkg/srs"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCardUnknownAlgorithm(t *testing.T) {
	card, err := srs.NewCard("URL://FIEJFIEJI", "DICKMAN")
	assert.Nil(t, card, "Expected card to be nil")
	assert.NotNil(t, err, "Expected error to be non-nil")
}

func TestNewCardLeitnerAlgorithm(t *testing.T) {
	card, err := srs.NewCard("URL://FIEJFIEJI", srs.AlgorithmLeitner)
	assert.NotNil(t, card, "Expected card to be non-nil")
	assert.Nil(t, err, "Expected error to be nil")
}
