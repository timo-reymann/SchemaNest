import { Snippet } from "@heroui/snippet";
import React from "react";

export function InlineSnippet({ children }: Readonly<{ children: string }>) {
  return (
    <Snippet hideSymbol className="pb-0 pt-0">
      {children}
    </Snippet>
  );
}
