package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var source = rand.NewSource(time.Now().UnixNano())
var random = rand.New(source)

// randomString will create a string that starts
// wth the prefix and ends with a six digit random number
func randomValue(prefix string) string {
	return fmt.Sprintf("%s-%s", prefix, randomNumber())
}

// generate a random number and convert it to a string
// this is done in one function so we have one place
// to look in case we need to change the format of the
// number
func randomNumber() string {
	result := random.Intn(8999999)
	result += 1000000
	return strconv.Itoa(result)
}
