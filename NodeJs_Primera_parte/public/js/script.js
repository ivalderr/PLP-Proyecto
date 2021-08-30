var socket = io();

var messages = document.getElementById('messages');
var loggedUsers = document.getElementById('logged-users');
var form = document.getElementById('form');
var input = document.getElementById('input');
var welcome = document.getElementById('msg-welcome');
var mainForm = document.getElementById('main-form');
var nicknameInput = document.getElementById('floatingInput');

let nickname = '';

form.addEventListener('submit', function(e) {
  e.preventDefault();
  if (input.value) {
    socket.emit('chat message', input.value, socket.id);
    input.value = '';
  }
});

mainForm.addEventListener('submit', function(e) {
  e.preventDefault();
  if (nicknameInput.value) {
    nickname = nicknameInput.value;
    let welcomeMsg = 'Hi '+nickname+', Welcome to the chat!';
    document.body.classList.remove("text-center")
    document.getElementById('msg-welcome').textContent = welcomeMsg;
    document.getElementById('sign-in').remove();
    document.getElementById('style-sign-in').remove();
    document.querySelectorAll('.hid').forEach( el => {
      el.style.visibility = 'visible';
    });

    socket.emit('join chat', nickname);
  }
});

socket.on('chat message', function(msg) {
  var item = document.createElement('li');
  item.textContent = msg;
  messages.appendChild(item);
  window.scrollTo(0, document.body.scrollHeight);
});

socket.on('update logged', (logged) => {
  loggedUsers.innerHTML = "";
  logged.forEach(lg =>{
    var item = document.createElement('li');
    item.textContent = lg.nickname;
    loggedUsers.appendChild(item);
  });
});
