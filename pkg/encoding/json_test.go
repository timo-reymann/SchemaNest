package encoding

import (
	"bytes"
	"testing"
)

func TestWriteJSON(t *testing.T) {
	tests := []struct {
		name     string
		data     any
		expected string
		wantErr  bool
	}{
		{"Valid JSON", map[string]string{"key": "value"}, `{"key":"value"}` + "\n", false},
		{"Empty Object", map[string]string{}, `{}` + "\n", false},
		{"Nil Data", nil, "null\n", false},
		{"Invalid Data", make(chan int), "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			err := WriteJSON(&buf, tt.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("WriteJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if buf.String() != tt.expected {
				t.Errorf("WriteJSON() = %q, want %q", buf.String(), tt.expected)
			}
		})
	}
}
