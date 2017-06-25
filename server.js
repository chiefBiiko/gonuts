/* socketbroker */

const http = require('http');
const net = require('net');

// [START http_server]
const HTTPserver = http.createServer((req, res) => {
   
}).listen(process.env.PORT || 8080, () => {
  console.log('http.server listening on port 8080');
});
// [END http_server]

// [START tcp_server]
const TCPserver = net.createServer(socket => {
  socket.setEncoding('utf8');
  
  // on connect
  console.log(`CLIENT: ${socket.remoteAddress}:${socket.remotePort}`);
    
  // Add a 'data' event handler to this instance of socket
  socket.on('data', data => {
    console.log(`FROM ${socket.remoteAddress}:${socket.remotePort}: ${data}`);
    socket.write(`You said "${data}"`);  // echo
  });
    
  // Add a 'error' event handler to this instance of socket
  socket.on('error', err => {
    console.error(`ERROR: ${err}`);
  });
    
  // Add a 'close' event handler to this instance of socket
  socket.on('close', had_error => {
    console.log(`CLOSED: ${socket.remoteAddress}:${socket.remotePort}`);
  });
  
}).listen(65080, '127.0.0.1', () => {  // forwarded port in VPC network on gcloud
  console.log('TCPserver listening on port 65080');
});
// [END tcp_server]





/*
// TODO: REWRITE REQUEST STUFF WITH A HTTP GET CLIENT
const request = require('request');

// [START external_ip]
// In order to use websockets on App Engine, you need to connect directly to
// application instance using the instance's public external IP. This IP can
// be obtained from the metadata server.
const METADATA_NETWORK_INTERFACE_URL = 'http://metadata/computeMetadata/v1/' +
    '/instance/network-interfaces/0/access-configs/0/external-ip';

function getExternalIp(cb) {
  // request options
  const options = {
    url: METADATA_NETWORK_INTERFACE_URL,
    headers: {
      'Metadata-Flavor': 'Google'
    }
  };
  // actual request
  request(options, (err, resp, body) => {
    if (err || resp.statusCode !== 200) {
      console.log('Error while talking to metadata server, assuming localhost');
      cb(null, 'localhost');
      return;
    }
    cb(null, body);
  });
}
// [END external_ip]
*/