// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

type ListPlanModifier struct {
	Custom *CustomPlanModifier `json:"custom,omitempty"`
}
