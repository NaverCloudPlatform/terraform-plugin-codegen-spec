// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

// Float32PlanModifiers type defines Float32PlanModifier types
type Float32PlanModifiers []Float32PlanModifier

// CustomPlanModifiers returns CustomPlanModifier for each Float32PlanModifier.
func (v Float32PlanModifiers) CustomPlanModifiers() CustomPlanModifiers {
	var customPlanModifiers CustomPlanModifiers

	for _, planModifier := range v {
		customPlanModifier := planModifier.Custom

		if customPlanModifier == nil {
			continue
		}

		customPlanModifiers = append(customPlanModifiers, customPlanModifier)
	}

	return customPlanModifiers
}

// Equal returns true if the given Float32PlanModifiers is the same
// length, and each of the Float32PlanModifier entries is equal.
func (v Float32PlanModifiers) Equal(other Float32PlanModifiers) bool {
	if v == nil && other == nil {
		return true
	}

	if v == nil || other == nil {
		return false
	}

	if len(v) != len(other) {
		return false
	}

	planModifiers := v.CustomPlanModifiers()

	otherPlanModifiers := other.CustomPlanModifiers()

	if len(planModifiers) != len(otherPlanModifiers) {
		return false
	}

	planModifiers.Sort()

	otherPlanModifiers.Sort()

	for k, planModifier := range planModifiers {
		if !planModifier.Equal(otherPlanModifiers[k]) {
			return false
		}
	}

	return true
}

// Float32PlanModifier type defines type and function that provides plan modification
// functionality.
type Float32PlanModifier struct {
	Custom *CustomPlanModifier `json:"custom,omitempty"`
}

// Equal returns true if the fields of the given Float32PlanModifier are equal.
func (v Float32PlanModifier) Equal(other Float32PlanModifier) bool {
	return v.Custom.Equal(other.Custom)

}
