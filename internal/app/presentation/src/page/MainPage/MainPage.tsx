import React from "react";
import { Button, Divider, Grid, Stack } from "@mui/material";
import { ItemProvider, ReceiptProvider, useReceipt } from "@/provider";
import { ItemSection, ReceiptSection } from "@/components";

function _MainPage(): JSX.Element {
  const { refetchReceipt, printReceipt } = useReceipt();

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
          addItem={() => {
            refetchReceipt();
          }}
        />
      </Grid>
    </Grid>
  );
}

export function MainPage(): JSX.Element {
  return (
    <ItemProvider>
      <ReceiptProvider>
        <_MainPage />
      </ReceiptProvider>
    </ItemProvider>
  );
}
