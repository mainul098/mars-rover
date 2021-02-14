package pkg

import (
	"errors"
	"fmt"
)

// Rover to explore the Mars
type Rover struct {
	coordinate Coordinate
	direction  Direction
	grid       Grid
	obstracles map[Coordinate]bool
}

// NewRover will give a new rover landed on Mars with the initial position
func NewRover(x int, y int, direction Direction, obstacles []Coordinate) *Rover {
	obstaclesMap := make(map[Coordinate]bool)
	for _, obstacle := range obstacles {
		obstaclesMap[obstacle] = true
	}

	return &Rover{
		coordinate: Coordinate{x, y},
		direction:  direction,
		grid:       NewGrid(),
		obstracles: obstaclesMap,
	}
}

// ExecuteCommand will excute the command strind and move the rover
func (r *Rover) ExecuteCommand(commands string) (string, error) {
	for _, c := range commands {
		switch string(c) {
		case "F":
			coordinate := r.grid.forward(r.coordinate, r.direction)
			if _, ok := r.obstracles[coordinate]; ok {
				return r.getPosition(true), nil
			}
			r.coordinate = coordinate
		case "B":
			coordinate := r.grid.backward(r.coordinate, r.direction)
			if _, ok := r.obstracles[coordinate]; ok {
				return r.getPosition(true), nil
			}
			r.coordinate = coordinate
		case "L":
			r.direction = r.grid.left(r.direction)
		case "R":
			r.direction = r.grid.right(r.direction)
		default:
			return "", errors.New("Invalid command")
		}
	}
	return r.getPosition(false), nil
}

func (r *Rover) getPosition(isStopped bool) string {
	if isStopped {
		return fmt.Sprintf("(%d, %d) %s STOPPED", r.coordinate.x, r.coordinate.y, r.direction)
	}
	return fmt.Sprintf("(%d, %d) %s", r.coordinate.x, r.coordinate.y, r.direction)
}
