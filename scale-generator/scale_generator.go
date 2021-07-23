// package scale implements a routine that generates the scale
// given a tonic and intervals
package scale

import "strings"

// Scale takes in tonic and interval information and returns
// the appropriate scale
func Scale(tonic string, interval string) []string {
    scale := []string{"C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab", "A", "Bb", "B"}
    if strings.Contains(":C:G:D:A:E:B:F#:a:e:b:f#:c#:g#:d#:",":"+tonic+":") {
        scale = []string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"}
    }
    tonic = strings.Title(tonic)
    for i, note := range scale {
        if note == tonic {
            scale = append(scale[i:], scale[:i]...)
            break
        }
    }
    if interval == "" { return scale }
    rScale := []string{}
    stepSize := map[string]int{"m": 1, "M": 2, "A": 3}
    i := 0
    for _, delta := range strings.Split(interval, "") {
        rScale = append(rScale, scale[i % 12])
        i += stepSize[delta] 
    }
    return rScale
}
