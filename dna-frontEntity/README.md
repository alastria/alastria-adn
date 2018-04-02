# Alastria DNA - Entity Console

Entity console is installed once a time for each entity that conforms the network.

## Configuration

If we need to change the port in which the app is deployed, we need to change the target field of the following files:

### /conf/browsersync-dist.conf.js

``` code
  var middleware = proxyMiddleware('/v1', { target: 'http://localhost:8081/', changeOrigin: 'localhost'});
```

### /conf/browsersync.conf.js

``` code
  var middleware = proxyMiddleware('/v1', { target: 'http://localhost:8081/', changeOrigin: 'localhost'});
```

The default port is **8081**

## Run Dna Alastria FrontEnd for the entity

You need to clone the folder if you need more entities

``` bash
cd ./dna-frontEntity
npm install
bower install
gulp serve
```