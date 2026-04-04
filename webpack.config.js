const fs = require('fs')
const path = require('path')
const HtmlWebpackPlugin = require('html-webpack-plugin')

/** Копирует src/dist/favicon.ico → output/favicon.ico (всегда docs/). */
class CopyFaviconIcoPlugin {
    apply(compiler) {
        compiler.hooks.afterEmit.tapAsync('CopyFaviconIcoPlugin', (compilation, callback) => {
            const from = path.resolve(__dirname, 'src/dist/favicon.ico')
            const to = path.join(compilation.options.output.path, 'favicon.ico')
            try {
                if (!fs.existsSync(from)) {
                    console.warn(`[webpack] CopyFaviconIcoPlugin: missing ${from} (skipping copy)`)
                    return callback()
                }
                // При webpack serve каталог output (docs/) может ещё не существовать на диске.
                fs.mkdirSync(compilation.options.output.path, { recursive: true })
                fs.copyFileSync(from, to)
            } catch (err) {
                return callback(err)
            }
            callback()
        })
    }
}

// Сборка фронта всегда в docs/ (GitHub Pages / единая папка артефактов).
module.exports = (env, argv) => {
    const prod = argv.mode === 'production'
    const outDir = 'docs'
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
        new HtmlWebpackPlugin({
            template: './src/index.html',
            // Иконка: ./favicon.ico из шаблона + копирование CopyFaviconIcoPlugin из src/dist/favicon.ico
        }),
        new CopyFaviconIcoPlugin(),
    ]
    }
}