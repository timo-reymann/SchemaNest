"use client";
import { useRouter } from "next/router";
import React, { useEffect, useState } from "react";
import { DiffEditor, Editor } from "@monaco-editor/react";
import { useTheme } from "next-themes";
import { Code } from "@heroui/code";
import { Tab, Tabs } from "@heroui/tabs";
import {
  CheckIcon,
  ClockCircleLinearIcon,
  EyeIcon,
  LinkIcon,
} from "@heroui/shared-icons";
import { Chip } from "@heroui/chip";
import {
  Table,
  TableBody,
  TableCell,
  TableColumn,
  TableHeader,
  TableRow,
} from "@heroui/table";
import { Snippet } from "@heroui/snippet";
import { Select, SelectItem } from "@heroui/select";
import NextLink from "next/link";

import LoadingSpinner from "@/components/LoadingSpinner";
import { useBoundStore } from "@/store/main";
import { title } from "@/components/primitives";
import { createHead } from "@/util/layoutHelpers";
import { JsonSchemaDetails } from "@/store/jsonSchemas.slice";
import { DiffIcon } from "@/components/icons";

const JsonSchemaViewer = React.memo(
  function JsonSchemaViewer({
    theme,
    content,
    schema,
  }: {
    theme: string;
    content: Object;
    schema?: Object;
  }) {
    return (
      <Editor
        beforeMount={(monaco) => {
          monaco.languages.json.jsonDefaults.setDiagnosticsOptions({
            validate: true,
            schemas: [
              {
                uri: "in-memory",
                fileMatch: ["*"],
                schema,
              },
            ],
          });
        }}
        defaultLanguage="json"
        height="50vh"
        options={{
          readOnly: schema === undefined,
          contextmenu: false,
        }}
        theme={theme === "dark" ? "vs-dark" : "vs-light"}
        value={JSON.stringify(content, null, "  ")}
      />
    );
  },
  (prev, next) => prev.content === next.content && prev.theme == next.theme,
);

const JsonSchemaDiffViewer = React.memo(
  function JsonSchemaDiffViewer({
    theme,
    currentSchema,
    otherSchema,
  }: {
    theme: string;
    currentSchema: Object;
    otherSchema: Object;
  }) {
    const modified = JSON.stringify(otherSchema, null, 2);
    const original = JSON.stringify(currentSchema, null, 2);

    return (
      <DiffEditor
        height="50vh"
        keepCurrentModifiedModel={true}
        keepCurrentOriginalModel={true}
        language="json"
        modified={modified}
        options={{
          readOnly: true,
          contextmenu: false,
        }}
        original={original}
        theme={theme === "dark" ? "vs-dark" : "vs-light"}
      />
    );
  },
  (prev, next) =>
    prev.theme === next.theme &&
    prev.currentSchema === next.currentSchema &&
    prev.otherSchema === next.otherSchema,
);

export function JsonSchemaLinks({
  identifier,
  version,
  schema,
}: {
  identifier: string;
  version: string;
  schema: Partial<JsonSchemaDetails>;
}) {
  const hostname = window.location.origin;

  const buildApiLink = (path: string) => `${hostname}/api${path}`;
  const [major, minor] = version.split(".");

  const links = [
    {
      label: "Get JSON schema for the currently selected version",
      method: "GET",
      url: buildApiLink(`/schema/json-schema/${identifier}/version/${version}`),
    },
    {
      label: "Get latest JSON schema version",
      method: "GET",
      url: buildApiLink(`/schema/json-schema/${identifier}/latest`),
    },
    {
      label: "Get schema for latest minor for current version",
      method: "GET",
      url: buildApiLink(`/schema/json-schema/${identifier}/channel/${major}.x`),
    },
    {
      label: "Get schema for latest patch for current version",
      method: "GET",
      url: buildApiLink(
        `/schema/json-schema/${identifier}/channel/${major}.${minor}.x`,
      ),
    },
    {
      label: "Create new schema definition for a given version",
      method: "POST",
      url: buildApiLink(`/schema/json-schema/${identifier}/version/{version}`),
    },
  ];

  return (
    <Table isStriped aria-label="Links for schema">
      <TableHeader>
        <TableColumn minWidth={200}>Operation</TableColumn>
        <TableColumn minWidth={200}>Endpoint</TableColumn>
      </TableHeader>
      <TableBody>
        {links.map((l) => (
          <TableRow key={l.label}>
            <TableCell>{l.label}</TableCell>
            <TableCell>
              <div className="flex gap-2 items-center">
                <Chip color={l.method === "GET" ? "primary" : "secondary"}>
                  {l.method}
                </Chip>
                <Snippet hideSymbol className="w-full">
                  <span>{l.url}</span>
                </Snippet>
              </div>
            </TableCell>
          </TableRow>
        ))}
      </TableBody>
    </Table>
  );
}

export function JsonSchemaComparer({
  currentSchema,
  identifier,
  schemaInfo,
  theme,
  otherSchema,
  otherVersion,
  onVersionChanged,
  currentVersion,
}: {
  currentSchema: Object;
  currentVersion: string;
  otherSchema: Object | null;
  otherVersion: string | undefined;
  onVersionChanged: (version: string) => void;
  identifier: string;
  schemaInfo: Partial<JsonSchemaDetails>;
  theme: string;
}) {
  return (
    <div>
      <Select
        isVirtualized
        className="mb-3"
        label="Select version to compare against"
        placeholder="..."
        selectedKeys={otherVersion ? [otherVersion] : []}
        onSelectionChange={(k) => onVersionChanged(k.currentKey!)}
      >
        {schemaInfo
          .versions!.filter((v) => v.version !== currentVersion)
          .map((version) => (
            <SelectItem key={version.version}>{version.version}</SelectItem>
          ))}
      </Select>
      {otherSchema && otherVersion ? (
        <>
          <div className="flex justify-around mt-3 mb-1">
            <div className="text-center font-light">{currentVersion}</div>
            <div className="text-center font-light">{otherVersion}</div>
          </div>
          <JsonSchemaDiffViewer
            currentSchema={currentSchema}
            otherSchema={otherSchema}
            theme={theme}
          />
        </>
      ) : (
        <p className="text-center">
          Please select a version to compare the schema with.
        </p>
      )}
    </div>
  );
}

