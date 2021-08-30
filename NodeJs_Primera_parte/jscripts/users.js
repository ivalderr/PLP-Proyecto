let users = [];

function joinUser(socketId , nickname) {
  const user = {
    socketId : socketId,
    nickname : nickname
  }
  users.push(user)
  return user;
}

function removeUser(id) {
  const getID = users => users.socketId === id;
  const index =  users.findIndex(getID);
  if (index !== -1) {
    return users.splice(index, 1)[0];
  }
}

function getUser(socketId) {
  return users.find( s => s.socketId === socketId );
}

function getAll() {
  return users;
}

module.exports = { joinUser, removeUser, getUser, getAll }
