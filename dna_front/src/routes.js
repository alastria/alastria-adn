angular
  .module('app')
  .config(routesConfig);

/** @ngInject */
function routesConfig($stateProvider, $urlRouterProvider, $locationProvider) {
  $locationProvider.html5Mode(true).hashPrefix('!');
  $urlRouterProvider.otherwise('/index');

  $stateProvider
    .state('app', {
      url: '/index',
      component: 'fountainTitle',
      params: {data: null}
    });
}
