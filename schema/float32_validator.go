// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

// Float32Validators type defines Float32Validator types
type Float32Validators []Float32Validator

// CustomValidators returns CustomValidator for each Float32Validator.
func (v Float32Validators) CustomValidators() CustomValidators {
	var customValidators CustomValidators

	for _, validator := range v {
		customValidator := validator.Custom

		if customValidator == nil {
			continue
		}

		customValidators = append(customValidators, customValidator)
	}

	return customValidators
}

// Equal returns true if the given Float32Validators is the same
// length, and each of the Float32Validator entries is equal.
func (v Float32Validators) Equal(other Float32Validators) bool {
	if v == nil && other == nil {
		return true
	}

	if v == nil || other == nil {
		return false
	}

	if len(v) != len(other) {
		return false
	}

	validators := v.CustomValidators()

	otherValidators := other.CustomValidators()

	if len(validators) != len(otherValidators) {
		return false
	}

	validators.Sort()

	otherValidators.Sort()

	for k, validator := range validators {
		if !validator.Equal(otherValidators[k]) {
			return false
		}
	}

	return true
}

// Float32Validator type defines type and function that provides validation
// functionality.
type Float32Validator struct {
	Custom *CustomValidator `json:"custom,omitempty"`
}

// Equal returns true if the fields of the given Float32Validator equal.
func (v Float32Validator) Equal(other Float32Validator) bool {
	return v.Custom.Equal(other.Custom)
}
