// regulator.js

'use strict';

function IntervenerController($scope, $log, $state, remresIntervener) {
  $scope.haveUser = false;

  $scope.registerIntervener = function (name) {
    var Name = {
      Id: name
    }
    remresIntervener.createUser(Name)
    .then(function (registered) {
      $scope.user = registered;
      $scope.haveUser = true;
      document.getElementById('newUser').value = '';
    },function(err) {
      $log.error('Error -> ' + err);
    })
  };

  function setInputEmpty() {

  }

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
