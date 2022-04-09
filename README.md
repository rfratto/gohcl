# gohcl-capsule

`gohcl-capsule` is a fork of github.com/hashicorp/hcl/v2/gohcl and
github.com/zclconf/go-cty/cty/gocty which includes support for capsule types
through a registration system.

It makes the following changes:

* nil slices will return an empty go-cty list or go-cty set instead of null.
