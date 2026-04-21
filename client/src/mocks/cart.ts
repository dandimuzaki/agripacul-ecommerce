export const mockCart = {
  id: 1,
  customer_id: 7,
  items: [
    {
      id: 1,
      sku_id: 101,
      sku_code: "VEG-spinach-001",
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
      quantity: 2,
      stock: 120,
      subtotal: 30000,
      is_selected: true,
      is_available: true
    },

    {
      id: 2,
      sku_id: 102,
      sku_code: "VEG-carrot-002",
      product: {
        id: 12,
        name: "Fresh Sweet Carrots",
        slug: "fresh-sweet-carrots",
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
      quantity: 1,
      stock: 60,
      subtotal: 22000,
      is_selected: true,
      is_available: true
    },

    {
      id: 3,
      sku_id: 103,
      sku_code: "VEG-tomato-003",
      product: {
        id: 13,
        name: "Fresh Red Tomatoes",
        slug: "fresh-red-tomatoes",
        main_image_url:
          "https://images.unsplash.com/photo-1546094096-0df4bcaaa337"
      },
      variants: [
        {
          name: "Weight",
          value: "500g"
        },
        {
          name: "Type",
          value: "Cherry Tomato"
        }
      ],
      price: {
        unit_price: 18000,
        sale_price: 16000,
        discount_percentage: 11
      },
      quantity: 3,
      stock: 80,
      subtotal: 54000,
      is_selected: false,
      is_available: true
    }
  ],

  summary: {
    total_items: 6,
    total_price: 106000,
    total_selected_items: 3,
    total_selected_price: 52000
  }
};