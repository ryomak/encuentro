var webpack = require('webpack');
module.exports = {
    baseUrl: '/',
    assetsDir: 'assets',
    pages: {
        index: {
            // entry for the page
            entry: 'src/main.ts',
            // the source template
            template: 'public/index.html',
            // output as dist/index.html
            filename: 'index.html',
            // when using title option,
            // template title tag needs to be <title><%= htmlWebpackPlugin.options.title %></title>
            title: 'encuentro',
        },
        //arrange: './src/pages/arrange/app.js',
    },

    css: {
        loaderOptions: {
            css: {
                // options here will be passed to css-loader
            },
            sass: {
                includePaths: ['./src/assets', './src'],
            },
        },
    },

    chainWebpack: config => {
        config.module
            .rule('vue')
            .use('vue-loader')
            .loader('vue-loader')
            .tap(options => {
                options.compilerOptions.preserveWhitespace = true;
                return options;
            });
        config.module
            .rule('worker')
            .test(/\.worker\.js$/)
            .use('worker-loader')
            .loader('worker-loader')
            .options({
                //inline: true,
                name: 'worker.[hash].js',
            })
            .end();
    },

    configureWebpack: {
        plugins: [new webpack.ProvidePlugin({})],
        output: {
            globalObject: 'this',
        },
    },

    devServer: {
        port: 8080,
        proxy: {
            '/api': {
                target: process.env.USE_LOCAL_SERVER ? 'http://localhost:3000' : 'https://app.com',
                changeOrigin: true,
            },
        },
    },
};
