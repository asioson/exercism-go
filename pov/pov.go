// package pov implements routines on directed graphs
package pov

import "fmt"

// Graph is an adjaceny list
type Graph map[string][]string

// New creates a new Graph
func New() *Graph {
	g := make(Graph)
	return &g
}

// AddNode adds a new node in Graph g
func (g *Graph) AddNode(nodeLabel string) {
	if _, ok := (*g)[nodeLabel]; !ok {
		(*g)[nodeLabel] = []string{}
	}
}

// AddArc adds a new arc in Graph g
func (g *Graph) AddArc(from, to string) {
	if _, ok := (*g)[from]; !ok {
		(*g)[from] = []string{to}
	} else {
		(*g)[from] = append((*g)[from], to)
	}
}

// ArcList generates a list of arcs
func (g *Graph) ArcList() []string {
	res := []string{}
	for from, adjNodes := range *g {
		for _, to := range adjNodes {
			res = append(res, fmt.Sprintf("%v -> %v", from, to))
		}
	}
	return res
}

// ChangeRoot takes a Graph g and generates a new Graph
// having newRoot as new root
func (g *Graph) ChangeRoot(oldRoot, newRoot string) *Graph {
	ug := make(Graph)
	ng := make(Graph)
	visited := map[string]bool{}
	for u, vList := range *g {
		for _, v := range vList {
			ug.AddArc(u, v)
			ug.AddArc(v, u)
		}
		visited[u] = false
	}
	q := []string{newRoot}
	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		if !visited[u] {
			visited[u] = true
			ng.AddNode(u)
		}
		for _, v := range ug[u] {
			if !visited[v] {
				q = append(q, v)
				ng.AddArc(u, v)
			}
		}
	}
	return &ng
}
