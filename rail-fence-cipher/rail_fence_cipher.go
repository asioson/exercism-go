// package railfence implements routines to encrypt and decrypt 
// messages using the rail fence method
package railfence


// Encode takes in a message and rail height and returns
// an encrypted message using the rail fence method
func Encode(mesg string, railH int) string {
    rails := make([]string, railH)
    down := true
    r := 0
    for _, ch := range mesg {
        rails[r] += string(ch)
        if down { r++ } else { r-- }
        if r == railH {
            r -= 2
            down = false
        }
        if r < 0 {
            r += 2
            down = true
        }
    }
    encrypted := ""
    for i := 0; i < railH; i++ {
        encrypted += rails[i]
    }
    return encrypted
}

// Decode takes in an encrypted message and rail height and 
// returns the decrypted message
func Decode(encrypted string, railH int) string {
    // step 1: determine the length of characters per rail
    railLen := make([]int, railH)
    down := true
    r := 0
    for i := 0; i < len(encrypted); i++ {
        railLen[r]++
        if down { r++ } else { r-- }
        if r == railH {
            r -= 2
            down = false
        }
        if r < 0 {
            r += 2
            down = true
        }
    }
    // step 2: extract the characters per rail
    railData := make([]string, railH)
    railData[0] = encrypted[:railLen[0]]
    k := railLen[0]
    for i := 1; i < railH; i++ {
        railData[i] = encrypted[k:k+railLen[i]]
        k += railLen[i]
    }
    // step 3: simulate the decoding process
    idx := make([]int, railH)
    mesg := ""
    down = true
    r = 0
    for i := 0; i < len(encrypted); i++ {
        mesg += string(railData[r][idx[r]])
        idx[r]++
        if down { r++ } else { r-- }
        if r == railH {
            r -= 2
            down = false
        }
        if r < 0 {
            r += 2
            down = true
        }
    }
    return mesg
}
