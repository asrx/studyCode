var debug = process.env.NODE_ENV !== "production"

module.exports = {
    devtool: debug ? "sourcemap" : null,
    entry: "./src/js/entry.js",
    output:{
        filename: "bundle.js"
    },
    module: {
        rules: [
          {
            test: /\.css$/i,
            use: ['style-loader', 'css-loader'],
          },
          {
            test: /\.m?js$/,
            exclude: /(node_modules|bower_components)/,
            use: {
              loader: 'babel-loader',
              options: {
                presets: ['@babel/preset-env']
              }
            }
          }
        ]
    },
    plugins: debug ? [] : []
}