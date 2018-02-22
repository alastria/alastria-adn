package main

import (
	"encoding/json"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/rs/xid"
)

var logger = shim.NewLogger("Management Chaincode")

// SimpleChaincode representing a class of chaincode
type ManagementChaincode struct{}

// type Target struct{
// 	Alias string `json: Alias`
// }
type Code struct {
	Name   string   `json:"Name"`
	Source string   `json:"Source"`
	Target []string `json: Target`
}

type CodeInfo struct {
	Id       string
	Source   string
	Target   map[string]bool
	Approved int  // autoincrement with the  approvement of a party
	Verified bool // If approved == map.length -> TRUE

}
type CodeStore struct {
	Name     string
	Source   string
	Target   map[string]bool
	Approved int  // autoincrement with the  approvement of a party
	Verified bool // If approved == map.length -> TRUE
}

var targetList []string
var uniqueID int

// Init to initiate the SimpleChaincode class
func (t *ManagementChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	uniqueID = 0
	level, err := shim.LogLevel("DEBUG")
	if err != nil {
		return shim.Error("Problems with loggin level")
	}
	logger.SetLevel(shim.LoggingLevel(level))

	logger.Debug("[Management Chaincode][Init]Instanciating chaincode...")
	return shim.Success([]byte("Init called"))
}

// Invoke a method specified in the SimpleChaincode class
func (t *ManagementChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	logger.Debug("[Management Chaincode][Invoke]Invoking chaincode...")
	function, args := stub.GetFunctionAndParameters()
	logger.Debug("[Management Chaincode][StoreCode]Args ID", args[0])
	if function == "storeCode" {
		return t.storeCode(stub, args)
	}
	if function == "registrar" {
		return t.registrar(stub, args)
	}
	if function == "getAlias" {
		return t.getAlias(stub, args)
	}
	if function == "getListCC" {
		return t.getListCC(stub)
	}
	return shim.Success([]byte("Invoke"))
}
func (t *ManagementChaincode) registrar(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect Argumetns")
	}
	logger.Debug("[Management Chaincode][Registrar]Calling registrar...")
	//Get the transaction ID
	txID := stub.GetTxID()
	logger.Debug("[Management Chaincode][StoreCode]Transaction ID", txID)
	// Get certs from transaction sender
	caller, err := stub.GetCreator()
	if err != nil {
		logger.Debug(err)
		return shim.Error(err.Error())
	}
	//Get alias to store
	alias := args[0]
	logger.Debug("[Management Chaincode][StoreCode]Storing cert, alias", args[0])
	err = stub.PutState(string(caller[:]), []byte(alias))
	logger.Debug("[Management Chaincode][StoreCode]Caller", string(caller[:]))
	if err != nil {
		logger.Error("[Management Chaincode][addTarget]Problem adding new target..", err)
		return shim.Error(err.Error())
	}
	logger.Debug("[Management Chaincode][StoreCode]Stored successful", args[0])

	return shim.Success([]byte(caller))
}

func (t *ManagementChaincode) getAlias(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	logger.Debug("[Management Chaincode][Registrar]Calling registrar...")
	//Get the transaction ID
	txID := stub.GetTxID()
	logger.Debug("[Management Chaincode][StoreCode]Transaction ID", txID)
	// Get certs from transaction sender
	caller, err := stub.GetCreator()
	if err != nil {
		return shim.Error(err.Error())
	}

	logger.Debug("[Management Chaincode][StoreCode]Getting alias from caller---", string(caller[:]))
	state, err := stub.GetState(string(caller[:]))
	if err != nil {
		logger.Error("[Management Chaincode][addTarget]Problem adding new target..", err)
		return shim.Error(err.Error())
	}
	logger.Debug("[Management Chaincode][StoreCode]State", string(state[:]))
	return shim.Success([]byte(""))
}

func (t *ManagementChaincode) storeCode(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	guid := xid.New().String()
	logger.Debug("GUUUUID", guid)
	logger.Debug("[Management Chaincode][StoreCode]Args ID", args[0])
	logger.Debug("[Management Chaincode][StoreCode]Calling storeCode...")
	//Get the transaction ID
	txID := stub.GetTxID()
	logger.Debug("[Management Chaincode][StoreCode]Transaction ID", txID)
	rawIn := json.RawMessage(args[0])
	bytes, err := rawIn.MarshalJSON()
	if err != nil {
		return shim.Error(err.Error())
	}
	var request Code
	err = json.Unmarshal(bytes, &request)
	//Value to store
	logger.Debug("[Management Chaincode][StoreCode]Args ID", request)
	store := CodeStore{}
	store.Name = request.Name
	store.Source = request.Source
	store.Approved = 0
	store.Verified = false

	i := 0
	for i < len(request.Target) {
		logger.Debug(request.Target[i])
		i++
	}
	i = 0
	store.Target = make(map[string]bool)
	for i < len(request.Target) {
		store.Target[request.Target[i]] = false
		i++
	}
	//Index
	data, err := json.Marshal(store)
	if err != nil {
		return shim.Error(err.Error())
	}

	guid3 := "K902902903"
	stub.PutState(guid3, data)

	res := addToListCC(stub, guid)
	if res == false {
		return shim.Error(err.Error())

	}
	return shim.Success([]byte(strconv.Itoa(uniqueID)))
}
func addTarget(stub shim.ChaincodeStubInterface, newTarget string) bool {
	state, err := stub.GetState("targetList")
	if err != nil {
		logger.Error("[Management Chaincode][addTarget]Problem adding new target..", err)
		return false
	}
	json.Unmarshal(state, &targetList)
	slice := append(targetList, newTarget)
	logger.Debug("[Management Chaincode][addTarget] Actual state of the slice ..", slice)
	toStore, err := json.Marshal(slice)
	if err != nil {
		return false
	}
	logger.Debug("[Management Chaincode][addTarget] Updating the state...")
	stub.PutState("targetList", toStore)
	logger.Debug("[Management Chaincode][addTarget] Updated the state")

	return true
}

func addToListCC(stub shim.ChaincodeStubInterface, newCode string) bool {
	state, err := stub.GetState("codeList")
	if err != nil {
		logger.Error("[Management Chaincode][addToListCC]Problem adding new Code..", err)
		return false
	}
	json.Unmarshal(state, &targetList)
	slice := append(targetList, newCode)
	logger.Debug("[Management Chaincode][addToListCC] Actual state of the slice ..", slice)
	toStore, err := json.Marshal(slice)
	if err != nil {
		return false
	}
	logger.Debug("[Management Chaincode][addToListCC] Updating the state...")
	stub.PutState("targetList", toStore)
	logger.Debug("[Management Chaincode][addToListCC] Updated the state")

	return true
}

func (t *ManagementChaincode) getListCC(stub shim.ChaincodeStubInterface) pb.Response {
	state, err := stub.GetState("codeList")
	if err != nil {
		logger.Error("[Management Chaincode][getListCC]Problem adding new Code..", err)
		return shim.Error(err.Error())
	}
	return shim.Success(state)
}
func main() {
	err := shim.Start(new(ManagementChaincode))
	if err != nil {
		logger.Debugf("Error: %s", err)
	}
}
