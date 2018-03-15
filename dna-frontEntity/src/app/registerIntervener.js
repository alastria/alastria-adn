// regulator.js

'use strict';

function IntervenerController($scope, $log, $window, remresIntervener) {
  $scope.haveUser = false;
  $scope.antistress = false;

  $scope.registerIntervener = function (name) {
    $scope.antistress = true;
    var Name = {
      Id: name
    };
    remresIntervener.createUser(Name)
    .then(function (registered) {
      $scope.user = registered;
      $scope.haveUser = true;
      $scope.antistress = false;
      $window.document.getElementById('registerUser').reset();
    }, function (err) {
      $log.error('Error -> ' + err);
    });
  };

  $scope.close = function () {
    $scope.haveUser = false;
  };
}

angular
  .module('app')
  .component('registerIntervener', {
    // templateUrl: 'app/registerIntervener.html',
    controller: IntervenerController
  });
