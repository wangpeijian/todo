var path = require('path')
var chalk = require('chalk')
var webpack = require('webpack')
var utils = require('./utils')
var config = require('../config')
var vueLoaderConfig = require('./vue-loader.conf')
var ProgressBarPlugin = require('progress-bar-webpack-plugin')
var os = require('os')
var HappyPack = require('happypack');
var happyThreadPool = HappyPack.ThreadPool({ size: os.cpus().length });
var ExtractTextPlugin = require('extract-text-webpack-plugin');
var svgoConfig = require('../config/svgo-config.json')

function resolve (dir) {
  return path.join(__dirname, '..', dir)
}

var cssLoader = ExtractTextPlugin.extract({
    use: [
        'happypack/loader?id=happy-css'
    ]
})

// inject happypack accelerate packing for vue-loader @17-08-18
Object.assign(vueLoaderConfig.loaders, {
    js: 'happypack/loader?id=happy-babel-vue',
    css: cssLoader
})

function createHappyPlugin (id, loaders) {
    return new HappyPack({
        id: id,
        loaders: loaders,
        threadPool: happyThreadPool,
        // make happy more verbose with HAPPY_VERBOSE=1
        verbose: process.env.HAPPY_VERBOSE === '1'
    })
}

module.exports = {
    entry: {
        app: './src/main.js'
    },
    output: {
        path: config.build.assetsRoot,
        filename: '[name].js',
        publicPath: process.env.NODE_ENV === 'production'
            ? config.build.assetsPublicPath
            : config.dev.assetsPublicPath
    },
    resolve: {
        extensions: ['.js', '.vue', '.json'],
        modules: [
            resolve('src'),
            resolve('node_modules')
        ],
        alias: {
            'vue$': 'vue/dist/vue.esm.js',
            '@': resolve('src')
        }
    },
    module: {
        noParse: /node_modules\/(element-ui\.js)/,
        rules: [
            {
                test: /\.svg$/,
                enforce: 'pre',
                loader: 'svgo-loader?' + JSON.stringify(svgoConfig),
                include: /static\/img/
            },
            {
                test: /\.vue$/,
                loader: 'vue-loader',
                options: vueLoaderConfig,
                exclude: /node_modules\/(?!(autotrack|dom-utils))|vendor\.dll\.js/
            },
            {
                test: /\.js$/,
                loader: 'happypack/loader?id=happy-babel-js',
                //loader: 'babel-loader',
                include: [resolve('src'), resolve('test')],
                exclude: /node_modules/
            },
            {
                test: /\.(png|jpe?g|gif|svg)(\?.*)?$/,
                loader: 'url-loader',
                //include: [resolve('static/img')],
                options: {
                    limit: 10000,
                    name: utils.assetsPath('img/[name].[hash:7].[ext]')
                }
            },
            {
                test: /\.(woff2?|eot|ttf|otf)(\?.*)?$/,
                loader: 'url-loader',
                options: {
                    limit: 10000,
                    name: utils.assetsPath('fonts/[name].[hash:7].[ext]')
                }
            }
        ]
    },

    plugins: [
        new ProgressBarPlugin({
            format: '  build [:bar] ' + chalk.green.bold(':percent') + ' (:elapsed seconds)'
        }),
        createHappyPlugin('happy-babel-js', ['babel-loader?cacheDirectory=true']),
        createHappyPlugin('happy-babel-vue', ['babel-loader?cacheDirectory=true']),
        createHappyPlugin('happy-css', ['css-loader', 'vue-style-loader']),
        // https://github.com/amireh/happypack/pull/131
        new HappyPack({
            loaders: [{
                path: 'vue-loader',
                query: {
                    loaders: {
                        scss: 'vue-style-loader!css-loader!sass-loader?indentedSyntax'
                    }
                }
            }]
        })
    ]
}
