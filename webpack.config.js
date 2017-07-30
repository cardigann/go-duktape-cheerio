module.exports = {
  entry: [
    './index'
  ],
  devtool: 'eval-source-map',
  output: {
    path: __dirname + "/dist",
    filename: "bundle.js"
  },
  module: {
    rules: [
      {
        test: /\.js$/,
        exclude: /(node_modules)/,
        use: {
          loader: 'babel-loader',
          options: {
            presets: ['es2015']
          }
        }
      }
    ]
  },
  node: {
    console: false,
    process: false,
    setImmediate: false,
  }
}
