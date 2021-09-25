// Package wordy implements a routine to parse and evaluate
// a math expression in words to an integer value
package wordy

import (
    "fmt"
    "strings"
)

// Answer takes in a question, parses it, and evaluates the
// math expression in words to an integer value
func Answer(question string) (int, bool) {
    n := len(question)
    if n < 8 {
        return 0, false
    }
    if question[:8] != "What is " {
        return 0, false
    }
    data := strings.Replace(question[8:n-1],"plus","+", -1)
    data = strings.Replace(data,"minus","-", -1)
    data = strings.Replace(data,"multiplied by","*", -1)
    data = strings.Replace(data,"divided by","/", -1)
    data = strings.Replace(data,"raised to the","^", -1)
    data = strings.Replace(data,"th power","", -1)

    parsedData := (strings.Split(data," "))
    m := len(parsedData)
    if m % 2 == 0 || parsedData[0] == "+" || parsedData[0] == "-" {
        return 0, false
    }
    x := 0
    _, err := fmt.Sscanf(parsedData[0], "%d", &x)
    if err == nil && m > 1 {
        y := 0
        for i := 1; i < m; i += 2 {
            op := parsedData[i]
            _, err = fmt.Sscanf(parsedData[i+1], "%d", &y)
            if err != nil {
                return 0, false
            }
            switch op {
            case "+" :  x += y
            case "-" :  x -= y
            case "*" :  x *= y
            case "/" :  
                if y == 0 {
                    return 0, false
                }
                x /= y
            case "^" :
                if y < 0 {
                    return 0, false
                } else if y == 0 {
                    x = 1
                } else {
                    tx := x
                    for j := 0; j < y-1; j++ {
                        x *= tx
                    }
                }
            default:
                return 0, false
            }
        }
    } 
    return x, true 
}
