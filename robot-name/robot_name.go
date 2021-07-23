// package robotname implements routines that generates 
// random robot names
package robotname

import "fmt"
import "math/rand"
import "strconv"

const MAX = 676000

var generated = [MAX]bool{false}

type Robot struct {
    name string
}

// Name generates a new robot name if the object is just
// instantiated and then it returns the name
func (r *Robot) Name() (string, error) {
    if r.name == "" {
        alpha := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
        x := rand.Intn(MAX)
        for generated[x] {
            x = rand.Intn(MAX)
        }
        generated[x] = true
        y := x / 1000
        r.name += string(alpha[y / 26])
        r.name += string(alpha[y % 26])
        x = x % 1000
        if x < 10 {
            r.name += "00" + strconv.Itoa(x)
        } else if x < 100 {
            r.name += "0" + strconv.Itoa(x)
        } else {
            r.name += strconv.Itoa(x)
        }
    }
    fmt.Print("")
    return r.name, nil
}

// Reset simply sets the robot name to empty string
func (r *Robot) Reset() {
    r.name = ""
}
