// regulator.js

'use strict';

function RegulatorController($scope) {
  var vm = this;
  $scope.dataLoaded = false;
  $scope.antistress = false;
  $scope.modal = false;
  $scope.btnExecute = true;

  vm.$onInit = function () {
    // TODO
  };

  $scope.openModal = function (element) {
    // Obtains the file name, save it in a variable and open a modal to assign the ChainCode
    $scope.$apply(function ($scope) {
      $scope.chaincode = element.files[0].name;
      $scope.modal = true;
    });
  };

  // function calculateProgress(recos) {
  //  //  obtain the percentage of ChainCodes signed
  //   for (var i = 0; i < recos.length; i++) {
  //     var total = recos[i].managersSigns.length;
  //     var count = 0;
  //     for (var j = 0; j < total; j++) {
  //       if (recos[i].managersSigns[j].signed) {
  //         count++;
  //       }
  //     }
  //     recos[i].count = count;
  //     recos[i].total = total;
  //     recos[i].progress = count / total * 100;
  //     if (recos[i].progress === 100) {
  //       $scope.btnExecute = false;
  //     }
  //   }
  // }

  $scope.close = function () {
    // close the model without saving
    $scope.modal = false;
  };
}

angular
  .module('app')
  .component('regulator', {
    templateUrl: 'app/regulator.html',
    controller: RegulatorController
  });
