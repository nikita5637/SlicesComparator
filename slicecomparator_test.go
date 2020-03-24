package main

import (
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

const (
	size = 1024 // slices size for testing
)

func TestMain(m *testing.M) {
	logger = logrus.New()

	loglevel, err := logrus.ParseLevel("warning")
	if err != nil {
		panic(err)
	}
	logger.SetLevel(loglevel)

	os.Exit(m.Run())
}

func Test_equals(t *testing.T) {
	equalA, equalB := GenerateEqualSlices(size)
	notEqualA, notEqualB := GenerateNotEqualSlices(size)

	testCases := []struct {
		name       string
		a, b       []int
		comparator func([]int, []int) bool
		wanted     bool
	}{
		{
			name:       "Equal slices compare with sort",
			comparator: equalElementsWithSort,
			a:          equalA,
			b:          equalB,
			wanted:     true,
		},
		{
			name:       "Equal slices compare",
			comparator: equalElements,
			a:          equalA,
			b:          equalB,
			wanted:     true,
		},
		{
			name:       "Not equal slices compare with sort",
			comparator: equalElementsWithSort,
			a:          notEqualA,
			b:          notEqualB,
			wanted:     false,
		},
		{
			name:       "Not equal slices compare",
			comparator: equalElements,
			a:          notEqualA,
			b:          notEqualB,
			wanted:     false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			a := append(tc.a[:0:0], tc.a...)
			b := append(tc.b[:0:0], tc.b...)

			logger.Infof("Compare the following slices\n\t\ta=%d\n\t\tb=%d\n", a, b)

			ts := time.Now()
			isEq := tc.comparator(a, b)
			logger.Infof("Completed in %d\n", time.Now().Sub(ts))

			assert.Equal(t, isEq, tc.wanted)
		})
	}
}
