// uploadLUACodeModal.js

'use strict';

angular
  .module('app')
  .directive('modalLuacode', function () {
    return {
      restrict: 'E',
      template: '<div class"fades" style="position: fixed;top: 0;right: 0;bottom: 0;left: 0;z-index: 1050;display: block;overflow-y: scroll;outline: 0; height: auto; padding-top:10%; background-color:rgba(0, 0, 0, .5);" ng-show="modalLUA">' +
                  '<div class="modal-dialog" role="document">' +
                    '<form name="uploadCC">' +
                      '<div class="modal-content">' +
                        '<div class="modal-header">' +
                        '<h3 class="modal-title">Upload LUA Code</h3>' +
                        '<button type="button" ng-click="close()" class="close" data-dismiss="modal" aria-label="Close">' +
                          '<span aria-hidden="true">&times;</span>' +
                        '</button>' +
                        '</div>' +
                        '<div class="modal-body">' +
                          '<label for="fileName">Code name:</label>' +
                          '<input id="fileName" type="text" class="form-control" ng-model="LUAname"placeholder="Code name" required><br/>' +
                          '<table class="table">' +
                            '<thead>' +
                            '<tr>' +
                            '<th style="width:10%; text-align:center">Assigned</th>' +
                            '<th style="width:90%; margin-left:30px">Responsible</th>' +
                            '</tr>' +
                            '</thead>' +
                            '<tbody>' +
                              '<tr ng-repeat="target in targets">' +
                                '<td style="width:10%; text-align:center"><input type="checkbox" ng-model="target.selected" value="{{target.Id}}" required></td>' +
                                '<td style="width:90%; margin-left:30px">{{target.Id}}</td>' +
                              '</tr>' +
                            '</tbody>' +
                          '</table>' +
                          '<label for="LUAcode" style="text-align:center">Code:</label>' +
                          '<textarea rows="10" id="LUAcode" class="form-control" style="box-sizing: border-box;resize: vertical;" ng-model="LUAsource" placeholder="Insert LUA code here..." required></textarea>' +
                        '</div>' +
                        '<div class="modal-footer">' +
                          '<button type="button" ng-click="close()" class="btn btn-secondary" data-dismiss="modal">Close</button>' +
                          '<button type="button" ng-click="sendLUACode(LUAname, LUAsource)" ng-disabled="uploadCC.$invalid" class="btn btn-primary">Assign ChainCode</button>' +
                        '</div>' +
                      '</div>' +
                    '</form>' +
                  '</div>' +
                '</div>'

    };
  });
