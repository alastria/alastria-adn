package models

import (
	"errors"
	"strconv"
	"time"
	"fmt"
)

var (
	LuaChaincodes map[string]*LuaChaincode
)

type LuaChaincode struct {
	LuaChaincodeId   string
	Name string
	SourceCode string
	Targets []User
	Validations []bool
}

func init() {
	LuaChaincodes = make(map[string]*LuaChaincode)
	var code LuaChaincode 
	code.LuaChaincodeId = "1"
	code.Name = "nombre1"
	code.SourceCode = "codigo"
	code.Targets = append(code.Targets, User{"user_id1"})
	code.Validations = append(code.Validations, false)
	LuaChaincodes[code.LuaChaincodeId] = &code
	//LuaChaincodes["hjkhsbnmn123"] = &LuaChaincode{"id2", "nombre2", "astaxie"}
}

func AddOne(luaChaincode LuaChaincode) (LuaChaincodeId string) {
	luaChaincode.LuaChaincodeId = "luachaincode" + strconv.FormatInt(time.Now().UnixNano(), 10)
	LuaChaincodes[luaChaincode.LuaChaincodeId] = &luaChaincode
	return luaChaincode.LuaChaincodeId
}

func GetOne(LuaChaincodeId string) (luaChaincode *LuaChaincode, err error) {
	if v, ok := LuaChaincodes[LuaChaincodeId]; ok {
		return v, nil
	}
	return nil, errors.New("LuaChaincodeId Not Exist")
}

func GetAll() map[string]*LuaChaincode {
	return LuaChaincodes
}

func Update(LuaChaincodeId string, targetId string, validation bool) (err error) {
	fmt.Printf("hhhhh")
	if v, ok := LuaChaincodes[LuaChaincodeId]; ok {
		for i:=0; i<len(v.Targets); i++ {
					fmt.Printf("%s ", i)
					fmt.Printf("%s 1", targetId)
					fmt.Printf("%s 2", v.Targets[i].Id)
			if v.Targets[i].Id == targetId{
				v.Validations[i] = true
				LuaChaincodes[LuaChaincodeId] = v
				fmt.Printf("hhhhh")
				break
			}
		}
		return nil
	}
	return errors.New("LuaChaincodeId Not Exist")
}

func Delete(LuaChaincodeId string) {
	delete(LuaChaincodes, LuaChaincodeId)
}

