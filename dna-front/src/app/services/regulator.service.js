// regulator.service.js

'use string';

angular
  .module('app')
  .factory('remresRegulator', function ($http) {
    var regulatorService = {};
    regulatorService.listLUAChainCode = function () {
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

    regulatorService.getChainCode = function (chaincodeId) {
      var promise = $http({
        method: 'GET',
        url: '/v1/luaChaincode/' + chaincodeId
      })
      .then(function (response) {
        return response.data;
      }, function (error) {
        return error.data;
      });
      return promise;
    };

    return regulatorService;
  });
