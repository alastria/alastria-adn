// regulator.js

'use strict';

function RegulatorController($scope, $log, $state, remresRegulator) {
  var vm = this;
  $scope.dataLoaded = false;
  $scope.antistress = false;
  $scope.modalAssign = false;
  $scope.modalShowCode = false;
  $scope.modalLUA = false;
  $scope.modalModifyCode = false;
  $scope.btnExecute = true;
  $scope.countCC = '';

  vm.$onInit = function () {
    changeClass();
    getLUAChainCodes();
    $scope.dataLoaded = true;
  };

  function getTargets() {
    $scope.antistress = true;
    remresRegulator.getAllUsers()
    .then(function (allTargets) {
      if (allTargets !== null) {
        $scope.targets = []
        for (var i=0; i<allTargets.length; i++) {
          $scope.targets.push({Id: allTargets[i]});
        }
        $scope.antistress = false;
      } else {
        $scope.targets = []
        $scope.antistress = false;
      }
    }, function (err) {
      $scope.antistress = false;
      $log.error('Error -> ' + err);
    });
  }

  function saveTargets() {
    $scope.targetArray = [];
    angular.forEach($scope.targets, function (target) {
      if (target.selected) {
        $scope.targetArray.push(target.Id);
      }
    });
    return $scope.targetArray;
  }

  function composeSendData(name, luaCode) {
    var sendBody = {
      Name: name,
      SourceCode: luaCode.replace('\n', '\\n'),
      Targets: saveTargets()
    };
    return sendBody;
  }

  $scope.sendLUACode = function (name, luaCode) {
    $scope.antistress = true;
    remresRegulator.createLuaChaincode(composeSendData(name, luaCode))
    .then(function (Id) {
      $log.info('ID ' + JSON.stringify(Id) + 'Created successfuly');
      getLUAChainCodes();
      $scope.modalLUA = false;
    }, function (err) {
      $scope.antistress = false;
      $log.error('Error -> ' + err);
    });
  };

  function changeClass() {
    if ($state.router.urlRouter.location === '/regulator') {
      $scope.class = 'noNav';
    }
  }

  function getLUAChainCodes() {
    $log.debug('Getting LUA ChainCode´s list');
    $scope.antistress = true;
    remresRegulator.listLUAChainCode()
    .then(function (LUAList) {
      if (LUAList !== null) {
        calculateProgress(LUAList);
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

  $scope.getInfoChaincode = function (chaincodeId) {
    $scope.antistress = true;
    $log.debug('Getting Chaincode info');
    remresRegulator.getChainCode(chaincodeId)
    .then(function (chaincodeData) {
      $scope.chaincode = chaincodeData;
      $scope.modalShowCode = true;
      $scope.antistress = false;
    }, function (err) {
      $scope.antistress = false;
      $log.error('Error -> ' + err);
    });
  };

  function calculateProgress(LUAList) {
   //  obtain the percentage of ChainCodes signed
    for (var i = 0; i < Object.keys(LUAList).length; i++) {
      var total = LUAList[Object.keys(LUAList)[i]].Targets.length;
      var count = 0;
      for (var j = 0; j < total; j++) {
        if (LUAList[Object.keys(LUAList)[i]].Validations[j]) {
          count++;
        }
      }
      LUAList[Object.keys(LUAList)[i]].count = count;
      LUAList[Object.keys(LUAList)[i]].total = total;
      LUAList[Object.keys(LUAList)[i]].progress = count / total * 100;
    }
  }

  $scope.openModifyModal = function (chaincodeId) {
    $scope.antistress = true;
    $log.debug('Getting Chaincode info');
    remresRegulator.getChainCode(chaincodeId)
    .then(function (chaincodeData) {
      getTargets();
      $scope.chaincode = chaincodeData;
      $scope.modalModifyCode = true;
      $scope.antistress = false;
    }, function (err) {
      $scope.antistress = false;
      $log.error('Error -> ' + err);
    });
  };

  $scope.executeChaincode = function (Id) {
    $scope.antistress = true;
    remresRegulator.executeChaincode(Id)
    .then(function (executed) {
      console.log(executed);
      $scope.antistress = false;
    }, function(err) {
      $scope.antistress = false;
      $log.error('Error -> ' + err);
    })
  };

  // $scope.modifyChaincode = function (ccID, ccName, ccCode) {
  //   var Id = ccID;

  //   remresRegulator.updateLuaChaincode(Id)
  //   .then(function (updated) {
  //     console.log(updated);
  //   })
  // };

  $scope.uploadlLUA = function () {
    $scope.antistress = true;
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
    } else if ($scope.modalModifyCode === true) {
      $scope.modalModifyCode = false;
    }
  };
}

angular
  .module('app')
  .component('regulator', {
    templateUrl: 'app/regulator.html',
    controller: RegulatorController
  });
