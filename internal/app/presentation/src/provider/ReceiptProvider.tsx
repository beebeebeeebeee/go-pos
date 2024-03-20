import {
  createContext,
  ReactNode,
  useCallback,
  useContext,
  useMemo,
} from "react";
import { useQuery } from "@tanstack/react-query";
import * as ReceiptController from "@wails/controller/ReceiptController";

export type ReceiptContextType = {
  receiptText: string;
  refetchReceipt: () => void;
  printReceipt: () => void;
};

export const ReceiptContext = createContext<ReceiptContextType | undefined>(
  undefined
);

export type ReceiptProviderProps = {
  children: ReactNode;
};

export function ReceiptProvider(props: ReceiptProviderProps): JSX.Element {
  const { children } = props;

  const { data, refetch } = useQuery<string>({
    queryKey: ["receipt"],
    queryFn: () => ReceiptController.GetReceipt(),
  });

  const printReceipt = useCallback(() => {
    ReceiptController.PrintReceipt();
  }, []);

  const value: ReceiptContextType = useMemo(
    () => ({ receiptText: data ?? "", refetchReceipt: refetch, printReceipt }),
    [data]
  );

  return (
    <ReceiptContext.Provider value={value}>{children}</ReceiptContext.Provider>
  );
}

export function useReceipt(): ReceiptContextType {
  const context = useContext(ReceiptContext);
  if (!context) {
    throw new Error("useReceipt must be used within a ReceiptProvider");
  }

  return context;
}
