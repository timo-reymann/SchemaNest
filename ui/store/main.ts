import { create } from "zustand/react";

import {
  createJsonSchemaSlice,
  JsonSchemasSlice,
} from "@/store/jsonSchemas.slice";

export const useBoundStore = create<JsonSchemasSlice>((...a) => ({
  ...createJsonSchemaSlice(...a),
}));
