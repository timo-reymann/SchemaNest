import React from "react";
import { DiffEditor } from "@monaco-editor/react";

export const JsonSchemaDiffViewer = React.memo(
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
