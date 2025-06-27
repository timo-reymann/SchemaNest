import { useTheme } from "next-themes";
import { Snippet } from "@heroui/snippet";
import React from "react";
// @ts-ignore
import { Prism as SyntaxHighlighter } from "react-syntax-highlighter";
// @ts-ignore
import {default as codeDark} from  "react-syntax-highlighter/dist/esm/styles/prism/dracula"
// @ts-ignore
import {default as codeLight} from "react-syntax-highlighter/dist/esm/styles/prism/one-light"

export function CodeHighlight({
  code,
  language,
}: Readonly<{ code: string; language: string }>) {
  const theme = useTheme();
  const style = theme.theme == "light" ? codeLight : codeDark;

  return (
    <Snippet
      className="pt-0 pb-0 pl-0 pr-0"
      hideSymbol={true}
      style={{ background: style['pre[class*="language-"]'].background }}
    >
      <SyntaxHighlighter
        className="lg:max-w-[500px] xl:max-w-[640px] sm:max-w-[500px] p-0"
        language={language}
        style={style}
        wrapLongLines={true}
      >
        {code}
      </SyntaxHighlighter>
    </Snippet>
  );
}
