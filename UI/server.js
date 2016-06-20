const webpack = require('webpack');
const express = require('express');
const path = require('path');
const app = express();


const config = require('./webpack-dev-server.config');
const compiler = webpack(config);

// require('./fakeAPI');

app.use(require('webpack-dev-middleware')(compiler, {
    noInfo: true,
    publicPath: config.output.publicPath
}));

app.use(require('webpack-hot-middleware')(compiler));

app.get('*', (req, res) => {
    res.sendFile(path.join(__dirname, 'src/www/index.html'));
});

app.listen(3000, '0.0.0.0', err => {
    if (err) {
        console.log(err);
        return;
    }
    console.log('Listening at localhost:3000');
});
