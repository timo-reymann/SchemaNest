import Head from "next/head";

export function createTitle(title: string) {
  return <title>{title} | SchemaNest</title>;
}

export function createHead(title: string) {
  return <Head>{createTitle(title)}</Head>;
}
