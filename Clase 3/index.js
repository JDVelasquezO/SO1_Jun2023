const express = require("express");
const app = express();
const port = 3000;
const cors = require("cors");
require("dotenv").config();

app.use(cors());
app.use(express.urlencoded({
    extended: true,
}));
app.use(express.json());

const { ObjectId } = require("mongodb");
const mongoose = require("mongoose").default;

const uri = `mongodb://${process.env.DB_HOST}:27017/DB`;

const cpuSchema = new mongoose.Schema({
    _id: ObjectId,
    percent: String
}, { collection: 'cpu' });

app.get("/", (req, res) => { 
    res.send("Hola a todos");
});

app.get("/getCpu", async (req, res) => {
    await mongoose.connect(uri);
    const Cpu = mongoose.model('cpu', cpuSchema);
    const percent = await Cpu.find({}, {_id:0})
        .sort({ $natural: -1 })
        .limit(1);

    return res.json(percent);
});

app.listen(port, () => {
    console.log("Server on port ", port);
});
