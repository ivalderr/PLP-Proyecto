const express = require('express');
const app = express();
const http = require('http');
const server = http.createServer(app);
const { Server } = require("socket.io");
const io = new Server(server);
require('dotenv').config()

app.use("/public", express.static('public'));
app.use(express.urlencoded({
  extended: true
}));
app.use(express.json());

app.get('/', (req, res) => {
  res.sendFile(__dirname + '/index.html');
});

app.post('/chatroom', (req, res) => {
  let now = new Date();
  let logTime =  now.getHours()+':'+now.getMinutes()+':'+ now.getSeconds();
  console.log(`[${logTime}] `+req.body.nickname+' just entered the chat.');
  res.sendFile(__dirname + '/chatroom.html');
});

app.get('/chatroom', (req, res) => {
  res.sendFile(__dirname + '/chatroom.html');
});

io.on('connection', (socket) => {
  // console.log('a user connected');
  socket.on('disconnect', () => {
    // console.log('user disconnected');
  });
  socket.on('chat message', (msg) => {
    io.emit('chat message', msg);
  });
});

server.listen(process.env.PORT, () => {
  console.log(`listening on *:${process.env.PORT}`);
});
