/* socket.io server */

'use strict';

const fs = require('fs');

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
  if (req.method === 'GET' && 
      /^.*\/reg.html$|^.*\/log.html$|^.*\/chat.html$/.test(req.url)) {
        
    var src;  // resource
    const ep = req.url.replace(/.(?=.*\/)|\//g, '');
    
    try {
      src = fs.createReadStream(`${__dirname}/${ep}`);
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
    
  } else if (req.url === '/reg' && req.method === 'POST') {
    
  } else if (req.url === '/log' && req.method === 'POST') {
    
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