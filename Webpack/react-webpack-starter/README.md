## **创建项目**



首先创建`react-webpack-starter`项目并使用 `npm init` 初始化。

##  

## **安装依赖**



### **安装react**



```
npm i react react-dom
```



**安装webpack相关**



```
npm i -D webpack webpack-cli webpack-dev-server html-webpack-plugin style-loader css-loader
```



安装`webpack-cli`后可以在命令行中执行webpack命令；`webpack-dev-server`提供了简易的web服务器，并且在修改文件之后自动执行webpack的编译操作并自动刷新浏览器，省去了重复的手动操作；`html-webpack-plugin`用于自动生成`index.html`文件，并且在`index.html`中自动添加对bundle文件的引用；`style-loader`和`css-loader`用于加载css文件。

### 

### **安装babel相关**



由于react中使用了`class`, `import`这样的es6的语法，为了提高网站的浏览器兼容性，我们需要用babel转换一下。



```
npm i -D @babel/core @babel/preset-env @babel/preset-react babel-loader 
```



其中`@babel/core`是babel的核心模块，包含了babel的核心功能；`@babel/preset-env`支持转换ES6以及更新的js语法，并且可根据需要兼容的浏览器类型选择加载的plugin从而精简生成的代码；`@babel/preset-react`包含了babel转换react所需要的plugin；`babel-loader`是webpack的babel加载器。



## 

## **配置webpack**



在项目根目录下面新建`webpack.config.js`，内容如下：

**webpack.config.js**



```json
const path = require('path');
const HtmlWebpackPlugin = require('html-webpack-plugin');

module.exports = {
  entry: './src/index.js',
  output: {
    path: path.resolve(__dirname, 'dist'),
    filename: 'bundle.js'
  },
  module: {
    rules: [
      {
        test: /\.js$/,
        exclude: /node_module/,
        use: 'babel-loader'
      },
      {
        test: /\.css$/,
        use: ['style-loader', 'css-loader'] // 注意排列顺序，执行顺序与排列顺序相反
      }
    ]
  },
  plugins: [
    new HtmlWebpackPlugin({
      template: './src/index.html'
    })
  ]
}
```



其中`HtmlWebpackPlugin`使用自定义的模版来生成html 文件，模版的内容如下：

**./src/index.html**



```html
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <meta http-equiv="X-UA-Compatible" content="ie=edge">
  <title>My React App</title>
</head>
<body>
  <div id="app"></div>
</body>
</html>
```



## **配置babel**



在项目根目录下面新建`.babelrc`文件，配置我们安装的两个babel preset：

**.babelrc**



```json
{
  "presets": [
    "@babel/preset-env",
    "@babel/preset-react"
  ]
}
```



## **生成react应用根节点**



**./src/index.js**



```js
import React from 'react';
import ReactDOM from 'react-dom';
import App from './components/App';

ReactDOM.render(<App />, document.getElementById('app'));
```



**./src/component/App.js**



```js
import React, { Component } from 'react';
import './App.css';

export default class App extends Component {
  render() {
    return (
      <div>
        my react webpack starter
      </div>
    )
  }
}
```



**./src/components/App.css**



```css
body{
  font-size: 60px;
  font-weight: bolder;
  color: red;
}
```



## **配置** **package.json**



最后，在`package.json`文件里面加上两个scripts，用来运行开发服务器和打包：

**package.json**



```json
"scripts": {
  "start": "webpack-dev-server --mode development --open --hot",
  "build": "webpack --mode production"
}
```



注意，我们启用了`webpack-dev-server`的模块热更新功能（HMR），进一步提高我们的开发效率。

