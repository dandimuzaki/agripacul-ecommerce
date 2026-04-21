import z from "zod";

export const loginSchema = z.object({
  email: z.string(),
  password: z.string(),
})

export type LoginFormValues = z.input<typeof loginSchema>;

export const registerSchema = z.object({
  full_name: z.string(),
  email: z.string(),
  password: z.string(),
})

export type RegisterFormValues = z.input<typeof registerSchema>;