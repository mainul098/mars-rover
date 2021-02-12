package pkg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewRover(t *testing.T) {
	tt := []struct {
		tag              string
		rover            *Rover
		expectedPosition string
	}{
		{
			tag:              "Get a new rover with the initial position",
			rover:            NewRover(4, 2, "NORTH"),
			expectedPosition: "(4, 2) NORTH",
		},
	}

	for _, tc := range tt {
		t.Run(tc.tag, func(t *testing.T) {
			position, err := tc.rover.GetCurrentPosition()
			assert.Nil(t, err)
			assert.Equal(t, position, tc.expectedPosition)
		})
	}
}
