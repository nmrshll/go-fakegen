package random

import (
	"fmt"

	"stash.ovh.net/enginebilling/tools/go/common/convert"
)

func randomURL() *string {
	return convert.StringToPtr(fmt.Sprintf("http://%s.com/%s", RandStringBytes(8), RandStringBytes(8)))
}

// URLGenerator generates a random *string and holds it in state
// Retrieve the current value with Get() or set/get a new one with Next()
type URLGenerator struct {
	url              *string
	fnOnGenerateHook func(oldURL, newURL string)
}

// NewURLGenerator creates a new URLGenerator and sets the first *string state value
func NewURLGenerator() *URLGenerator {
	return &URLGenerator{randomURL(), nil}
}

// Same retrieves the current string held by the URLGenerator
func (gen *URLGenerator) Same() *string {
	return gen.url
}

// Next sets a new random *string value and returns it
func (gen *URLGenerator) Next() *string {
	oldURL := *gen.url
	gen.url = randomURL()
	if gen.fnOnGenerateHook != nil {
		gen.fnOnGenerateHook(oldURL, *gen.url)
	}
	return gen.url
}

// OnGenerate adds a hook function that runs when the URLGenerator generates an URL
func (gen *URLGenerator) OnGenerate(fnHook func(oldURL, newURL string)) *URLGenerator {
	gen.fnOnGenerateHook = fnHook
	return gen
}
