import { create } from "zustand/react";

import {
  createJsonSchemaSlice,
  JsonSchemasSlice,
} from "@/store/jsonSchemas.slice";
import { ConfigSlice, createConfigSlice } from "@/store/config.slice";

export const useBoundStore = create<JsonSchemasSlice & ConfigSlice>((...a) => ({
  ...createJsonSchemaSlice(...a),
  ...createConfigSlice(...a),
}));
