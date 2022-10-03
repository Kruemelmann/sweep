//<hostname>:<port>
var currentLocation = window.location.hostname+":"+window.location.port;
var client = new WebSocket('ws://'+currentLocation+'/ws');
var ctx = document.getElementById('canvas-container');

client.onopen = () => {
    console.log('WebSocket Client Connected');
};
client.onmessage = (message) => {
    console.log('reload image');
    loadImage()
};

async function loadImage() {
    let url = 'http://'+currentLocation+'/frame';

    fetch(url)
        .then(response => response.blob())
        .then(imageBlob => {
            renderImage(ctx, imageBlob)
        });
}

function renderImage(canvas, blob) {
  const ctx = canvas.getContext('2d')
  const img = new Image()
  img.onload = (event) => {
      URL.revokeObjectURL(event.target.src)
      ctx.canvas.width  = window.innerWidth;
      ctx.canvas.height = window.innerHeight;
      ctx.drawImage(event.target, 0, 0)
      ctx.drawImage(img, 0, 0, img.width, img.height,  // source rectangle
                   0, 0, canvas.width, canvas.height); // destination rectangle

  }
  img.src = URL.createObjectURL(blob)
}
