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
			tag:              "Rotate the rover in left for one full cycle using LLLL command",
			command:          "LLLL",
			expectedPosition: "(4, 2) EAST",
		},
		{
			tag:              "Rotate the rover in right for one full cycle using RRRR command",
			command:          "RRRR",
			expectedPosition: "(4, 2) EAST",
		},
		{
			tag:              "Move the rover using FLFFFRFLB command",
			command:          "FLFFFRFLB",
			expectedPosition: "(6, 4) NORTH",
		},
		{
			tag:              "Move the rover in negevate direction using RFFFFR command",
			command:          "RFFFFR",
			expectedPosition: "(4, -2) WEST",
		},
	}

	for _, tc := range tt {
		t.Run(tc.tag, func(t *testing.T) {
			r := NewRover(4, 2, EAST, nil)
			position, err := r.ExecuteCommand(tc.command)
			assert.Nil(t, err)
			assert.Equal(t, tc.expectedPosition, position)
		})
	}
}

func TestExecuteRoverCommandWithObstacles(t *testing.T) {
	tt := []struct {
		tag              string
		command          string
		obstacles        []Coordinate
		expectedPosition string
	}{
		{
			tag:     "Move the rover with no obstacles",
			command: "LFFLFFBR",
			obstacles: []Coordinate{
				{2, 4},
				{3, 5},
				{7, 4},
			},
			expectedPosition: "(3, 4) WEST STOPPED",
		},
	}

	for _, tc := range tt {
		t.Run(tc.tag, func(t *testing.T) {
			r := NewRover(4, 2, EAST, tc.obstacles)
			position, err := r.ExecuteCommand(tc.command)
			assert.Nil(t, err)
			assert.Equal(t, tc.expectedPosition, position)
		})
	}
}
