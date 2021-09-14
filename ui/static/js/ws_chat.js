var input = document.getElementById("input");
var output = document.getElementById("output");
var socket = new WebSocket("ws://localhost:8000/ws"); // ws://funclicker.online/ws

socket.onopen = function () {
    output.innerHTML += "Status: Connected\n";
};

socket.onmessage = function (e) {
    output.innerHTML += "Server: " + e.data + "\n";
};

function send() {
    socket.send(input.value);
    input.value = "";
    console.log(socket.CONNECTING)
}