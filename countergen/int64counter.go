package countergen

import (
	"stash.ovh.net/enginebilling/tools/go/common/convert"
)

// Int64Counter generates a random *int64 and holds it in state
// Retrieve the current value with Same() or set/get a new one with Next()
type Int64Counter struct {
	i64               *int64
	fnOnIncrementHook func(oldURL, newURL int64)
}

// NewInt64Counter creates a new Int64Counter and sets the first *int64 state value
func NewInt64Counter() *Int64Counter {
	return &Int64Counter{convert.Int64ToPtr(1), nil}
}

// Same retrieves the current int64 held by the Int64Counter
func (gen *Int64Counter) Same() *int64 {
	return gen.i64
}

// Next sets a new random *int64 value and returns it
func (gen *Int64Counter) Next() *int64 {
	oldI64 := *gen.i64
	gen.i64 = convert.Int64ToPtr(*gen.i64 + 1)
	if gen.fnOnIncrementHook != nil {
		gen.fnOnIncrementHook(oldI64, *gen.i64)
	}
	return gen.i64
}

// OnIncrement adds a hook function that runs when the Int64Counter generates an URL
func (gen *Int64Counter) OnIncrement(fnHook func(oldURL, newURL int64)) *Int64Counter {
	gen.fnOnIncrementHook = fnHook
	return gen
}
