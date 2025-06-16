import { Form } from "@heroui/form";
import { Input } from "@heroui/input";
import { Button } from "@heroui/button";
import { FormEvent, useEffect, useState } from "react";
import { addToast } from "@heroui/toast";
import { useRouter } from "next/navigation";

import { useBoundStore } from "@/store/main";
import LoadingSpinner from "@/components/LoadingSpinner";
import { createHead } from "@/util/layoutHelpers";

const semverRegex = /^(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)$/;

function validateVersion(value: string): string | null {
  if (value.trim().length === 0) {
    return "Please enter a version";
  }
  if (!semverRegex.test(value)) {
    return "Please enter a valid semantic version";
  }

  return null;
}

export default function Upload() {
  const router = useRouter();
  const loadConfig = useBoundStore((s) => s.loadConfig);
  const configLoading = useBoundStore((s) => s.configLoading);
  const config = useBoundStore((s) => s.config);
  const uploadSchema = useBoundStore((s) => s.uploadSchema);
  const [status, setStatus] = useState<"uploading" | "error" | "success">();

  useEffect(() => {
    loadConfig();
  }, []);

  if (configLoading) {
    return <LoadingSpinner label="Loading configuration ..." />;
  }

  const upload = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    setStatus("uploading");
    let data = Object.fromEntries(new FormData(e.currentTarget));

    try {
      // @ts-ignore
      await uploadSchema(data);
      setStatus("success");
      addToast({
        title: "Successfully uploaded schema",
        description: "The schema has been uploaded.",
        color: "success",
        timeout: 3_000,
      });
      router.push("/schemas");
    } catch (e) {
      addToast({
        title: "Failed to upload schema",
        description: e?.toString() || "An unknown error occurred.",
        color: "danger",
        timeout: 10_000,
      });
      setStatus("error");
    }
  };

  return (
    <>
      {createHead("Upload")}
      <Form className="w-full" onSubmit={upload}>
        {config?.apiKeyAuthEnabled && (
          <Input
            isRequired
            errorMessage="Please enter a API key"
            label="API-Key"
            labelPlacement="outside"
            name="apiKey"
            placeholder="API-Key for uploading"
            type="password"
          />
        )}
        <Input
          isRequired
          errorMessage="Please enter a valid identifier"
          label="Identifier"
          labelPlacement="outside"
          name="identifier"
          placeholder="Identifier for the schema"
          type="text"
        />
        <Input
          isRequired
          label="Version"
          labelPlacement="outside"
          name="version"
          placeholder="x.y.z"
          type="text"
          validate={validateVersion}
        />
        <Input
          isRequired
          errorMessage="Schema file is required for upload"
          label="Schema file"
          labelPlacement="outside"
          name="schema"
          placeholder="Schema file"
          type="file"
        />
        <Button
          color="primary"
          isLoading={status === "uploading"}
          type="submit"
        >
          Upload
        </Button>
      </Form>
    </>
  );
}
