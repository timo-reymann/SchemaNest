"use client";
import { Card, CardBody, CardFooter, CardHeader } from "@heroui/card";
import { Code } from "@heroui/code";
import { Pagination } from "@heroui/pagination";
import React, { useEffect, useState } from "react";
import NextLink from "next/link";
import { useRouter } from "next/navigation";
import { Input } from "@heroui/input";

import { createHead } from "@/util/layoutHelpers";
import { JsonSchemaInfo } from "@/store/jsonSchemas.slice";
import { useBoundStore } from "@/store/main";
import { JsonSchemaIcon, SearchIcon } from "@/components/icons";
import LoadingSpinner from "@/components/LoadingSpinner";

const MAX_PER_PAGE = 25;

export default function SchemasPage() {
  const router = useRouter();
  const fetchSchemas = useBoundStore((s) => s.fetchSchemas);
  const filterAndPaginateSchemas = useBoundStore(
    (s) => s.filterAndPaginateSchemas,
  );
  const schemas = useBoundStore((s) => s.schemas);
  const schemasLoading = useBoundStore((s) => s.schemaLoading);

  const [results, setResults] = useState<JsonSchemaInfo[]>([]);
  const [pageCount, setPageCount] = useState(1);
  const [page, setPage] = useState(1);
  const [initialPage, setInitialPage] = useState(1);
  const [search, setSearch] = useState("");

  useEffect(() => {
    const params = new URLSearchParams(window.location.search);
    const page = params.get("page");

    if (page) {
      const parsedPag = Number.parseInt(page);

      setInitialPage(parsedPag);
      setPage(parsedPag);
    }
  }, []);

  useEffect(() => {
    fetchSchemas();
  }, [fetchSchemas]);

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
    console.log("update page");
    const params = new URLSearchParams(window.location.search);

    params.set("page", pageNum.toString());
    const newUrl = `${window.location.pathname}?${params.toString()}`;

    router.replace(newUrl, { scroll: false });

    setPage(pageNum);
  };

  const updateSearch = (search: string) => {
    setSearch(search);
    setPage(1);
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
          <>
            <div className="col col-span-full text-center">No results</div>
          </>
        )}
        {results.map((s, idx) => (
          <NextLink
            key={idx}
            href={`/schemas/json-schema/${s.identifier}/${s.latestVersion.major}.${s.latestVersion.minor}.${s.latestVersion.patch}`}
          >
            <Card className="max-w[340px]">
              <CardHeader className="justify-between">
                <div className="flex gap-5">
                  <JsonSchemaIcon size={56} />
                  <div className="flex flex-col gap-1 items-start justify-center">
                    <h4 className="text-small font-semibold leading-none text-default-600">
                      {s.identifier}
                    </h4>
                    <h5 className="text-small tracking-tight text-default-400">
                      v{s.latestVersion.major}.{s.latestVersion.minor}.
                      {s.latestVersion.patch}
                    </h5>
                  </div>
                </div>
              </CardHeader>
              <CardBody className="px-3 py-0 text-small text-default-400">
                <p>No description provided</p>
              </CardBody>
              <CardFooter className="flex justify-between">
                <div className="flex gap-1">
                  <Code color="primary">JSON-Schema</Code>
                </div>
              </CardFooter>
            </Card>
          </NextLink>
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
