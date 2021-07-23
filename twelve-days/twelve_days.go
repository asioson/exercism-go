// package twelve implements routines that produces the individual verses and entire lyrics
// of the song 'Twelve days of Christmas'
package twelve

// Verse takes in a number x and returns verse x of the song
func Verse(x int) string {
    d := []string{"", "first", "second", "third", "fourth", "fifth",
            "sixth","seventh","eighth","ninth","tenth","eleventh","twelfth"}
    verse := "On the " + d[x] + " day of Christmas my true love gave to me: "
    gifts := ""
    switch (x) {
    case 12 :
        gifts += "twelve Drummers Drumming, "
        fallthrough
    case 11 :
        gifts += "eleven Pipers Piping, "
        fallthrough
    case 10 :
        gifts += "ten Lords-a-Leaping, "
        fallthrough
    case 9 :
        gifts += "nine Ladies Dancing, "
        fallthrough
    case 8 :
        gifts += "eight Maids-a-Milking, "
        fallthrough
    case 7 :
        gifts += "seven Swans-a-Swimming, "
        fallthrough
    case 6 :
        gifts += "six Geese-a-Laying, "
        fallthrough
    case 5 :
        gifts += "five Gold Rings, "
        fallthrough
    case 4 :
        gifts += "four Calling Birds, "
        fallthrough
    case 3 :
        gifts += "three French Hens, "
        fallthrough
    case 2 :
        gifts += "two Turtle Doves, "
        fallthrough
    case 1 :
        if (x > 1) {
            gifts += "and " 
        }
        gifts += "a Partridge in a Pear Tree."
    }
    
    return verse + gifts
}

// Song returns the entire lyrics of the song
func Song() string {
    song := Verse(1)
    for i := 2; i <= 12; i++ {
        song += "\n" + Verse(i)
    }
    return song
}
