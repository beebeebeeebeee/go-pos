type IItemId = string;
type IOderId = string;
type IOderItemId = string;

type IItem = {
  id: IItemId;
  name: string;
  price: number;
};

type IOrder = {
  id: IOderId;
  orderItems: OrderItem[];
  totalPrice: number;
};

type IOrderItem = {
  id: IOderItemId;
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
    type ItemID = IItemId;
    type OrderID = IOderId;
    type OrderItemID = IOderItemId;
    type Item = IItem;
    type Order = IOrder;
    type OrderItem = IOrderItem;
    type Printer = IPrinter;
  }
}
