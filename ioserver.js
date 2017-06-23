/* socket.io server */

const server = require('http').createServer();
const io = require('socket.io')(server);
io.on('connection', client => {
  console.log('CLIENT:' + client);
  client.on('event', data => {
    console.log('DATA:' + data);
  });
  client.on('error', err => {
    console.error('ERROR:' + err.message);
  });
  client.on('disconnect', () => {});
});
server.listen(65080, () => {
  console.log('server listening on port 65080');
});