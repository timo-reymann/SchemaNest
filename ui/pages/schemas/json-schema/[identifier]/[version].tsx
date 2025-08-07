"use client";
import { useRouter } from "next/router";
import React, { useEffect, useState } from "react";
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
import NextLink from "next/link";

import LoadingSpinner from "@/components/LoadingSpinner";
import { useBoundStore } from "@/store/main";
import { title } from "@/components/primitives";
import { createHead } from "@/util/layoutHelpers";
import { JsonSchemaDetails } from "@/store/jsonSchemas.slice";
import { DiffIcon } from "@/components/icons";
import { JsonSchemaSetupInstructions } from "@/components/json-schema/JsonSchemaSetupInstructions";
import { JsonSchemaViewer } from "@/components/json-schema/JsonSchemaViewer";
import { JsonSchemaLinks } from "@/components/json-schema/JsonSchemaLinks";
import { JsonSchemaComparer } from "@/components/json-schema/JsonSchemaComparer";
import GearIcon from "next/dist/next-devtools/dev-overlay/icons/gear-icon";

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
  const fetchLatestVersion = useBoundStore((s) => s.fetchLatestVersion);
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

    if (version === "latest") {
      (async () => {
        try {
          const latest = await fetchLatestVersion(identifier);

          await router.push(
            `/schemas/json-schema/${encodeURIComponent(identifier)}/${latest.major}.${latest.minor}.${latest.patch}`,
          );
        } catch {
          setExists(false);
        }
      })();
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
        comparedVersion,
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
      content: <JsonSchemaLinks identifier={identifier} version={version} />,
    },
    {
      title: "Compare",
      id: "compare",
      icon: <DiffIcon size={18} />,
      content: (
        <>
          {details.versions?.length! > 1 ? (
            <JsonSchemaComparer
              currentSchema={schema}
              currentVersion={version}
              otherSchema={comparedSchema}
              otherVersion={comparedVersion}
              schemaInfo={details}
              theme={theme!}
              onVersionChanged={setComparedVersion}
            />
          ) : (
            <p className="text-warning text-center">
              No other version to compare against.
            </p>
          )}
        </>
      ),
    },
    {
      title: "Validate document",
      id: "validate-document",
      icon: <CheckIcon fontSize={16} />,
      content: (
        <>
          <p className="mb-2">
            Below you can find an editor which you can use to paste JSON or
            write from scratch and get validation inline.
            <br />
            To check a finding hover over the yellow underlined characters.
          </p>
          <JsonSchemaViewer content={{}} schema={schema} theme={theme!} />
        </>
      ),
    },
    {
      title: "Tool Setup",
      id: "setup",
      icon: <GearIcon />,
      content: (
        <JsonSchemaSetupInstructions
          identifier={identifier}
          schema={details}
          version={version}
        />
      ),
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
