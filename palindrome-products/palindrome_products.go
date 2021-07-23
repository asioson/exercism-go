// package palindrome implements routines that determines
// the smallest and largest palindrome products in a 
// given range

package palindrome

import "fmt"

type Product struct {
    Product int
    Factorizations [][2]int
}

func IsPalindrome(x int) bool {
    if x < 10 { return true }
    y := 0 
    tx := x
    for tx > 0 {
        r := tx % 10
        tx = tx / 10
        y = y * 10 + r
    }
    return x == y
}

// Products function takes fmin and fmax and returns the largest and smallest 
// palindromes, along with the factors of each within the range [fmin, fmax]. 
// If the largest or smallest palindrome has more than one pair of factors 
// within the range, the function returns all the pairs.

func Products(fmin, fmax int) (Product, Product, error) {
    pmin := Product{}
    pmax := Product{}
    if fmin > fmax { return pmin, pmax, fmt.Errorf("fmin > fmax...") } 
    first := true
    for x := fmin; x <= fmax; x++ {
        for y := x; y <= fmax; y++ {
            if IsPalindrome(y * x) {
                if first {
                    first = false
                    pmin.Product = x * y
                    pmin.Factorizations = append(pmin.Factorizations,[2]int{x,y})
                    pmax.Product = x * y
                    pmax.Factorizations = append(pmax.Factorizations,[2]int{x,y})
                } else {
                    if y * x < pmin.Product {
                        pmin.Product = x * y
                        pmin.Factorizations = [][2]int{}
                        pmin.Factorizations = append(pmin.Factorizations,[2]int{x,y})
                    } else if x * y == pmin.Product {
                        pmin.Factorizations = append(pmin.Factorizations,[2]int{x,y})
                    } else if x * y > pmax.Product {
                        pmax.Product = x * y
                        pmax.Factorizations = [][2]int{}
                        pmax.Factorizations = append(pmax.Factorizations,[2]int{x,y})
                    } else if x * y == pmax.Product {
                        pmax.Factorizations = append(pmax.Factorizations,[2]int{x,y})
                    }
                }
            }
        }
    }
    if len(pmin.Factorizations) == 0  {
        return pmin, pmax, fmt.Errorf("no palindromes...")
    }
    return pmin, pmax, nil
}

