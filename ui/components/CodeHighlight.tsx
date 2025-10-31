import { useTheme } from "next-themes";
import { Snippet } from "@heroui/snippet";
import React from "react";
import ShikiHighlighter from "react-shiki";

export function CodeHighlight({
  code,
  language,
}: Readonly<{ code: string; language: string }>) {
  const theme = useTheme();
  const style = theme.theme == "light" ? "github-light" : "github-dark";

  return (
    <Snippet className="pt-0 pb-0 pl-0 pr-0" hideSymbol={true}>
      <ShikiHighlighter
        className="max-w-full"
        language={language}
        showLanguage={false}
        theme={style}
      >
        {code}
      </ShikiHighlighter>
    </Snippet>
  );
}
