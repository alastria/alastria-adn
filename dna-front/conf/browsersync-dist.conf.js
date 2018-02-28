const conf = require('./gulp.conf');
var proxyMiddleware = require('http-proxy-middleware');

module.exports = function () {
  var middleware = proxyMiddleware('/v1', {target: 'http://localhost:8083/', changeOrigin: 'localhost'});

  return {
    server: {
      baseDir: [
        conf.paths.dist
      ],
      middleware: middleware
    },
    open: false
  };
};