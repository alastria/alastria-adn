#  Alastria DNA - Admin Console

Admin console is installed once a time by the network admin.

## Configuration

If we need to change the port in which the app is deployed, we need to change the target field of the following files:
### /conf/browsersync-dist.conf.js

var middleware = proxyMiddleware('/v1', { **target: 'http://localhost:8083/'**, changeOrigin: 'localhost'});

### /conf/browsersync.conf.js

var middleware = proxyMiddleware('/v1', { **target: 'http://localhost:8083/'**, changeOrigin: 'localhost'});

The default port is **8083**

## Run Dna Alastria FrontEnd for the admin
``` bash
cd ./dna-frontAdmin
npm install
bower install
gulp serve
```