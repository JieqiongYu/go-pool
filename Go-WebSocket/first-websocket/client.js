var input = document.getElementById("input");
var output = document.getElementById("output");

var socket = new WebSocket("ws://" + window.location.host + "/echo");
socket.onopen = function() {
    output.innerHTML += "Status: Connected\n";
};

socket.onmessage = function(e) {
    output.innerHTML += "Message from Server: " + e.data + "\n";
};

function send() {
    socket.send (
        JSON.stringify(
            {
                message: input.value
            }
        )
    );
    input.value="";
}