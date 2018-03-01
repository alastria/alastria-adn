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
      component: 'regulator',
      params: {data: null}
    })
    .state('intervener', {
      url: '/intervener',
      component: 'intervener',
      params: {data: null}
    })
    .state('register', {
      url: '/register',
      component: 'registerIntervener',
      params: {data: null}
    });
}
