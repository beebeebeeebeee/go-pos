import React, { useCallback } from "react";
import { Button, Grid, Stack } from "@mui/material";
import { entity } from "@wails/models";
import {
  ItemProvider,
  OrderProvider,
  ReceiptProvider,
  useOrder,
  useReceipt,
} from "@/provider";
import { ItemSection, ReceiptSection } from "@/components";

function _MainPage(): JSX.Element {
  const { updateOrderItem } = useOrder();
  const { refetchReceipt, printReceipt } = useReceipt();

  const addItem = useCallback(
    async (item: entity.Item) => {
      await updateOrderItem(item, 1);
    },
    [updateOrderItem]
  );

  return (
    <Grid container spacing={2}>
      <Grid item>
        <Stack spacing={2}>
          <Button
            onClick={() => {
              void printReceipt();
            }}
            variant="contained"
          >
            print receipt
          </Button>
          <ReceiptSection />
        </Stack>
      </Grid>
      <Grid item xs>
        <ItemSection
          addItem={async (item) => {
            await addItem(item);
          }}
        />
      </Grid>
    </Grid>
  );
}

export function MainPage(): JSX.Element {
  return (
    <OrderProvider>
      <ItemProvider>
        <ReceiptProvider>
          <_MainPage />
        </ReceiptProvider>
      </ItemProvider>
    </OrderProvider>
  );
}
