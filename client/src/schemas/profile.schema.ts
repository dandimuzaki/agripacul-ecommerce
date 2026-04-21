import z from "zod";

export const profileSchema = z.object({
  full_name: z.string().optional(),
  date_of_birth: z.coerce.date(),
  phone_number: z.string().optional(),
  profile_image: z.file().optional(),
});

export type ProfileFormValues = z.input<typeof profileSchema>;