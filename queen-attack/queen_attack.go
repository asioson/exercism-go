// package queenattack implements a routine that checks
// if queens on the board can attack each other
package queenattack

import "fmt"

func valid(x string) bool {
    if len(x) != 2 { return false }
    if x[0] < 'a' || x[0] > 'h' { return false }
    if x[1] < '1' || x[1] > '8' { return false }
    return true
}

// CanQueenAttack checks if the queens positioned
// at w and b can attack each other
func CanQueenAttack(w, b string) (bool,error) {
    if valid(w) && valid(b) && w != b {
        if w[0] == b[0] {
            return true, nil 
        } else if w[1] == b[1] {
            return true, nil 
        } else if w[0] - b[0] == w[1] - b[1] {
            return true, nil 
        } else if w[0] - b[0] == b[1] - w[1] {
            return true, nil 
        }
        return false, nil
    } 
    return false, fmt.Errorf("invalid positions") 
}
