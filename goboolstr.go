package goboolstr

import (
	"fmt"
	"strings"
)

type BoolOrString struct {
	rawBool   bool
	rawString string
}

func isTruthy(val string) bool {
	lc := strings.ToLower(val)
	return lc == "true" || lc == "yes" || lc == "on" || lc == "1"
}

func FromBool(from bool) BoolOrString {
	b := BoolOrString{}
	b.FromBool(from)
	return b
}

func (b *BoolOrString) FromBool(from bool) {
	b.rawBool = from
	b.rawString = fmt.Sprintf("%v", from)
}

func FromString(from string) BoolOrString {
	b := BoolOrString{}
	b.FromString(from)
	return b
}

func (b *BoolOrString) FromString(from string) {
	b.rawBool = isTruthy(from)
	b.rawString = from
}

// will always marshal as a native bool (since the whole purpose
// of the type is to get a string to become a bool)
func (b BoolOrString) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%v", b.rawBool)), nil
}

func (b *BoolOrString) UnmarshalJSON(in []byte) error {
	trimmed := strings.Trim(string(in), "\"")
	b.rawBool = isTruthy(trimmed)
	b.rawString = trimmed
	return nil
}
