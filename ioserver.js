/* socket.io server */
'use strict';

const fs = require('fs');

function handler (req, res) {
  fs.readFile(__dirname + '/index.html', (err, data) => {
    if (err) {
      res.writeHead(500);
      console.log('Error loading index.html');
      return res.end('Error loading index.html');
    }
    res.writeHead(200);
    res.end(data);
  });
}

const server = require('http').createServer(handler);
const io = require('socket.io')(server);

io.on('connection', socket => {
  
  console.log('CLIENT: ' + socket.id);
  
  socket.emit('serverMsg', 'Hi client!');
  
  socket.on('clientmsg', data => {
    console.log('DATA:' + data);
    socket.emit('serverMsg', data);
  });
  
  socket.on('error', err => {
    console.error('ERROR:' + err.message);
  });
  
  socket.on('disconnect', () => {});
  
});

server.listen(65080, '127.0.0.1', () => {
  console.log('server listening on port 65080');
});