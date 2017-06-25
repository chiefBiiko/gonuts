/* socket.io server */

'use strict';

const fs = require('fs');
const path = require('path');

const db = {
  chiefBiiko: {
    contacts: [],
    password: '419'
  }
};

var sockets = [];

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
      src = fs.createReadStream(path.join(__dirname, 'entry.html'));
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
    console.log(`Serving html to ${req.headers.origin}`);
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
      const data = JSON.parse(body);
      if (req.url === '/register') {
        if (Object.keys(db).includes(data.name)) {
          res.statusCode = 423;  // locked
          res.end();
        } else {
          db[data.name] = {
            contacts: [],
            password: data.password,
            socketid: ''
          };
          
          let src;  // resource
          try {
            src = fs.createReadStream(path.join(__dirname, 'chat.html'));
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
          console.log(`Serving html to ${req.headers.origin}`);
          res.writeHead(200, {'Content-Type': 'text/html'});
          src.pipe(res);
          
        }
      } else if (req.url === '/login') {
        if (Object.keys(db).includes(data.name) &&
            db[data.name].password === data.password) {
          
          let src;  // resource
          try {
            src = fs.createReadStream(path.join(__dirname, 'chat.html'));
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
          console.log(`Serving html to ${req.headers.origin}`);
          res.writeHead(200, {'Content-Type': 'text/html'});
          src.pipe(res);
          
        } else {
          res.statusCode = 401;
          res.end();
        }
      }
    });
  } else {
    res.statusCode = 404;
    res.end();
  }
  
}

const server = require('http').createServer(handler);
const io = require('socket.io')(server);

io.on('connection', socket => {
  
  socket.on('whoami', x => {
    sockets.push([x, socket]);
    console.log(`Client: ${x} with socketid ${socket.id}`);
    socket.emit('message', `Hi ${x}`);
  });
  
  socket.on('message', msg => {
    const username = sockets.filter(s => s[1] === socket)[0][0];
    console.log(`From ${username}: ${msg}`);
    socket.emit('message', `From ${username}: ${msg}`);
  });
  
  socket.on('error', err => {
    console.error(`Error: ${err.message} in socket ${socket.id}`);
  });
  
  socket.on('disconnect', () => {
    const username = sockets.filter(s => s[1] === socket)[0][0];
    console.log(`DISCONNECT:\n ${username}`);
    sockets = sockets.filter(s => s[1] !== socket);
  });
  
});

server.listen(65080, '127.0.0.1', () => {
  console.log('server listening on port 65080');
});