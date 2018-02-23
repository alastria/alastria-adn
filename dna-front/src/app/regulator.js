// regulator.js

'use strict';

function RegulatorController($scope, $log, remresRegulator) {
  var vm = this;
  $scope.dataLoaded = false;
  $scope.antistress = false;
  $scope.modalAssign = false;
  $scope.modalLUA = false;
  $scope.btnExecute = true;
  $scope.countCC = '';

  vm.$onInit = function () {
    getLUAChainCodes();
  };

  $scope.openModal = function (element) {
    // Obtains the file name, save it in a variable and open a modal to assign the ChainCode
    $scope.$apply(function ($scope) {
      $scope.chaincode = element.files[0].name;
      $scope.modalAssign = true;
    });
  };

  function getLUAChainCodes() {
    $log.debug('Getting LUA ChainCode´s list');
    remresRegulator.listLUAChainCode()
    .then(function (LUAList) {
      calculateProgress(LUAList);
      $scope.countCC = Object.keys(LUAList).length; // obtain number of chaincodes
      $scope.chaincodes = LUAList;
    }, function (err) {
      $log.error('Getting LUA ChainCode´s list');
    });
  };

  $scope.showCode = function (chaincodeId) {
    $log.debug('Getting Chaincode info');
    remresRegulator.getChainCode(chaincodeId)
    .then( function (chaincodeData) {
      console.log(chaincodeData);
      swal({
        title: '<u style="text-align:left">' + chaincodeData.Name + '</u>',
        html:
          'Code: ' +
          '<textarea class="form-control">' + chaincodeData.SourceCode + '</textarea>',
        showCloseButton: true,
      });
    }, function (error) {
      console.log(error);
      
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

  $scope.openModalLUA = function () {
    // Open Uploader of LUA code
    $scope.modalLUA = true;
  };

  // $scope.asignCC = function () {
  // };
  $scope.close = function () {
    // close model without saving
    if ($scope.modalLUA === true){
      $scope.modalLUA = false;
    } else if ($scope.modalAssign === true) {
      $scope.modalAssign = false;
    }
  };
}

angular
  .module('app')
  .component('regulator', {
    templateUrl: 'app/regulator.html',
    controller: RegulatorController
  });
