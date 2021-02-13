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
			expectedPosition: "(4, 2) NORTH",
		},
		{
			tag:              "Move forward",
			command:          "F",
			expectedPosition: "(5, 2) NORTH",
		},
		{
			tag:              "Move backward",
			command:          "B",
			expectedPosition: "(3, 2) NORTH",
		},
		{
			tag:              "Move left",
			command:          "L",
			expectedPosition: "(4, 2) EAST",
		},
		{
			tag:              "Move right",
			command:          "R",
			expectedPosition: "(4, 2) WEST",
		},
	}

	for _, tc := range tt {
		t.Run(tc.tag, func(t *testing.T) {
			r := NewRover(4, 2, "NORTH")
			err := r.ExecuteCommand(tc.command)
			assert.Nil(t, err)

			position, err := r.GetCurrentPosition()
			assert.Nil(t, err)
			assert.Equal(t, position, tc.expectedPosition)
		})
	}
}
