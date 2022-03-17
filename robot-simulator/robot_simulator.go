// package robot implements routines needed to simulate robot movements
package robot

// the constants N, E, S, and W are defined in strict order to simplify
// left and right turns
const (
	N = iota
	E
	S
	W
)

// String returns the string equivalent of director dir
func (dir Dir) String() string {
	name := map[Dir]string{N: "North", E: "East", S: "South", W: "West"}
	if n, ok := name[dir]; ok {
		return n
	}
	return ""
}

// Step1 routines

// Right makes the robot turn right
func Right() {
	Step1Robot.Dir = (Step1Robot.Dir + 1) % 4
}

// Left makes the robot turn left
func Left() {
	Step1Robot.Dir = (Step1Robot.Dir + 3) % 4
}

// Advance makes the robot move one step in the direction it is facing
func Advance() {
	switch Step1Robot.Dir {
	case N:
		Step1Robot.Y++
	case E:
		Step1Robot.X++
	case S:
		Step1Robot.Y--
	case W:
		Step1Robot.X--
	}
}

// Step 2 routines

// Right makes the robot turn right
func (robot *Step2Robot) Right() {
	robot.Dir = (robot.Dir + 1) % 4
}

// Left makes the robot turn left
func (robot *Step2Robot) Left() {
	robot.Dir = (robot.Dir + 3) % 4
}

// Advance makes the robot move one step in the direction it is facing
func (robot *Step2Robot) Advance(rect Rect) bool {
	switch robot.Dir {
	case N:
		if robot.Pos.Northing < rect.Max.Northing {
			robot.Pos.Northing++
			return true
		}
	case E:
		if robot.Pos.Easting < rect.Max.Easting {
			robot.Pos.Easting++
			return true
		}
	case S:
		if robot.Pos.Northing > rect.Min.Northing {
			robot.Pos.Northing--
			return true
		}
	case W:
		if robot.Pos.Easting > rect.Min.Easting {
			robot.Pos.Easting--
			return true
		}
	}
	return false
}

// IsOutside checks if the robot is outside the extent of the area
func (robot *Step2Robot) IsOutside(rect Rect) bool {
	if robot.Pos.Northing < rect.Min.Northing {
		return true
	}
	if robot.Pos.Northing > rect.Max.Northing {
		return true
	}
	if robot.Pos.Easting < rect.Min.Easting {
		return true
	}
	if robot.Pos.Easting > rect.Max.Easting {
		return true
	}
	return false
}

type Action int

// commands that the robot may process
const (
	advance = iota
	left
	right
	end
	unknown
)

// StartRobot takes commands from a script and inserts the appropriate command in action channel
func StartRobot(command chan Command, action chan Action) {
	for c := range command {
		switch c {
		case 'A':
			action <- advance
		case 'L':
			action <- left
		case 'R':
			action <- right
		}
	}
	action <- end
}

// Room sets up the sequence of robot actions to report channel
func Room(extent Rect, robot Step2Robot, action chan Action, report chan Step2Robot) {
	for a := range action {
		switch a {
		case advance:
			robot.Advance(extent)
		case left:
			robot.Left()
		case right:
			robot.Right()
		case end:
			report <- robot
			return
		}
	}
}

// Step 3
type Action3 struct {
	name    string
	command Action
}

// StartRobot3 takes commands from a script and inserts the appropriate command in action channel
// for a specific named robot
func StartRobot3(name, script string, action chan Action3, log chan string) {
	if name == "" {
		return
	}
	for _, a := range script {
		switch a {
		case 'A':
			action <- Action3{name, advance}
		case 'L':
			action <- Action3{name, left}
		case 'R':
			action <- Action3{name, right}
		default:
			action <- Action3{name, unknown}
			return
		}
	}
	action <- Action3{name, end}
}

// Process inspects the robots in the robot list and returns respective IDs of robots and locations occupied by the robots
func Process(extent Rect, robots []Step3Robot, log chan string) (map[string]int, map[Pos]bool, bool) {
	ID := map[string]int{}
	occupied := map[Pos]bool{}
	for i, robot := range robots {
		if robot.Name == "" {
			log <- "Robot has no name"
			return ID, occupied, false
		} else if ID[robot.Name] != 0 {
			log <- "Robot name '" + robot.Name + "' is already taken"
			return ID, occupied, false
		} else if occupied[robot.Pos] {
			log <- "Another robot occupies this position"
			return ID, occupied, false
		} else if robot.IsOutside(extent) {
			log <- "Robot is outside the valid area"
			return ID, occupied, false
		}
		ID[robot.Name] = i + 1
		occupied[robot.Pos] = true
	}
	return ID, occupied, true
}

// Room3 sets up the sequence of robot actions to report channel for robots in robot list
func Room3(extent Rect, robots []Step3Robot, action chan Action3, rep chan []Step3Robot, log chan string) {
	defer close(rep)
	ID, occupied, successful := Process(extent, robots, log)
	if !successful {
		return
	}
	for a := range action {
		id, ok := ID[a.name]
		if !ok {
			log <- "Robot '" + a.name + "' is unknown"
			return
		}
		idx := id - 1
		switch a.command {
		case advance:
			r := robots[idx].Step2Robot
			prevPos := r.Pos
			if !r.Advance(extent) {
				log <- "Robot moving to invalid area"
			} else if occupied[r.Pos] {
				log <- "Robot moving to an occupied position"
			} else {
				robots[idx].Step2Robot = r
				occupied[r.Pos] = true
				occupied[prevPos] = false
			}
		case left:
			robots[idx].Step2Robot.Left()
		case right:
			robots[idx].Step2Robot.Right()
		case end:
			delete(ID, a.name)
			if len(ID) == 0 {
				rep <- robots
				return
			}
		case unknown:
			log <- "Unknown robot command"
			return
		}
	}
}
