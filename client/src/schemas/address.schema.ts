import z from "zod";

export const addressSchema = z.object({
  recipient_name: z.string().min(3, "Recipient name is required"),
  label: z.string().min(3, "Label is required"),
  phone_number: z.string().min(3, "Phone number name is required"),
  province_id: z.number().optional(),
  regency_id: z.number().optional(),
  district_id: z.number().optional(),
  subdistrict_id: z.number().optional(),
  detail_address: z.string().optional(),
  postal_code: z.string().optional()
});

export type AddressFormValues = z.infer<typeof addressSchema>;