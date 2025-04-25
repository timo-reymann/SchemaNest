package json_schema

import (
	"context"
	"github.com/timo-reymann/SchemaNest/pkg/persistence/database"
)

type JsonSchemaEntity struct {
	Id         *int64
	Identifier string
}

type JsonSchemaRepository interface {
	Insert(ctx context.Context, entity *JsonSchemaEntity) error
	List(ctx context.Context) ([]*JsonSchemaEntity, error)
}

type JsonSchemaRepositoryImpl struct {
	DB *database.DBConnection
}

func (j *JsonSchemaRepositoryImpl) Insert(ctx context.Context, entity *JsonSchemaEntity) error {
	id, err := j.DB.Insert("INSERT INTO json_schema (identifier) VALUES (?)", entity.Identifier)
	if err != nil {
		return err
	}

	entity.Id = id
	return nil
}

func (j *JsonSchemaRepositoryImpl) List(ctx context.Context) ([]*JsonSchemaEntity, error) {
	results := make([]*JsonSchemaEntity, 0)

	err := j.DB.Query(context.Background(), func(scan func(dest ...any) error) (bool, error) {
		res := &JsonSchemaEntity{}
		err := scan(&res.Id, &res.Identifier)
		if err == nil {
			results = append(results, res)
		}
		return true, err
	}, "SELECT id, identifier FROM json_schema")
	if err != nil {
		return nil, err
	}

	return results, nil
}
