let mix = require('laravel-mix');

mix.js('./src/js/index.js', 'js')
    .sass('./src/scss/screen.scss', 'css')
    .setPublicPath('./static')
    .vue({ version: 3 })

mix.setResourceRoot('../')

mix.copyDirectory('./images', './static/images')
mix.copyDirectory('./fonts', './static/fonts')

if (mix.inProduction()) {
    mix.version()
} else {
    mix.sourceMaps();
}