CREATE TABLE json_schema
(
    id         SERIAL PRIMARY KEY,
    identifier TEXT NOT NULL UNIQUE
);

CREATE TABLE json_schema_version
(
    id             SERIAL PRIMARY KEY,
    version_major  INTEGER NOT NULL,
    version_minor  INTEGER NOT NULL,
    version_patch  INTEGER NOT NULL,
    content        TEXT    NOT NULL,
    description    TEXT,
    json_schema_id INTEGER NOT NULL REFERENCES json_schema (id) ON DELETE CASCADE,
    CONSTRAINT json_schema_version_uqk_version
        UNIQUE (version_major, version_minor, version_patch, json_schema_id)
);
