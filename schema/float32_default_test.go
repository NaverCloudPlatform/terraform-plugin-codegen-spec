// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/NaverCloudPlatform/terraform-plugin-codegen-spec/schema"
)

func TestFloat32Default_Equal(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		float32Default *schema.Float32Default
		other          *schema.Float32Default
		expected       bool
	}{
		"both_nil": {
			expected: true,
		},
		"float32_default_nil_other_not_nil": {
			other:    &schema.Float32Default{},
			expected: false,
		},
		"float32_default_static_nil_other_not_nil": {
			float32Default: &schema.Float32Default{},
			other: &schema.Float32Default{
				Static: pointer(float32(1.234)),
			},
			expected: false,
		},
		"float32_default_static_not_nil_other_nil": {
			float32Default: &schema.Float32Default{
				Static: pointer(float32(1.234)),
			},
			other:    &schema.Float32Default{},
			expected: false,
		},
		"match": {
			float32Default: &schema.Float32Default{
				Static: pointer(float32(1.234)),
			},
			other: &schema.Float32Default{
				Static: pointer(float32(1.234)),
			},
			expected: true,
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := testCase.float32Default.Equal(testCase.other)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
