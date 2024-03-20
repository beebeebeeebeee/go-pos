import React, { useEffect, useMemo, useState } from "react";
import * as PrinterController from "@wails/controller/PrinterController";
import {
  AppBar,
  Box,
  CssBaseline,
  FormControl,
  MenuItem,
  Select,
  Stack,
  Toolbar,
  Typography,
} from "@mui/material";
import AdbIcon from "@mui/icons-material/Adb";
import PrintIcon from "@mui/icons-material/Print";
import { entity } from "@wails/models";
import { MainPage } from "@/page";

export function App(): JSX.Element {
  const [printers, setPrinters] = useState<entity.Printer[]>([]);

  const selectedPrinter = useMemo(() => {
    return printers.find((e) => e.selected)?.name;
  }, [printers]);

  async function selectPrinter(printer: string): Promise<void> {
    const pl = await PrinterController.SetSelectedPrinter(
      printers.find((e) => e.name === printer)!
    );
    setPrinters(pl);
  }

  useEffect(() => {
    (async () => {
      const pl = await PrinterController.GetPrinterList();
      setPrinters(pl);
    })();
  }, []);

  return (
    <Box>
      <CssBaseline />
      <AppBar
        position="fixed"
        sx={{
          px: 2,
        }}
      >
        <Toolbar disableGutters variant="dense">
          <Stack direction="row" alignItems="center" spacing={1}>
            <AdbIcon />
            <Typography
              variant="h6"
              noWrap
              sx={{
                mr: 2,
                display: { xs: "none", md: "flex" },
                fontFamily: "monospace",
                fontWeight: 700,
                letterSpacing: ".3rem",
                color: "inherit",
                textDecoration: "none",
              }}
            >
              LOGO
            </Typography>
          </Stack>
          <Stack
            sx={{
              width: "100%",
            }}
            spacing={2}
            direction="row"
            alignItems="center"
            justifyContent="flex-end"
          >
            <PrintIcon />
            <FormControl sx={{ width: "10rem" }}>
              <Select
                labelId="demo-simple-select-label"
                id="demo-simple-select"
                value={selectedPrinter}
                label="Age"
                onChange={(e) => selectPrinter(e.target.value)}
              >
                {printers.map((e) => (
                  <MenuItem value={e.name}>{e.name}</MenuItem>
                ))}
              </Select>
            </FormControl>
          </Stack>
        </Toolbar>
      </AppBar>
      <Box sx={{ px: 1 }}>
        <Toolbar />
        <MainPage />
      </Box>
    </Box>
  );
}
