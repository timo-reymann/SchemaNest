import {StateCreator} from "zustand/vanilla";
import {getJSON} from "@/util/fetch";

export interface JsonSchemaInfo {
    identifier: string;
    latestVersion: {
        major: number;
        minor: number;
        patch: number;
    };
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
    fetchSchemaByVersion: (identifier: string, version: string) => Promise<Object>
}

const mockSchemas = [
    {
        identifier: "trump",
        description:
            "Absolutely the Best, Most Incredible Schema Ever Created, Believe Me!",
        type: "JSON-Schema",
    },
    {
        identifier: "lazy",
        type: "JSON-Schema",
    },
    {
        identifier: "bigly",
        description: "Big data? Big validation. It’s going to be YUGE.",
        type: "JSON-Schema",
    },
    {
        identifier: "covfefe",
        description: "Nobody knows what it means, but it's strong. Very strong.",
        type: "JSON-Schema",
    },
    {
        identifier: "fake_news",
        description: "This schema blocks fake fields. Only true data allowed!",
        type: "JSON-Schema",
    },
    {
        identifier: "maga",
        description: "Making APIs Great Again, one schema at a time.",
        type: "JSON-Schema",
    },
    {
        identifier: "wall",
        description: "It stops bad input. And trust me, it works!",
        type: "JSON-Schema",
    },
    {
        identifier: "unbelievable",
        type: "JSON-Schema",
    },
    {
        identifier: "ratings",
        description: "Five stars. Ten out of ten. The best rated schema ever.",
        type: "JSON-Schema",
    },
    {
        identifier: "tremendous",
        description: "Everyone's talking about it. It's just tremendous.",
        type: "JSON-Schema",
    },
    {
        identifier: "best_words",
        description: "This schema? It uses the best words. The very best.",
        type: "JSON-Schema",
    },
    {
        identifier: "strong",
        description: "So strong, your backend will feel it.",
        type: "JSON-Schema",
    },
    {
        identifier: "gold_standard",
        type: "JSON-Schema",
    },
    {
        identifier: "truth",
        description: "This schema only validates truth. No lies. No fake.",
        type: "JSON-Schema",
    },
    {
        identifier: "success",
        type: "JSON-Schema",
    },
    {
        identifier: "genius",
        description: "Some people call it genius. Others call it stable.",
        type: "JSON-Schema",
    },
    {
        identifier: "patriot",
        type: "JSON-Schema",
    },
    {
        identifier: "firehose",
        description: "It handles data like a firehose. No leaks. Total control.",
        type: "JSON-Schema",
    },
    {
        identifier: "winner",
        description: "This schema? It’s a total winner. Never loses.",
        type: "JSON-Schema",
    },
    {
        identifier: "perfect",
        description: "Perfect validation. Like a perfect call.",
        type: "JSON-Schema",
    },
    {
        identifier: "flawless",
        type: "JSON-Schema",
    },
    {
        identifier: "deal_maker",
        description: "Use this schema in a deal? Automatic win.",
        type: "JSON-Schema",
    },
    {
        identifier: "powerful",
        type: "JSON-Schema",
    },
    {
        identifier: "yuge_success",
        description: "It brings tremendous success. Just look at the logs.",
        type: "JSON-Schema",
    },
    {
        identifier: "elite",
        type: "JSON-Schema",
    },
    {
        identifier: "real_deal",
        description: "Not like those phony specs. This is the real deal.",
        type: "JSON-Schema",
    },
    {
        identifier: "stable_genius",
        description: "Designed by a stable genius. Obviously.",
        type: "JSON-Schema",
    },
    {
        identifier: "nasty_bug_blocker",
        type: "JSON-Schema",
    },
    {
        identifier: "totally_legal",
        description: "All fields legal. Totally. I’ve seen the docs.",
        type: "JSON-Schema",
    },
    {
        identifier: "executive",
        description: "Executive-level schema validation. The best people use it.",
        type: "JSON-Schema",
    },
    {
        identifier: "trump",
        description:
            "Absolutely the Best, Most Incredible Schema Ever Created, Believe Me!",
        type: "JSON-Schema",
    },
    {
        identifier: "lazy",
        type: "JSON-Schema",
    },
    {
        identifier: "bigly",
        description: "Big data? Big validation. It’s going to be YUGE.",
        type: "JSON-Schema",
    },
    {
        identifier: "covfefe",
        description: "Nobody knows what it means, but it's strong. Very strong.",
        type: "JSON-Schema",
    },
    {
        identifier: "fake_news",
        description: "This schema blocks fake fields. Only true data allowed!",
        type: "JSON-Schema",
    },
    {
        identifier: "maga",
        description: "Making APIs Great Again, one schema at a time.",
        type: "JSON-Schema",
    },
    {
        identifier: "wall",
        description: "It stops bad input. And trust me, it works!",
        type: "JSON-Schema",
    },
    {
        identifier: "unbelievable",
        type: "JSON-Schema",
    },
    {
        identifier: "ratings",
        description: "Five stars. Ten out of ten. The best rated schema ever.",
        type: "JSON-Schema",
    },
    {
        identifier: "tremendous",
        description: "Everyone's talking about it. It's just tremendous.",
        type: "JSON-Schema",
    },
    {
        identifier: "best_words",
        description: "This schema? It uses the best words. The very best.",
        type: "JSON-Schema",
    },
    {
        identifier: "strong",
        description: "So strong, your backend will feel it.",
        type: "JSON-Schema",
    },
    {
        identifier: "gold_standard",
        type: "JSON-Schema",
    },
    {
        identifier: "truth",
        description: "This schema only validates truth. No lies. No fake.",
        type: "JSON-Schema",
    },
    {
        identifier: "success",
        type: "JSON-Schema",
    },
    {
        identifier: "genius",
        description: "Some people call it genius. Others call it stable.",
        type: "JSON-Schema",
    },
    {
        identifier: "patriot",
        type: "JSON-Schema",
    },
    {
        identifier: "firehose",
        description: "It handles data like a firehose. No leaks. Total control.",
        type: "JSON-Schema",
    },
    {
        identifier: "winner",
        description: "This schema? It’s a total winner. Never loses.",
        type: "JSON-Schema",
    },
    {
        identifier: "perfect",
        description: "Perfect validation. Like a perfect call.",
        type: "JSON-Schema",
    },
    {
        identifier: "flawless",
        type: "JSON-Schema",
    },
    {
        identifier: "deal_maker",
        description: "Use this schema in a deal? Automatic win.",
        type: "JSON-Schema",
    },
    {
        identifier: "powerful",
        type: "JSON-Schema",
    },
    {
        identifier: "yuge_success",
        description: "It brings tremendous success. Just look at the logs.",
        type: "JSON-Schema",
    },
    {
        identifier: "elite",
        type: "JSON-Schema",
    },
    {
        identifier: "real_deal",
        description: "Not like those phony specs. This is the real deal.",
        type: "JSON-Schema",
    },
    {
        identifier: "stable_genius",
        description: "Designed by a stable genius. Obviously.",
        type: "JSON-Schema",
    },
    {
        identifier: "nasty_bug_blocker",
        type: "JSON-Schema",
    },
    {
        identifier: "totally_legal",
        description: "All fields legal. Totally. I’ve seen the docs.",
        type: "JSON-Schema",
    },
    {
        identifier: "executive",
        description: "Executive-level schema validation. The best people use it.",
        type: "JSON-Schema",
    },
];

const mockSchemaWithDetails = mockSchemas.map((v) => ({
    ...v,
    identifier: v.identifier,
    latestVersion: {major: 1, minor: 0, patch: 0},
}));

export const createJsonSchemaSlice: StateCreator<JsonSchemasSlice> = (
    set,
    get,
) => ({
    schemas: [],
    schemaLoading: true,
    fetchSchemas: async () => {
        try {
            const schemas = await getJSON<JsonSchemaInfo[]>("/api/schema/json-schema")
            set({
                schemas,
                schemaLoading: false,
            });
        } finally {
            set({schemaLoading: false});
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
        })
        try {
            let result = await getJSON(`/api/schema/json-schema/${identifier}/version/${version}`)
            return result as Object
        } finally {
            set({
                schemaLoading: false,
            })
        }
    }
});
