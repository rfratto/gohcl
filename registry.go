package gohcl

import (
	"fmt"
	"reflect"

	"github.com/zclconf/go-cty/cty"
)

var (
	registeredGoTypes = map[reflect.Type]cty.Type{}
	registeredTypes   = map[cty.Type]reflect.Type{}
)

// RegisterCapsuleType registers a capsule type for use with encoding/decoding
// cty and HCL values. RegisterCapsuleType must not be called with different
// capsule types that have the same underlying encapsulated type.
//
// RegisterCapsuleType panics if ty is not a capsule type.
func RegisterCapsuleType(ty cty.Type) {
	if !ty.IsCapsuleType() {
		panic("RegisterCapsuleType called with non capsule type " + ty.GoString())
	}

	if _, reg := registeredTypes[ty]; reg {
		// Ignore already-registered type
		return
	}
	if otherType, goTypeConsumed := registeredGoTypes[ty.EncapsulatedType()]; goTypeConsumed {
		panic(fmt.Sprintf("Go type %s is already registered with capsule type %s", ty.EncapsulatedType().String(), otherType.GoString()))
	}

	registeredTypes[ty] = ty.EncapsulatedType()
	registeredGoTypes[ty.EncapsulatedType()] = ty
}
