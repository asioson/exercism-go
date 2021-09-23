// Package cipher implements routines for simple ciphers
package cipher

import "unicode"

type CipherKey []int


// Encode takes in a mesg and returns an encrypted string
// using the given cipher key
func (key CipherKey) Encode(mesg string) string {
    n := len(key)
    code := ""
    i := 0
    for _, c := range mesg {
        if !unicode.IsLetter(c) {
            continue
        }
        if unicode.IsUpper(c) {
            c += 'a' - 'A'
        }
        offset := int(c - 'a') + key[i % n]
        if offset < 0 {
            offset += 26
        }
        c = rune(offset % 26) + 'a'
        i++
        code += string(c)
    }
    return code
}


// Decode takes in an encrypted string and produce the plaintext
// using the given cipher key
func (key CipherKey) Decode(code string) string {
    n := len(key)
    mesg := ""
    i := 0
    for _, c := range code {
        offset := int(c - 'a') - key[i % n]
        if offset < 0 {
            offset += 26
        }
        c = rune(offset % 26) + 'a'
        i++
        mesg += string(c)
    }
    return mesg
}

// NewCaesar generates a new Caesar cipher key
func NewCaesar() CipherKey {
    return CipherKey{3}
}


// NewShift generates a new Shift cipher key
func NewShift(n int) CipherKey {
    if n <= -26 || n >= 26 || n == 0 {
        return nil
    }
    return CipherKey{n}
}


// NewVigenere generates a new Vigenere cipher key
func NewVigenere(s string) CipherKey {
    n := len(s)
    if n == 0 {
        return nil
    }
    allZero := true
    key := make(CipherKey, n)
    for i, c := range s {
        if !unicode.IsLetter(c) || unicode.IsUpper(c) {
            return nil
        }
        key[i] = int(c - 'a')
        if key[i] != 0 {
            allZero = false
        }
    }
    if allZero {
        return nil
    }
    return key
}
