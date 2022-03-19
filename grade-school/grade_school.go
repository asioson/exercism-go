// package school implements routines to manage the roster of students
package school

import (
	"sort"
)

type Grade struct {
	level int      // grade level
	names []string // name of students
}

type School struct {
	roster map[int]*Grade
}

// New returns a new School instance
func New() *School {
	return &School{roster: map[int]*Grade{}}
}

// Add inserts student name in specified grade level
func (s *School) Add(student string, grade int) {
	if list, ok := s.roster[grade]; !ok {
		s.roster[grade] = &Grade{level: grade, names: []string{student}}
	} else {
		list.names = append(list.names, student)
	}
}

// Grade returns a sorted list of students in a grade level
func (s *School) Grade(level int) []string {
	if list, ok := s.roster[level]; ok {
		sort.Strings(list.names)
		return list.names
	} else {
		return []string{}
	}
}

// Enrollment returns a roster of students in a grade level
func (s *School) Enrollment() []Grade {
	levels := make([]int, len(s.roster))
	i := 0
	for l := range s.roster {
		levels[i] = l
		i++
	}
	sort.Ints(levels)
	enrollment := []Grade{}
	for _, l := range levels {
		enrollment = append(enrollment, Grade{level: l, names: s.Grade(l)})
	}
	return enrollment
}
