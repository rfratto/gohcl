# gohcl

`gohcl` is a fork of github.com/hashicorp/hcl/v2/gohcl and
github.com/zclconf/go-cty/cty/gocty which includes improvements over the
existing package at the expense of making backwards-incompatible changes to the
API.

It makes the following changes:

* Capsule types can be used through a registration system.
* nil slices will return an empty go-cty list or go-cty set instead of null.
* It is valid for capsule values to be nil.
* `hcl` struct tags are used globally instead of a combination of `hcl` and
  `cty` tags.
* Encoding a value into a HCL body will no longer clear the existing contents
  of the body.
* Capsule types can be encoded as HCL attributes through the use of capsule
  extension operations; allowing a capsule type to return raw HCL tokens which
  represent the capsule value.
* Supports Go types like time.Duration
