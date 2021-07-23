// package reverse implements a routine that reverses the
// characters of an input string
package reverse

// Reverse takes a word and returns a version with characters 
// in reverse order
func Reverse(word string) string {
    revWord := ""
    for _, c := range word {
        revWord = string(c) + revWord
    }
    return revWord  
}
