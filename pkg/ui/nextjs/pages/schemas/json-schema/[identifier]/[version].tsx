"use client";
import {createHead} from "@/util/layoutHelpers";
import {useParams} from "next/navigation";
import {subtitle, title} from "@/components/primitives";
import {useBoundStore} from "@/store/main";
import React, {useEffect, useState} from "react";
import LoadingSpinner from "@/components/LoadingSpinner";
import {Editor} from "@monaco-editor/react";
import {useTheme} from "next-themes";
import {Code} from "@heroui/code";

export default function JsonSchemaVersionPage() {
    const params = useParams<{ identifier: string, version: string }>()

    const [schema, setSchema] = useState({})
    const [exists, setExists] = useState(false)
    const fetchSchemaByVersion = useBoundStore(s => s.fetchSchemaByVersion)
    const loading = useBoundStore(s => s.schemaLoading)
    const {theme} = useTheme()

    useEffect(() => {
        if (params == null) {
            return
        }
        (async () => {
            try {
                let schema = await fetchSchemaByVersion(params.identifier, params.version)
                setSchema(schema)
                setExists(true)
            } catch {
                setExists(false)
            }
        })()
    }, [fetchSchemaByVersion, params]);

    if (params == null) {
        return <></>
    }

    if (loading) {
        return <LoadingSpinner label="Loading schema ..."/>
    }

    if (!exists) {
        return <div className="flex h-full w-full">
            <p className=" mx-auto text-center">
                JSON-Schema <Code>{params.identifier}</Code> in version <Code>{params.version}</Code> not found.
            </p>
        </div>
    }

    return (
        <>
            {createHead(`${params.identifier} - JSON Schema`)}
            <h1 className={title()}>{params.identifier}</h1>
            <h2 className={subtitle()}>{params.version}</h2>
            <Editor height="50vh"
                    options={{
                        readOnly: true
                    }}
                    theme={theme === "dark" ? "vs-dark" : "vs-light"}
                    defaultLanguage="javascript"
                    value={JSON.stringify(schema, null, "  ")}/>
        </>
    );
}
