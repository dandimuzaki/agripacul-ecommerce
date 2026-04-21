import { CheckoutPreviewResponse } from "@/types/checkout";

export const mockCheckoutPreview: CheckoutPreviewResponse = {
    selected_items: [
      {
        id: 1,
        sku_id: 101,
        sku_code: "VEG-spinach-500g-org",
        product: {
          id: 11,
          name: "Fresh Organic Spinach",
          slug: "fresh-organic-spinach",
          main_image_url:
            "https://images.unsplash.com/photo-1576045057995-568f588f82fb"
        },
        variants: [
          {
            name: "Weight",
            value: "500g"
          },
          {
            name: "Farming Method",
            value: "Organic"
          }
        ],
        price: {
          unit_price: 15000,
          sale_price: 12000,
          discount_percentage: 20
        },
        quantity: 3,
        stock: 80,
        subtotal: 36000
      },
      {
        id: 2,
        sku_id: 102,
        sku_code: "VEG-carrot-1kg-prem",
        product: {
          id: 12,
          name: "Sweet Farm Carrots",
          slug: "sweet-farm-carrots",
          main_image_url:
            "https://images.unsplash.com/photo-1447175008436-1701707537b7"
        },
        variants: [
          {
            name: "Weight",
            value: "1kg"
          },
          {
            name: "Grade",
            value: "Premium"
          }
        ],
        price: {
          unit_price: 22000,
          sale_price: 20000,
          discount_percentage: 9
        },
        quantity: 2,
        stock: 60,
        subtotal: 40000
      }
    ],

    shipping_option_list: [
      {
        name: "JNE",
        code: "jne",
        service: "REG",
        description: "Regular Service",
        cost: 15000,
        etd: "2-3 days"
      },
      {
        name: "J&T Express",
        code: "jnt",
        service: "EZ",
        description: "Reguler",
        cost: 18000,
        etd: "2-4 days"
      },
      {
        name: "SiCepat",
        code: "sicepat",
        service: "BEST",
        description: "Best Service",
        cost: 20000,
        etd: "1-2 days"
      }
    ],

    promotion_list: [
      {
        id: 1,
        name: "Welcome Voucher",
        end_date: "2026-06-01T00:00:00+07:00",
        discount_type: "amount",
        discount_value: 10000,
        minimum_order_value: 50000,
        maximum_discount: 10000,
        usage_left: 100,
        voucher_code: "WELCOME10"
      },
      {
        id: 2,
        name: "Farm Fresh Sale",
        end_date: "2026-05-01T00:00:00+07:00",
        discount_type: "percentage",
        discount_value: 15,
        minimum_order_value: 100000,
        maximum_discount: 30000,
        usage_left: 200,
        voucher_code: null
      }
    ],

    totals: {
      subtotal: 76000,
      discount_amount: 10000,
      shipping_cost: 0,
      grand_total: 66000
    }
};