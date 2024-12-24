const path = require('path');
const MiniCssExtractPlugin = require('mini-css-extract-plugin');
const CopyWebpackPlugin = require('copy-webpack-plugin');

module.exports = {
    mode: "development",
    entry: './assets/src/index.ts', // 入口文件
    output: {
        filename: 'index.js', // 输出文件名
        path: path.resolve(__dirname, 'dist'), // 输出路径
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
    plugins: [
        new MiniCssExtractPlugin({
            filename: 'css/[name].css',
        }),
        new CopyWebpackPlugin({
            patterns: [
                {from: 'assets/src/img/*', to: 'img/[name][ext]'},
                {from: 'assets/src/css/*.css', to: 'css/[name].css'},
                {from: 'assets/src/config/*.json', to: 'config/[name].json'},
            ],
        }),
    ],
};
