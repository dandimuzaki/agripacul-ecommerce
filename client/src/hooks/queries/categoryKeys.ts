export const categoryKeys = {
  all: ["category"] as const,

  detail: (id: number) =>
    [...categoryKeys.all, "detail", id] as const
}