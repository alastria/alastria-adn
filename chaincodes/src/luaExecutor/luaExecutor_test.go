package main

import (
	"fmt"
	"testing"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func checkInit(t *testing.T, stub *shim.MockStub, args [][]byte) {
	res := stub.MockInit("1", args)
	if res.Status != shim.OK {
		fmt.Println("Init failed", string(res.Message))
		t.FailNow()
	}
}

func checkInvoke(t *testing.T, stub *shim.MockStub, code string) {
	res := stub.MockInvoke("1", [][]byte{[]byte("invoke"), []byte(code)})
	if res.Status != shim.OK {
		fmt.Println("Invoking lua code")
		t.FailNow()
	}
}

func checkInvokeFail(t *testing.T, stub *shim.MockStub) {
	res := stub.MockInvoke("1", [][]byte{[]byte("invoke")})
	if res.Status == shim.OK {
		fmt.Println("Invoking lua code")
		t.FailNow()
	}
}
func checkInvokeFail2(t *testing.T, stub *shim.MockStub, code string) {
	res := stub.MockInvoke("1", [][]byte{[]byte("invoke"), []byte(code)})
	if res.Status == shim.OK {
		fmt.Println("Invoking lua code")
		t.FailNow()
	}
}
func checkQuery(t *testing.T, stub *shim.MockStub) {
	res := stub.MockInvoke("1", [][]byte{[]byte("query")})
	if res.Status != shim.OK {
		fmt.Println("Invoking lua code")
		t.FailNow()
	}
}

func Test_Init(t *testing.T) {
	scc := new(LExecutionChaincode)
	stub := shim.NewMockStub("ex02", scc)
	checkInit(t, stub, [][]byte{[]byte("init")})
}

func Test_Invoke(t *testing.T) {
	scc := new(LExecutionChaincode)
	stub := shim.NewMockStub("ex02", scc)
	checkInit(t, stub, [][]byte{[]byte("init")})
	luaFunction := `function execute()result = ServiceCall('http://35.176.99.163:8050/services/?token=yeahbaby23&database=0&command=manyrecords&index=0', 'GET')        return result    end `
	checkInvoke(t, stub, luaFunction)
	checkQuery(t, stub)
	checkInvokeFail(t, stub)

	luaFunctionFail := `function result = ServiceCall('http://35.176.99.163:8050/services/?token=yeahbaby23&database=0&command=manyrecords&index=0', 'GET')        return result    end `
	checkInvokeFail2(t, stub, luaFunctionFail)
}
