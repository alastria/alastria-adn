package models

import (
	"errors"
)

type LuaChaincodePost struct {
	Name string
	SourceCode string
	Targets []string
}

type LuaChaincode struct {
	LuaChaincodeId   string
	Name string
	SourceCode string
	Targets []string
	Validations []bool
}

func init() {
}

func AddOne(luaChaincode LuaChaincodePost) (LuaChaincodeId string) {
    return fabric_add_code(luaChaincode.Name, luaChaincode.SourceCode, luaChaincode.Targets)
}

func GetOne(LuaChaincodeId string) (luaChaincode LuaChaincode, err error) {
    return fabric_query_code(LuaChaincodeId), nil
}

func GetAll() []LuaChaincode {
	return fabric_query_codes()
}

func Update(LuaChaincodeId string, targetId string, validation bool) (err error) {
	return errors.New("LuaChaincodeId Not Exist")
}

func Delete(LuaChaincodeId string) {
}

