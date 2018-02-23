// modalAssignCC.js

'use strict';

angular
  .module('app')
  .directive('modalAssigncc', function () {
    return {
      restrict: 'E',
      template: '<div class"fades" style="position: fixed;top: 0;right: 0;bottom: 0;left: 0;z-index: 1050;display: block;overflow: hidden;-webkit-overflow-scrolling: touch;outline: 0; padding-top:10%; background-color:rgba(0, 0, 0, .5)" ng-show="modalAssign">' +
                  '<div class="modal-dialog" role="document">' +
                    '<div class="modal-content">' +
                      '<div class="modal-header">' +
                      '<h3 class="modal-title">Assign ChainCode</h3>' +
                      '<button type="button" ng-click="close()" class="close" data-dismiss="modal" aria-label="Close">' +
                        '<span aria-hidden="true">&times;</span>' +
                      '</button>' +
                      '</div>' +
                      '<div class="modal-body">' +
                        '<table class="table">' +
                          '<thead>' +
                          '<tr>' +
                          '<th style="width:10%; text-align:center">Assigned</th>' +
                          '<th style="width:90%; margin-left:30px">Responsible</th>' +
                          '</tr>' +
                          '</thead>' +
                          '<tbody>' +
                            // '<tr ng-repeat="">' +
                            '<tr>' +
                              // '<td style="width:10%; text-align:center"><input type="checkbox" ng-model="" value="{{}}"></td>' +
                              '<td style="width:10%; text-align:center"><input type="checkbox"></td>' +
                              '<td style="width:90%; margin-left:30px">{{}}el tipo que sea</td>' +
                            '</tr>' +
                            '<tr>' +
                              // '<td style="width:10%; text-align:center"><input type="checkbox" ng-model="" value="{{}}"></td>' +
                              '<td style="width:10%; text-align:center"><input type="checkbox"></td>' +
                              '<td style="width:90%; margin-left:30px">{{}}el tipo que sea</td>' +
                            '</tr>' +
                          '</tbody>' +
                        '</table>' +
                      '</div>' +
                      '<div class="modal-footer">' +
                        '<button type="button" ng-click="close()" class="btn btn-secondary" data-dismiss="modal">Close</button>' +
                        '<button type="button" ng-click="" class="btn btn-primary">Assign ChainCode</button>' +
                      '</div>' +
                    '</div>' +
                  '</div>' +
                '</div>'

    };
  });
