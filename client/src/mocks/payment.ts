import { PaymentMethod } from "@/types/payment";

export const mockPaymentMethods: PaymentMethod[] = [
      {
        id: 1,
        name: "QRIS",
        is_active: true,
        icon_url: "/icons/qris.png"
      },
      {
        id: 2,
        name: "GoPay",
        is_active: true,
        icon_url: "/icons/gopay.png"
      },
      {
        id: 3,
        name: "OVO",
        is_active: true,
        icon_url: "/icons/ovo.png"
      },
      {
        id: 4,
        name: "DANA",
        is_active: true,
        icon_url: "/icons/dana.png"
      },
      {
        id: 5,
        name: "BCA Virtual Account",
        is_active: true,
        icon_url: "/icons/bca.png"
      },
      {
        id: 6,
        name: "Mandiri Virtual Account",
        is_active: true,
        icon_url: "/icons/mandiri.png"
      },
      {
        id: 7,
        name: "COD (Cash on Delivery)",
        is_active: true,
        icon_url: "/icons/cod.png"
      }
    ];