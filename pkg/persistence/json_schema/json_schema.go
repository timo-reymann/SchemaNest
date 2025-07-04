package json_schema

import (
	"context"
	"database/sql"
	"github.com/timo-reymann/SchemaNest/pkg/persistence/database"
)

func scanSingleRowForSchema(entity *JsonSchemaEntity) func(scan func(dest ...any) error) (bool, error) {
	mapper := func(scan func(dest ...any) error) (bool, error) {
		err := scan(&entity.Id, &entity.Identifier)
		return false, err
	}
	return mapper
}

type JsonSchemaEntity struct {
	Id         *int64
	Identifier string
}

type JsonSchemaEntityLatestVersion struct {
	Major int
	Minor int
	Patch int
}

type JsonSchemaEntityWithBasicInfo struct {
	JsonSchemaEntity
	LatestVersion JsonSchemaEntityLatestVersion
	Description   string
}

type JsonSchemaRepository interface {
	Insert(ctx context.Context, entity *JsonSchemaEntity) error
	List(ctx context.Context) ([]*JsonSchemaEntityWithBasicInfo, error)
	Get(ctx context.Context, identifier string) (*JsonSchemaEntity, error)
}

type JsonSchemaRepositoryImpl struct {
	DB *database.DBConnection
}

func (j *JsonSchemaRepositoryImpl) Get(ctx context.Context, identifier string) (*JsonSchemaEntity, error) {
	entity := &JsonSchemaEntity{}
	err := j.DB.Query(ctx, scanSingleRowForSchema(entity), `
	SELECT id, identifier
	FROM json_schema
	WHERE identifier = ?
	`, identifier)
	if err != nil {
		return nil, err
	}

	return entity, err
}

func (j *JsonSchemaRepositoryImpl) Insert(ctx context.Context, entity *JsonSchemaEntity) error {
	err := j.DB.Insert(
		`
			INSERT INTO json_schema (identifier)
			VALUES (?)
			`,
		entity.Identifier,
	)
	if err != nil {
		return err
	}

	return nil
}

func (j *JsonSchemaRepositoryImpl) List(ctx context.Context) ([]*JsonSchemaEntityWithBasicInfo, error) {
	results := make([]*JsonSchemaEntityWithBasicInfo, 0)

	err := j.DB.Query(
		ctx,
		func(scan func(dest ...any) error) (bool, error) {
			res := &JsonSchemaEntityWithBasicInfo{
				LatestVersion: JsonSchemaEntityLatestVersion{},
			}
			var description sql.NullString
			err := scan(&res.Id, &res.Identifier, &description, &res.LatestVersion.Major, &res.LatestVersion.Minor, &res.LatestVersion.Patch)
			if description.Valid {
				res.Description = description.String
			}
			if err == nil {
				results = append(results, res)
			}
			return true, err
		},
		`
		SELECT s.id, s.identifier, v.description, v.version_major as latest_version_major, v.version_minor as latest_version_minor, v.version_patch as latest_version_minor
		FROM json_schema s
				 INNER JOIN json_schema_version v ON s.id = v.json_schema_id
		WHERE v.id = (SELECT lv.id
					  FROM json_schema_version lv
					  WHERE lv.json_schema_id = s.id
					  ORDER BY lv.version_major DESC, lv.version_minor
														   DESC, lv.version_patch
														   DESC
					  LIMIT 1);
`,
	)
	if err != nil {
		return nil, err
	}

	return results, nil
}
