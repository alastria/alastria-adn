
function HeaderController($scope, $log, $document, remresIntervener) {
  this.$onInit = function () {
    getCurrentUser();
  };

  function getCurrentUser() {
    $log.debug('Getting Organization name');
    remresIntervener.getCurrentUser()
    .then(function (res) {
      $scope.orgName = res;
      $document[0].title = res + " Console";
    }, function (err) {
      $log.error('Error -> ' + err);
    });
  }
}

angular
  .module('app')
  .component('fountainHeader', {
    templateUrl: 'app/header.html',
    controller: HeaderController
  });
