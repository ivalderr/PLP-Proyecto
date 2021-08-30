const express = require('express');
const app = express();
const http = require('http');
const server = http.createServer(app);
const { Server } = require("socket.io");
const io = new Server(server);
const { joinUser, removeUser, getUser, getAll } = require('./jscripts/users');
const { getTime } = require('./jscripts/utils');
const { getChatRoomView } = require('./views/chatRoomView.js');

require('dotenv').config()

app.use(express.static('public'));
app.use(express.json());
app.use(express.urlencoded({
  extended: true
}));

app.get('/', (req, res) => {
  res.send(getChatRoomView());
});

io.on('connection', (socket) => {
  socket.on('join chat', (nickname) => {
    console.log(getTime(), nickname, 'just joined the chat.');
    joinUser(socket.id, nickname);
    io.emit('update logged', getAll());
  });
  socket.on('chat message', (msg, socketId) => {
    console.log(getTime(), getUser(socketId).nickname+" said:", msg);
    io.emit('chat message', getUser(socketId).nickname + ': ' + msg);
  });
  socket.on('disconnect', () => {
    if(getUser(socket.id)) {
      console.log(getTime(), getUser(socket.id).nickname, 'leaved the chat.');
      removeUser(socket.id);
      io.emit('update logged', getAll());
    }
  });
});

server.listen(process.env.PORT, () => {
  console.log(getTime(), `listening on *:${process.env.PORT}`);
});
