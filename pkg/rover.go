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
	for _, c := range commands {
		if err := r.performAction(string(c)); err != nil {
			return "", err
		}
	}
	return fmt.Sprintf("(%d, %d) %s", r.coordinate.x, r.coordinate.y, r.direction), nil
}

func (r *Rover) performAction(command string) error {
	switch string(command) {
	case "F":
		r.coordinate = r.grid.forward(r.coordinate, r.direction)
	case "B":
		r.coordinate = r.grid.backward(r.coordinate, r.direction)
	case "L":
		r.direction = r.grid.left(r.direction)
	case "R":
		r.direction = r.grid.right(r.direction)
	default:
		return errors.New("Invalid command")
	}
	return nil
}
