import { createContext, ReactNode, useContext, useMemo } from "react";
import { entity } from "@wails/models";
import { useQuery } from "@tanstack/react-query";
import * as ItemController from "@wails/controller/ItemController";

export type ItemContextType = {
  items: entity.Item[];
};

export const ItemContext = createContext<ItemContextType | undefined>(
  undefined
);

export type ItemProviderProps = {
  children: ReactNode;
};

export function ItemProvider(props: ItemProviderProps): JSX.Element {
  const { children } = props;

  const { data } = useQuery<entity.Item[]>({
    queryKey: ["items"],
    queryFn: () => ItemController.GetItemList(),
  });

  const value: ItemContextType = useMemo(() => ({ items: data ?? [] }), [data]);

  return <ItemContext.Provider value={value}>{children}</ItemContext.Provider>;
}

export function useItem(): ItemContextType {
  const context = useContext(ItemContext);
  if (!context) {
    throw new Error("useItem must be used within a ItemProvider");
  }

  return context;
}
