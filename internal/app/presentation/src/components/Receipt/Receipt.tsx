import { Box } from "@mui/material";

type ReceiptProps = {
  text: string;
};

export function Receipt(props: ReceiptProps) {
  const { text } = props;
  return (
    <Box
      sx={{
        fontFamily: "AssetsMono",
        fontSize: "0.75rem",
        whiteSpace: "pre-wrap",
        width: "15rem",
      }}
    >
      {text}
    </Box>
  );
}
