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
    console.log(`Serving html to ${req.headers.referer}`);  // ???
    res.writeHead(200, {'Content-Type': 'text/html'});
    res.end(data);
  });
}

const server = require('http').createServer(handler);
const io = require('socket.io')(server);

io.on('connection', socket => {
  
  console.log(`Client: ${socket.id}`);
  
  socket.emit('message', `Hi client ${socket.id}`);
  
  socket.on('message', msg => {
    console.log(`From ${socket.id}: ${msg}`);
    socket.emit('message', `From ${socket.id}: ${msg}`);
  });
  
  socket.on('error', err => {
    console.error(`Error: ${err.message}`);
  });
  
  socket.on('disconnect', info => {
    console.log(info);
  });
  
});

server.listen(65080, '127.0.0.1', () => {
  console.log('server listening on port 65080');
});