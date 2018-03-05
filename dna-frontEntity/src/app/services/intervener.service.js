// regulator.service.js

'use string';

angular
  .module('app')
  .factory('remresIntervener', function ($http) {
    var intervenerService = {};

    intervenerService.listLUAChainCode = function () {
      var promise = $http({
        method: 'GET',
        url: '/v1/luaChaincode'
      })
      .then(function (response) {
        return response.data;
      }, function (Error) {
        return Error.data;
      });
      return promise;
    };

    intervenerService.createUser = function (objectData) {
      var promise = $http({
        method: 'POST',
        url: '/v1/user',
        data: objectData
      })
      .then(function (response) {
        return response.data;
      }, function (Error) {
        return Error.data;
      });
      return promise;
    };

    intervenerService.getChainCode = function (chaincode) {
      var promise = $http({
        method: 'GET',
        url: '/v1/luaChaincode/' + chaincode
      })
      .then(function (response) {
        return response.data;
      }, function (Error) {
        return Error.data;
      });
      return promise;
    };

    intervenerService.valdateChaincode = function (chaincode) {
      var promise = $http({
        method: 'PUT',
        url: '/v1/luaChaincode/' + chaincode
      })
      .then(function (response) {
        return response.data;
      }, function (Error) {
        return Error.data;
      });
      return promise;
    };

    return intervenerService;
  });
