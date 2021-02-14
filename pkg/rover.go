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
func NewRover(x int, y int, direction Direction) *Rover {
	return &Rover{
		coordinate: Coordinate{x, y},
		direction:  direction,
		grid:       NewGrid(),
	}
}

// ExecuteCommand will excute the command strind and move the rover
func (r Rover) ExecuteCommand(commands string) (string, error) {
	for _, c := range commands {
		if err := r.performAction(string(c)); err != nil {
			return "", err
		}
	}
	return fmt.Sprintf("(%d, %d) %s", r.coordinate.x, r.coordinate.y, r.direction), nil
}

func (r *Rover) performAction(command string) error {
	var err error
	switch string(command) {
	case "F":
		if r.coordinate, err = r.grid.forward(r.coordinate, r.direction); err != nil {
			return err
		}
	case "B":
		if r.coordinate, err = r.grid.backward(r.coordinate, r.direction); err != nil {
			return err
		}
	case "L":
		if r.direction, err = r.grid.left(r.direction); err != nil {
			return err
		}
	case "R":
		d, err := r.grid.right(r.direction)
		if err != nil {
			return err
		}
		r.direction = d
	default:
		return errors.New("Invalid command")
	}
	return nil
}
