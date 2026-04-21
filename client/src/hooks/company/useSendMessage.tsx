"use client";

import { useMutation, useQueryClient } from "@tanstack/react-query"
import { contactKeys } from "../queries/contactKeys";
import { contactService } from "@/services/contact.service";
import { ContactFormValues, contactSchema } from "@/schemas/contact.schema";
import { toast } from "sonner";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";

export const useSendMessage = () => {
  const queryClient = useQueryClient()

  const form = useForm<ContactFormValues>({
    resolver: zodResolver(contactSchema),
    defaultValues: {}
  })

  const mutation = useMutation({
    mutationFn: (payload: ContactFormValues) =>
      contactService.sendMessage(payload),

    onSuccess: () => {
      // refresh message list cache
      queryClient.invalidateQueries({
        queryKey: contactKeys.all
      })

      toast.success("Message sent successfully")
    },

    onError: (error: any) => {
      toast.error(error.message || "Failed to create message")
    }
  })

  return {form, ...mutation}
}