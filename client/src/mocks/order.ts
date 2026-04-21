export const mockOrder = {
      id: 1001,
      display_status: "delivered",
      steps: [
        {
          key: "created",
          label: "Order Placed",
          done: true,
          at: "2024-03-20T09:15:00Z",
          description: "Your order is successfully created and please proceed to payment"
        },
        {
          key: "paid",
          label: "Payment Confirmed",
          done: true,
          at: "2024-03-20T09:30:00Z",
          description: "Your order payment is confirmed and is being reviewed"
        },
        {
          key: "confirmed",
          label: "Order Confirmed",
          done: true,
          at: "2024-03-20T10:00:00Z",
          description: "Your order is confirmed and is being prepared"
        },
        {
          key: "shipped",
          label: "Order Shipped",
          done: true,
          at: "2024-03-21T08:00:00Z",
          description: "Your order is shipped"
        },
        {
          key: "delivered",
          label: "Order Delivered",
          done: true,
          at: "2024-03-22T14:30:00Z",
          description: "Your order is delivered"
        },
        {
          key: "completed",
          label: "Order Completed",
          done: true,
          at: "2024-03-23T10:00:00Z",
          description: "Your order is completed"
        }
      ],
      cancellation: null,
      created_at: "2024-03-20T09:15:00Z",
      billing_address: {
        recipient_name: "Dian Pramesti",
        label: "Home",
        phone_number: "0815",
        province: "Jawa Barat",
        regency: "Kota Bandung",
        district: "Cihampelas",
        subdistrict: "Cipaganti",
        postal_code: "40115",
        detail_address: "Jl. Cihampelas No. 167, Gang Mawar"
      },
      items: [
        {
          name: "Bamboo Dish Brush Set",
          main_image_url: "https://images.unsplash.com/photo-1584269600464-37b1b58a9fe7?w=400",
          variants: [
            {name: "Set Type", value: "Starter Set (2 brushes)"},
            {name: "Bristle Type", value: "Sisal (Firm)"}
          ],
          quantity: 1,
          price: 45000,
          total_price: 45000
        },
        {
          name: "Stainless Steel Straw Set",
          main_image_url: "https://images.unsplash.com/photo-1584269600464-37b1b58a9fe7?w=400",
          variants: [
            {name: "Set Type", value: "Basic Set (4 straws)"},
            {name: "Straw Type", value: "Mixed (2 straight, 2 bent)"}
          ],
          quantity: 2,
          price: 35000,
          total_price: 70000
        }
      ],
      shipping: {
        name: "JNE",
        code: "jne",
        service: "REG",
        cost: 12000,
        etd: 2
      },
      totals: {
        subtotal: 115000,
        discount_amount: 17250,
        shipping_cost: 12000,
        grand_total: 109750
      },
      notes: "Please leave package at security desk",
      tracking_number: "JNE1234567890"
    }

export const mockOrders = [
    {
      id: 1001,
      customer: {
        name: "Dandi",
        email: "dandimuzaki@gmail.com"
      },
      created_at: "2024-03-20T09:15:00Z",
      updated_at: "2024-03-20T09:15:00Z",
      display_status: "delivered",
      grand_total: 109750,
      first_item: {
          name: "Bamboo Dish Brush Set",
          main_image_url: "https://images.unsplash.com/photo-1584269600464-37b1b58a9fe7?w=400",
          variants: [
            {name: "Set Type", value: "Starter Set (2 brushes)"},
            {name: "Bristle Type", value: "Sisal (Firm)"}
          ],
          quantity: 1,
          price: 45000,
          total_price: 45000
        },
      item_count: 2,
      cancellation: false,
      tracking_number: "JNE1234567890"
    },
    {
      id: 1002,
      customer: {
        name: "Dandi",
        email: "dandimuzaki@gmail.com"
      },
      created_at: "2024-03-25T10:30:00Z",
      updated_at: "2024-03-20T09:15:00Z",
      display_status: "shipped",
      grand_total: 109750,
      first_item: {
          name: "Organic Cotton Tote Bag",
          main_image_url: "https://images.unsplash.com/photo-1544716278-ca5e3f4abd8c?w=400",
          variants: [
            {name: "Size", value: "Medium (35x40cm)"},
            {name: "Color", value: "Natural"}
          ],
          quantity: 1,
          price: 45000,
          total_price: 45000
        },
      item_count: 2,
      cancellation: false,
      tracking_number: "JNE1234567890"
    }
  ]