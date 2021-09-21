// Package chessboard implements routines for a chess game
package chessboard

// Declare a type named Rank which stores if a square is occupied by a piece - this will be a slice of bools
type Rank []bool

// Declare a type named Chessboard contains a map of eight Ranks, accessed with values from 1 to 8
type Chessboard map[int]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank int) (ret int) {
    if rank < 1 || rank > 8 {
        return 0
    }
    ret = 0
    for _, status := range cb[rank] {
        if status {
            ret++
        }
    }
    return ret
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) (ret int) {
    if file < 1 || file > 8 {
        return 0
    }
    ret = 0
    for _, rank := range cb {
        if rank[file-1] {
            ret++
        }
    }
    return ret
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) (ret int) {
    ret = 0
    for _, rank := range cb {
        ret += len(rank)
    }
    return ret
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) (ret int) {
    ret = 0
    for _, rank := range cb {
        for _, occupied := range rank {
            if occupied {
                ret++
            }
        }
    }
    return ret

}
