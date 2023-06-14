const express = require("express");
const app = express();
const http = require("http");
const server = http.createServer(app);

const { Server } = require("socket.io");
const io = new Server(server);

const mysql = require("mysql");
const db = mysql.createConnection({
    host: 'localhost',
    user: 'root',
    database: 'test'
});

db.connect(err => {
    if (err) console.log(err);
});

app.get("/", (req, res) => {
    res.sendFile(__dirname + "/index.html");
});

io.on('connection', (socket) => {
    console.log("Se conectó un usuario");
    socket.on('disconnect', () => {
        console.log("Se desconectó un usuario");
    });

    socket.on("task", (task) => {
        // console.log("Tarea: ", task);
        io.emit('task', task);
        db.query("INSERT INTO task (task_desc) VALUES (?)", task);
    });
});

server.listen(3000, () => {
    console.log("Server on port 3000");
});