var webpack = require('webpack');
module.exports = {
    baseUrl: '/',
    assetsDir: 'assets',
    devServer: {
        port: 8111,
        proxy: {
            '/api': {
                target: process.env.USE_LOCAL_SERVER ? 'http://localhost:3000' : 'https://aaa.com',
                changeOrigin: true,
            },
        },
    },
    configureWebpack: {
        output: {
            globalObject: 'this',
        },
    },

};
