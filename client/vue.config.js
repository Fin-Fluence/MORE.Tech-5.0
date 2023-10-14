const { defineConfig } = require('@vue/cli-service')

module.exports = defineConfig({
  transpileDependencies: true,
  configureWebpack: {
    externals: {
      "ymaps3": "ymaps3"
    }
  }
})
