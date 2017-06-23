/* socketclient */

const net = require('net');

const client = new net.Socket();
client.connect(65080, '127.0.0.1', () => {
  console.log('CONNECTED');
  client.write('Hello, server! Love, Client.');
});

client.on('data', function(data) {
  console.log('RECEIVED: ' + data);
  client.destroy(); // kill client after server's response
});

client.on('close', function() {
  console.log('DISCONNECTED');
});