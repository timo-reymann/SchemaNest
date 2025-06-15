import {
  Navbar as HeroUINavbar,
  NavbarBrand,
  NavbarContent,
  NavbarItem,
} from "@heroui/navbar";
import { Link } from "@heroui/link";
import { button as buttonStyles } from "@heroui/theme";
import { usePathname } from "next/navigation";
import NextLink from "next/link";

import { siteConfig } from "@/config/site";
import { ThemeSwitch } from "@/components/theme-switch";
import { GithubIcon, Logo } from "@/components/icons";

const navItems = [
  {
    label: "Overview",
    href: "/schemas",
  },
  {
    label: "API Documentation",
    href: "/api-docs",
  },
];

export const Navbar = () => {
  const pathName = usePathname();

  return (
    <HeroUINavbar maxWidth="2xl" position="static">
      <NavbarContent className="basis-full" justify="start">
        <NavbarBrand as="li" className="gap-3 max-w-fit">
          <NextLink
            className="flex justify-start items-center gap-1"
            color="foreground"
            href="/"
          >
            <Logo />
            <p className="font-bold text-inherit">SchemaNest</p>
          </NextLink>
        </NavbarBrand>
        <ul className="flex gap-4 justify-start ml-2">
          {navItems.map((nav) => (
            <NavbarItem key={nav.href} isActive={nav.href === pathName}>
              <Link
                color={nav.href === pathName ? "primary" : "foreground"}
                href={nav.href}
              >
                {nav.label}
              </Link>
            </NavbarItem>
          ))}
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
