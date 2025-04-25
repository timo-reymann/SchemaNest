package api

import "testing"

func TestNewServeMux(t *testing.T) {
	mux, err := NewServeMux(&SchemaNestApiContext{})
	if err != nil {
		t.Fatal(err)
	}

	if mux == nil {
		t.Fatal("Expected a new ServeMux, got nil")
	}
}
