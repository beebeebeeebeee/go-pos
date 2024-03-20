import { Button, Grid, Stack, Typography } from "@mui/material";
import { entity } from "@wails/models";
import { formatPrice } from "@/util/string.util";
import { useItem } from "@/provider";

type ItemSectionProps = {
  addItem: (item: entity.Item) => void | Promise<void>;
};

export function ItemSection(props: ItemSectionProps): JSX.Element {
  const { addItem } = props;

  const { items } = useItem();

  return (
    <Grid container spacing={2}>
      {items?.map((item) => (
        <Grid item xs={2.4}>
          <Button
            sx={{
              width: "100%",
              height: "6em",
              bgcolor: "#e1d8b6",
              color: "#42402b",
            }}
            onClick={() => void addItem(item)}
          >
            <Stack>
              <Typography>{item.name}</Typography>
              <Typography>{formatPrice(item.price)}</Typography>
            </Stack>
          </Button>
        </Grid>
      ))}
    </Grid>
  );
}