export default function JsonSchemaVersionPage() {
  const { theme } = useTheme();
  const router = useRouter();
  const { identifier, version } = router.query as {
    identifier: string;
    version: string;
  };

  const fetchSchemaByVersion = useBoundStore((s) => s.fetchSchemaByVersion);
  const fetchSchemaDetailsByVersion = useBoundStore(
    (s) => s.fetchSchemaDetailsByVersion,
  );
  const loading = useBoundStore((s) => s.schemaLoading);
  const [tab, setTab] = useState("view");
  const [schema, setSchema] = useState<Object>({});
  const [details, setDetails] = useState<Partial<JsonSchemaDetails>>({});
  const [exists, setExists] = useState(false);
  const [comparedVersion, setComparedVersion] = useState<string>();
  const [comparedSchema, setComparedSchema] = useState<Object | null>(null);

  useEffect(() => {
    if (!router.isReady || !identifier || !version) {
      return;
    }
    setComparedSchema(null);
    setComparedVersion(undefined);
    (async () => {
      try {
        let details = await fetchSchemaDetailsByVersion(identifier);

        setDetails(details);
      } catch {
        setExists(false);
      }

      try {
        let schema = await fetchSchemaByVersion(identifier, version);

        setSchema(schema);
        setExists(true);
      } catch {
        setExists(false);
      }
    })();
  }, [fetchSchemaByVersion, router.isReady, version]);

  useEffect(() => {
    if (router.query.tab) {
      setTab(router.query.tab as string);
    }
  }, [router.isReady, router.query.tab]);

  useEffect(() => {
    if (!router.isReady) {
      return;
    }
    const newQuery = { ...router.query, tab };
    const url = {
      pathname: router.pathname,
      query: newQuery,
    };

    router.push(url, undefined, { shallow: true });
  }, [tab]);

  useEffect(() => {
    if (!identifier || !comparedVersion) {
      return;
    }
    (async () => {
      const otherSchema = await fetchSchemaByVersion(
        identifier,
        comparedVersion!,
      );

      setComparedSchema(otherSchema);
    })();
  }, [comparedVersion, version]);

  if (loading) {
    return <LoadingSpinner label="Loading schema ..." />;
  }

  if (!exists) {
    return (
      <div className="flex h-full w-full">
        <p className=" mx-auto text-center">
          JSON-Schema <Code>{identifier}</Code> in version{" "}
          <Code>{version}</Code> not found.
        </p>
      </div>
    );
  }

  const tabs = [
    {
      title: "View",
      id: "view",
      content: <JsonSchemaViewer content={schema} theme={theme!} />,
      icon: <EyeIcon fontSize={16} />,
    },
    {
      title: "Endpoints",
      id: "endpoints",
      icon: <LinkIcon fontSize={16} />,
      content: (
        <JsonSchemaLinks
          identifier={identifier}
          schema={details}
          version={version}
        />
      ),
    },
    {
      title: "Compare",
      id: "compare",
      icon: <DiffIcon size={18} />,
      content: (
        <JsonSchemaComparer
          currentSchema={schema}
          currentVersion={version}
          identifier={identifier}
          otherSchema={comparedSchema}
          otherVersion={comparedVersion}
          schemaInfo={details}
          theme={theme!}
          onVersionChanged={setComparedVersion}
        />
      ),
    },
    {
      title: "Validate document",
      id: "validate-document",
      icon: <CheckIcon fontSize={16} />,
      content: <JsonSchemaViewer content={{}} schema={schema} theme={theme!} />,
    },
    {
      title: "History",
      id: "other-versions",
      icon: <ClockCircleLinearIcon />,
      content: (
        <ul className="list-disc">
          {details.versions!.map((v) => (
            <li key={v.version}>
              <NextLink
                passHref
                href={{
                  pathname: "/schemas/json-schema/[identifier]/[version]",
                  query: {
                    identifier: identifier,
                    version: v.version,
                    tab: "view",
                  },
                }}
              >
                {v.version}
              </NextLink>
            </li>
          ))}
        </ul>
      ),
    },
  ];

  return (
    <>
      {createHead(`${identifier} - JSON Schema`)}
      <div className="flex items-center gap-2">
        <h1 className={title()}>{identifier}</h1>
        <Chip>v{version}</Chip>
      </div>
      {details.description && <p className="mt-2">{details.description}</p>}
      <Tabs
        aria-label="Options"
        className="mt-3 mx-auto block"
        classNames={{
          tabList:
            "gap-6 w-full relative rounded-none p-0 border-b border-divider mx-auto mb-2",
          cursor: "w-full ",
          tab: "max-w-fit px-0 h-12",
        }}
        color="primary"
        items={tabs}
        selectedKey={tab}
        variant="underlined"
        onSelectionChange={(t) => setTab(t.toString())}
      >
        {(item) => (
          <Tab
            key={item.id}
            title={
              <div className="flex items-center space-x-2">
                {item.icon}
                <span>{item.title}</span>
              </div>
            }
          >
            {item.content}
          </Tab>
        )}
      </Tabs>
    </>
  );
}
