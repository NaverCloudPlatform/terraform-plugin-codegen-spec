// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

// Float32Default defines a value, or a custom type for a default 32-bit floating point value.
type Float32Default struct {
	// Custom defines a schema definition, and optional imports.
	Custom *CustomDefault `json:"custom,omitempty"`

	// Static defines a specific 32-bit floating point value.
	Static *float32 `json:"static,omitempty"`
}

// Equal returns true if all fields of the given Float32Default are equal.
func (d *Float32Default) Equal(other *Float32Default) bool {
	if d == nil && other == nil {
		return true
	}

	if d == nil || other == nil {
		return false
	}

	if !d.Custom.Equal(other.Custom) {
		return false
	}

	if d.Static == nil && other.Static != nil {
		return false
	}

	if d.Static != nil && other.Static == nil {
		return false
	}

	return *d.Static == *other.Static
}
