  'use strict';

  angular.module('app')
  .directive('preloader', function () {
    return {
      restrict: 'E',
      template: '<div ng-show="antistress">' +
                  '<div class="overlayContainer" style="display:block">' +
                    '<div class="overlayBackground"></div>' +
                    '<div class="overlayContent" style="top: 447px; left: 620px;">' +
                      '<div class="spinner">' +
                        '<div class="double-bounce1"></div>' +
                        '<div class="double-bounce2"></div>' +
                      '</div>' +
                    '</div>' +
                  '</div>' +
                '</div>'
    };
  });
