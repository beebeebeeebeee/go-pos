import {
  createContext,
  ReactNode,
  useCallback,
  useContext,
  useMemo,
} from "react";
import { entity } from "@wails/models";
import { useQuery } from "@tanstack/react-query";
import * as OrderController from "@wails/controller/OrderController";

export type OrderContextType = {
  order?: entity.Order;
  updateOrderItem: (item: entity.Item, adjustment: number) => Promise<void>;
};

export const OrderContext = createContext<OrderContextType | undefined>(
  undefined
);

export type OrderProviderProps = {
  children: ReactNode;
};

export function OrderProvider(props: OrderProviderProps): JSX.Element {
  const { children } = props;

  const { data, refetch } = useQuery<entity.Order>({
    queryKey: ["order"],
    queryFn: () => OrderController.GetLatestOrder(),
  });

  const updateOrderItem = useCallback(
    async (item: entity.Item, adjustment: number) => {
      if (data == null) return;

      await OrderController.UpdateOrderItem(data.id, item.id, adjustment);
      await refetch();
    },
    [data, refetch()]
  );

  const value: OrderContextType = useMemo(
    () => ({ order: data, updateOrderItem }),
    [data, updateOrderItem]
  );

  return (
    <OrderContext.Provider value={value}>{children}</OrderContext.Provider>
  );
}

export function useOrder(): OrderContextType {
  const context = useContext(OrderContext);
  if (!context) {
    throw new Error("useOrder must be used within a OrderProvider");
  }

  return context;
}
