package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// this is brought out into an interface so during
// the unit tests we can replace the random things with
// something much less random
type random interface {
	value(prefix string) string
	suffix() string
}

type randomGenerator struct {
	source rand.Source
	rand   *rand.Rand
}

var rgInstance random

// standard initialization which for runtime is going to be
// fine
func init() {
	rg := randomGenerator{}
	rg.source = rand.NewSource(time.Now().UnixNano())
	rg.rand = rand.New(rg.source)

	rgInstance = rg
}

// a uniform way to create a field value that starts with the
// prefix
func (rg randomGenerator) value(prefix string) string {
	return fmt.Sprintf("%s-%s", prefix, rg.suffix())
}

// a uniform way to create a random suffix
// this will allow us to mock this number as well
// as update it in one place if we change our minds about
// the length and format.
func (rg randomGenerator) suffix() string {
	result := rg.rand.Intn(8999999)
	result += 1000000
	return strconv.Itoa(result)
}
