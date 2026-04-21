"use client";

import { useMutation, useQueryClient } from "@tanstack/react-query"
import { addressKeys } from "../queries/addressKeys";
import { addressService } from "@/services/address.service";
import { AddressFormValues } from "@/schemas/address.schema";
import { toast } from "sonner";
import { AxiosError } from "axios";

export const useCreateAddress = (options?: {
  onSuccess?: () => void
}) => {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: (payload: AddressFormValues) =>
      addressService.createAddress(payload),

    onSuccess: () => {
      toast.success("Address created successfully")

      // Refresh address list cache
      queryClient.invalidateQueries({
        queryKey: addressKeys.lists()
      })

      options?.onSuccess?.()
    },

    onError: (error: AxiosError<{ message: string }>) => {
      toast.error(error.response?.data?.message || "Failed to create address")
    }
  })
}