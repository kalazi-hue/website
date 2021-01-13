const path = require('path')
const IS_PROD = ['production', 'test'].includes(process.env.NODE_ENV) // 热更新

// 在vue-config.js 中加入
// 开启gzip压缩
const CompressionWebpackPlugin = require('compression-webpack-plugin');
// 判断开发环境
const isProduction = process.env.NODE_ENV === 'production';

const resolve = dir => {
  return path.join(__dirname, dir)
}

// 项目部署基础
const BASE_URL = process.env.NODE_ENV === 'production'
  ? '/'
  : '/'

module.exports = {
  //webpack配置
  configureWebpack:config => {
    // 开启gzip压缩
    if (isProduction) {
      config.plugins.push(new CompressionWebpackPlugin({
        algorithm: 'gzip',
        test: /\.js$|\.html$|\.json$|\.css/,
        threshold: 10240,
        minRatio: 0.8
      }));
      // 开启分离js
      config.optimization = {
        runtimeChunk: 'single',
        splitChunks: {
          chunks: 'all',
          maxInitialRequests: Infinity,
          minSize: 20000,
          cacheGroups: {
            vendor: {
              test: /[\\/]node_modules[\\/]/,
              name (module) {
                // get the name. E.g. node_modules/packageName/not/this/part.js
                // or node_modules/packageName
                const packageName = module.context.match(/[\\/]node_modules[\\/](.*?)([\\/]|$)/)[1]
                // npm package names are URL-safe, but some servers don't like @ symbols
                return `npm.${packageName.replace('@', '')}`
              }
            }
          }
        }
      };
      // 取消webpack警告的性能提示
      config.performance = {
        hints:'warning',
            //入口起点的最大体积
            maxEntrypointSize: 50000000,
            //生成文件的最大体积
            maxAssetSize: 30000000,
            //只给出 js 文件的性能提示
            assetFilter: function(assetFilename) {
          return assetFilename.endsWith('.js');
        }
      }
    }
  },
  publicPath: BASE_URL,
  devServer: {
    host: 'localhost',
    port: 8000, // 端口号
    hotOnly: false,
    https: false, // https:{type:Boolean}
    open: true, //配置自动启动浏览器
    proxy: {
      '/api': {
        target: 'http://www.8dafu.com', // api接口基础路径
        changeOrigin: true, // 是否支持跨域
        ws: true,
        pathRewrite: {
          '^/api': ''
        }
      }
    }
  },
  // 如果你不需要使用eslint，把lintOnSave设为false即可
  lintOnSave: true,
  css:{
    loaderOptions:{
      less:{
        javascriptEnabled:true
      }
    },
    extract: true,// 是否使用css分离插件 ExtractTextPlugin
    sourceMap: false,// 开启 CSS source maps
    modules: false// 启用 CSS modules for all css / pre-processor files.
  },
  chainWebpack: config => {
    config.resolve.alias
      .set('@', resolve('src')) // key,value自行定义，比如.set('@@', resolve('src/components'))
      .set('@c', resolve('src/components'))
  },
  // 打包时不生成.map文件
  productionSourceMap: false,
  // 这里写你调用接口的基础路径，来解决跨域，如果设置了代理，那你本地开发环境的axios的baseUrl要写为 '' ，即空字符串
  // devServer: {
  //   proxy: 'localhost:3000'
  // }
  css: {
    extract: IS_PROD,
    sourceMap: false,
    modules: false
  }
}
