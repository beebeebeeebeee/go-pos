import {useEffect, useMemo, useState} from 'react';
import * as ReceiptController from "@wails/controller/ReceiptController";
import * as PrinterController from "@wails/controller/PrinterController";
import {
    AppBar,
    Box,
    Button,
    Container,
    CssBaseline,
    FormControl,
    Grid, InputLabel,
    MenuItem,
    Select,
    Toolbar,
    Typography
} from "@mui/material";
import {Receipt} from "@/components/Receipt";
import AdbIcon from "@mui/icons-material/Adb";

export function App() {
    const [printers, setPrinters] = useState<any[]>([])
    const [receiptText, setReceiptText] = useState("");

    const selectedPrinter = useMemo(()=>{
        return printers.find(e=>e.Selected)?.Name
    },[printers])

    async function getReceipt():Promise<void>   {
        const receipt = await ReceiptController.GetReceipt()
        setReceiptText(receipt)
    }

    async function printReceipt() : Promise<void>{
         await ReceiptController.PrintReceipt()
    }

    async function selectPrinter(printer: string) : Promise<void>{
        const pl = await PrinterController.SetSelectedPrinter(printers.find(e=>e.Name === printer))
        setPrinters(pl)
    }

    useEffect(() => {
        (async()=>{
            const pl = await PrinterController.GetPrinterList()
            setPrinters(pl)
        })();

    }, []);

    return (
        <Box>
            <CssBaseline/>
            <AppBar position="fixed" sx={{
                px: 2
            }}>
                <Toolbar disableGutters variant="dense">
                    <AdbIcon sx={{display: {xs: 'none', md: 'flex'}, mr: 1}}/>
                    <Typography
                        variant="h6"
                        noWrap
                        sx={{
                            mr: 2,
                            display: {xs: 'none', md: 'flex'},
                            fontFamily: 'monospace',
                            fontWeight: 700,
                            letterSpacing: '.3rem',
                            color: 'inherit',
                            textDecoration: 'none',
                        }}
                    >
                        LOGO
                    </Typography>
                    <FormControl fullWidth>
                        <Select
                            labelId="demo-simple-select-label"
                            id="demo-simple-select"
                            value={selectedPrinter}
                            label="Age"
                            onChange={(e)=>selectPrinter(e.target.value)}
                        >
                            {printers.map(e => <MenuItem value={e.Name}>{e.Name}</MenuItem>)}
                        </Select>
                    </FormControl>

                </Toolbar>
            </AppBar>
            <Box sx={{px: 1}}>
                <Toolbar/>
                <Grid container spacing={2}>
                    <Grid item>
                        <Button
                            onClick={() => {
                                void printReceipt()
                            }}
                            variant="contained"
                        >
                            print receipt
                        </Button>
                        <Receipt text={receiptText}/>
                    </Grid>
                    <Grid item xs>
                        <Grid container spacing={2}>
                            {
                                Array.from({length: 20}, (_, i) => <Grid item xs={2.4}>
                                    <Button
                                        sx={{
                                            width: '100%',
                                            height: '6em',
                                            bgcolor: '#e1d8b6'
                                        }}
                                        onClick={() => {
                                            void getReceipt()
                                        }}
                                    >
                                        Item {i + 1}
                                    </Button>
                                </Grid>)
                            }
                        </Grid>
                    </Grid>
                </Grid>
            </Box>
        </Box>
    );
}

