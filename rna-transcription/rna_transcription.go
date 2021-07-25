// package strand implements a routine that transcribe a dna sequence
// to rna sequence
package strand

// ToRNA takes in a dna sequence and returns the corresponding
// rna sequence
func ToRNA(dna string) string {
    rnaMap := map[rune]string { 'G':"C", 'C':"G", 'T':"A", 'A':"U" }
    rna := ""
    for _, nucleotide := range dna {
        rna += rnaMap[nucleotide]
    }
    return rna
}
