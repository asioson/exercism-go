// package accumulate implements a routine that performs
// a function in each element of an input list to 
// produce another list of the transformed elements
package accumulate

// Accumulate takes in a list input and a function f and 
// returns a list resulting to the application of f
// to each of the list's contents
func Accumulate(input []string, f func(string) string) []string {
    result := []string{}
	for _, s := range input {
		result = append(result, f(s))
	}
	return result
}
