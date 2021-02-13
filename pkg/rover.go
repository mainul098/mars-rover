package pkg

import "fmt"

// Rover to explore the Mars
type Rover struct {
	x         int
	y         int
	direction string
}

// NewRover will give a new rover landed on Mars with the initial position
func NewRover(x int, y int, direction string) *Rover {
	return &Rover{
		x:         x,
		y:         y,
		direction: direction,
	}
}

// ExecuteCommand will excute the command strind and move the rover
func (r *Rover) ExecuteCommand(command string) error {
	for _, c := range command {
		if string(c) == "F" {
			r.forward()
		}
	}
	return nil
}

func (r *Rover) forward() error {
	r.x += 1
	return nil
}

// GetCurrentPosition will give the current position for the rover
func (r *Rover) GetCurrentPosition() (string, error) {
	return fmt.Sprintf("(%d, %d) %s", r.x, r.y, r.direction), nil
}
