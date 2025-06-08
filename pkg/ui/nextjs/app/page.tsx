import {subtitle, title} from "@/components/primitives";
import {Link} from "@heroui/link";
import {button as buttonStyles} from "@heroui/theme";
import {siteConfig} from "@/config/site";
import {Logo} from "@/components/icons";

export default function Home() {
    return (
        <section className="flex flex-col items-center justify-center gap-4 py-8 md:py-10">
            <Logo size={300} />
            <div className="inline-block max-w-3xl text-center justify-center">
                <div>
                    <span className={title()}>
                    Where schemas&nbsp;
                </span>
                    <span className={title({color: "violet"})}>
                    grow,&nbsp;
                </span>
                    <span className={title({color: "yellow"})}>
                     thrive,&nbsp;<br />
                </span>
                    <span className={title()}>
                    and&nbsp;
                </span>
                    <span className={title({color: "blue"})}>
                    scale&nbsp;
                </span>
                    <span className={title()}>
                    with your team.
                </span>
                </div>
                <div className={subtitle({class: "mt-4"})}>
                    Quick, simple and API-first.
                </div>
            </div>

            <div className="flex gap-3">
                <Link isExternal
                      className={buttonStyles({
                          color: "primary",
                          radius: "full",
                          variant: "shadow",
                      })}
                      href={siteConfig.links.apiDocs}>
                    API-Documentation
                </Link>
                <Link className={buttonStyles({
                          color: "default",
                          radius: "full",
                          variant: "bordered",
                      })}
                      href="/schemas">
                    Schemas
                </Link>
            </div>

        </section>
    );
}
