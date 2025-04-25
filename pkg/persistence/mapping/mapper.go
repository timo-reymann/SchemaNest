package mapping

func MapEntitiesToModel[Entity any, Model any](mapper func(e *Entity) *Model, entities []*Entity) []*Model {
	models := make([]*Model, len(entities))
	for idx, entity := range entities {
		model := mapper(entity)
		models[idx] = model
	}
	return models
}
