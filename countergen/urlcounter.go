package countergen

import (
	"fmt"

	"stash.ovh.net/enginebilling/tools/go/common/convert"
)

// URLCounter generates a random *string and holds it in state
// Retrieve the current value with Get() or set/get a new one with Next()
type URLCounter struct {
	strGen            StringCounter
	fnOnIncrementHook func(oldURL, newURL string)
}

func (gen *URLCounter) url() *string {
	return convert.StringToPtr(fmt.Sprintf("http://%s.com/%s", *gen.strGen.Same(), *gen.strGen.Same()))
}

// NewURLCounter creates a new URLCounter and sets the first *string state value
func NewURLCounter() *URLCounter {
	return &URLCounter{
		strGen: NewStringCounter(),
	}
}

// Same retrieves the current string held by the URLCounter
func (gen *URLCounter) Same() *string {
	return gen.url()
}

// Next sets a new random *string value and returns it
func (gen *URLCounter) Next() *string {
	oldURL := *gen.url()
	gen.strGen.Next()
	if gen.fnOnIncrementHook != nil {
		gen.fnOnIncrementHook(oldURL, *gen.url())
	}
	return gen.url()
}

// OnIncrement adds a hook function that runs when the URLCounter generates an URL
func (gen *URLCounter) OnIncrement(fnHook func(oldURL, newURL string)) *URLCounter {
	gen.fnOnIncrementHook = fnHook
	return gen
}
