export async function getJSON<T>(path: string): Promise<T> {
  const response = await fetch(path);
  const json = response.json();

  return json as T;
}

export async function postJSON(
  path: string,
  body: Object,
  headers: HeadersInit | undefined,
): Promise<Response> {
  return await fetch(path, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      ...(headers || {}),
    },
    body: JSON.stringify(body),
  });
}
