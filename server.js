/* socketbroker */

const net = require('net');
const http = require('http');
// TODO: REWRITE REQUEST STUFF WITH A HTTP GET CLIENT
const request = require('request');


// [START external_ip]
/* In order to use websockets on App Engine, you need to connect directly to
   application instance using the instance's public external IP. This IP can
   be obtained from the metadata server. */
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

// [START http_server]
const GETserver = http.createServer((req, res) => {
  if (err) {
    res.writeHead(500);
    res.end(err.message);
  }
  res.end('CLIENT_HTML INCLUDING WS CLIENT CODE');
});
GETserver.listen(process.env.PORT || 8080, () => {
  global.console.log('http.server listening on port 8080');
});
// [END http_server]

// [START tcp_server]
const TCPserver = net.createServer(socket => {
  
});
server.listen(65080, () => {  // forwarded port in VPC network on gcloud
  global.console.log('net.server listening on port 65080');
});
// [END tcp_server]