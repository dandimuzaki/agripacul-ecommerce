"use client";

import { useEffect, useMemo } from "react";
import { useForm, useWatch } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { useCheckoutPreview } from "./useCheckoutPreview";
import { ShippingOption, ShippingOptionTemp } from "@/types/checkout";
import { CheckoutFormValuesTemp, checkoutSchemaTemp } from "@/schemas/checkout.schema";
import { useQueryShipping } from "./useQueryShipping";
import { useAddress } from "../address/useAddress";

export const useCheckoutForm = () => {
  // Get default shipping address
  const { defaultShippingAddress } = useAddress()

  const form = useForm<CheckoutFormValuesTemp>({
    resolver: zodResolver(checkoutSchemaTemp),
  })

  const shippingAddress = useWatch({
    control: form.control,
    name: "shipping_address_id",
  });

  const shippingAddressId = shippingAddress ? Number(shippingAddress) : undefined

  const paymentMethodId = useWatch({
    control: form.control,
    name: "selected_payment_method_id",
  });

  const selectedShippingOptionId = useWatch({
    control: form.control,
    name: "selected_shipping_option_id",
  });

  // Get shipping option list
  const { data: shippingOptions, isLoading } = useQueryShipping({
    shipping_address_id: shippingAddressId
  })

  const normalizeShippingOption = (options: ShippingOption[]): ShippingOptionTemp[] => 
    options.map((o) => ({
      name: o.name,
      service: o.service,
      description: o.description,
      code: o.code,
      cost: o.cost,
      etd: o.etd,
      id: `${o.code}-${o.service}`
    })).sort((a,b) => a.cost - b.cost)

  const options = useMemo(() => 
    normalizeShippingOption(shippingOptions || []),
    [shippingOptions]
  )
  
  const selectedShippingOption = options.find(o => o.id === selectedShippingOptionId);

  // Get cheapest shipping option
  const cheapestOption = options[0]

  // set default address ONCE
  useEffect(() => {
    if (!shippingAddressId && defaultShippingAddress) {
      form.setValue("shipping_address_id", defaultShippingAddress.id);
    }
  }, [defaultShippingAddress, shippingAddressId, form]);

  // set cheapest shipping when options ready
  useEffect(() => {
    if (
      !selectedShippingOptionId && cheapestOption
    ) {
      form.setValue("selected_shipping_option", cheapestOption);
      form.setValue("selected_shipping_option_id", cheapestOption.id);
    }
  }, [form, cheapestOption, selectedShippingOptionId]);

  const { data } = useCheckoutPreview({
    shipping_address_id: shippingAddressId,
    selected_shipping_option_id: selectedShippingOptionId,
    selected_payment_method_id: paymentMethodId,
  }, selectedShippingOption);

  console.log("paymentMethodId", paymentMethodId)

  return { checkout: data, isLoading, form, shippingAddressId, options };
};