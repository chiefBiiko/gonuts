/* socket.io server */

'use strict';

const fs = require('fs');

const db = {
  /*
  chiefBiiko: {
    contacts: ['chief', 'capo'],
    password: '',
    socketid: null
  }
  */
};

function handler(req, res) {
  
  // error handler for request stream
  req.on('error', err => {
    console.error(err);
    res.statusCode = 400;
    res.end();
  });
  // error handler for response stream
  res.on('error', err => console.error(err));
  // readstream from resource on GET
  if (req.method === 'GET') {
    let src;  // resource
    try {
      src = fs.createReadStream(`${__dirname}/entry.html`);
    } catch (err) {  // error handling when initializing readable
      console.error(err);
      res.statusCode = 500;
      res.end();
    }
    // error handling when serving readable
    src.on('error', err => {
      console.error(err);
      res.statusCode = 500;
      res.end();
    });
    // pipe readable to response
    console.log(`Serving html to ${req.headers.referer}`);
    res.writeHead(200, {'Content-Type': 'text/html'});
    src.pipe(res);
  } else if (req.method === 'POST') {
    // gather POST data in body
    let body = '';
    req.on('data', data => {
      body += data;
      if (body.length > 1e6) {  // ~~~ 1MB -> FLOOD ATTACK
        req.connection.destroy();
        res.statusCode = 413;  // request entity too large
        res.end();
      }
    });
    // use POST body data on end
    req.on('end', () => {
      if (req.url === '/reg') {
        console.log(body);
      } else if (req.url === '/log') {
        console.log(body);
      }
      res.statusCode = 200;
      res.end();  // rather stream out chat.html
    });
  } else {
    res.statusCode = 404;
    res.end();
  }
  
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
  
  socket.on('disconnect', info => console.log(info));
  
});

server.listen(65080, '127.0.0.1', () => {
  console.log('server listening on port 65080');
});