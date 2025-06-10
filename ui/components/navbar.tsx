import {
  Navbar as HeroUINavbar,
  NavbarBrand,
  NavbarContent,
  NavbarItem,
} from "@heroui/navbar";
import { Link } from "@heroui/link";
import NextLink from "next/link";
import clsx from "clsx";
import { button as buttonStyles, link as linkStyles } from "@heroui/theme";

import { siteConfig } from "@/config/site";
import { ThemeSwitch } from "@/components/theme-switch";
import { GithubIcon, Logo } from "@/components/icons";

export const Navbar = () => {
  return (
    <HeroUINavbar maxWidth="2xl" position="static">
      <NavbarContent className="basis-full" justify="start">
        <NavbarBrand as="li" className="gap-3 max-w-fit">
          <NextLink className="flex justify-start items-center gap-1" href="/">
            <Logo />
            <p className="font-bold text-inherit">SchemaNest</p>
          </NextLink>
        </NavbarBrand>
        <ul className="lg:flex gap-4 justify-start ml-2">
          <NavbarItem key="/schemas">
            <NextLink
              className={clsx(
                linkStyles({ color: "foreground" }),
                "data-[active=true]:text-primary data-[active=true]:font-medium",
              )}
              color="foreground"
              href="/schemas"
            >
              Overview
            </NextLink>
          </NavbarItem>
        </ul>
      </NavbarContent>

      <NavbarContent className="sm:flex basis-1/5 sm:basis-full" justify="end">
        <NavbarItem className="hidden sm:flex gap-2">
          <ThemeSwitch />
        </NavbarItem>
        <Link
          isExternal
          className={buttonStyles({ variant: "bordered", radius: "full" })}
          href={siteConfig.links.github}
        >
          <GithubIcon size={20} />
          GitHub
        </Link>
      </NavbarContent>
    </HeroUINavbar>
  );
};
