// package foodchain implements routines that generates
// the song 'I Know an Old Lady Who Swallowed a Fly'
package foodchain

// song phrases
var object = []string{
    "",
    "",
    "spider to catch the fly.",
    "bird to catch the spider that wriggled and jiggled and tickled inside her.",
    "cat to catch the bird.",
    "dog to catch the cat.",
    "goat to catch the dog.",
    "cow to catch the goat.",
}

var know = []string{
    "",
    "",
    "It wriggled and jiggled and tickled inside her.",
    "How absurd to swallow a bird!",
    "Imagine that, to swallow a cat!",
    "What a hog, to swallow a dog!",
    "Just opened her throat and swallowed a goat!",
    "I don't know how she swallowed a cow!",
    "She's dead, of course!",
}

var animal = []string{
    "",
    "fly",
    "spider",
    "bird",
    "cat",
    "dog",
    "goat",
    "cow",
    "horse",
}

// Verse(i) returns verse i of the song
func Verse(i int) string {
    if i < 1 || i > 8 { return "" }
    verse := "I know an old lady who swallowed a " + animal[i] + "."
    if i <= 7 {
        if i > 1 {
            verse += "\n" + know[i]
            for j := i; j > 1; j-- {
                verse += "\nShe swallowed the " + object[j]
            }
        }
        verse += "\nI don't know why she swallowed the fly. Perhaps she'll die."
    } else {
        verse += "\nShe's dead, of course!"
    }
    return verse
}

// Verses(i,j) returns verses i to j of the song
func Verses(i, j int) string {
    verses := ""
    if i <= j {
        verses = Verse(i)
        for k := i+1; k <= j; k++ {
            verses += "\n\n" + Verse(k)
        }
    }
    return verses
}

// Song() returns verses 1 to 8 of the song
func Song() string {
    return Verses(1,8) 
}
