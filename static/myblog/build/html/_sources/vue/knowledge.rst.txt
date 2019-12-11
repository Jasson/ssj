Vue 积累 
########################################


跨域
----------------------------------------------------------------------------------


跨域::

  assetsSubDirectory: 'static',
  assetsPublicPath: '/',
  proxyTable: {
    // '/v2.0/': {
    //   // target: 'http://127.0.0.1:8443',
    //   target: 'https://deviot.langxw.com.cn',
    //   changeOrigin: true,
    //   pathRewrite: {
    //     '^/v2.0/': '/v2.0/'
    //   }
    //
    // },
    '/langxwid/': {
      target: 'http://127.0.0.1:8000',
      // target: 'https://deviot.langxw.com.cn',
      changeOrigin: true,
      pathRewrite: {
        '^/langxwid/': '/langxwid/'
      }
    }
  },
