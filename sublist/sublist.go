// package sublist implements a routine that checks
// the relation of two lists to each other
package sublist

type Relation string

// Sublist takes two integer lists and checks if
// the lists are equal, 1st is sublist of second,
// 1st is superlist of second, or unequal
func Sublist(listOne []int, listTwo []int) Relation {
    n1 := len(listOne)
    n2 := len(listTwo)
    relation := ""
    if n1 == n2 {
        relation = "equal"
        for i := 0; i < n1; i++ {
            if listOne[i] != listTwo[i] {
                relation = "unequal"
                break
            }
        }
    } else if n1 == 0 {
        relation = "sublist"
    } else if n2 == 0 {
        relation = "superlist"
    } else if n1 < n2 {
        relation = "unequal"
        for i := 0; i < n2-n1+1; i++ {
            if listOne[0] == listTwo[i] {
                found := true
                for j := 0; j < n1; j++ {
                    if listOne[j] != listTwo[j+i] {
                        found = false
                        break
                    }
                }
                if found {
                    relation = "sublist"
                    break
                }
            }
        }
    } else {
        relation = "unequal"
        for i := 0; i < n1-n2+1; i++ {
            if listTwo[0] == listOne[i] {
                found := true
                for j := 0; j < n2; j++ {
                    if listTwo[j] != listOne[j+i] {
                        found = false
                        break
                    }
                }
                if found {
                    relation = "superlist"
                    break
                }
            }
        }
    }
    return Relation(relation) 
}
