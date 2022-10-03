//<hostname>:<port>
var currentLocation = window.location.hostname+":"+window.location.port;

var client = new WebSocket('ws://'+currentLocation+'/ws');

client.onopen = () => {
    console.log('WebSocket Client Connected');
};
client.onmessage = (message) => {
    console.log('reload image');
};
