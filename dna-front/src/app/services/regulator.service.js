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

    regulatorService.createLuaChaincode = function (objectData) {
      var promise = $http({
        method: 'POST',
        url: '/v1/luaChaincode/',
        data: objectData
      })
      .then(function (response) {
        return response.data;
      }, function (error) {
        return error.data;
      });
      return promise;
    };

    // regulatorService.updateLuaChaincode = function (Id, objectData) {
    //   var promise = $http({
    //     method: 'PUT',
    //     url: '/v1/luaChaincode/' + Id,
    //     data: objectData
    //   })
    //   .then(function (response) {
    //     return response.data;
    //   }, function (error) {
    //     return error.data;
    //   });
    //   return promise;
    // };

    regulatorService.getAllUsers = function () {
      var promise = $http({
        method: 'GET',
        url: '/v1/user/'
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
