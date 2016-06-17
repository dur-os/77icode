var webpack = require('webpack');
var express = require('express');
var path = require('path');
var app = express();


var config = require('./webpack-dev-server.config');
var compiler = webpack(config);

// require('./fakeAPI');

app.use(require('webpack-dev-middleware')(compiler, {
    noInfo: true,
    publicPath: config.output.publicPath
}));

app.use(require('webpack-hot-middleware')(compiler));

app.get('*', function(req, res) {
res.sendFile(path.join(__dirname, 'src/www/index.html'));
});

app.listen(3000, '0.0.0.0', function(err) {
if (err) {
    console.log(err);
    return;
}
console.log('Listening at localhost:3000');
});
