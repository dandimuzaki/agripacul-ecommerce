import { skuDetails, skuStatus } from "@/types/sku";

export const skus: skuDetails[] = [
  {
            "id": 90,
            "product_id": 28,
            "sku_code": "LOOF-KIT-3PC-512-SzojI",
            "price": 20000,
            "stock": 300,
            "min_stock": 40,
            "status": "active" as skuStatus,
            "weight": 60,
            "variants": [
                {
                    "name": "Pack Size",
                    "value": "Starter Pack - 3 pieces"
                }
            ],
            "images": null
        },
        {
            "id": 91,
            "product_id": 28,
            "sku_code": "LOOF-KIT-5PC-122-Dap5P",
            "price": 30000,
            "stock": 250,
            "min_stock": 30,
            "status": "active" as skuStatus,
            "weight": 100,
            "variants": [
                {
                    "name": "Pack Size",
                    "value": "Family Pack - 5 pieces"
                }
            ],
            "images": null
        },
        {
            "id": 92,
            "product_id": 28,
            "sku_code": "LOOF-KIT-10PC-311-oEhFi",
            "price": 55000,
            "stock": 150,
            "min_stock": 20,
            "status": "active" as skuStatus,
            "weight": 200,
            "variants": [
                {
                    "name": "Pack Size",
                    "value": "Stock Up Pack - 10 pieces"
                }
            ],
            "images": null
        }
]

export const sku = {
    "id": 90,
    "product": {
        "id": 1,
        "name": "Tomato Cherry"
    },
    "sku_code": "LOOF-KIT-3PC-512-SzojI",
    "price": 20000,
    "stock": 300,
    "min_stock": 40,
    "status": "active" as skuStatus,
    "weight": 60,
    "variants": [
        {
            "name": "Pack Size",
            "value": "Starter Pack - 3 pieces"
        }
    ],
    "images": null
}