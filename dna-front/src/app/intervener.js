// intervener.js

'use strict';

function IntervenerController ($scope, $log, remresIntervener) {
  var vm = this;
  $scope.antistress = false;
  $scope.dataLoaded = false;

  vm.$onInit = function () {
    getLUAChainCodes();
    $scope.antistress = false;
    $scope.dataLoaded = true;
  }

  function getLUAChainCodes() {
    $scope.antistress = true;
    $log.debug('Getting LUA ChainCode´s list');
    remresIntervener.listLUAChainCode()
    .then(function (LUAList) {
      $scope.chaincodes = LUAList;
      $scope.antistress = false;
    }, function (err) {
      $scope.antistress = false;
      $log.error('Error -> ' + err);
    });
  }

  $scope.getInfoChaincode = function (chaincodeId) {
    $scope.antistress = true;
    $log.debug('Getting Chaincode info');
    remresIntervener.getChainCode(chaincodeId)
    .then(function (chaincodeData) {
      $scope.chaincode = chaincodeData;
      $scope.modalShowCode = true;
      $scope.antistress = false;
    }, function (err) {
      $scope.antistress = false;
      $log.error('Error -> ' + err);
    });
  };

  $scope.close = function () {
    // close model without saving
    if ($scope.modalShowCode === true) {
      $scope.modalShowCode = false;
    }
  };

}

angular
  .module('app')
  .component('intervener', {
    templateUrl: 'app/intervener.html',
    controller: IntervenerController
  });