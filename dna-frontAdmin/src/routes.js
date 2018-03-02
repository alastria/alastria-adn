angular
  .module('app')
  .config(routesConfig);

/** @ngInject */
function routesConfig($stateProvider, $urlRouterProvider, $locationProvider) {
  $locationProvider.html5Mode(true).hashPrefix('!');
  $urlRouterProvider.otherwise('/regulator');

  $stateProvider
    .state('regulator', {
      url: '/regulator',
      component: 'regulator'
    });
}
