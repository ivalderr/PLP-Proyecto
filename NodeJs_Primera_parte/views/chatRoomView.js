function getChatRoomView() {
 let str = '';

  str+='<!DOCTYPE html>';
  str+='<html>';
  str+='<head>';
  str+='<meta http-equiv="content-type" content="text/html; charset=UTF-8">';
  str+='<meta charset="utf-8">';
  str+='<meta name="viewport" content="width=device-width, initial-scale=1">';
  str+='<meta name="description" content="">';
  str+='<meta name="generator" content="Hugo 0.87.0">';
  str+='<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.1.0/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-KyZXEAg3QhqLMpG8r+8fhAXLRk2vvoC2f3B09zVXn8CA5QIVfZOJ3BCsw2P0p/We" crossorigin="anonymous">';
  str+='<script src="https://cdn.jsdelivr.net/npm/bootstrap@5.1.0/dist/js/bootstrap.bundle.min.js" integrity="sha384-U1DAWAznBHeqEIlVSCgzq+c9gqGAJn5c/t99JyeKa9xxaYpSvHU5awsuZVVFIhvj" crossorigin="anonymous"></script>';
  str+='<title>YourChatRoom</title>';
  str+='<link href="css/style.css" rel="stylesheet">';
  str+='<link href="css/signin.css" rel="stylesheet" id="style-sign-in">';
  str+='</head>';
  str+='<body class="text-center">';
  str+='<div class="container">';
  str+='<div class="row" id="sign-in">';
  str+='<div class="col">';
  str+='<main class="form-signin">';
  str+='<form method="post" action="" id="main-form" pb-autologin="true" autocomplete="off">';
  str+='<img class="mb-4" src="images/logo_negro.png" alt="">';
  str+='<h1 class="h3 mb-3 fw-normal">Your Chat Room</h1>';
  str+='<div class="form-floating">';
  str+='<input type="text" class="form-control" name="nickname" id="floatingInput" placeholder="nickname" pb-role="username">';
  str+='<label for="floatingInput">Nickname</label>';
  str+='</div>';
  str+='<button class="w-100 btn btn-lg btn-primary" type="submit" pb-role="submit">Join Chat</button>';
  str+='<p class="mt-5 mb-3 text-muted">Maestría en Ciencias de la Computación © 2021</p>';
  str+='</form>';
  str+='</main>';
  str+='</div>';
  str+='</div>';
  str+='<div class="row">';
  str+='<div class="col">';
  str+='<h2 class="hid text-center" style="visibility:hidden" id="msg-welcome"></h2>';
  str+='</div>';
  str+='</div>';
  str+='<div class="row">';
  str+='<div class="col">';
  str+='<ul class="hid" style="visibility:hidden" id="messages"></ul>';
  str+='</div>';
  str+='<div class="col">';
  str+='<ul class="hid" style="visibility:hidden" id="logged-users"></ul>';
  str+='</div>';
  str+='</div>';
  str+='</div>';
  str+='<form class="hid" style="visibility:hidden" id="form" action="">';
  str+='<input class="form-control hid" style="visibility:hidden" id="input" autocomplete="off" />';
  str+='<button class="hid" style="visibility:hidden">Send</button>';
  str+='</form>';
  str+='<script src="/socket.io/socket.io.js"></script>';
  str+='<script src="js/script.js"></script>';
  str+='</body>';
  str+='</html>';

  return str;
}

module.exports = { getChatRoomView }
