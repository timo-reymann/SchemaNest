import React from "react";
import { Editor } from "@monaco-editor/react";

export const JsonSchemaViewer = React.memo(
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
      <>
        <p className="mb-2">
          Below you can find an editor which you can use to paste JSON or write
          from scratch and get validation inline.
          <br />
          To check a finding hover over the yellow underlined characters.
        </p>
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
      </>
    );
  },
  (prev, next) => prev.content === next.content && prev.theme == next.theme,
);
