// package protein implements routines that translates a codon to protein 
// and an RNA sequence to sequence of proteins
package protein

import "fmt"

var ErrStop = fmt.Errorf("Stop")
var ErrInvalidBase = fmt.Errorf("Invalid Base")

// FromCodon translates a codon to protein
func FromCodon(seq string) (string, error) {
    protein := map[string]string{
    "AUG":"Methionine",
    "UUU":"Phenylalanine",
    "UUC":"Phenylalanine",
    "UUA":"Leucine",
    "UUG":"Leucine",
    "UCU":"Serine",
    "UCC":"Serine",
    "UCA":"Serine",
    "UCG":"Serine",
    "UAU":"Tyrosine",
    "UAC":"Tyrosine",
    "UGU":"Cysteine",
    "UGC":"Cysteine",
    "UGG":"Tryptophan",
    "UAA":"STOP",
    "UAG":"STOP",
    "UGA":"STOP",
    }
    if result, ok := protein[seq]; ok {
        if result == "STOP" {
            return "", ErrStop
        } else {
            return result, nil
        }
    }
    return "", ErrInvalidBase
}

// FromRNA translates an RNA sequence to protein sequence
func FromRNA(seq string) ([]string, error) {
    proteinSeq := []string{}
    n := len(seq)
    for i := 0; i < n; i += 3 {
        protein, err := FromCodon(seq[i:i+3])
        if err == ErrStop { break }
        if err != nil {
            return proteinSeq, err
        }
        proteinSeq = append(proteinSeq, protein)
    }
    return proteinSeq, nil
}
