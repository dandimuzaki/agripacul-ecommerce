import z from "zod";

export const categorySchema = z.object({
  name: z.string().min(3, "Name is required"),
  icon: z.file().optional(),
});

export type CategoryFormValues = z.infer<typeof categorySchema>;