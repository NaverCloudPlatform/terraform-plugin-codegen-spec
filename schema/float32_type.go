// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

// Float32Type is a representation of a 32-bit floating point number.
type Float32Type struct {
	// CustomType is a customization of the Float32Type.
	CustomType *CustomType `json:"custom_type,omitempty"`
}
