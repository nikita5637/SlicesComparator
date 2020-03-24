package main

import (
	"errors"
	"math"
	"math/rand"
	"reflect"
)

var (
	ErrValueIsExists   = errors.New("Value is exists in the map")
	entropyCoefficient = 10
)

type Slice struct {
	name         string
	size         int
	freeIndexes  map[int]struct{}
	existsValues map[int]struct{}
	values       []int
}

func (s *Slice) init(size int) {
	logger.Debugf("[slice '%s'] Initializing slice structure with size '%d'", s.name, size)

	s.size = size

	s.freeIndexes = make(map[int]struct{})
	for i := 0; i < s.size; i++ {
		s.freeIndexes[i] = struct{}{}
	}

	s.existsValues = make(map[int]struct{})

	s.values = make([]int, size)

	logger.Debugf("[slice '%s'] Initializing done", s.name)
}

func (s *Slice) insertValue(v int) error {
	if _, ok := s.existsValues[v]; ok {
		logger.Debugf("[slice '%s'] %s", s.name, ErrValueIsExists)
		return ErrValueIsExists
	}

	mapKeys := reflect.ValueOf(s.freeIndexes).MapKeys()
	indexPosition := rand.Intn(len(mapKeys))
	indexValue := int(mapKeys[indexPosition].Int())

	s.values[indexValue] = v
	delete(s.freeIndexes, indexValue)

	s.existsValues[v] = struct{}{}

	logger.Debugf("[slice '%s'] Inserted value '%d' into index '%d'\n", s.name, v, indexValue)
	return nil
}

// Generate two slices with the same elements
// but random order.
func GenerateEqualSlices(size int) ([]int, []int) {
	if size <= 0 || size > math.MaxInt32 {
		logger.Fatalf("Slice size must be between '1' and '%d'", math.MaxInt32)
	}

	a := Slice{
		name: "a",
	}
	a.init(size)

	b := Slice{
		name: "b",
	}
	b.init(size)

	for i := 0; i < size; i++ {
		newRandomValue := rand.Intn(size * entropyCoefficient)
		logger.Debugf("Generating new random value '%d'\n", newRandomValue)
		for {
			if err := a.insertValue(newRandomValue); err == nil {
				break
			}
			logger.Debugf("Generating new random value '%d'\n", newRandomValue)
			newRandomValue = rand.Intn(size * entropyCoefficient)
		}

		b.insertValue(newRandomValue)
	}

	return a.values, b.values
}

// Generate two slices with the different elements.
func GenerateNotEqualSlices(size int) ([]int, []int) {
	if size <= 0 || size > math.MaxInt32 {
		logger.Fatalf("Slice size must be between '1' and '%d'", math.MaxInt32)
	}

	a := Slice{
		name: "a",
	}
	a.init(size)

	b := Slice{
		name: "b",
	}
	b.init(size)

	for i := 0; i < size; i++ {
		newRandomValueA := rand.Intn(size * entropyCoefficient)
		newRandomValueB := rand.Intn(size * entropyCoefficient)
		logger.Debugf("Generating new random values '%d' and '%d'\n", newRandomValueA, newRandomValueB)
		a.insertValue(newRandomValueA)
		b.insertValue(newRandomValueB)
	}

	return a.values, b.values
}
