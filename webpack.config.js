const path = require('path')
const HtmlWebpackPlugin = require('html-webpack-plugin')

// SITE_OUT=docs → GitHub Pages (mage Site). Иначе dist/ для dev и yarn build.
module.exports = (env, argv) => {
    const prod = argv.mode === 'production'
    const outDir = process.env.SITE_OUT === 'docs' ? 'docs' : 'dist'
    return {
    entry: './src/js/main.js',
    output: {
        filename: 'main.js',
        path: path.resolve(__dirname, outDir),
        publicPath: prod ? './' : 'auto',
    },
    devServer: {
        // По умолчанию webpack-dev-server берёт local-ip и не слушает 127.0.0.1 — на localhost:8080
        // остаётся другой процесс (старый сайт). 0.0.0.0 — один dev-сервер и на localhost, и по IP WSL.
        host: '0.0.0.0',
        allowedHosts: 'all',
        static: path.resolve(__dirname, 'src/dist'),
        port: 8080,
        hot: true
    },
    module: {
        rules: [
            {
                test: /\.(s?css)$/,
                use: [
                    {
                        loader: 'style-loader'
                    },
                    {
                        loader: 'css-loader'
                    },
                    {
                        loader: 'postcss-loader',
                        options: {
                            postcssOptions: {
                                plugins: () => [
                                    require('autoprefixer')
                                ]
                            }
                        }
                    },
                    {
                        loader: 'sass-loader'
                    }
                ]
            }
        ]
    },
    plugins: [
        new HtmlWebpackPlugin({ template: './src/index.html' })
    ]
    }
}