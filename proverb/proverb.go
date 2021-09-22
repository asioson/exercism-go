// Package proverb implements a routine that produces
// proverbial rhyme
package proverb

// Proverb takes in a slice of items (each is a string) and generates 
// proverbial rhymes out of the items.
func Proverb(rhyme []string) []string {
    var proverb []string
    n := len(rhyme)
    if n > 0 {
        for i := 0; i < n - 1; i++ {
            line := "For want of a " + rhyme[i] + " the " + rhyme[i+1] + " was lost."
            proverb = append(proverb, line)
        }
        proverb = append(proverb, "And all for the want of a " + rhyme[0] + ".")
    }
	return proverb
}
