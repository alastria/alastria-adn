package main

import (
	"encoding/json"
	"fmt"
	"strings"
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

func checkRegistrar(t *testing.T, stub *shim.MockStub, value string) {
	res := stub.MockInvoke("1", [][]byte{[]byte("registrar"), []byte(value)})
	if res.Status != shim.OK {
		fmt.Println("Registrar", value, "failed", string(res.Message))
		t.FailNow()
	}
}
func checkRegistrarFail(t *testing.T, stub *shim.MockStub, value string) {
	res := stub.MockInvoke("1", [][]byte{[]byte("registrar"), []byte(value), []byte(value)})
	if res.Status == shim.OK {
		fmt.Println("Registrar", value, "failed", string(res.Message))
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
func checkStoreCodeFail(t *testing.T, stub *shim.MockStub, name string, value string) {
	res := stub.MockInvoke("1", [][]byte{[]byte(name), []byte(value), []byte("value")})
	if res.Status == shim.OK {
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
	fmt.Printf(string(value[:]))
	fmt.Printf(string(res.Payload[:]))
	if string(res.Payload[:]) != string(value[:]) {
		fmt.Println("Getcode", id, "value not expected")
		t.FailNow()
	}
}
func checkGetCodeFail(t *testing.T, stub *shim.MockStub, id string, value []byte) {
	res := stub.MockInvoke("1", [][]byte{[]byte("getCode")})
	if res.Status == shim.OK {
		fmt.Println("Getcode", id, "failed", string(res.Message))
		t.FailNow()
	}

}
func checkListCC(t *testing.T, stub *shim.MockStub, expectedValue []byte) {
	res := stub.MockInvoke("1", [][]byte{[]byte("getListCC"), []byte("")})
	if res.Status != shim.OK {
		fmt.Println("GetlistCC", "failed", string(res.Message))
		t.FailNow()
	}
	if res.Payload == nil {
		fmt.Println("GetlistCC", "failed", string(res.Message))
		t.FailNow()
	}

	if string(res.Payload[:]) != string(expectedValue[:]) {
		fmt.Println("Getcode", "value not expected")
		t.FailNow()
	}
}
func checkListCCFail(t *testing.T, stub *shim.MockStub, expectedValue []byte) {
	res := stub.MockInvoke("1", [][]byte{[]byte("getListCC"), []byte("")})

	if string(res.Payload[:]) == string(expectedValue[:]) {
		fmt.Println("Getcode", "value not expected")
		t.FailNow()
	}
}

func checkgetAllTarget(t *testing.T, stub *shim.MockStub, expectedValue []byte) {
	res := stub.MockInvoke("1", [][]byte{[]byte("getAllTargets")})
	if res.Status != shim.OK {
		fmt.Println("getAllTargets", "failed", string(res.Message))
		t.FailNow()
	}

}

func checkApproveFail(t *testing.T, stub *shim.MockStub, idCC string) {
	res := stub.MockInvoke("1", [][]byte{[]byte("approveCode"), []byte("")})

	if res.Status == shim.OK {
		fmt.Println("GetlistCC", "failed", string(res.Message))
		t.FailNow()
	}
}
func checkApprove(t *testing.T, stub *shim.MockStub, idCC string) {
	res := stub.MockInvoke("1", [][]byte{[]byte("approveCode"), []byte(idCC)})

	if res.Status != shim.OK {
		fmt.Println("GetlistCC", "failed", string(res.Message))
		t.FailNow()
	}
}
func checkExecuteFail(t *testing.T, stub *shim.MockStub, idCC string) {
	res := stub.MockInvoke("1", [][]byte{[]byte("exectuteCC"), []byte(idCC)})

	if res.Status == shim.OK {
		fmt.Println("GetlistCC", "failed", string(res.Message))
		t.FailNow()
	}
}

func checkGetAllCC(t *testing.T, stub *shim.MockStub, expectedValue string) {
	res := stub.MockInvoke("1", [][]byte{[]byte("getAllChaincodes")})
	if res.Status != shim.OK {
		fmt.Println("GetlcheckGetAllCCistCC", "failed", string(res.Message))
		t.FailNow()
	}
	if strings.TrimSpace(string(res.GetPayload()[:])) != strings.TrimSpace(expectedValue) {
		fmt.Println(strings.TrimSpace(string(res.GetPayload()[:])))
		fmt.Println(expectedValue)
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
	checkInit(t, stub, [][]byte{[]byte("init"), []byte("lua")})
}

func Test_Invoke(t *testing.T) {
	scc := new(ManagementChaincode)
	stub := shim.NewMockStub("ex02", scc)
	checkInit(t, stub, [][]byte{[]byte("init"), []byte("lua")})
	checkInvoke(t, stub, "registrar", "bank1")
}

func Test_Registrar(t *testing.T) {
	scc := new(ManagementChaincode)
	stub := shim.NewMockStub("ex02", scc)
	checkInit(t, stub, [][]byte{[]byte("init"), []byte("lua")})
	checkRegistrar(t, stub, "bank1")
	//FAIL
	checkRegistrarFail(t, stub, "bank1") //Incorrect Number of arguments
}

func Test_StoreCode(t *testing.T) {
	scc := new(ManagementChaincode)
	stub := shim.NewMockStub("ex02", scc)
	checkInit(t, stub, [][]byte{[]byte("init"), []byte("lua")})
	checkStoreCode(t, stub, "storeCode", `{"Name": "mycc2", "Source": "ks20230", "Target":["org1", "org2"]}`)

	//FAIL
	checkStoreCodeFail(t, stub, "storeCode", `{"Name": "mycc2", "Source": "ks20230", "Target":["org1", "org2"]}`) //Incorrect Number of arguments
}

func Test_getCode(t *testing.T) {
	scc := new(ManagementChaincode)
	stub := shim.NewMockStub("ex02", scc)
	checkInit(t, stub, [][]byte{[]byte("init"), []byte("lua")})
	checkStoreCode(t, stub, "storeCode", `{"Name": "mycc2", "Source": "ks20230", "Target":["org1", "org2"]}`)
	codeProcessed := ProcessCode(`{"Name": "mycc2", "Source": "ks20230", "Target":["org1", "org2"]}`)
	checkGetCode(t, stub, "0", codeProcessed)
	checkGetCodeFail(t, stub, "0", codeProcessed)
}

func Test_getListCC(t *testing.T) {
	scc := new(ManagementChaincode)
	stub := shim.NewMockStub("ex02", scc)
	//ACCEPTED
	checkInit(t, stub, [][]byte{[]byte("init"), []byte("lua")})
	checkStoreCode(t, stub, "storeCode", `{"Name": "mycc1", "Source": "aksdjladkjsladsks20230", "Target":["org1", "org2"]}`)
	checkStoreCode(t, stub, "storeCode", `{"Name": "mycc2", "Source": "aslaksñlaksñlkaks20230", "Target":["org1", "org2"]}`)
	checkStoreCode(t, stub, "storeCode", `{"Name": "mycc3", "Source": "asmasaslaskljaslkjasks20230", "Target":["org1", "org2"]}`)
	expectedValue := []string{"0", "1", "2"}
	sliceExpected, err := json.Marshal(expectedValue)
	if err != nil {
		fmt.Printf("Problem converting expected value to byte array")
	}
	checkListCC(t, stub, sliceExpected)

	//FAILED
	checkStoreCode(t, stub, "storeCode", `{"Name": "mycc4", "Source": "asmasaslaskljaslkjasks20230", "Target":["org1", "org2"]}`)
	checkListCCFail(t, stub, sliceExpected)
}

func Test_ApproveCode(t *testing.T) {
	scc := new(ManagementChaincode)
	stub := shim.NewMockStub("ex03", scc)
	//FAIL
	checkInit(t, stub, [][]byte{[]byte("init"), []byte("lua")})
	checkStoreCode(t, stub, "storeCode", `{"Name": "mycc1a", "Source": "function execute()result = ServiceCall('http://35.176.99.163:8050/services/?token=yeahbaby23&database=0&command=manyrecords&index=0', 'GET')        return result    end ", "Target":["org1"]}`)
	checkApproveFail(t, stub, "0") // target not registered yet

	checkRegistrar(t, stub, "bank1")
	codeProcessed := ProcessCode(`{"Name": "mycc1a", "Source": "function execute()result = ServiceCall('http://35.176.99.163:8050/services/?token=yeahbaby23&database=0&command=manyrecords&index=0', 'GET')        return result    end ", "Target":["org1"]}`)
	checkGetCode(t, stub, "0", codeProcessed)
	checkApprove(t, stub, "0") // target not registered yet
	checkExecuteFail(t, stub, "0")
}

func Test_AllTarget(t *testing.T) {
	scc := new(ManagementChaincode)
	stub := shim.NewMockStub("ex03", scc)
	//FAIL
	checkInit(t, stub, [][]byte{[]byte("init"), []byte("lua")})
	checkRegistrar(t, stub, "bank1")
	expectedValue := []string{"bank1"}
	value, err := json.Marshal(expectedValue)
	if err != nil {
		fmt.Println("Error")
	}
	checkgetAllTarget(t, stub, value)

}

func Test_GetAllCC(t *testing.T) {
	scc := new(ManagementChaincode)
	stub := shim.NewMockStub("ex03", scc)
	//FAIL
	checkInit(t, stub, [][]byte{[]byte("init"), []byte("lua")})
	checkStoreCode(t, stub, "storeCode", `{"Name": "mycc1a", "Source": "function execute()result = ServiceCall('http://35.176.99.163:8050/services/?token=yeahbaby23&database=0&command=manyrecords&index=0', 'GET')return result end ", "Target":["org1"]}`)
	checkStoreCode(t, stub, "storeCode", `{"Name": "mycc2a", "Source": "function execute()result = ServiceCall('http://35.176.99.163:8050/services/?token=yeahbaby23&database=0&command=manyrecords&index=0', 'GET')return result end ", "Target":["org1"]}`)
	checkStoreCode(t, stub, "storeCode", `{"Name": "mycc3a", "Source": "function execute()result = ServiceCall('http://35.176.99.163:8050/services/?token=yeahbaby23&database=0&command=manyrecords&index=0', 'GET')return result end ", "Target":["org1"]}`)
	expectedValue := `{"0":{"Name":"mycc1a","Source":"function execute()result = ServiceCall('http://35.176.99.163:8050/services/?token=yeahbaby23\u0026database=0\u0026command=manyrecords\u0026index=0', 'GET')return result end ","Target":{"org1":false},"Approved":0,"Verified":false},"1":{"Name":"mycc2a","Source":"function execute()result = ServiceCall('http://35.176.99.163:8050/services/?token=yeahbaby23\u0026database=0\u0026command=manyrecords\u0026index=0', 'GET')return result end ","Target":{"org1":false},"Approved":0,"Verified":false},"2":{"Name":"mycc3a","Source":"function execute()result = ServiceCall('http://35.176.99.163:8050/services/?token=yeahbaby23\u0026database=0\u0026command=manyrecords\u0026index=0', 'GET')return result end ","Target":{"org1":false},"Approved":0,"Verified":false}}`
	valueExpected := strings.TrimSpace(expectedValue)
	checkGetAllCC(t, stub, valueExpected)

}
