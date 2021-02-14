package pkg

// Direction over the grid
type Direction string

// Directions for the allowed movement on the Grid
const (
	EAST  Direction = "EAST"
	NORTH Direction = "NORTH"
	WEST  Direction = "WEST"
	SOUTH Direction = "SOUTH"
)

// Rotation define the left or right direction for a axis on Grid
type Rotation struct {
	left  Direction
	right Direction
}

// Coordinate points on the Grid
type Coordinate struct {
	x int
	y int
}

// Grid for the rover coordinate plane with the obstacles
type Grid struct {
	rotations map[Direction]Rotation
	axes      map[Direction]Coordinate
	obstacles []Coordinate
}

// NewGrid to move the Rover
func NewGrid(obstacles []Coordinate) Grid {
	// define the rotation allwed on each Direction
	rotations := map[Direction]Rotation{
		EAST:  {NORTH, SOUTH},
		NORTH: {WEST, EAST},
		WEST:  {SOUTH, NORTH},
		SOUTH: {EAST, WEST},
	}

	// define the axis cooredinates
	axes := map[Direction]Coordinate{
		EAST:  {1, 0},
		NORTH: {0, 1},
		WEST:  {-1, 0},
		SOUTH: {0, -1},
	}

	return Grid{
		rotations: rotations,
		axes:      axes,
		obstacles: obstacles,
	}
}

func (g Grid) forward(coordinate Coordinate, direction Direction) (Coordinate, bool) {
	axis := g.axes[direction]
	newCoordinate := Coordinate{coordinate.x + axis.x, coordinate.y + axis.y}
	for _, obstacle := range g.obstacles {
		if newCoordinate.x == obstacle.x && newCoordinate.y == obstacle.y {
			return coordinate, true
		}
	}

	return newCoordinate, false
}

func (g Grid) backward(coordinate Coordinate, direction Direction) (Coordinate, bool) {
	axis := g.axes[direction]
	newCoordinate := Coordinate{coordinate.x - axis.x, coordinate.y - axis.y}
	for _, obstacle := range g.obstacles {
		if newCoordinate.x == obstacle.x && newCoordinate.y == obstacle.y {
			return coordinate, true
		}
	}
	return newCoordinate, false
}

func (g Grid) left(direction Direction) Direction {
	return g.rotations[direction].left
}

func (g Grid) right(direction Direction) Direction {
	return g.rotations[direction].right
}
