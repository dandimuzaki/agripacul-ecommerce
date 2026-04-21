import { api } from "@/lib/api";
import { ContactFormValues } from "@/schemas/contact.schema";

export const contactService = {
  sendMessage(data: ContactFormValues) {
    return api.post("/contact", data)
  }
}