package main

import (
	"encoding/json"
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

func checkInvoke(t *testing.T, stub *shim.MockStub, name string, value string) {
	res := stub.MockInvoke("1", [][]byte{[]byte("query"), []byte(name)})
	if res.Status != shim.OK {
		fmt.Println("Query", name, "failed", string(res.Message))
		t.FailNow()
	}
}

func checkRegistrar(t *testing.T, stub *shim.MockStub, name string, value string) {
	res := stub.MockInvoke("1", [][]byte{[]byte("registrar"), []byte(name)})
	if res.Status != shim.OK {
		fmt.Println("Registrar", name, "failed", string(res.Message))
		t.FailNow()
	}
}

func checkStoreCode(t *testing.T, stub *shim.MockStub, name string, value string) {
	res := stub.MockInvoke("1", [][]byte{[]byte(name), []byte(value)})
	if res.Status != shim.OK {
		fmt.Println("StoreCode", name, "failed", string(res.Message))
		t.FailNow()
	}
}
func checkGetCode(t *testing.T, stub *shim.MockStub, id string, value []byte) {
	res := stub.MockInvoke("1", [][]byte{[]byte("getCode"), []byte(id)})
	if res.Status != shim.OK {
		fmt.Println("Getcode", id, "failed", string(res.Message))
		t.FailNow()
	}
	if res.Payload == nil {
		fmt.Println("Getcode", id, "failed to get value")
		t.FailNow()
	}

	fmt.Println("value", []byte(value))
	if string(res.Payload[:]) != string(value[:]) {
		fmt.Println("Getcode", id, "value not expected")
		t.FailNow()
	}
}
func ProcessCode(value string) []byte {
	rawIn := json.RawMessage(value)
	bytes, err := rawIn.MarshalJSON()
	if err != nil {

		return nil
	}
	var request Code
	err = json.Unmarshal(bytes, &request)
	//Value to store
	logger.Debug("[Management Chaincode][StoreCode]Content to store", request)
	store := CodeStore{}
	store.Name = request.Name
	store.Source = request.Source
	store.Approved = 0
	store.Verified = false
	store.Target = make(map[string]bool)
	i := 0
	for i < len(request.Target) {
		store.Target[request.Target[i]] = false
		i++
	}
	data, err := json.Marshal(store)
	return data
}
func Test_Init(t *testing.T) {
	scc := new(ManagementChaincode)
	stub := shim.NewMockStub("ex02", scc)
	checkInit(t, stub, [][]byte{[]byte("init")})
}

func Test_Invoke(t *testing.T) {
	scc := new(ManagementChaincode)
	stub := shim.NewMockStub("ex02", scc)
	checkInit(t, stub, [][]byte{[]byte("init")})
	checkInvoke(t, stub, "registrar", "bank1")
}

func Test_Registrar(t *testing.T) {
	scc := new(ManagementChaincode)
	stub := shim.NewMockStub("ex02", scc)
	checkInit(t, stub, [][]byte{[]byte("init")})
	checkRegistrar(t, stub, "registrar", "bank1")
}

func Test_StoreCode(t *testing.T) {
	scc := new(ManagementChaincode)
	stub := shim.NewMockStub("ex02", scc)
	checkInit(t, stub, [][]byte{[]byte("init")})
	checkStoreCode(t, stub, "storeCode", `{"Name": "mycc2", "Source": "ks20230", "Target":["org1", "org2"]}`)
}

func Test_getCode(t *testing.T) {
	scc := new(ManagementChaincode)
	stub := shim.NewMockStub("ex02", scc)
	checkInit(t, stub, [][]byte{[]byte("init")})
	checkStoreCode(t, stub, "storeCode", `{"Name": "mycc2", "Source": "ks20230", "Target":["org1", "org2"]}`)
	codeProcessed := ProcessCode(`{"Name": "mycc2", "Source": "ks20230", "Target":["org1", "org2"]}`)
	checkGetCode(t, stub, "0", codeProcessed)
}

func Test_getCode(t *testing.T) {
	scc := new(ManagementChaincode)
	stub := shim.NewMockStub("ex02", scc)
	checkInit(t, stub, [][]byte{[]byte("init")})
	checkStoreCode(t, stub, "storeCode", `{"Name": "mycc2", "Source": "ks20230", "Target":["org1", "org2"]}`)
	codeProcessed := ProcessCode(`{"Name": "mycc2", "Source": "ks20230", "Target":["org1", "org2"]}`)
	checkGetCode(t, stub, "0", codeProcessed)
}
