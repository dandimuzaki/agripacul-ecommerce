import z from "zod";

export const contactSchema = z.object({
  first_name: z.string().min(3, "First name is required"),
  last_name: z.string().min(3, "Last name is required"),
  email: z.string().min(3, "Email is required"),
  phone_number: z.string().min(3, "Phone number is required"),
  subject: z.string().min(3, "Subject is required"),
  body: z.string().min(3, "Message is required"),
});

export type ContactFormValues = z.infer<typeof contactSchema>;