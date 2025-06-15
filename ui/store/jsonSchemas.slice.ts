import { StateCreator } from "zustand/vanilla";

import { getJSON, postJSON } from "@/util/fetch";

export interface JsonSchemaInfo {
  identifier: string;
  description: string | undefined;
  latestVersion: {
    major: number;
    minor: number;
    patch: number;
  };
}

export interface JsonSchemaDetails {
  description: string | undefined;
  versions: [
    {
      version: string;
    },
  ];
}

export interface Error {
  error: string;
}

export interface JsonSchemasSlice {
  schemas: JsonSchemaInfo[];
  schemaLoading: boolean;
  fetchSchemas: () => Promise<void>;
  filterAndPaginateSchemas: (
    search: string,
    page: number,
    perPage: number,
  ) => {
    slice: JsonSchemaInfo[];
    totalPages: number;
  };
  fetchSchemaByVersion: (
    identifier: string,
    version: string,
  ) => Promise<Object>;
  fetchSchemaDetailsByVersion: (
    identifier: string,
  ) => Promise<JsonSchemaDetails>;

  uploadSchema(input: {
    apiKey: string | undefined;
    identifier: string;
    version: string;
    schema: File;
  }): Promise<void>;
}

export const createJsonSchemaSlice: StateCreator<JsonSchemasSlice> = (
  set,
  get,
) => ({
  schemas: [],
  schemaLoading: false,
  fetchSchemas: async () => {
    set({
      schemas: [],
      schemaLoading: true,
    });
    try {
      const schemas = await getJSON<JsonSchemaInfo[]>(
        "/api/schema/json-schema",
      );

      set({
        schemas,
        schemaLoading: false,
      });
    } catch {
      set({ schemaLoading: false });
    }
  },
  filterAndPaginateSchemas: (search, page, perPage) => {
    let filtered = get().schemas;

    if (search.trim() != "") {
      filtered = filtered.filter((s) =>
        s.identifier.toLowerCase().includes(search.toLowerCase().trim()),
      );
    }

    const start = (page - 1) * perPage;
    const end = start + perPage;

    return {
      slice: filtered.slice(start, end),
      totalPages: Math.ceil(filtered.length / perPage),
    };
  },
  fetchSchemaByVersion: async (identifier, version) => {
    set({
      schemaLoading: true,
    });
    try {
      const res = await getJSON<Object & Error>(
        `/api/schema/json-schema/${encodeURIComponent(identifier)}/version/${version}`,
      );

      if (res.error) {
        throw res.error;
      }

      return res as Object;
    } finally {
      set({
        schemaLoading: false,
      });
    }
  },
  fetchSchemaDetailsByVersion: async (identifier) => {
    set({
      schemaLoading: true,
    });
    try {
      return await getJSON<JsonSchemaDetails>(
        `/api/schema/json-schema/${encodeURIComponent(identifier)}`,
      );
    } finally {
      set({
        schemaLoading: false,
      });
    }
  },
  async uploadSchema(input: {
    apiKey: string | undefined;
    identifier: string;
    version: string;
    schema: File;
  }): Promise<void> {
    let response = await postJSON(
      `/api/schema/json-schema/${encodeURIComponent(input.identifier)}/version/${input.version}`,
      JSON.parse(await input.schema.text()),
      input.apiKey ? { Authorization: `Bearer ${input.apiKey}` } : {},
    );

    if (response.status != 201) {
      throw (await response.json()).error;
    }
  },
});
