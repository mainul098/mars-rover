package rover

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
			tag:              "Rotate the rover in left for all directions",
			command:          "LLLL",
			expectedPosition: "(4, 2) EAST",
		},
		{
			tag:              "Rotate the rover in right for for all directions",
			command:          "RRRR",
			expectedPosition: "(4, 2) EAST",
		},
		{
			tag:              "Move the rover in positivie directions",
			command:          "FLFFFRFLB",
			expectedPosition: "(6, 4) NORTH",
		},
		{
			tag:              "Move the rover in negevate directions",
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
			tag:     "Rover stopped with obstacles",
			command: "LFFLFFR",
			obstacles: []Coordinate{
				{2, 4},
				{3, 5},
				{7, 4},
			},
			expectedPosition: "(3, 4) WEST STOPPED",
		},
		{
			tag:     "Rover moves without any obstacles to stop",
			command: "LFFLFFR",
			obstacles: []Coordinate{
				{1, 4},
				{3, 5},
				{7, 4},
			},
			expectedPosition: "(2, 4) NORTH",
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
