// package variablelengthquantity implements a VLQ encoder and decoder
package variablelengthquantity

import "fmt"


// EncodeVarint takes a slice of uint32 values and 
// encodes those to equivalent slice of VLQ encoded bytes
func EncodeVarint(input []uint32) []byte {
    vlq := []byte{} 
    for _, v := range(input) {
        flag := false
        item := []byte{}
        for x := v; x > 0; x = x >> 7 {
            y := x & 0x7f
            if flag {
                y = y | 0x80
            }
            item = append([]byte{byte(y)}, item ...)
            flag = true
        }
        if len(item) == 0 {
            vlq = append(vlq, 0)
        } else {
            vlq = append(vlq, item ...)
        }
    }
    if len(vlq) == 0 {
        vlq = append(vlq, 0)
    }
    return vlq 
}


// DecodeVarint takes a slice of VLQ encoded bytes and 
// decodes those to the equivalent slice of uint32 values
func DecodeVarint(input []byte) ([]uint32, error) {
    decVlq := []uint32{}
    x := uint32(0)
    complete := true 
    for i, v := range(input) {
        fmt.Println(i, v)
        if v & 0x80 == 0 {
            x = (x << 7) | uint32(v)
            decVlq = append(decVlq, x)
            complete = true
            x = 0
        } else {
           x = (x << 7) | uint32(v & 0x7f)
           complete = false
        }
    }
    if !complete {
        return decVlq, fmt.Errorf("incomplete sequence") 
    }
    return decVlq, nil
}
