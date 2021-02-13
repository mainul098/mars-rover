package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecuteRoverCommand(t *testing.T) {
	tt := []struct {
		tag              string
		command          string
		expectedPosition string
	}{
		{
			tag:              "No movement",
			command:          "",
			expectedPosition: "(4, 2) EAST",
		},
		{
			tag:              "Rotate the rover in left for one full cycle.Command LLLL",
			command:          "LLLL",
			expectedPosition: "(4, 2) EAST",
		},
		{
			tag:              "Rotate the rover in right for one full cycle.Command RRRR",
			command:          "RRRR",
			expectedPosition: "(4, 2) EAST",
		},
		{
			tag:              "Move the rover with FLFFFRFLB command",
			command:          "FLFFFRFLB",
			expectedPosition: "(6, 4) NORTH",
		},
	}

	for _, tc := range tt {
		t.Run(tc.tag, func(t *testing.T) {
			r := NewRover(4, 2, EAST)
			position, err := r.ExecuteCommand(tc.command)
			assert.Nil(t, err)
			assert.Equal(t, tc.expectedPosition, position)
		})
	}
}
