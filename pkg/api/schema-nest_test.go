package api

import (
	"context"
	"github.com/timo-reymann/SchemaNest/pkg/persistence/json_schema"
	"testing"
)

type MockJsonSchemaRepository struct {
	Schemas   []*json_schema.JsonSchemaEntityWithBasicInfo
	InsertErr error
}

func (m *MockJsonSchemaRepository) List(ctx context.Context) ([]*json_schema.JsonSchemaEntityWithBasicInfo, error) {
	return m.Schemas, nil
}

func (m *MockJsonSchemaRepository) Insert(ctx context.Context, entity *json_schema.JsonSchemaEntity) error {
	return m.InsertErr
}

type MockJsonSchemaVersionRepository struct {
	Versions []*json_schema.JsonSchemaVersionEntity
}

func (m *MockJsonSchemaVersionRepository) GetForLatestMajorVersion(ctx context.Context, identifier string, versionMajor int64) (*json_schema.JsonSchemaVersionEntity, error) {
	for _, version := range m.Versions {
		if version.VersionMajor == versionMajor && version.JsonSchemaId == 1 {
			return version, nil
		}
	}
	return nil, nil
}

func (m *MockJsonSchemaVersionRepository) GetForLatestMinorVersion(ctx context.Context, identifier string, versionMajor int64, versionMinor int64) (*json_schema.JsonSchemaVersionEntity, error) {
	for _, version := range m.Versions {
		if version.VersionMajor == versionMajor && version.VersionMinor == versionMinor && version.JsonSchemaId == 1 {
			return version, nil
		}
	}
	return nil, nil
}

func (m *MockJsonSchemaVersionRepository) ListForJsonSchema(ctx context.Context, identifier string) ([]*json_schema.JsonSchemaVersionEntity, error) {
	var results []*json_schema.JsonSchemaVersionEntity
	for _, version := range m.Versions {
		if version.JsonSchemaId == 1 {
			results = append(results, version)
		}
	}
	return results, nil
}

func (m *MockJsonSchemaVersionRepository) GetForJsonSchemaAndVersion(ctx context.Context, identifier string, versionMajor int64, versionMinor int64, versionPatch int64) (*json_schema.JsonSchemaVersionEntity, error) {
	for _, version := range m.Versions {
		if version.VersionMajor == versionMajor && version.VersionMinor == versionMinor && version.VersionPatch == versionPatch {
			return version, nil
		}
	}
	return nil, nil
}

func (m *MockJsonSchemaVersionRepository) Insert(ctx context.Context, entity *json_schema.JsonSchemaVersionEntity) error {
	m.Versions = append(m.Versions, entity)
	return nil
}

func (m *MockJsonSchemaVersionRepository) GetLatestVersion(ctx context.Context, identifier string) (*json_schema.JsonSchemaVersionEntity, error) {
	for _, version := range m.Versions {
		if version.JsonSchemaId == 1 {
			return version, nil
		}
	}
	return nil, nil
}

func createTestContext() SchemaNestApiContext {
	return SchemaNestApiContext{
		JsonSchemaRepository:        &MockJsonSchemaRepository{},
		JsonSchemaVersionRepository: &MockJsonSchemaVersionRepository{},
	}
}

func TestNewSchemaNestApi(t *testing.T) {
	mockContext := createTestContext()
	api := NewSchemaNestApi(&mockContext)
	if api == nil {
		t.Errorf("Expected NewSchemaNestApi to return a non-nil value")
	}
}
