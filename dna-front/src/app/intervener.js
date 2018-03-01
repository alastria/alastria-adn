// intervener.js

'use strict';

function IntervenerController ($scope, $log, remresIntervener) {
  var vm = this;
  $scope.antistress = false;
  $scope.msgApproved = false;

  vm.$onInit = function () {
    getLUAChainCodes();
  }

  function getLUAChainCodes() {
    $log.debug('Getting LUA ChainCode´s list');
    $scope.antistress = true;
    remresIntervener.listLUAChainCode()
    .then(function (LUAList) {
      if (LUAList !== null){
        $scope.countCC = Object.keys(LUAList).length; // obtain number of chaincodes
        $scope.chaincodes = LUAList;
        $scope.antistress = false;
      } else {
        $scope.countCC = 0;
        $scope.antistress = false;
      }
    }, function (err) {
      $scope.antistress = false;
      $log.error('Error -> ' + err);
    });
  }

  $scope.approveChaincode = function (Id) {
    $scope.antistress = true;
    remresIntervener.valdateChaincode(Id)
    .then(function (approved) {
      if (approved === 'update success!' ) {
        $scope.approve = 'Chaincode ' + Id + ' approved succesfully';
      }
      $scope.msgApproved = true;
      $scope.antistress = false;
    }, function(err) {
      $scope.antistress = false;
      $log.error('Error -> ' + err);
    });
  };

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
    if ($scope.modalShowCode === true) {
      $scope.modalShowCode = false;
    } else if ($scope.msgApproved === true) {
      $scope.msgApproved = false;
    }
  };

}

angular
  .module('app')
  .component('intervener', {
    templateUrl: 'app/intervener.html',
    controller: IntervenerController
  });