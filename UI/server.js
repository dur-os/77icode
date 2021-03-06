const webpack = require('webpack');
const express = require('express');
const path = require('path');
const app = express();
const port = 3100;

const config = require('./webpack-dev-server.config');
const compiler = webpack(config);

app.use(require('webpack-dev-middleware')(compiler, {
  noInfo: true,
  publicPath: config.output.publicPath
}));

app.use(require('webpack-hot-middleware')(compiler));

app.get('*', (req, res) => {
  res.sendFile(path.join(__dirname, 'src/www/index.html'));
});

/* eslint no-console: ["error", { allow: ["warn", "error"] }] */
app.listen(port, '0.0.0.0', err => {
  if (err) {
    console.error(err);
    return;
  }
  console.warn(`Listening at localhost:${port}`);
});
