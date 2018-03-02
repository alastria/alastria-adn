// showResultModal.js

'use strict';

angular
  .module('app')
  .directive('modalResult', function () {
    return {
      restrict: 'E',
      template: '<div class"fades" style="position: fixed;top: 0;right: 0;bottom: 0;left: 0;z-index: 1050;display: block;overflow: hidden;-webkit-overflow-scrolling: touch;outline: 0; height: auto; padding-top:10%; background-color:rgba(0, 0, 0, .5);" ng-show="modalShowExecution">' +
                  '<div class="modal-dialog" role="document">' +
                    '<div class="modal-content">' +
                      '<div class="modal-header">' +
                      '<h3 class="modal-title">Result Execution Chaincode</h3>' +
                      '<button type="button" ng-click="close()" class="close" data-dismiss="modal" aria-label="Close">' +
                        '<span aria-hidden="true">&times;</span>' +
                      '</button>' +
                      '</div>' +
                      '<div class="modal-body">' +
                        '<table class="table">' +
                        '<thead>' +
                          '<tr>' +
                            '<th style="width:30%; text-align:center">Organization Name</th>' +
                            '<th style="width:70%; margin-left:30px">Result</th>' +
                          '</tr>' +
                        '</thead>' +
                        '<tbody>' +
                          '<tr ng-repeat="result in results">' +
                            '<td style="width:30%; text-align:center">{{result.orgName}}</td>' +
                            '<td style="width:70%; margin-left:30px">{{result.executionResult}}</td>' +
                          '</tr>' +
                        '</tbody>' +
                      '</table>' +
                      '</div>' +
                      '<div class="modal-footer">' +
                        '<button type="button" ng-click="close()" class="btn btn-secondary" data-dismiss="modal">Close</button>' +
                      '</div>' +
                    '</div>' +
                  '</div>' +
                '</div>'
    };
  });
