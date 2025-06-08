import {Navbar as HeroUINavbar, NavbarBrand, NavbarContent, NavbarItem,} from "@heroui/navbar";
import {Kbd} from "@heroui/kbd";
import {Link} from "@heroui/link";
import {Input} from "@heroui/input";
import NextLink from "next/link";
import clsx from "clsx";
import {siteConfig} from "@/config/site";
import {ThemeSwitch} from "@/components/theme-switch";
import {GithubIcon, Logo, SearchIcon,} from "@/components/icons";
import {button as buttonStyles} from "@heroui/theme";
import { link as linkStyles } from "@heroui/theme";

export const Navbar = () => {
    const searchInput = (
        <Input
            aria-label="Search"
            classNames={{
                inputWrapper: "bg-default-100",
                input: "text-sm",
            }}
            endContent={
                <Kbd className="hidden lg:inline-block" keys={["command"]}>
                    K
                </Kbd>
            }
            labelPlacement="outside"
            placeholder="Search..."
            startContent={
                <SearchIcon className="text-base text-default-400 pointer-events-none flex-shrink-0"/>
            }
            type="search"
        />
    );

    return (
        <HeroUINavbar maxWidth="2xl" position="static">
            <NavbarContent className="basis-full" justify="start">
                <NavbarBrand as="li" className="gap-3 max-w-fit">
                    <NextLink className="flex justify-start items-center gap-1" href="/">
                        <Logo/>
                        <p className="font-bold text-inherit">SchemaNest</p>
                    </NextLink>
                </NavbarBrand>
                <ul className="hidden lg:flex gap-4 justify-start ml-2">

                        <NavbarItem key="/schemas">
                            <NextLink
                                className={clsx(
                                    linkStyles({ color: "foreground" }),
                                    "data-[active=true]:text-primary data-[active=true]:font-medium",
                                )}
                                color="foreground"
                                href="/schemas"
                            >
                                Schemas
                            </NextLink>
                        </NavbarItem>

                </ul>
            </NavbarContent>

            <NavbarContent
                className="sm:flex basis-1/5 sm:basis-full" justify="end">
                <NavbarItem className="hidden sm:flex gap-2">
                    <ThemeSwitch/>
                </NavbarItem>
                <Link isExternal
                      className={buttonStyles({variant: "bordered", radius: "full"})}
                      href={siteConfig.links.github}>
                    <GithubIcon size={20}/>
                    GitHub
                </Link>
            </NavbarContent>

        </HeroUINavbar>
    );
};
