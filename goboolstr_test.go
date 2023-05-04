package goboolstr

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

type testStruct struct {
	Val BoolOrString
}

func TestMarshalJSON(t *testing.T) {
	tests := []struct {
		description string
		expectErr   bool
		expected    []byte
		input       testStruct
	}{
		{
			description: "true from bool",
			expectErr:   false,
			expected:    []byte("{\"Val\":true}"),
			input:       testStruct{Val: FromBool(true)},
		},
		{
			description: "true from string",
			expectErr:   false,
			expected:    []byte("{\"Val\":true}"),
			input:       testStruct{Val: FromString("true")},
		},
		{
			description: "false from bool",
			expectErr:   false,
			expected:    []byte("{\"Val\":false}"),
			input:       testStruct{Val: FromBool(false)},
		},
		{
			description: "false from string",
			expectErr:   false,
			expected:    []byte("{\"Val\":false}"),
			input:       testStruct{Val: FromString("false")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			out, err := json.Marshal(tt.input)
			if tt.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.expected, out)
			}
		})
	}
}

func TestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		description string
		expectErr   bool
		expected    testStruct
		rawJSON     []byte
	}{
		{
			description: "true string",
			expectErr:   false,
			expected: testStruct{
				Val: BoolOrString{
					rawBool:   true,
					rawString: "true",
				},
			},
			rawJSON: []byte("{\"val\": \"true\"}"),
		},
		{
			description: "yes string",
			expectErr:   false,
			expected: testStruct{
				Val: BoolOrString{
					rawBool:   true,
					rawString: "yes",
				},
			},
			rawJSON: []byte("{\"val\": \"yes\"}"),
		},
		{
			description: "on string",
			expectErr:   false,
			expected: testStruct{
				Val: BoolOrString{
					rawBool:   true,
					rawString: "on",
				},
			},
			rawJSON: []byte("{\"val\": \"on\"}"),
		},
		{
			description: "1 string",
			expectErr:   false,
			expected: testStruct{
				Val: BoolOrString{
					rawBool:   true,
					rawString: "1",
				},
			},
			rawJSON: []byte("{\"val\": \"1\"}"),
		},
		{
			description: "true bool",
			expectErr:   false,
			expected: testStruct{
				Val: BoolOrString{
					rawBool:   true,
					rawString: "true",
				},
			},
			rawJSON: []byte("{\"val\": true}"),
		},
		{
			description: "false string",
			expectErr:   false,
			expected: testStruct{
				Val: BoolOrString{
					rawBool:   false,
					rawString: "false",
				},
			},
			rawJSON: []byte("{\"val\": \"false\"}"),
		},
		{
			description: "false bool",
			expectErr:   false,
			expected: testStruct{
				Val: BoolOrString{
					rawBool:   false,
					rawString: "false",
				},
			},
			rawJSON: []byte("{\"val\": false}"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			res := testStruct{}
			err := json.Unmarshal(tt.rawJSON, &res)

			if tt.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.expected, res)
			}
		})
	}
}
