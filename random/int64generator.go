package random

import (
	"math/rand"

	"stash.ovh.net/enginebilling/tools/go/common/convert"
)

// Int64Generator generates a random *int64 and holds it in state
// Retrieve the current value with Same() or set/get a new one with Next()
type Int64Generator struct {
	i64              *int64
	fnOnGenerateHook func(oldURL, newURL int64)
}

// NewInt64Generator creates a new Int64Generator and sets the first *int64 state value
func NewInt64Generator() *Int64Generator {
	return &Int64Generator{convert.Int64ToPtr(rand.Int63()), nil}
}

// Same retrieves the current int64 held by the Int64Generator
func (gen *Int64Generator) Same() *int64 {
	return gen.i64
}

// Next sets a new random *int64 value and returns it
func (gen *Int64Generator) Next() *int64 {
	oldI64 := *gen.i64
	gen.i64 = convert.Int64ToPtr(rand.Int63())
	if gen.fnOnGenerateHook != nil {
		gen.fnOnGenerateHook(oldI64, *gen.i64)
	}
	return gen.i64
}

// OnGenerate adds a hook function that runs when the Int64Generator generates an URL
func (gen *Int64Generator) OnGenerate(fnHook func(oldURL, newURL int64)) *Int64Generator {
	gen.fnOnGenerateHook = fnHook
	return gen
}
