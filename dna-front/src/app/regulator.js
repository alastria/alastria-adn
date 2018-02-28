// regulator.js

'use strict';

function RegulatorController($scope, $log, remresRegulator) {
  var vm = this;
  $scope.dataLoaded = false;
  $scope.antistress = false;
  $scope.modalAssign = false;
  $scope.modalShowCode = false;
  $scope.modalLUA = false;
  $scope.btnExecute = true;
  $scope.countCC = '';

  vm.$onInit = function () {
    getLUAChainCodes();
  };

  $scope.openModal = function (element) {
    // Obtains the file name, save it in a variable and open a modal to assign the ChainCode
    $scope.$apply(function ($scope) {
      $scope.myChaincode = element.files[0].name;
      $scope.modalAssign = true;
    });
  };

  function getTargets() {
    remresRegulator.getAllUsers()
    .then(function (allTargets) {
      $scope.targets = []
      for(var i=0; i<allTargets.length; i++){
        $scope.targets.push({Id: allTargets[i]});
      }
    }, function (err) {
      $log.error('Error -> ' + err)
    })
  };

  function saveTargets() {
    $scope.targetArray = [];

    angular.forEach($scope.targets, function(target) {
      if (target.selected) {
        $scope.targetArray.push(target.Id)
      };
    });
    return $scope.targetArray;
  };
  
  function composeSendData(name, luaCode) {
    var sendBody = {
      'Name': name,
      'SourceCode': luaCode.replace("\n"," "),
      'Targets': saveTargets()
    };
    return sendBody
  };

  $scope.sendLUACode = function (name, luaCode) {
    remresRegulator.createLuaChaincode(composeSendData(name, luaCode))
    .then(function (Id) {
      console.log(Id);

      getLUAChainCodes();
      $scope.modalLUA = false;
    }, function(err){
      $log.error('Error -> ' + err)
    })
    
  };

  function getLUAChainCodes() {
    $log.debug('Getting LUA ChainCode´s list');
    remresRegulator.listLUAChainCode()
    .then(function (LUAList) {
      calculateProgress(LUAList);
      $scope.countCC = Object.keys(LUAList).length; // obtain number of chaincodes
      $scope.chaincodes = LUAList;
    }, function (err) {
      $log.error('Error -> ' + err);
    });
  };

  $scope.showCode = function (chaincodeId) {
    $log.debug('Getting Chaincode info');
    remresRegulator.getChainCode(chaincodeId)
    .then( function (chaincodeData) {
      console.log(chaincodeData);
      $scope.chaincode = chaincodeData;
      $scope.modalShowCode = true;
    }, function (error) {
      $log.error('Error -> ' + err);
    })
  };

  function calculateProgress(LUAList) {
   //  obtain the percentage of ChainCodes signed
    for (var i = 0; i < Object.keys(LUAList).length; i++){
      var total = LUAList[Object.keys(LUAList)[i]].Targets.length;
      var count = 0;
      for (var j = 0; j < total; j++) {
        if(LUAList[Object.keys(LUAList)[i]].Validations[j]){
          count++;
        }
      }
      LUAList[Object.keys(LUAList)[i]].count = count;
      LUAList[Object.keys(LUAList)[i]].total = total;
      LUAList[Object.keys(LUAList)[i]].progress = count / total * 100;
    }
  };

  $scope.uploadlLUA = function () {
    // Open Uploader of LUA code
    getTargets();
    $scope.modalLUA = true;
  };

  $scope.close = function () {
    // close model without saving
    if ($scope.modalLUA === true) {
      $scope.modalLUA = false;
    } else if ($scope.modalAssign === true) {
      $scope.modalAssign = false;
    } else if ($scope.modalShowCode === true) {
      $scope.modalShowCode = false;
    }
  };
}

angular
  .module('app')
  .component('regulator', {
    templateUrl: 'app/regulator.html',
    controller: RegulatorController
  });
