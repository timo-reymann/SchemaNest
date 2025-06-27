export const useBuildApiLink = () => {
  const hostname = window.location.origin;

  return (path: string) => `${hostname}/api${path}`;
};
