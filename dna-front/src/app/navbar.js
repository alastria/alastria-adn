// navbar-js

'use strict';

function NavbarController ($scope, $state, $location) {
  var vm = this;
  $scope.interevenerState = true;

  vm.$onInit = function () {
    isIntervener()
  }

  $scope.isActive = function (viewLocation) { 
    return viewLocation === $location.path();
  };

  function isIntervener() {
    var state = $state.router.urlRouter.location;
    if (state === '/regulator') {
      $scope.interevenerState = false;
    }
  }

}

angular
  .module('app')
  .component('navbar', {
    templateUrl: 'app/navbar.html',
    controller: NavbarController
  });