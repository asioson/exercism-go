// package stringset implements routines on sets with strings as elements
package stringset

// Implement Set as a collection of unique string values.
//
// For Set.String, use '{' and '}', output elements as double-quoted strings
// safely escaped with Go syntax, and use a comma and a single space between
// elements. For example, a set with 2 elements, "a" and "b", should be formatted as {"a", "b"}.
// Format the empty set as {}.

type Set struct {
	items *[]string
	idx   map[string]bool
}

// New creates a new Set object
func New() Set {
	s := new(Set)
	s.items = new([]string)
	s.idx = make(map[string]bool, 0)
	return *s
}

// NewFromSlice creates a new Set object with elements coming from the slice of strings
func NewFromSlice(l []string) Set {
	s := New()
	for _, u := range l {
		s.Add(u)
	}
	return s
}

// String returns the elements of the set in string format
func (s Set) String() string {
	outStr := "{"
	if len([]string(*s.items)) > 0 {
		outStr += "\"" + (*s.items)[0] + "\""
	}
	for i := 1; i < len((*s.items)); i++ {
		outStr += ", \"" + (*s.items)[i] + "\""
	}
	outStr += "}"
	return outStr
}

// IsEmpty return true if the set is empty, false otherwise
func (s Set) IsEmpty() bool {
	return len(*s.items) == 0
}

// Has returns true if elem is a member of the set, false otherwise
func (s Set) Has(elem string) bool {
	_, ok := s.idx[elem]
	return ok
}

// Add inserts a new element elem in the set
func (s Set) Add(elem string) {
	if !s.Has(elem) {
		*(s.items) = append(*(s.items), elem)
		s.idx[elem] = true
	}
}

// Subset checks if s1 is a subset of s2
func Subset(s1, s2 Set) bool {
	for _, u := range *(s1.items) {
		if !s2.Has(u) {
			return false
		}
	}
	return true
}

// Disjoint checks if s1 and s2 are disjoint (i.e., no common element)
func Disjoint(s1, s2 Set) bool {
	return Intersection(s1, s2).IsEmpty()
}

// Equal checks if s1 and s2 has the same elements
func Equal(s1, s2 Set) bool {
	return Subset(s1, s2) && Subset(s2, s1)
}

// Intersection returns a set of elements common to s1 and s2
func Intersection(s1, s2 Set) Set {
	s3 := New()
	for _, u := range *(s1.items) {
		if s2.Has(u) {
			s3.Add(u)
		}
	}
	return s3
}

// Difference returns a set of elements in s1 not found in s2
func Difference(s1, s2 Set) Set {
	s3 := New()
	for _, u := range *(s1.items) {
		if !s2.Has(u) {
			s3.Add(u)
		}
	}
	return s3
}

// Union returns a set of elements from s1 and s2
func Union(s1, s2 Set) Set {
	s3 := New()
	for _, u := range *(s1.items) {
		s3.Add(u)
	}
	for _, u := range *(s2.items) {
		s3.Add(u)
	}
	return s3
}
