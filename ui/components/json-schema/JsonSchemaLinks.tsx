import {
  Table,
  TableBody,
  TableCell,
  TableColumn,
  TableHeader,
  TableRow,
} from "@heroui/table";
import { Chip } from "@heroui/chip";
import { Snippet } from "@heroui/snippet";
import React from "react";

import { useBuildApiLink } from "@/util/apiLink";

export function JsonSchemaLinks({
  identifier,
  version,
}: Readonly<{
  identifier: string;
  version: string;
}>) {
  const buildApiLink = useBuildApiLink();
  const [major, minor] = version.split(".");

  const links = [
    {
      label: "Get JSON schema for the currently selected version",
      method: "GET",
      url: buildApiLink(
        `/schema/json-schema/${encodeURIComponent(identifier)}/version/${version}`,
      ),
    },
    {
      label: "Get latest JSON schema version",
      method: "GET",
      url: buildApiLink(
        `/schema/json-schema/${encodeURIComponent(identifier)}/latest`,
      ),
    },
    {
      label: "Get schema for latest minor for current version",
      method: "GET",
      url: buildApiLink(
        `/schema/json-schema/${encodeURIComponent(identifier)}/channel/${major}.x`,
      ),
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
