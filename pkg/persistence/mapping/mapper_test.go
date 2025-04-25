package mapping

import "testing"

func TestMapEntitiesToModel(t *testing.T) {
	// Define local structs
	type Entity struct {
		ID   int
		Name string
	}
	type Model struct {
		ID   int
		Name string
	}

	// Define a mapper function
	mapper := func(e *Entity) *Model {
		return &Model{
			ID:   e.ID,
			Name: e.Name,
		}
	}

	// Input entities
	entities := []*Entity{
		{ID: 1, Name: "Entity1"},
		{ID: 2, Name: "Entity2"},
	}

	// Call MapEntitiesToModel
	models := MapEntitiesToModel(mapper, entities)

	// Expected models
	expected := []*Model{
		{ID: 1, Name: "Entity1"},
		{ID: 2, Name: "Entity2"},
	}

	// Validate results
	if len(models) != len(expected) {
		t.Fatalf("Expected %d models, got %d", len(expected), len(models))
	}

	for i, model := range models {
		if model.ID != expected[i].ID || model.Name != expected[i].Name {
			t.Errorf("Model at index %d does not match. Expected %+v, got %+v", i, expected[i], model)
		}
	}
}
