export interface PaymentMethod {
  id: number;
  name: string;
  is_active: boolean;
  icon_url: string;
}

export interface PaymentType {
  id: number,
  name: string,
  methods: PaymentMethod[]
}