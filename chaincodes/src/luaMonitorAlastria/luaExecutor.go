/*
Copyright IBM Corp. 2016 All Rights Reserved.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
		 http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

//WARNING - this chaincode's ID is hard-coded in chaincode_example04 to illustrate one way of
//calling chaincode from a chaincode. If this example is modified, chaincode_example04.go has
//to be modified as well with the new ID of chaincode_example02.
//chaincode_example05 show's how chaincode ID can be passed in as a parameter instead of
//hard-coding.

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"crypto/tls"

	"github.com/yuin/gopher-lua"
	//"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)


var rsaCertPEM = [] byte(`-----BEGIN CERTIFICATE-----
MIIEKzCCAhOgAwIBAgIRAKlBaL+CEYXR1jXqJVDAtvIwDQYJKoZIhvcNAQELBQAw
MDEbMBkGA1UEChMSQ29uc29yY2lvIEFsYXN0cmlhMREwDwYDVQQDEwhhbGFzdHJp
YTAeFw0xODAxMjMxMTE2MjJaFw0xOTAxMjMxMTE2MjJaMEkxEDAOBgNVBAoTB21v
bml0b3IxNTAzBgNVBAMTLENsaWVudCBjZXJ0aWZpY2F0ZSBmb3IgVExTL1NTTCBh
Z2VudCBtb25pdG9yMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAtoV2
clGh7lCSuR6VhTsbHbkhQqwoD0RmjQkft916Bv6hzQAQTJ4hve39V1mnMkMzxxmr
YGRX4LMirYw/laWuSGKNwpkbZlpgTm2x/AUIBPqXzEGbT25FuVIHbs7BJRNlTvQ7
RWOnldilO4mjr9MVX9G7kYofIbu6xXNReR31SHRIm/B314AWoxVoWcHtdvndcgkF
YtG50cB4aRauXGc+p+Tx6TPo1aexWTJIKJaJp8WeHuMr0KcPuNHyyCNW/PqvDYGZ
4eb28IKYidjozI4gESycsVenYMMuWyx5tzZ9q/IyoEHPcXBce/O6zSKtzgxyttY6
4d1G1LISrG6iVKUcHQIDAQABoycwJTAOBgNVHQ8BAf8EBAMCBaAwEwYDVR0lBAww
CgYIKwYBBQUHAwIwDQYJKoZIhvcNAQELBQADggIBAENgh/01kKvDSZg55Mve3o/p
E30F+P5Ys2WL5dea92h/x1Ccl96nGKtSny1NUJLnwHQRjt/A6CLEreI+9BxQenmH
G14l02YsylEtCQd3lG0bMSRLextW4pLKJDMgGzozw6iiytjU01y2xnl16PgTmeEm
5fAYMsQo5NguuzZB/Xi0b56DG9gd6qviqcQsLSaN5Tn9LhbhTngYJc7p02DRTNZh
1ugEc5Z3l4NMx0105z693PAOiVBX2ObKtcPDyyEmRMTTGLcZ3F+/osr6dcwRDIGO
M1UlkADr025Ae3EL2vdxB7MXgOSLK0S8I6pnSRBF063Xfa6i1Cp9Zdl5Mh1z77yi
iSzsLaCs+stYSWmV5As/pzmGttMJqyd5r7Eo3ZIDpkdk29e7rST6cAdv7mdBjYyw
fMgdzkA06MfZuW2DlKccp8SbeJiyhq75NxMtq4AXpufeZjy93O134ssjcY5r2BgU
4+820AH+O39Yg3HNa+iga4vcjqg25a0wF1+V+bkOn3IwW8CDuG6YmOlNSS5RM6cT
XaAtPt0La9V/9lgLPPz81tVggdg6/KH/WnsBlTWewCGdYbkpJQ3I9bpiUDo59yb2
15eGePlChmYyCyAr/vU21myni1EqZ4VwOZg2Y9pR+Nv8RT2ljq/NIgqpf7eV0WCQ
L/Bbhvz+bqpiiZnXk5Xj
-----END CERTIFICATE-----`)

var rsaKeyPEM = [] byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAtoV2clGh7lCSuR6VhTsbHbkhQqwoD0RmjQkft916Bv6hzQAQ
TJ4hve39V1mnMkMzxxmrYGRX4LMirYw/laWuSGKNwpkbZlpgTm2x/AUIBPqXzEGb
T25FuVIHbs7BJRNlTvQ7RWOnldilO4mjr9MVX9G7kYofIbu6xXNReR31SHRIm/B3
14AWoxVoWcHtdvndcgkFYtG50cB4aRauXGc+p+Tx6TPo1aexWTJIKJaJp8WeHuMr
0KcPuNHyyCNW/PqvDYGZ4eb28IKYidjozI4gESycsVenYMMuWyx5tzZ9q/IyoEHP
cXBce/O6zSKtzgxyttY64d1G1LISrG6iVKUcHQIDAQABAoIBACjPwuJg+nJNNeGK
wygdRTzqLlO4JuTzCHM0vRDhxu1VdlxeTUa0fRr44hLsCwSkHinAxZ8yEKw/odto
ZrdRapzo3IXMsmG6h5pB0PBnN5nVZqfXa6DhKVn8y4itVmax5Y838SDc3ZYV8SxU
5HLmIftg7C3o0nk6ftKzrF4GotfmN5px8xNykH3RRBTpoM7cdCLYCVBifwcUQBc+
HkNKu/dGUW9hYU6DNxdlAziECsPh4NEOPRsDPUFe1e2Mm0CXCbuekGLYLXR/7WyO
7L40p24Jj0sexuQUOJQfPHDmIlxk2D68mgxsc0eh2QB8ieQS25rG8iYJMJYmvfvn
/4kBaoECgYEAzCAxhaXlqCa4DZUkzEJKmOJ4lcx9aYvoXOTqkStyRyrIG2+jWx/8
bN4IgMbm++DRGNIuOYjN6q3YCcFbPLK4qpSBqHyr45Ann+CLEA/3zR43HarrBpQD
5reGeL8gGFVyK286Ytv5HgQmYql8wR7ptaafxwYlFbS2wn+0xLVz3JECgYEA5OfA
P0fC0ouRRqzBXdoQqTusCP+I6E0N5L2J+DNrkEcpVWobdcczDGSXMqdDQTEN8bwc
FZUXyT8sg7sYB5yR2Rn+pm0VDnkaZ6en5bbeIW2UGOLABygp8BxDFLUZTwmT5bod
qrAdoUZX5wAEo2fhpvzzmPtH/bSx50aU1wAzvM0CgYB6ASf4TL7mcTYDEEitOVYi
6QHP7yhqZHAezcgRupkURlLSaziFJ6oVW+RTLA81LxtrLXzpcIY7JWsB8arZrazI
b/jLPrDyU+ALJAeaMyEWXV/uAJF3HhLy3HCoTPwe7ztNEK8iFX5hXXOf9tOVDif8
JbpMF7Uksx2lRPVDuXylsQKBgQDdXxMZH7lWDQvOIbxPm1iJkd/qQ9aPchWBpZwM
oQ3hVjCvHhK4cJD55z4iCoiMP3iQ068FYE7EnYfbTdELa2vRXcZcBRpcCNp0bDt/
eL6R7XtQJdo4EaudPnfRSuLTARw7Cyctr2y4T2MrT2Us8oXUDMU4qasRvPJgDG5d
DyD58QKBgFye0S1BDK72JZY2hmLeB7A5mO/ScUWdZHbUODvpEl57myAtf9qEX8sX
8YwLKVu4OkuIepz65ewyKdL7w8ijmwNcvcIN4wd1+MtXyA5Mbg8tyyrHkRagLjo8
g1UEr6GryQJZKfpyreDxSEOcLbzbB8odU7g2dernrIcbS/OJfV6t
-----END RSA PRIVATE KEY-----`)


var logger = shim.NewLogger("Lua Execution Chaincode")

// SimpleChaincode example simple Chaincode implementation
type LExecutionChaincode struct {
}

type customEvent struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}

func (t *LExecutionChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("ex02 Init")
	_, args := stub.GetFunctionAndParameters()
	fmt.Printf("%s", args)
	var err error
	level, err := shim.LogLevel("DEBUG")
	if err != nil {
		return shim.Error("Problems with loggin level")
	}
	logger.SetLevel(shim.LoggingLevel(level))

	logger.Debug("[Management Chaincode][Init]Instanciating chaincode...")
	if len(args) != 0 {
		return shim.Error("Incorrect number of arguments. Expecting 0")
	}

	// Write the state to the ledger
	err = stub.PutState("LuaResult", []byte("hola"))
	var event = customEvent{"putState", "Successfully put state Lua code: empty string"}
	eventBytes, err := json.Marshal(&event)
	err = stub.SetEvent("evtSender", eventBytes)

	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

func (t *LExecutionChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	logger.Debug("[Lua Execution Chaincode][Invoke]Invoking chaincode...")
	function, args := stub.GetFunctionAndParameters()
	if function == "invoke" {
		// execute lua code in chaincode
		return t.invoke(stub, args)
	} else if function == "query" {
		// the old "Query" is now implemtned in invoke
		return t.query(stub, args)
	}

	return shim.Error("Invalid invoke function name. Expecting \"invoke\" \"delete\" \"query\"")
}

// Transaction makes payment of X units from A to B
func (t *LExecutionChaincode) invoke(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	if len(args) < 1 {
		logger.Error("[Lua Execution Chaincode][invoke]Error with arguments...")
		return shim.Error("Incorrect number of arguments. Expecting at least 1")
	}

	luaFuncCode := args[0]
	L := lua.NewState()
	defer L.Close()
	L.SetGlobal("ServiceCall", L.NewFunction(ServiceCall))
	if err := L.DoString(luaFuncCode); err != nil {
		logger.Error("[Lua Execution Chaincode][invoke]Error storing new function...")
		logger.Error(err.Error())
        return shim.Error(err.Error())

	}

	if err := L.CallByParam(lua.P{
		Fn:      L.GetGlobal("execute"), // name of Lua function
		NRet:    1,                      // number of returned values
		Protect: true,                   // return err or panic
	}); err != nil {
		logger.Error("[Lua Execution Chaincode][invoke]Error executing lua code")
		logger.Error(err.Error())
        return shim.Error(err.Error())
	}

	// Get the returned value from the stack and cast it to a lua.LString
	luaFuncResult, ok := L.Get(-1).(lua.LString)
	if ok {
		logger.Debug("[Lua Execution Chaincode][invoke] Execution OK")
		//fmt.Println(luaFuncResult)
	}

	// save the result
	// Write the state back to the ledger
	err = stub.PutState("LuaResult", []byte(luaFuncResult))

	/*
	   var event = customEvent{"putState", "Successfully put state lua func result: " + string(luaFuncResult)}
	   eventBytes, err := json.Marshal(&event)
	   err = stub.SetEvent("evtSender", eventBytes)
	*/

	if err != nil {
		logger.Error("[Lua Execution Chaincode][invoke]Error storing result...")
		logger.Error(err.Error())
		return shim.Error(err.Error())
	}
	jsonResp := "{\"LuaResult\":\"" + string(luaFuncResult) + "\"}"
	logger.Debug("[Lua Execution Chaincode][invoke]Query Response:\n", jsonResp)

	return shim.Success([]byte(jsonResp))
}

