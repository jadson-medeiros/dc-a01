import app from './src/app.js';

const port = process.env.PORT || 3000;
const host = '0.0.0.0';

app.listen(port, host, () => {
  console.log('listening on port ' + port + ' and host ' + host)
})