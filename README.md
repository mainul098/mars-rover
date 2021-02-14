# TDD Kata on Mars Exploration
Execise the TDD Kata in a TDD way (red-green-refactor mantra).
## Problem Statement
Develop an API that trasnslate the commands setn from earcth to instructions that are undersatood by the rover. 

- The rover will be set to an initial coordinates and direction by argument `(x,y, direction)` eg `(4, 2, EAST)`
- The rovee is given a command string which contains miltiple commands. 
    - The valid commands are:
        - F -> Move forward on the current heading
        - B -> Move backward on the current heading
        - L -> Rotate left by `90-degrees`
        - R -> Rotate right by `90-degrees`
    - An example command might be `​FLFFFRFLB`
- Output the position of the Rover after executing the command

## TDD Steps
- [x] Initizalize the rover with initial coordinate and directions
- [x] Implement commands that move the rover forward/backward `(F,B)`
- [x] Implement commands that rotae the rover left/right `(L,R)`
- [x] Implement the obstacles detection of the rover on move forward/backward


