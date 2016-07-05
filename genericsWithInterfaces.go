// genericsWithInterfaces demonstrates how interfaces can be used to achieve generic-like behavior.
package main

import "fmt"

// Comparable is our 'generic' type
type Comparable interface {
	CompareTo(Comparable) int
	Metric() int
}

// SortedList is a 'generic' list that holds anything that satisfies the Comparable interface
type SortedList []Comparable

func (s *SortedList) Sort() {
	// TODO
}

// DeepSkyObject is a type that implements Comparable.
type DeepSkyObject struct {
	Name     string
	Distance int
}

// CompareTo compares two DeepSkyObjects by distance.
func (d1 DeepSkyObject) CompareTo(d2 Comparable) int {
	if d1.Distance < d2.Metric() {
		return -1
	} else if d1.Distance == d2.Metric() {
		return 0
	}
	return 1
}

// Metric() returns a value for comparison
func (d DeepSkyObject) Metric() int {
	return d.Distance
}

func main() {
	a := DeepSkyObject{"Hyades", 46}
	b := DeepSkyObject{"Andromeda", 778000}
	c := SortedList{a, b}
	fmt.Printf("%s compared to %s: %d\n", a.Name, b.Name, c[0].CompareTo(c[1]))
}
