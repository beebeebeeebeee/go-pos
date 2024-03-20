type IItem = {
  id: string;
  name: string;
  price: number;
};

type IOrder = {
  id: string;
  orderItems: OrderItem[];
  totalPrice: number;
};

type IOrderItem = {
  id: string;
  item: Item;
  quantity: number;
  totalPrice: number;
};

type IPrinter = {
  name: string;
  isDefault: boolean;
  selected: boolean;
  options: Record<string, string>;
};

declare module "@wails/models" {
  namespace entity {
    type Item = IItem;
    type Order = IOrder;
    type OrderItem = IOrderItem;
    type Printer = IPrinter;
  }
}
