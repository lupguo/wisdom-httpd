const path = require('path');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const TerserWebpackPlugin = require('terser-webpack-plugin');

module.exports = {
    entry: './assets/src/index.ts', // TypeScript 入口文件
    output: {
        filename: 'main.js', // 输出文件名
        path: path.resolve(__dirname, 'dist'), // 输出路径
        clean: true, // 清理输出目录
    },
    resolve: {
        extensions: ['.ts', '.js'], // 支持的文件扩展名
    },
    module: {
        rules: [
            {
                test: /\.ts$/,
                use: 'ts-loader',
                exclude: /node_modules/,
            },
            {
                test: /\.css$/,
                use: ['style-loader', 'css-loader'],
            },
            {
                test: /\.(eot|svg|ttf|woff|woff2|png|jpg|gif)$/i,
                type: 'asset/resource',
                generator: {
                    filename: 'imgs/[name]_[hash][ext]',
                },
            },
        ],
    },
    optimization: {
        minimize: true,
        minimizer: [new TerserWebpackPlugin()],
    },
    plugins: [
        new HtmlWebpackPlugin({
            template: './assets/html/index.html', // 使用的 HTML 模板
        }),
    ],
    devtool: 'source-map', // 生成源映射文件
    mode: "development"
};
