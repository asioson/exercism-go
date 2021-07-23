// package tree implements routines to build a tree structure of records
package tree

import "sort"
import "fmt"

// structures required in the problem
type Record struct {
    ID int
    Parent int
}

type Node struct {
    ID       int
    Children []*Node
}

// Build takes in a record containing information about the parent of
// each node mentioned. The function returns a tree structure given
// the information contained in the records.
func Build(records []Record) (*Node, error) {
    sort.Slice(records, func(i, j int) bool {
        if records[i].ID == records[j].ID {
            return records[i].Parent < records[j].Parent
        }
        return records[i].ID < records[j].ID
    })
    node := map[int]*Node{}
    for idx, r := range records {
        if r.ID != idx {
            return nil, fmt.Errorf("Build: missing node")
        }
        if (r.ID <= r.Parent) && (r.Parent != 0) {
            return nil, fmt.Errorf("Build: parent ID cannot be greater than child's")
        }
        node[r.ID] = &Node{ID: r.ID}
        if r.ID != 0 {
            node[r.Parent].Children = append(node[r.Parent].Children, node[r.ID])
        }
    }
    return node[0], nil
}
