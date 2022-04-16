package gohcl

import (
	"math/big"
	"reflect"

	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/set"
)

var valueType = reflect.TypeOf(cty.Value{})
var typeType = reflect.TypeOf(cty.Type{})

var setType = reflect.TypeOf(set.Set{})

var bigFloatType = reflect.TypeOf(big.Float{})
var bigIntType = reflect.TypeOf(big.Int{})

var emptyInterfaceType = reflect.TypeOf(interface{}(nil))

var stringType = reflect.TypeOf("")

// structTagIndices interrogates the fields of the given type (which must
// be a struct type, or we'll panic) and returns a map from the cty
// attribute names declared via struct tags to the indices of the
// fields holding those tags.
//
// This function will panic if two fields within the struct are tagged with
// the same cty attribute name.
func structTagIndices(st reflect.Type) map[string]int {
	var (
		ft = getFieldTags(st)

		ret = make(map[string]int, st.NumField())
	)

	// We only look through attributes and blocks: labels, body, and remain don't
	// map to a cty concept (and optionals fill the attributes set).
	for name, index := range ft.Attributes {
		ret[name] = index
	}
	for name, index := range ft.Blocks {
		ret[name] = index
	}

	return ret
}

// optionalFields returns optional fields from st.
func optionalFields(st reflect.Type) map[string]bool {
	var (
		ft = getFieldTags(st)

		ret = make(map[string]bool, st.NumField())
	)

	// We only look through attributes and blocks: labels, body, and remain don't
	// map to a cty concept (and optionals fill the attributes set).
	for name := range ft.Attributes {
		ret[name] = ft.Optional[name]
	}
	for name := range ft.Blocks {
		ret[name] = ft.Optional[name]
	}

	return ret
}
