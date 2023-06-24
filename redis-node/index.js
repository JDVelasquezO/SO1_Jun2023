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

const server = http.createServer(app);
const io = socketio(server, {
    cors: {
        origin: "*"
    },
});

const client = createClient({
    url: 'redis://localhost:6379/15'
}); //creates a new client
client.on('error', err => console.log('Redis Client Error', err));

Promise.all([client.connect()]).then(() => {
    io.on('connection', socket => {
        console.log("Connected");
    
        socket.on('sendkey', async data => {
            key = data.key;
            let res = await client.hGetAll("Love-1967-Forever_Changes");
            // console.log(JSON.stringify(res, null, 2));
            io.emit("sendkey", res);
        });
    });
    
    server.listen(4000, () => console.log('Server levantado en el puerto 4000'));
});