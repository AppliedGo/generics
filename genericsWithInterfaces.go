// genericsWithInterfaces demonstrates how interfaces can be used to achieve generic-like behavior.
package main

import "fmt"

// Comparable is our 'generic' type
type Comparable interface {
	CompareTo(Comparable) int
	Metric() int
}

// CString is a comparable string.
type CString string

// CompareTo compares the lengths of s and c.
func (s CString) CompareTo(c Comparable) int {
	if s.Metric() < c.Metric() {
		return -1
	} else if s.Metric() == c.Metric() {
		return 0
	}
	return 1
}

// Metric returns the length of the string.
func (s CString) Metric() int {
	return len(s)
}

// DeepSkyObject also iplements Comparable.
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
	fmt.Printf("%s compared to %s: %d\n", a.Name, b.Name, a.CompareTo(b))
}
