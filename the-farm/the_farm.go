// package thefarm implements SillyNephewError type and DivideFood() method
package thefarm

import "fmt"
import "errors"

type SillyNephewError struct {
    cows int
}

func (e *SillyNephewError) Error() string {
    return fmt.Sprintf("silly nephew, there cannot be %d cows", e.cows)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
    if cows < 0 {
        return 0, &SillyNephewError{cows:cows}
    } else if cows == 0 {
        return 0, errors.New("Division by zero")
    } else {
        fodder, err := weightFodder.FodderAmount()
        if fodder < 0 {
            return 0, errors.New("Negative fodder")
        }
        if err == nil {
            return fodder / float64(cows), nil
        } else {
            if err == ErrScaleMalfunction {
                if (fodder > 0) {
                    return (fodder * 2) / float64(cows), nil 
                }
            }
            return 0, err 
        }
    }
}
