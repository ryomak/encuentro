module.exports = {

  modules: [
    '@nuxtjs/axios',
    '@nuxtjs/proxy',
    // Or if you have custom bootstrap CSS...
    ['bootstrap-vue/nuxt'],
  ],
  plugins: [
    '~/plugins/axios'
  ],

  /*
  ** Headers of the page
  */
  head: {
    title: 'encuentro',
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { hid: 'description', name: 'description', content: 'Nuxt.js project' }
    ],
    link: [
      { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }
    ]
  },
  /*
  ** Customize the progress bar color
  */
  loading: { color: '#3B8070' },
  /*
  ** Build configuration
  */
  build: {
    /*
    ** Run ESLint on save
    */
    extend (config, { isDev, isClient }) {
      if (isDev && isClient) {
        config.module.rules.push({
          enforce: 'pre',
          test: /\.(js|vue)$/,
          loader: 'eslint-loader',
          exclude: /(node_modules)/
        })
      }
    }
  },
  axios:{
    proxy:true,
  },
  proxy:{
    '/api/': process.env.USE_LOCAL_SERVER ? 'http://localhost:3000' : 'http://localhost:3000',
  }
}
