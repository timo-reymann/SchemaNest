import { Spinner } from "@heroui/spinner";
import React from "react";

export default function LoadingSpinner({ label }: { label: string }) {
  return (
    <div className="flex h-full w-full">
      <Spinner
        className="justify-center mx-auto"
        color="default"
        label={label}
        labelColor="foreground"
        size="lg"
      />
    </div>
  );
}
