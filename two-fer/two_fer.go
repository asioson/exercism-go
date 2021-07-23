// Package twofer implements a routine that returns a sentence given a name
package twofer

// ShareWith takes a name X and returns a sentence in 
// the form: One for X, one for me.
// If name is empty, X is replaced with 'you'
func ShareWith(name string) string {
    if name == "" {
        name = "you"
    }
    return "One for " + name + ", one for me."
}
