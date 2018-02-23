// uploadLUACodeModal.js

'use strict';

angular
  .module('app')
  .directive('modalLuacode', function () {
    return {
      restrict: 'E',
      template: '<div class"fades" style="position: fixed;top: 0;right: 0;bottom: 0;left: 0;z-index: 1050;display: block;overflow: hidden;-webkit-overflow-scrolling: touch;outline: 0; height: auto; padding-top:10%; background-color:rgba(0, 0, 0, .5);" ng-show="modalLUA">' +
                  '<div class="modal-dialog" role="document">' +
                    '<div class="modal-content">' +
                      '<div class="modal-header">' +
                      '<h3 class="modal-title">Upload LUA Code</h3>' +
                      '<button type="button" ng-click="close()" class="close" data-dismiss="modal" aria-label="Close">' +
                        '<span aria-hidden="true">&times;</span>' +
                      '</button>' +
                      '</div>' +
                      '<div class="modal-body">' +
                        '<label for="fileName">File name:</label>' +
                        '<input id="fileName" type="text" class="form-control" placeholder="File name">' +
                        '<label for="LUAcode" style="text-align:center">Code:</label>' +
                        '<textarea rows="10" id="LUAcode" class="form-control" style="box-sizing: border-box;resize: none" placeholder="Insert LUA code here..."></textarea>' +
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
