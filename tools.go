package toolkit

import (
	"math/rand"
	"strings"
)

const randomStringSource = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_+"

// Tools is the type use to instantiate this module. Any variable of this type will have access
// to all the methods with the reciever *Tools
type Tools struct{}

// RandomString returns a string of random characters of length n, using randomStringSource
// as the source for the string
func (t *Tools) RandomString(n int) string {
	stringBuilder := strings.Builder{}

	stringBuilder.Grow(n)

	for i := 0; i < n; i++ {
		stringBuilder.WriteByte(randomStringSource[rand.Intn(len(randomStringSource))])
	}

	return stringBuilder.String()
}
