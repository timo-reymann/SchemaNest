import { RedocStandalone } from "redoc";
import resolveConfig from "tailwindcss/resolveConfig";
import { useTheme } from "next-themes";

import { createHead } from "@/util/layoutHelpers";
import tailwindConfig from "@/tailwind.config";

const fullConfig = resolveConfig(tailwindConfig);
const colors = fullConfig.theme.colors;

export default function ApiDocsPage() {
  const sidebarBackground = colors.blue[800];
  const theme = useTheme();
  const isDark = theme.theme === "dark";

  return (
    <>
      {createHead("API-Documentation")}
      <div>
        <RedocStandalone
          options={{
            disableSearch: true,
            hideHostname: true,
            requiredPropsFirst: true,
            sortRequiredPropsFirst: true,
            theme: {
              rightPanel: {
                backgroundColor: colors.gray[800],
              },
              colors: {
                primary: {
                  main: colors.blue[400],
                },
                text: {
                  primary: isDark ? colors.slate[100] : colors.black,
                },
                success: {
                  main: colors.green[500],
                },
              },
              schema: {
                nestedBackground: "inherit",
              },
              typography: {
                fontFamily: "inherit",
                headings: {
                  fontFamily: "inherit",
                },
              },
              sidebar: {
                backgroundColor: sidebarBackground,
                textColor: colors.slate[100],
              },
            },
          }}
          specUrl="/api-spec.yml"
        />
      </div>
    </>
  );
}
