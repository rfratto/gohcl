package gohcl

import (
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

type tokensExtension struct{}

// CapsuleTokenExtensionKey is the key used for capsule types which can be
// encoded to HCL. If a capsule type implements this extension with a value of
// CapsuleTokenExtension, the result of the function call will be used to set
// tokens for an attribute.
//
// If a capsule type does *not* implement this extension, encoding will panic.
var CapsuleTokenExtensionKey tokensExtension

// CapsuleTokenExtension is a function used by attributes of capsule types
// which can be encoded as HCL.
type CapsuleTokenExtension func(cty.Value) hclwrite.Tokens
