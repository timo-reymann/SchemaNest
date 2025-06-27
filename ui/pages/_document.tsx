import {Head, Html, Main, NextScript} from "next/document";
import clsx from "clsx";

import {fontSans} from "@/config/fonts";

export default function Document() {
    return (
        <Html suppressHydrationWarning>
            <Head>
                <link href="/favicon.ico" rel="icon"/>
                <meta name="viewport" content="width=device-width, initial-scale=1"/>
            </Head>
            <body
                className={clsx(
                    "min-h-screen text-foreground bg-background font-sans antialiased",
                    fontSans.variable,
                )}
            >
            <Main/>
            <NextScript/>
            </body>
        </Html>
    );
}
