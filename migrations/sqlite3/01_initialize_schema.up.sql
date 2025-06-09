create table json_schema
(
    id         integer
        constraint json_schema_pk
            primary key autoincrement,
    identifier TEXT not null
        constraint json_schema_uqk_identifier
            unique
);

create table json_schema_version
(
    id             integer not null
        constraint json_schema_version_pk
            primary key autoincrement,
    version_major  integer not null,
    version_minor  integer not null,
    version_patch  integer not null,
    content        TEXT    not null,
    description    TEXT,
    json_schema_id integer not null
        constraint json_schema_version__json_schema_fk
            references json_schema
            on delete cascade,
    constraint json_schema_version_uqk_version
        unique (version_major, version_minor, version_patch, json_schema_id)
);
