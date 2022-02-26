// package forth implements a forth simulator
package forth

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// CleanUp takes in forth code, takes care of user-defined names and returns cleaned up code
func CleanUp(input []string) ([][]string, error) {
	f := map[string][]string{}
	code := [][]string{}
	for _, line := range input {
		token := strings.Split(strings.ToLower(line), " ")
		n := len(token)
		if n > 0 && token[0] == ":" {
			name := token[1]
			if unicode.IsDigit(rune(name[0])) {
				return code, fmt.Errorf("cannot redefine numbers")
			}
			tf := map[string][]string{}
			if v, ok := f[name]; ok {
				tf[name] = v
			}
			f[name] = make([]string, 0)
			for _, t := range token[2 : n-1] {
				if v, ok := tf[t]; ok {
					f[name] = append(f[name], v...)
				} else if v, ok := f[t]; ok {
					f[name] = append(f[name], v...)
				} else {
					f[name] = append(f[name], t)
				}
			}
		} else {
			c := []string{}
			for _, t := range token {
				if v, ok := f[t]; ok {
					c = append(c, v...)
				} else {
					c = append(c, t)
				}
			}
			code = append(code, c)
		}
	}
	return code, nil
}

// Forth takes in forth code and returns the result of computations
func Forth(input []string) ([]int, error) {
	stack := []int{}
	inp, err := CleanUp(input)
	if err != nil {
		return stack, err
	}
	for _, code := range inp {
		for _, item := range code {
			if unicode.IsDigit(rune(item[0])) {
				v, _ := strconv.Atoi(item)
				stack = append(stack, v)
			} else {
				n := len(stack)
				switch item {
				case "+":
					if n < 2 {
						return stack, fmt.Errorf("not enough operands for +")
					}
					b, a := stack[n-1], stack[n-2]
					stack = append(stack[:n-2], a+b)
				case "-":
					if n < 2 {
						return stack, fmt.Errorf("not enough operands for -")
					}
					b, a := stack[n-1], stack[n-2]
					stack = append(stack[:n-2], a-b)
				case "*":
					if n < 2 {
						return stack, fmt.Errorf("not enough operands for *")
					}
					b, a := stack[n-1], stack[n-2]
					stack = append(stack[:n-2], a*b)
				case "/":
					if n < 2 {
						return stack, fmt.Errorf("not enough operands for /")
					}
					b, a := stack[n-1], stack[n-2]
					if b == 0 {
						return stack, fmt.Errorf("division by zero")
					}
					stack = append(stack[:n-2], a/b)
				case "swap":
					if n < 2 {
						return stack, fmt.Errorf("not enough operands to swap")
					}
					b, a := stack[n-1], stack[n-2]
					stack = append(stack[:n-2], b, a)
				case "over":
					if n < 2 {
						return stack, fmt.Errorf("not enough operands")
					}
					stack = append(stack, stack[n-2])
				case "dup":
					if n < 1 {
						return stack, fmt.Errorf("nothing to dup")
					}
					stack = append(stack, stack[n-1])
				case "drop":
					if n < 1 {
						return stack, fmt.Errorf("nothing to drop")
					}
					stack = stack[:n-1]
				default:
					return []int{}, fmt.Errorf("unknown identifier: %s", item)
				}
			}
		}
	}
	return stack, nil
}
