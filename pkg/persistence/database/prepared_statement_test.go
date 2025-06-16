package database

import "testing"

func TestReplacePlaceholders(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "no placeholders",
			input:    "SELECT * FROM users",
			expected: "SELECT * FROM users",
		},
		{
			name:     "single placeholder",
			input:    "SELECT * FROM users WHERE id = ?",
			expected: "SELECT * FROM users WHERE id = $1",
		},
		{
			name:     "multiple placeholders",
			input:    "SELECT * FROM users WHERE name = ? AND age > ?",
			expected: "SELECT * FROM users WHERE name = $1 AND age > $2",
		},
		{
			name:     "many placeholders",
			input:    "INSERT INTO users (a,b,c,d,e,f,g,h,i,j) VALUES (?,?,?,?,?,?,?,?,?,?)",
			expected: "INSERT INTO users (a,b,c,d,e,f,g,h,i,j) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)",
		},
		{
			name:     "question mark in string literal",
			input:    "SELECT * FROM users WHERE name = ? AND description LIKE '%?%'",
			expected: "SELECT * FROM users WHERE name = $1 AND description LIKE '%?%'",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ReplacePlaceholders(tt.input)
			if result != tt.expected {
				t.Errorf("ReplacePlaceholders(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}
