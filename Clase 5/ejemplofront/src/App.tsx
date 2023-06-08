import viteLogo from '/vite.svg'
import './App.css'
import {useEffect, useState} from "react";
import {getCpu} from "./api/endpoints.ts";

function App() {

    const [ percent, setPercent ] = useState("0");

    useEffect(() => {
        const interval = setInterval(() => {
            (
                async () => {
                    const req = await getCpu();
                    const res = await req.json();
                    console.log(res);

                    setPercent(res[0].percent);
                }
            )();
        }, 500);

        return () => clearInterval(interval);
    })

    return (
    <>
      <div>
        <a href="https://vitejs.dev" target="_blank">
          <img src={viteLogo} className="logo" alt="Vite logo" />
        </a>
      </div>
      <h1>{ percent }%</h1>
    </>
    );
}

export default App
