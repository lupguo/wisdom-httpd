const path = require('path');
const MiniCssExtractPlugin = require('mini-css-extract-plugin');
const CopyWebpackPlugin = require('copy-webpack-plugin');
const TerserPlugin = require('terser-webpack-plugin');
const webpack = require('webpack')

module.exports = {
    entry: './assets/src/index.ts', // 入口文件
    output: {
        filename: 'bundle.js', // 输出文件名
        path: path.resolve(__dirname, 'dist'), // 输出路径
        publicPath: '/dist', // 这对 HMR 是必要的
    },
    devServer: {
        static: {
            directory: path.join(__dirname, 'dist'), // 提供静态文件的目录
        },
        hot: true, // 启用热模块替换
        port: 3000,
        open: true, // 启动后自动打开浏览器
    },
    resolve: {
        extensions: ['.ts', '.js'], // 支持的文件扩展名
    },
    module: {
        rules: [
            {
                test: /\.ts$/, // 处理 TypeScript 文件
                use: 'ts-loader',
                exclude: /node_modules/,
            },
            {
                test: /\.css$/, // 处理 CSS 文件
                use: ['style-loader', 'css-loader'],
                generator: {
                    filename: 'css/[name][ext]', // 输出图片路径
                },
            },
            {
                test: /\.(eot|svg|ttf|woff|woff2|png|jpg|gif)$/i, // 处理图片文件
                type: 'asset/resource',
                use: ['file-loader'],
                generator: {
                    filename: 'imgs/[name][hash][ext]', // 输出图片路径
                },
            },
        ],
    },
    optimization: {
        minimize: true,
        minimizer: [new TerserPlugin()],
    },

    mode: 'development', // 开发模式
    plugins: [
        new webpack.HotModuleReplacementPlugin(), // Add HMR plugin
        new MiniCssExtractPlugin({
            filename: 'css/[name].css',
        }),
        new CopyWebpackPlugin({
            patterns: [
                {from: 'assets/src/imgs/*', to: 'imgs/[name][ext]'},
                {from: 'assets/src/css/*.css', to: 'css/[name].css'},
                {from: 'assets/src/config/*.json', to: 'config/[name].json'},
            ],
        }),
    ],
};
