const conf = require('./gulp.conf');
// var proxyMiddleware = require('http-proxy-middleware');

module.exports = function () {
  // var middleware = proxyMiddleware('/', {target: 'http://localhost:8080/', changeOrigin: 'localhost'});

  return {
    server: {
      baseDir: [
        conf.paths.dist
      ],
      // middleware: middleware
    },
    open: false
  };
};