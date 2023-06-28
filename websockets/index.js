const { createClient } = require('redis');
const express = require('express');
const http = require('http');
const socketio =  require('socket.io');
const cors = require('cors');

const app = express();
app.use(cors());
app.use(express.urlencoded({
    extended: true,
}));
app.use(express.json());

app.get("/", (req, res) => {
    res.sendFile(__dirname + "/index.html");
});

const server = http.createServer(app);
const io = socketio(server, {
    cors: {
        origin: "*"
    },
});

const client = createClient({
    url: 'redis://10.157.19.235:6379/15'
}); //creates a new client
client.on('error', err => console.log('Redis Client Error', err));

Promise.all([client.connect()]).then(() => {
    io.on('connection', socket => {
        console.log("Connected");

        socket.on('sendkey', data => {
            key = data.key;
            setInterval(async () => {
                let res = await client.hGetAll(key);
                io.emit("sendkey", res);
            }, 500);
            // console.log(JSON.stringify(res, null, 2));
        });
    });

    server.listen(4000, () => console.log('Server levantado en el puerto 4000'));
});
