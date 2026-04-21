export const inventoryKeys = {
  all: ["inventory"] as const,

  lists: () => [...inventoryKeys.all, "lists"] as const,

  list: (query: any) =>
    [...inventoryKeys.lists(), query] as const,

  detail: (id: number) =>
    [...inventoryKeys.all, "detail", id] as const,

  log: (id: number) =>
    [...inventoryKeys.all, "log", id] as const
}