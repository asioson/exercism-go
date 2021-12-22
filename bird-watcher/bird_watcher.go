// package birdwatcher implements methods that keep track of birds
package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
    birdsSeen := 0
    for i := 0; i < len(birdsPerDay); i++ {
        birdsSeen += birdsPerDay[i]
    }
    return birdsSeen
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
    birdsSeen := 0
    offset := (week - 1) * 7
    for i := 0; i < 7; i++ {
        birdsSeen += birdsPerDay[offset + i]
    }
    return birdsSeen;
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
    for i := 0; i < len(birdsPerDay); i++ {
        birdsPerDay[i] += ((i + 1) % 2)
    }
    return birdsPerDay 
}
