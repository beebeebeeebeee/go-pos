import {useState} from 'react';
import logo from './assets/images/logo-universal.png';
import './App.css';
import {Printer} from "../wailsjs/go/main/App";

function App() {
    const [resultText, setResultText] = useState("Please enter your name below 👇");

    async function greet() {
       const printer = await Printer()
        setResultText(printer)
    }

    return (
        <div id="App">
            <img src={logo} id="logo" alt="logo"/>
            <div id="result" className="result">{resultText}</div>
            <button className="btn" onClick={greet}>Greet</button>
        </div>
    )
}

export default App
