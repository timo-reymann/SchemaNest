package json_schema

import (
	"context"
	"github.com/timo-reymann/SchemaNest/pkg/persistence/database"
)

type JsonSchemaVersionEntity struct {
	Id           *int64
	VersionMajor int64
	VersionMinor int64
	VersionPatch int64
	Content      string
	JsonSchemaId int64
}

type JsonSchemaVersionRepository interface {
	Insert(ctx context.Context, entity *JsonSchemaVersionEntity) error
	ListForJsonSchema(ctx context.Context, identifier string) ([]*JsonSchemaVersionEntity, error)
	GetForJsonSchemaAndVersion(ctx context.Context, identifier string, versionMajor int64, versionMinor int64, versionPatch int64) (*JsonSchemaVersionEntity, error)
	GetForLatestMajorVersion(ctx context.Context, identifier string, versionMajor int64) (*JsonSchemaVersionEntity, error)
	GetForLatestMinorVersion(ctx context.Context, identifier string, versionMajor int64, versionMinor int64) (*JsonSchemaVersionEntity, error)
	GetLatestVersion(ctx context.Context, identifier string) (*JsonSchemaVersionEntity, error)
}

type JsonSchemaVersionRepositoryImpl struct {
	DB *database.DBConnection
}

func scanSingleRow(entity *JsonSchemaVersionEntity) func(scan func(dest ...any) error) (bool, error) {
	mapper := func(scan func(dest ...any) error) (bool, error) {
		err := scan(&entity.Id, &entity.VersionMajor, &entity.VersionMinor, &entity.VersionPatch, &entity.Content, &entity.JsonSchemaId)
		return false, err
	}
	return mapper
}

func (j *JsonSchemaVersionRepositoryImpl) Insert(ctx context.Context, entity *JsonSchemaVersionEntity) error {
	err := j.DB.Insert(
		`
		INSERT INTO json_schema_version (version_major, version_minor, version_patch, content, json_schema_id) 
		VALUES (?, ?, ?, ?, ?)
        `,
		entity.VersionMajor, entity.VersionMinor, entity.VersionPatch, entity.Content, entity.JsonSchemaId,
	)
	if err != nil {
		return err
	}
	return nil
}

func (j *JsonSchemaVersionRepositoryImpl) ListForJsonSchema(ctx context.Context, identifier string) ([]*JsonSchemaVersionEntity, error) {
	results := make([]*JsonSchemaVersionEntity, 0)

	err := j.DB.Query(
		ctx,
		func(scan func(dest ...any) error) (bool, error) {
			res := &JsonSchemaVersionEntity{}
			err := scan(&res.Id, &res.VersionMajor, &res.VersionMinor, &res.VersionPatch, &res.Content, &res.JsonSchemaId)
			if err == nil {
				results = append(results, res)
			}
			return true, err
		},
		`
		SELECT id, version_major, version_minor, version_patch, content, json_schema_id 
		FROM json_schema_version 
		WHERE json_schema_id = (SELECT id FROM json_schema WHERE identifier = ?) 
		ORDER BY version_major, version_minor, version_patch
		`,
		identifier)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (j *JsonSchemaVersionRepositoryImpl) GetForJsonSchemaAndVersion(ctx context.Context, identifier string, versionMajor int64, versionMinor int64, versionPatch int64) (*JsonSchemaVersionEntity, error) {
	entity := &JsonSchemaVersionEntity{}
	err := j.DB.Query(
		ctx,
		scanSingleRow(entity),
		`
		SELECT id, version_major, version_minor, version_patch, content, json_schema_id 
		FROM json_schema_version 
		WHERE json_schema_id = (SELECT id FROM json_schema WHERE identifier = ?) 
		AND version_major = ? 
		AND version_minor = ?
		AND version_patch = ?
		`,
		identifier, versionMajor, versionMinor, versionPatch,
	)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (j *JsonSchemaVersionRepositoryImpl) GetForLatestMajorVersion(ctx context.Context, identifier string, versionMajor int64) (*JsonSchemaVersionEntity, error) {
	entity := &JsonSchemaVersionEntity{}

	err := j.DB.Query(
		ctx,
		scanSingleRow(entity),
		`
		SELECT id, version_major, version_minor, version_patch, content, json_schema_id 
		FROM json_schema_version
		WHERE json_schema_id = (SELECT id
                        FROM json_schema
                        WHERE identifier = ?)
		AND version_major = ?
		ORDER BY version_major DESC,
        		 version_minor DESC,
         		 version_patch DESC
		LIMIT 1`,
		identifier, versionMajor,
	)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (j *JsonSchemaVersionRepositoryImpl) GetForLatestMinorVersion(ctx context.Context, identifier string, versionMajor int64, versionMinor int64) (*JsonSchemaVersionEntity, error) {
	entity := &JsonSchemaVersionEntity{}

	err := j.DB.Query(
		ctx,
		scanSingleRow(entity),
		`
		SELECT id, version_major, version_minor, version_patch, content, json_schema_id 
		FROM json_schema_version
		WHERE json_schema_id = (SELECT id
                        FROM json_schema
                        WHERE identifier = ?)
		AND version_major = ?
		AND version_minor = ?
		ORDER BY version_major DESC,
        		 version_minor DESC,
         		 version_patch DESC
		LIMIT 1`,
		identifier, versionMajor, versionMinor,
	)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (j *JsonSchemaVersionRepositoryImpl) GetLatestVersion(ctx context.Context, identifier string) (*JsonSchemaVersionEntity, error) {
	entity := &JsonSchemaVersionEntity{}

	err := j.DB.Query(
		ctx,
		scanSingleRow(entity),
		`
		SELECT id, version_major, version_minor, version_patch, content, json_schema_id 
		FROM json_schema_version
		WHERE json_schema_id = (SELECT id
                        FROM json_schema
                        WHERE identifier = ?)
		ORDER BY version_major DESC,
        		 version_minor DESC,
         		 version_patch DESC
		LIMIT 1`,
		identifier,
	)
	if err != nil {
		return nil, err
	}

	return entity, nil
}
