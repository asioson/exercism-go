// package etl implements a routine to transform
// a map from count to list of letters to
// a map from letter to count
package etl

import "strings"

// Transform takes in a map from an old scoring system 
// to a map of new scoring system
func Transform(data map[int][]string) map[string]int {
    result := map[string]int{}
    for k, list := range data {
        for _, v := range list {
            result[strings.ToLower(v)] = k
        }
    }
    return result
}

