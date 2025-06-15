"use client";
import { Card, CardBody, CardFooter, CardHeader } from "@heroui/card";
import { Code } from "@heroui/code";
import { Pagination } from "@heroui/pagination";
import React, { useEffect, useState } from "react";
import { useRouter } from "next/router";
import { Input } from "@heroui/input";
import { Popover, PopoverContent, PopoverTrigger } from "@heroui/popover";
import Link from "next/link";

import { createHead } from "@/util/layoutHelpers";
import { JsonSchemaInfo } from "@/store/jsonSchemas.slice";
import { useBoundStore } from "@/store/main";
import { HelpIcon, JsonSchemaIcon, SearchIcon } from "@/components/icons";
import LoadingSpinner from "@/components/LoadingSpinner";

const MAX_PER_PAGE = 25;

function SchemaCard({ schema }: Readonly<{ schema: JsonSchemaInfo }>) {
  const [showHelp, setShowHelp] = useState(false);

  return (
    <Card className="max-w[340px]">
      <CardHeader className="justify-between">
        <div className="flex gap-5">
          <JsonSchemaIcon size={56} />
          <div className="flex flex-col gap-1 items-start justify-center">
            <h4 className="text-small font-semibold leading-none text-default-600">
              {schema.identifier}
            </h4>
            <h5 className="text-small tracking-tight text-default-400">
              v{schema.latestVersion.major}.{schema.latestVersion.minor}.
              {schema.latestVersion.patch}
            </h5>
          </div>
        </div>
      </CardHeader>
      <CardBody className="px-3 py-0 text-small text-default-400">
        {schema.description ? (
          <p>{schema.description}</p>
        ) : (
          <div className="flex gap-3">
            <i>No description available</i>
            <Popover
              showArrow
              classNames={{
                base: ["before:bg-default-200"],
                content: ["py-3 px-4 border border-default-200 max-w-80"],
              }}
              isOpen={showHelp}
              placement="bottom-start"
              onOpenChange={(open) => setShowHelp(open)}
            >
              <PopoverTrigger>
                <HelpIcon
                  onClick={(e) => {
                    e.preventDefault();
                    e.stopPropagation();
                    setShowHelp(!showHelp);
                  }}
                  onMouseOut={() => setShowHelp(false)}
                  onMouseOver={() => setShowHelp(true)}
                />
              </PopoverTrigger>
              <PopoverContent>
                <div className="px-1 py-2">
                  <div className="font-bold">
                    No Description Provided by Schema
                  </div>
                  <div className="text-small">
                    The latest schema upload does not contain a{" "}
                    <Code>description</Code> property.
                  </div>
                </div>
              </PopoverContent>
            </Popover>
          </div>
        )}
      </CardBody>
      <CardFooter className="flex justify-between">
        <div className="flex gap-1">
          <Code color="primary">JSON-Schema</Code>
        </div>
      </CardFooter>
    </Card>
  );
}

export default function SchemasPage() {
  const router = useRouter();
  const { page: pageQuery } = router.query as { page: string | undefined };

  const fetchSchemas = useBoundStore((s) => s.fetchSchemas);
  const filterAndPaginateSchemas = useBoundStore(
    (s) => s.filterAndPaginateSchemas,
  );
  const schemasLoading = useBoundStore((s) => s.schemaLoading);

  const [results, setResults] = useState<JsonSchemaInfo[]>([]);
  const [pageCount, setPageCount] = useState(1);
  const [page, setPage] = useState(1);
  const [initialPage, setInitialPage] = useState(1);
  const [search, setSearch] = useState("");

  useEffect(() => {
    if (pageQuery) {
      const parsedPag = Number.parseInt(pageQuery);

      setInitialPage(parsedPag);
      setPage(parsedPag);
    }
  }, [pageQuery]);

  useEffect(() => {
    (async () => {
      await fetchSchemas();
    })();
  }, []);

  useEffect(() => {
    let { slice: schemasToShow, totalPages } = filterAndPaginateSchemas(
      search,
      page,
      MAX_PER_PAGE,
    );

    setPageCount(totalPages);
    setResults(schemasToShow);
  }, [schemasLoading, page, search]);

  const updatePage = (pageNum: number) => {
    const newQuery = { ...router.query, page: pageNum };
    const url = {
      pathname: router.pathname,
      query: newQuery,
    };

    router.push(url, undefined, { shallow: true });
    setPage(pageNum);
  };

  const updateSearch = (search: string) => {
    setSearch(search);
    updatePage(1);
  };

  if (schemasLoading) {
    return <LoadingSpinner label="Loading schemas ..." />;
  }

  return (
    <>
      {createHead("Schema Overview")}
      <div className="flex justify-center">
        <Input
          aria-label="Search"
          className="w-96 mb-8"
          labelPlacement="outside"
          placeholder="Search..."
          startContent={
            <SearchIcon className="text-base text-default-400 pointer-events-none flex-shrink-0" />
          }
          type="search"
          onValueChange={updateSearch}
        />
      </div>
      <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 2xl:grid-cols-5 gap-4 p-4">
        {results.length === 0 && (
          <div className="col col-span-full text-center">No results</div>
        )}
        {results.map((s, idx) => (
          <Link
            key={idx}
            href={{
              pathname: "/schemas/json-schema/[identifier]/[version]",
              query: {
                identifier: s.identifier,
                version: `${s.latestVersion.major}.${s.latestVersion.minor}.${s.latestVersion.patch}`,
              },
            }}
          >
            <SchemaCard schema={s} />
          </Link>
        ))}
      </div>

      {results.length !== 0 && (
        <div className="flex justify-center p-4">
          <Pagination
            initialPage={initialPage}
            page={page}
            total={pageCount}
            onChange={updatePage}
          />
        </div>
      )}
    </>
  );
}
