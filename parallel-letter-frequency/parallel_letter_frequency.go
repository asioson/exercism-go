// package letter implements routines that does sequential processing 
// and concurrent processing of an array or strings
// (c) 2021 allan.sioson@gmail.com
package letter

import "sync"

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
    m := FreqMap{}
    for _, r := range s {
        m[r]++
    }
    return m
}

// CuncurrentFrequency takes in an array of strings and do frequency count of
// characters in each in parallel using a shared map (sync.Map).  After the
// the parallel go routines are done with their respective frequency counts
// the contents of the shared map is transferred to an instance of FreqMap
// which is then returned as a result
func ConcurrentFrequency(s []string) FreqMap {
    sm := sync.Map{}
    wg := sync.WaitGroup{}
    mutex := sync.Mutex{}
    freqCount :=  func(lc *sync.Mutex, wg *sync.WaitGroup, t string) {
        defer wg.Done()
        for _, c := range t {
            lc.Lock()
            v, ok := sm.Load(c) 
            if ok {
                count, _ := v.(int)
                sm.Store(c,count+1)
            } else {
                sm.Store(c,1)
            }
            lc.Unlock()
        }
    }
    wg.Add(len(s))
    for i := 0; i < len(s); i++ {
        go freqCount(&mutex, &wg, s[i])
    }
    wg.Wait()
    m := FreqMap{}
    sm.Range(func(k, v interface {}) bool {
        count, _ := v.(int)
        key, _ := k.(rune)
        m[key] = count 
        return true
    })
    return m
}
