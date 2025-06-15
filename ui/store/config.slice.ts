import { StateCreator } from "zustand/vanilla";

import { getJSON } from "@/util/fetch";

export interface Configuration {
  apiKeyAuthEnabled: boolean;
}

export interface ConfigSlice {
  config: Configuration | null;
  configLoading: boolean;
  loadConfig: () => Promise<Configuration>;
}

export const createConfigSlice: StateCreator<ConfigSlice> = (set, get) => ({
  config: null,
  configLoading: false,
  loadConfig: async () => {
    if (get().config == null) {
      set({
        configLoading: true,
      });

      const config = await getJSON<Configuration>("/ui-config");

      set({
        config: config,
        configLoading: false,
      });
    }

    return Promise.resolve(get().config!);
  },
});
