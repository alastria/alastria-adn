// regulator.js

'use strict';

function IntervenerController($scope, $log, $document, remresIntervener) {
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
      $document.getElementById('newUser').value = '';
      $scope.antistress = false;
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
    templateUrl: 'app/registerIntervener.html',
    controller: IntervenerController
  });
