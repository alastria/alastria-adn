angular
  .module('app')
  .config(routesConfig);

/** @ngInject */
function routesConfig($stateProvider, $urlRouterProvider, $locationProvider) {
  $locationProvider.html5Mode(true).hashPrefix('!');
  $urlRouterProvider.otherwise('/register');

  $stateProvider
  .state('register', {
    url: '/register',
    component: 'registerIntervener',
    params: {data: null}
  })
    .state('intervener', {
      url: '/intervener',
      component: 'intervener',
      params: {data: null}
    });
}
