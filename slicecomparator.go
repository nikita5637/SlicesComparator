// Equivalence slice testing package
//	Conditions:
//		Equivalent slice length
//		The order of the elements in slices is unknown
//		Elements are not repeated in each individual slice
package main

import (
	"github.com/sirupsen/logrus"
	"reflect"
	"sort"
)

const (
	loglevel = "info"
)

var (
	logger *logrus.Logger
)

func equalElements(a, b []int) bool {
	// difference between sums of slices
	totalDiff := 0
	for i := range a {
		totalDiff += a[i]
		totalDiff -= b[i]
	}

	if totalDiff != 0 {
		return false
	}

	for _, valA := range a {
		founded := false
		for _, valB := range b {
			if valA == valB {
				founded = true
				break
			}
		}
		if !founded {
			return false
		}
	}

	return true
}

func equalElementsWithSort(a, b []int) bool {
	// difference between sums of slices
	totalDiff := 0
	for i := range a {
		totalDiff += a[i]
		totalDiff -= b[i]
	}

	if totalDiff != 0 {
		return false
	}

	sort.Ints(a)
	sort.Ints(b)

	return reflect.DeepEqual(a, b)
}

func main() {
	logger = logrus.New()
	level, err := logrus.ParseLevel(loglevel)
	if err != nil {
		panic(err)
	}

	logger.SetLevel(level)

	a, b := GenerateEqualSlices(128)
	logger.Infof("Compare the following slices\n\t\ta=%d\n\t\tb=%d\n", a, b)
	is_eq := equalElementsWithSort(a, b)

	logger.Infof("A == B: %t\n", is_eq)
}
