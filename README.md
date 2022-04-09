# gohcl

`gohcl` is a fork of github.com/hashicorp/hcl/v2/gohcl and
github.com/zclconf/go-cty/cty/gocty which includes improvements over the
existing package at the expense of making backwards-incompatible changes to the
API.

It makes the following changes:

* Capsule types can be used through a registration system.
* nil slices will return an empty go-cty list or go-cty set instead of null.
* It is valid for capsule values to be nil.
