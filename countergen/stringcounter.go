package countergen

import (
	"strconv"

	"stash.ovh.net/enginebilling/tools/go/common/convert"
)

// StringCounter generates a random *string and holds it in state
// Retrieve the current value with Same() or set/get a new one with Next()
type StringCounter struct {
	count uint64
}

// NewStringCounter creates a new StringCounter and sets the first *string state value
func NewStringCounter() StringCounter {
	return StringCounter{
		count: uint64(1),
	}
}

func (gen *StringCounter) str() *string {
	return convert.StringToPtr("somestring_" + strconv.FormatUint(gen.count, 10))
}

// Same retrieves the current string held by the StringCounter
func (gen *StringCounter) Same() *string {
	return gen.str()
}

// Next sets a new random *string value and returns it
func (gen *StringCounter) Next() *string {
	gen.count++
	return gen.str()
}
