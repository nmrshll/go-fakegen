package random

import (
	"stash.ovh.net/enginebilling/tools/go/common/convert"
)

// StringGenerator generates a random *string and holds it in state
// Retrieve the current value with Same() or set/get a new one with Next()
type StringGenerator struct {
	str *string
}

// NewStringGenerator creates a new StringGenerator and sets the first *string state value
func NewStringGenerator() StringGenerator {
	return StringGenerator{convert.StringToPtr(RandStringBytes(8))}
}

// Same retrieves the current string held by the StringGenerator
func (gen *StringGenerator) Same() *string {
	return gen.str
}

// Next sets a new random *string value and returns it
func (gen *StringGenerator) Next() *string {
	gen.str = convert.StringToPtr(RandStringBytes(8))
	return gen.str
}
