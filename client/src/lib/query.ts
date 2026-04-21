import { AppRouterInstance } from "next/dist/shared/lib/app-router-context.shared-runtime";

export function updateQuery(
  router: AppRouterInstance,
  key: string,
  value: string
) {
  const params = new URLSearchParams(window.location.search);

  if (value) {
    params.set(key, value);
  } else {
    params.delete(key);
  }

  router.push(`?${params.toString()}`);
}