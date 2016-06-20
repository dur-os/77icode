const webpack = require('webpack');
const path = require('path');
const buildPath = path.resolve(__dirname, 'build');
const ExtractTextPlugin = require("extract-text-webpack-plugin");

const config = {
  // Entry points to the project
  entry: [
    'webpack-hot-middleware/client',
    path.join(__dirname, '/src/app/app.js')
  ],
  output: {
    path: buildPath, // Path of output file
    filename: 'app.js'
  },
  plugins: [
    // Enables Hot Modules Replacement
    new webpack.HotModuleReplacementPlugin(),
    // Allows error warnings but does not stop compiling.
    new webpack.NoErrorsPlugin(),
    new ExtractTextPlugin("app.css")
    // Moves files
  ],
  module: {
    loaders: [{
      test: /\.js$/,
      loader: 'babel',
      include: path.join(__dirname, 'src', 'app')
    }, {
      test: /\.css$/,
      loader: ExtractTextPlugin.extract("style-loader","css-loader")
      //include: path.join(__dirname, 'src', 'app')
    }, {
      test: /\.less$/,
      loader:  ExtractTextPlugin.extract("style-loader", "css-loader!less-loader")
      //include: path.join(__dirname, 'src', 'app')
    }],
  },
};

module.exports = config;
