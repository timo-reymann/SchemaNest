import {useTheme} from "next-themes";
import {Snippet} from "@heroui/snippet";
import React from "react";
// @ts-ignore
import SyntaxHighlighter from "react-syntax-highlighter"
// @ts-ignore
import {darcula, idea} from "react-syntax-highlighter/dist/esm/styles/hljs";

export function CodeHighlight({code, language}: Readonly<{ code: string, language: string }>) {
    const theme = useTheme()
    const style = theme.theme == "light" ? idea : darcula

    return <Snippet style={{background: style.hljs.background}}
                    hideSymbol={true}>
        <SyntaxHighlighter language={language}
                           style={theme.theme == "light" ? idea : darcula}>
            {code}
        </SyntaxHighlighter>
    </Snippet>
}