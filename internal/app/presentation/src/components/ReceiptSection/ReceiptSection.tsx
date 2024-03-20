import { Box } from "@mui/material";
import { useReceipt } from "@/provider";

export function ReceiptSection(): JSX.Element {
  const { receiptText } = useReceipt();

  return (
    <Box
      sx={{
        fontFamily: "AssetsMono",
        fontSize: "0.75rem",
        whiteSpace: "pre-wrap",
        textAlign: "center",
        width: "15rem",
      }}
    >
      {receiptText}
    </Box>
  );
}
