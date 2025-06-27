import {
    Navbar as HeroUINavbar,
    NavbarBrand,
    NavbarContent,
    NavbarItem,
    NavbarMenu,
    NavbarMenuItem,
    NavbarMenuToggle,
} from "@heroui/navbar";
import {Link} from "@heroui/link";
import {button as buttonStyles} from "@heroui/theme";
import {usePathname} from "next/navigation";
import NextLink from "next/link";

import {siteConfig} from "@/config/site";
import {ThemeSwitch} from "@/components/theme-switch";
import {GithubIcon, Logo} from "@/components/icons";
import {useState} from "react";

const navItems = [
    {
        label: "Schemas",
        href: "/schemas",
    },
    {
        label: "Upload",
        href: "/upload",
    },
    {
        label: "API Documentation",
        href: "/api-docs",
    },
];

export const Navbar = () => {
    const pathName = usePathname();
    const [isMenuOpen, setIsMenuOpen] = useState(false)

    return (
        <HeroUINavbar maxWidth="2xl" position="static" isMenuOpen={isMenuOpen} onMenuOpenChange={setIsMenuOpen}>
            <NavbarContent className="basis-full" justify="start">
                <NavbarBrand as="li" className="gap-3 max-w-fit">
                    <NextLink
                        className="flex justify-start items-center gap-1"
                        color="foreground"
                        href="/"
                    >
                        <Logo/>
                        <p className="font-bold text-inherit">SchemaNest</p>
                    </NextLink>
                </NavbarBrand>
                <ul className="flex gap-4 justify-start ml-2 hidden md:flex">
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

            <NavbarContent justify="end">
                <NavbarItem>
                    <ThemeSwitch/>
                </NavbarItem>
                <Link
                    isExternal
                    className={buttonStyles({variant: "bordered", radius: "full"})}
                    href={siteConfig.links.github}
                >
                    <GithubIcon size={20}/>
                    GitHub
                </Link>
            </NavbarContent>
            <NavbarMenuToggle
                aria-label={isMenuOpen ? "Close menu" : "Open menu"}
                className="sm:hidden"
            />
            <NavbarMenu>
                {navItems.map((nav) => (
                    <NavbarMenuItem key={nav.href}
                                    isActive={nav.href === pathName}>
                        <Link
                            onPress={_ => setIsMenuOpen(false)}
                            color={nav.href === pathName ? "primary" : "foreground"}
                            href={nav.href}>
                            {nav.label}
                        </Link>
                    </NavbarMenuItem>
                ))}
            </NavbarMenu>
        </HeroUINavbar>
    );
};
