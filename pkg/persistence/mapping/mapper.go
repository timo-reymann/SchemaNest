package mapping

func MapEntitiesToModel[Entity any, Model any](mapper func(e *Entity) *Model, entities []*Entity) []*Model {
	models := make([]*Model, len(entities))
	for idx, entity := range entities {
		model := mapper(entity)
		models[idx] = model
	}
	return models
}

func MapEntitiesToModelWithValues[Entity any, Model any](
	mapper func(e *Entity) *Model,
	entities []*Entity,
) ([]*Model, []Model) {
	ptrs := make([]*Model, len(entities))
	values := make([]Model, len(entities))

	for i, entity := range entities {
		modelPtr := mapper(entity)
		ptrs[i] = modelPtr
		if modelPtr != nil {
			values[i] = *modelPtr
		}
	}

	return ptrs, values
}
