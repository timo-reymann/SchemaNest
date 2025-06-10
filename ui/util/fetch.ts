export async function getJSON<T>(path: string): Promise<T> {
  const response = await fetch(path);
  const json = response.json();

  return json as T;
}
