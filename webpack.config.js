const path = require('path');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const TerserWebpackPlugin = require('terser-webpack-plugin');
const webpack = require('webpack');

module.exports = (env, argv) => {
    const isProd = argv.mode === 'production'; // 根据命令行参数判断是否为生产模式

    const configFile = isProd ? 'global.prod.json' : 'global.dev.json';

    return {
        entry: './assets/src/index.ts',
        output: {
            filename: 'main.js',
            path: path.resolve(__dirname, isProd ? 'dist/prod' : 'dist/dev'),
            clean: true,
        },
        resolve: {
            extensions: ['.ts', '.js'],
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
            minimize: isProd, // 仅在生产模式下启用压缩
            minimizer: [new TerserWebpackPlugin()],
        },
        plugins: [
            new HtmlWebpackPlugin({
                template: './assets/html/index.html',
            }),
            new webpack.DefinePlugin({
                'process.env.NODE_ENV': JSON.stringify(argv.mode), // 使用命令行传入的 mode
                'process.env.CONFIG_FILE': JSON.stringify(configFile),
            }),
        ],
        devtool: isProd ? false : 'source-map',
        mode: argv.mode || 'development', // 使用命令行传入的 mode，默认为 development
    };
};
