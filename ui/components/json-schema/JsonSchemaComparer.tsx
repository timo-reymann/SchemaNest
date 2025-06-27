import { Select, SelectItem } from "@heroui/select";
import React from "react";

import { JsonSchemaDetails } from "@/store/jsonSchemas.slice";
import { JsonSchemaDiffViewer } from "@/components/json-schema/JsonSchemaDiffViewer";

export function JsonSchemaComparer({
  currentSchema,
  schemaInfo,
  theme,
  otherSchema,
  otherVersion,
  onVersionChanged,
  currentVersion,
}: Readonly<{
  currentSchema: Object;
  currentVersion: string;
  otherSchema: Object | null;
  otherVersion: string | undefined;
  onVersionChanged: (version: string) => void;
  schemaInfo: Partial<JsonSchemaDetails>;
  theme: string;
}>) {
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
