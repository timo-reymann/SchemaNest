import { Code } from "@heroui/code";
import { Accordion, AccordionItem } from "@heroui/accordion";
import React from "react";

import { CodeHighlight } from "@/components/CodeHighlight";
import { JsonSchemaDetails } from "@/store/jsonSchemas.slice";
import { useBuildApiLink } from "@/util/apiLink";
import { InlineSnippet } from "@/components/InlineSnippet";

export function JsonSchemaSetupInstructions({
  identifier,
  version,
}: Readonly<{
  schema: Partial<JsonSchemaDetails>;
  identifier: string;
  version: string;
}>) {
  const buildApiLink = useBuildApiLink();
  const [major, _, __] = version.split(".");
  const schemaLinkForTools = buildApiLink(
    `/schema/json-schema/${encodeURIComponent(identifier)}/channel/${major}.x`,
  );

  const vsCodeSnippetJson = `"json.schemas": [
    {
        "fileMatch": [
            "/literalPath.json",
            "**/config-match.json"
        ],
        "url": "${schemaLinkForTools}"
    }
]`;
  const vsCodeSnippetYaml = `"yaml.schemas": {
    "${schemaLinkForTools}": [
        "/literalPath.yaml",
        "**/config-match.yaml"
    ],
}`;

  const preCommitSnippet = `hooks:
  - repo: "https://github.com/python-jsonschema/check-jsonschema"
    rev: main
    hooks:
      - id: check-jsonschema
        args:
          - --schemafile
          - http://localhost:8080/api/schema/json-schema/%40test%2Ffoo/channel/1.x
        name: Validate files for schema @test/foo
        files: ^.*/my_config\.(yaml|json)$
`;

  return (
    <>
      <p>
        To make using the JSON-Schema <Code>{identifier}</Code> with other tools
        easier, you can find guides for the most common ecocsystems below.
      </p>
      <Accordion>
        <AccordionItem
          key="idea"
          subtitle="Using editor settings to get completion and validation in your favorite IDE"
          title="IntellIJ IDEA IDEs"
        >
          <article>
            <h2 className="text-xl font-bold">YAML & JSON</h2>
            <ol className="list-decimal ml-5">
              <li>
                In the Settings dialog go to
                <Code>
                  Language & Framework &gt; Schemas and DTDs &gt; JSON Schema
                  Mappings
                </Code>
              </li>
              <li>
                In the central pane, that shows all your previously configured
                Schemas, click <Code>Add</Code> on the toolbar and specify the
                name of the mapping.
              </li>
              <li>
                In the <Code>Name</Code> field enter{" "}
                <InlineSnippet>{identifier}</InlineSnippet>
              </li>
              <li>
                In the <Code>URL or file</Code> field enter{" "}
                <InlineSnippet>{schemaLinkForTools}</InlineSnippet>.
              </li>
              <li>
                Create a list of files or folders that you want to be validated
                against this Schema. Based on the list, IntelliJ IDEA internally
                detects the files to be validated.
                <br />
                The list may contain the names of specific files, the names of
                entire directories, and file name patterns.
                <br />
                Use the following rules to specify file name patterns:
                <ul className="list-disc ml-8">
                  <li>
                    <Code>role-*</Code>
                    matches all files with the names that start with role-.
                  </li>
                  <li>
                    <Code>role-*/**/*.yaml</Code>
                    matches all .yaml files with names that contain role, /, and
                    /.
                  </li>
                  <li>
                    <Code>role-**.yaml</Code>
                    matches all .yaml files with names that start with role-.
                  </li>
                </ul>
                To add an item to the list, click the <Code>
                  Add mapping
                </Code>{" "}
                icon.
              </li>
            </ol>
          </article>
        </AccordionItem>
        <AccordionItem
          key="vscode"
          subtitle="Using editor settings to get completion and validation"
          title="Visual Studio Code"
        >
          <div className="xl:grid grid-cols-2 ">
            <article>
              <h2 className="text-xl font-bold">JSON</h2>
              <ol className="list-decimal ml-5">
                <li>
                  Open your{" "}
                  <a href="https://code.visualstudio.com/docs/configure/settings">
                    Settings
                  </a>
                </li>
                <li>
                  Add a mapping to <Code>json.schemas</Code>:<br />
                  <CodeHighlight
                    code={vsCodeSnippetJson}
                    language={"javascript"}
                  />
                </li>
              </ol>
            </article>
            <article>
              <h2 className="text-xl font-bold">YAML</h2>
              <ol className="list-decimal ml-5">
                <li>
                  Open your{" "}
                  <a href="https://code.visualstudio.com/docs/configure/settings">
                    Settings
                  </a>
                </li>
                <li>
                  Add a mapping to <Code>yaml.schemas</Code>:<br />
                  <CodeHighlight
                    code={vsCodeSnippetYaml}
                    language={"javascript"}
                  />
                </li>
              </ol>
            </article>
          </div>
        </AccordionItem>
        <AccordionItem
          key="pre-commit"
          subtitle="Validate JSON and YAML files before commiting"
          title="pre-commit"
        >
          <p>
            Modify your <Code>.pre-commit-config.yml</Code>:
          </p>
          <CodeHighlight code={preCommitSnippet} language="yaml" />
        </AccordionItem>
      </Accordion>
    </>
  );
}
