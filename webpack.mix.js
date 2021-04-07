let mix = require('laravel-mix');

mix.js('./src/js/index.js', 'js')
    .sass('./src/scss/screen.scss', 'css')
    .setPublicPath('./static')
    .vue({ version: 3 })

mix.setResourceRoot('../')

mix.copyDirectory('./images', './static/images')

if (mix.inProduction()) {
    mix.version()
} else {
    mix.sourceMaps();
}