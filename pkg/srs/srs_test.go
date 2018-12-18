package srs_test

import (
	"srs-playlists/pkg/srs"
	"testing"
)

func GetCardsDueNow_test(t *testing.T) {
	_ = &srs.SRS{
		Cfg: srs.SRSConfig{
			Algorithm: srs.ALGORITHM_LEITNER,
		},
	}

}
