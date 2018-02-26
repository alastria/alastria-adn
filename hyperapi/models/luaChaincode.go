package models

import (
)

type LuaChaincodePost struct {
	Name string
	SourceCode string
	Targets []string
}

type LuaChaincode struct {
    LuaChaincodeId   string `required:"false" description:"object id"`
    Name string `required:"true" description:"name"`
	SourceCode string `required:"true" description:"source code "`
    Targets []string `required:"true" description:"target list"`
	Validations []bool `required:"true" description:"validation list"`
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

func Update(LuaChaincodeId string) (err error) {
    fabric_validate_code(LuaChaincodeId)
	return nil
}

func Delete(LuaChaincodeId string) {
}

