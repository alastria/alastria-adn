// showCodeModal.js

'use strict';

angular
  .module('app')
  .directive('modalModify', function () {
    return {
      restrict: 'E',
      template: '<div class"fades" style="position: fixed;top: 0;right: 0;bottom: 0;left: 0;z-index: 1050;display: block;overflow-y: scroll;outline: 0; height: auto; padding-top:10%; background-color:rgba(0, 0, 0, .5);" ng-show="modalModifyCode">' +
                  '<div class="modal-dialog" role="document">' +
                    '<div class="modal-content">' +
                      '<div class="modal-header">' +
                      '<h3 class="modal-title">Upload LUA Code</h3>' +
                      '<button type="button" ng-click="close()" class="close" data-dismiss="modal" aria-label="Close">' +
                        '<span aria-hidden="true">&times;</span>' +
                      '</button>' +
                      '</div>' +
                      '<div class="modal-body">' +
                        '<label for="idChaincode">Chaincode Id</label>' +
                        '<input id="idChaincode" type="text" class="form-control" placeholder="Chaincode Id" value="{{chaincode.LuaChaincodeId}}" ng-if="chaincode.LuaChaincodeId" readonly><br/>' +
                        '<label for="fileName">File Name</label>' +
                        '<input id="fileName" type="text" class="form-control" placeholder="File name" ng-model="chaincode.Name" ng-if="chaincode.Name"><br/>' +
                        '<table class="table" ng-if="targets">' +
                          '<thead>' +
                          '<tr>' +
                          '<th style="width:10%; text-align:center">Assigned</th>' +
                          '<th style="width:90%; margin-left:30px">Responsible</th>' +
                          '</tr>' +
                          '</thead>' +
                          '<tbody>' +
                            '<tr ng-repeat="target in targets">' +
                              '<td style="width:10%; text-align:center"><input type="checkbox" ng-model="target.selected" value="{{target.Id}}"></td>' +
                              '<td style="width:90%; margin-left:30px">{{target.Id}}</td>' +
                            '</tr>' +
                          '</tbody>' +
                        '</table>' +
                        '<label for="LUAcode" style="text-align:center" ng-if="chaincode.SourceCode">Code:</label>' +
                        '<textarea rows="10" id="LUAcode" class="form-control" style="box-sizing: border-box;resize: none" ng-model="chaincode.SourceCode" placeholder="Insert LUA code here..." ng-if="chaincode.SourceCode"></textarea>' +
                      '</div>' +
                      '<div class="modal-footer">' +
                        '<button type="button" ng-click="close()" class="btn btn-secondary" data-dismiss="modal">Close</button>' +
                        '<button type="button" ng-click="modifyChaincode(chaincode.LuaChaincodeId, chaincode.Name, chaincode.SourceCode)" class="btn btn-primary">Modify</button>' +
                      '</div>' +
                    '</div>' +
                  '</div>' +
                '</div>'
    };
  });