// query callback representing the query of a chaincode
func (t *LExecutionChaincode) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	// Get the state from the ledger
	LuaResult, err := stub.GetState("LuaResult")
	if err != nil {
		logger.Error("[Lua Execution Chaincode][query]Error querying result...")
		jsonResp := "{\"Error\":\"Failed to get state for " + "LuaResult" + "\"}"
		return shim.Error(jsonResp)
	}

	if LuaResult == nil {
		logger.Error("[Lua Execution Chaincode][query]Nil amount for LuaResult.")
		jsonResp := "{\"Error\":\"Nil amount for LuaResult\"}"
		return shim.Error(jsonResp)
	}

	jsonResp := "{\"LuaResult\":\"" + string(LuaResult) + "\"}"
	logger.Debug("[Lua Execution Chaincode][query]Query Response:\n", jsonResp)
	return shim.Success([]byte(jsonResp))
}

func ServiceCall(L *lua.LState) int {
	logger.Debug("[Lua Execution Chaincode][ServiceCall]Calling to function ")
	url := L.ToString(1)
	method := L.ToString(2)

	cert, err := tls.X509KeyPair([] byte(rsaCertPEM), [] byte(rsaKeyPEM))
	if err != nil {
		fmt.Printf(err.Error())
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
                InsecureSkipVerify: true,
				Certificates: []tls.Certificate{cert},
			},
		},
	}

	if method == "GET" {
		response, _ := client.Get(url)
		defer response.Body.Close()
		contents, _ := ioutil.ReadAll(response.Body)
		L.Push(lua.LString(contents))
	}
	if method == "POST" {
		response, _ := client.Get(url)
		defer response.Body.Close()
		contents, _ := ioutil.ReadAll(response.Body)
		L.Push(lua.LString(contents))
	}
	return 1
}

func main() {
	err := shim.Start(new(LExecutionChaincode))
	if err != nil {
		logger.Error("Error starting Simple chaincode: %s", err)
	}
}
