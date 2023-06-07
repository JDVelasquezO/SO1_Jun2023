const express = require("express");
const app = express();
const port = 3000;

app.get("/", (req, res) => {
    
    getMovies();
    console.log("Carrie");
    
    res.send("Hola a todos");
});

function getMovies() {
    setTimeout(() => {
        console.log("Scarface");
    }, 2000);
}

app.listen(port, () => {
    console.log("Server on port ", port);
});
