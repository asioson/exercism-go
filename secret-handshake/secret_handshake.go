// package secret implements handshake routine
package secret

// Handshake takes in an in input integer and returns
// an array of events
func Handshake(input uint) (result []string) {
    masks  := []uint{1,2,4,8,16}
    events := []string{"wink","double blink","close your eyes","jump","reverse"}
	reverse := false
	for i, mask := range masks {
		if input & mask != 0 {
			if events[i] == "reverse" {
				reverse = true
			} else {
				result = append(result, events[i])
			}
		}
	}
	if reverse {
        last := len(result) - 1
        for i := 0; i < len(result) / 2; i++ {
            result[i], result[last-i] = result[last-i], result[i]
		}
	}
	return result
}
