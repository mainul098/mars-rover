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
}

// NewRover will give a new rover landed on Mars with the initial position
func NewRover(x int, y int, direction Direction, obstacles []Coordinate) *Rover {
	return &Rover{
		coordinate: Coordinate{x, y},
		direction:  direction,
		grid:       NewGrid(obstacles),
	}
}

// ExecuteCommand will excute the command strind and move the rover
func (r *Rover) ExecuteCommand(commands string) (string, error) {
	var isStopped bool
	for _, c := range commands {
		switch string(c) {
		case "F":
			r.coordinate, isStopped = r.grid.forward(r.coordinate, r.direction)
			if isStopped {
				return r.getPosition(isStopped), nil
			}
		case "B":
			r.coordinate, isStopped = r.grid.backward(r.coordinate, r.direction)
			if isStopped {
				return r.getPosition(isStopped), nil
			}
		case "L":
			r.direction = r.grid.left(r.direction)
		case "R":
			r.direction = r.grid.right(r.direction)
		default:
			return "", errors.New("Invalid command")
		}
	}
	return r.getPosition(isStopped), nil
}

func (r *Rover) getPosition(isStopped bool) string {
	if isStopped {
		return fmt.Sprintf("(%d, %d) %s STOPPED", r.coordinate.x, r.coordinate.y, r.direction)
	}
	return fmt.Sprintf("(%d, %d) %s", r.coordinate.x, r.coordinate.y, r.direction)
}
