// package series implements routines needed to produce
// string series given a starting input string
package series

// All takes in length n and string s to produce a string
// array with string elements with length n.
func All(n int, s string) []string {
   strList := []string{}
   for i := 0; i < len(s) - n + 1; i++ {
       strList = append(strList,s[i:i+n])
   }
   return strList
}


// UnsafeFirst returns the first element of a series
func UnsafeFirst(n int, s string) string {
    k := len(s)
    if k < n {
        return ""
    }
    return s[:n] 
}
