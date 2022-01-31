// package dominoes implements routines to generate a Domino chain given
// a set of Dominoes
package dominoes

type Domino [2]int

type Cand struct {
	pos, link, size int
}

// NextCandidate generates information about the next possible Dominoes given
// partial solution
func NextCandidates(input []Domino, soln []Cand, k int) []Cand {
	cand := []Cand{}
	for i, d := range input {
		found := false
		for _, v := range soln {
			if v.pos == i {
				found = true
				break
			}
		}
		if !found {
			last := input[soln[k-1].pos]
			if last[soln[k-1].link] == d[0] {
				cand = append(cand, Cand{pos: i, link: 1})
			} else if last[soln[k-1].link] == d[1] {
				cand = append(cand, Cand{pos: i, link: 0})
			}
		}
	}
	return cand
}

// Solve determines an Eulerian path given the slice of Dominoes
// using backtracking approach
func Solve(input []Domino) ([]Cand, bool) {
	n := len(input)
	soln := []Cand{}
	stack := []Cand{{pos: 0, link: 1, size: 1}}
	for len(stack) > 0 {
		item := stack[0]
		stack = stack[1:]
		k := item.size
		soln = append(soln[:k-1], item)
		if k == n {
			return soln, true
		} else {
			for _, c := range NextCandidates(input, soln, k) {
				stack = append([]Cand{{pos: c.pos, link: c.link, size: k + 1}}, stack...)
			}
		}
	}
	return soln, false
}

// MakeChain produces a Domino chain (Euler circuit) when possible
func MakeChain(input []Domino) ([]Domino, bool) {
	chain := []Domino{}
	n := len(input)
	if n == 0 {
		return chain, true
	}
	if sequence, ok := Solve(input); ok {
		for _, s := range sequence {
			chain = append(chain, Domino{input[s.pos][1-s.link], input[s.pos][s.link]})
		}
		if chain[0][0] == chain[n-1][1] {
			return chain, true
		}
	}
	return chain, false
}
