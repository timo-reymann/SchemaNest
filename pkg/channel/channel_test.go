package channel

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		expected  *Channel
		expectErr bool
	}{
		{
			name:      "Valid channel with numbers",
			input:     "1.2",
			expected:  &Channel{Major: "1", Minor: "2"},
			expectErr: false,
		},
		{
			name:      "Invalid channel with 'x' as major",
			input:     "x.2",
			expected:  nil,
			expectErr: true,
		},
		{
			name:      "Valid channel with 'x' as minor",
			input:     "1.x",
			expected:  &Channel{Major: "1", Minor: ""},
			expectErr: false,
		},
		{
			name:      "Invalid channel with 'x' for both parts",
			input:     "x.x",
			expected:  nil,
			expectErr: true,
		},
		{
			name:      "Invalid channel format",
			input:     "1",
			expected:  nil,
			expectErr: true,
		},
		{
			name:      "Invalid major version",
			input:     "a.2",
			expected:  nil,
			expectErr: true,
		},
		{
			name:      "Invalid minor version",
			input:     "1.b",
			expected:  nil,
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Parse(tt.input)
			if (err != nil) != tt.expectErr {
				t.Errorf("Parse() error = %v, expectErr = %v", err, tt.expectErr)
				return
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Parse() = %v, expected = %v", result, tt.expected)
			}
		})
	}
}

func TestChannel_MajorVersion(t *testing.T) {
	tests := []struct {
		name     string
		channel  *Channel
		expected int
	}{
		{
			name:     "Valid major version",
			channel:  &Channel{Major: "1", Minor: "0"},
			expected: 1,
		},
		{
			name:     "Empty major version",
			channel:  &Channel{Major: "", Minor: "0"},
			expected: 0,
		},
		{
			name:     "Invalid major version defaults to 0",
			channel:  &Channel{Major: "x", Minor: "0"},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.channel.MajorVersion()
			if result != tt.expected {
				t.Errorf("MajorVersion() = %d, expected %d", result, tt.expected)
			}
		})
	}
}

func TestChannel_MinorVersion(t *testing.T) {
	tests := []struct {
		name     string
		channel  *Channel
		expected int
	}{
		{
			name:     "Valid minor version",
			channel:  &Channel{Major: "1", Minor: "2"},
			expected: 2,
		},
		{
			name:     "Empty minor version",
			channel:  &Channel{Major: "1", Minor: ""},
			expected: 0,
		},
		{
			name:     "Wildcard minor version",
			channel:  &Channel{Major: "1", Minor: "x"},
			expected: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.channel.MinorVersion()
			if result != tt.expected {
				t.Errorf("MinorVersion() = %d, expected %d", result, tt.expected)
			}
		})
	}
}
