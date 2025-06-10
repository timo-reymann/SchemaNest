import "@/styles/globals.css";
import { Viewport } from "next";
import clsx from "clsx";
import { AppProps } from "next/app";

import { Providers } from "@/styles/providers";
import { fontSans } from "@/config/fonts";
import { Navbar } from "@/components/navbar";

export const viewport: Viewport = {
  themeColor: [
    {
      media: "(prefers-color-scheme: light)",
      color: "white",
    },
    {
      media: "(prefers-color-scheme: dark)",
      color: "black",
    },
  ],
};

export default function MyApp({ Component, pageProps }: AppProps) {
  return (
    <div
      className={clsx(
        "min-h-screen text-foreground bg-background font-sans antialiased",
        fontSans.variable,
      )}
    >
      <Providers themeProps={{ attribute: "class", defaultTheme: "dark" }}>
        <div className="relative flex flex-col h-screen">
          <Navbar />
          <main className="container mx-auto pt-1 px-3 flex-grow">
            <Component {...pageProps} />
          </main>
        </div>
      </Providers>
    </div>
  );
}
