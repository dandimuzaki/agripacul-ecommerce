export const addressKeys = {
  all: ["address"] as const,

  lists: () => [...addressKeys.all, "lists"] as const,

  province: ["province"] as const,

  regency: (id: number) => ["regency", id] as const,

  district: (id: number) => ["district", id] as const,

  subdistrict: (id: number) => ["subidstrict", id] as const,

  detail: (id: number) =>
    [...addressKeys.all, "detail", id] as const
}