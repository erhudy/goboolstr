package goboolstr

import (
	"encoding/json"
	"fmt"
	"strings"
)

// +protobuf=true
// +protobuf.options.(gogoproto.goproto_stringer)=false
type BoolOrString struct {
	Type    Type   `protobuf:"bool,1,opt,name=type,casttype=Type"`
	BoolVal bool   `protobuf:"bool,2,opt,name=boolVal"`
	StrVal  string `protobuf:"string,3,opt,name=strVal"`
}

type Type bool

func True() BoolOrString {
	return FromBool(true)
}

func False() BoolOrString {
	return FromBool(false)
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

func (b *BoolOrString) AsBool() bool {
	return b.BoolVal
}

func (b *BoolOrString) FromBool(from bool) {
	b.BoolVal = from
	b.StrVal = fmt.Sprintf("%v", from)
}

func FromString(from string) BoolOrString {
	b := BoolOrString{}
	b.FromString(from)
	return b
}

func (b *BoolOrString) AsString() string {
	return b.StrVal
}

func (b *BoolOrString) FromString(from string) {
	b.BoolVal = isTruthy(from)
	b.StrVal = from
}

// will always marshal as a native bool (since the whole purpose
// of the type is to get a string to become a bool)
func (b BoolOrString) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%v", b.BoolVal)), nil
}

func (b *BoolOrString) UnmarshalJSON(in []byte) error {
	trimmed := strings.Trim(string(in), "\"")
	b.BoolVal = isTruthy(trimmed)
	b.StrVal = trimmed
	return nil
}

func (b BoolOrString) String() string {
	enc, _ := json.Marshal(b)
	return string(enc)
}
