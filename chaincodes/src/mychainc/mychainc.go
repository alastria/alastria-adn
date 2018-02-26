package main

import (
	"encoding/json"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
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

type CodeStore struct {
	Name     string
	Source   string
	Target   map[string]bool
	Approved int  // autoincrement with the  approvement of a party
	Verified bool // If approved == map.length -> TRUE
}

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
	if function == "registrar" {
		return t.registrar(stub, args)
	}
	if function == "storeCode" {
		return t.storeCode(stub, args)
	}
	if function == "getCode" {
		return t.getCode(stub, args)
	}
	if function == "approveCode" {
		return t.approveCode(stub, args)
	}
	if function == "getListCC" {
		return t.getListCC(stub)
	}
	if function == "getAllTargets" {
		return t.getAllTargets(stub)
	}
	return shim.Success([]byte("Invoke"))
}
func (t *ManagementChaincode) registrar(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("[Management Chaincode][Registrar]Incorrect Arguments")
	}
	logger.Debug("[Management Chaincode][Registrar]Calling registrar...")
	//Get the transaction ID
	txID := stub.GetTxID()
	logger.Debug("[Management Chaincode][Registrar]Transaction ID", txID)
	// Get certs from transaction sender
	caller, err := stub.GetCreator()
	if err != nil {
		logger.Error("[Management Chaincode][Registrar] Problem getting caller ")
		return shim.Error(err.Error())
	}
	//Get alias to store
	alias := args[0]
	logger.Debug("[Management Chaincode][Registrar]Storing cert, alias", args[0])
	err = stub.PutState(string(caller[:]), []byte(alias))
	logger.Debug("[Management Chaincode][Registrar]Caller", string(caller[:]))
	if err != nil {
		logger.Error("[Management Chaincode][Registrar]Problem adding new target..", err)
		return shim.Error(err.Error())
	}
	addTarget(stub, alias)
	logger.Debug("[Management Chaincode][Registrar]Stored successful", args[0])

	return shim.Success([]byte(caller))
}
func (t *ManagementChaincode) storeCode(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		logger.Error("[Management Chaincode][StoreCode]Incorrect Arguments")
		return shim.Error("Error")
	}
	guid := strconv.Itoa(uniqueID)
	logger.Debug("ID", guid)
	logger.Debug("[Management Chaincode][StoreCode]Args ID", args[0])
	logger.Debug("[Management Chaincode][StoreCode]Calling storeCode...")
	//Process the input
	rawIn := json.RawMessage(args[0])
	bytes, err := rawIn.MarshalJSON()
	if err != nil {
		logger.Error("[Management Chaincode][StoreCode]Error Marshaling Object")
		return shim.Error(err.Error())
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
	//Index
	data, err := json.Marshal(store)
	if err != nil {
		logger.Error("[Management Chaincode][StoreCode]Error Marshaling Object")
		return shim.Error(err.Error())
	}
	exists, error := stub.GetState(guid)
	var stored CodeStore
	err = json.Unmarshal(exists, &stored)
	if stored.Name == "" {
		res := addToListCC(stub, guid) //ADD TO LIST CC
		if res == false {
			logger.Error("[Management Chaincode][StoreCode]Error Addint to ListCC")
			return shim.Error(err.Error())
		}
		if error != nil {
			return shim.Error(err.Error())
		}
	}
	//Store code in bc
	err = stub.PutState(guid, data)
	if err != nil {
		logger.Error("[Management Chaincode][StoreCode] Error storing code")
		return shim.Error(err.Error())
	}
	logger.Debug("[Management Chaincode][StoreCode] Code stored with ID and data" + guid)
	uniqueID++
	return shim.Success([]byte(guid))
}
func (t *ManagementChaincode) getCode(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		logger.Error("[Management Chaincode][getCode]Incorrect Arguments")
		return shim.Error("[Management Chaincode][getCode]Incorrect Arguments")
	}
	logger.Debug("[Management Chaincode][getCode]Get code with ID", args[0])
	guid := string(args[0])
	//Store code in bc
	code, err := stub.GetState(guid)
	logger.Debug(code)
	// Store contract ID in list
	if err != nil {
		logger.Error("[Management Chaincode][getCode]Error getting code with id" + guid)
		return shim.Error(err.Error())
	}
	if code == nil {
		logger.Error("[Management Chaincode][getCode]Code not exist" + guid)
		return shim.Error(err.Error())
	}
	return shim.Success(code)
}

func (t *ManagementChaincode) approveCode(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		logger.Error("[Management Chaincode][getCode]Incorrect Arguments")
		return shim.Error("[Management Chaincode][getCode]Incorrect Arguments")
	}
	logger.Debug("[Management Chaincode][approveCode]Approving code with ID", args[0])
	//Process the input Getting Id
	guid := string(args[0])
	alias := getAlias(stub)
	if alias == "" {
		return shim.Error("[Management Chaincode][approveCode]Error getting alias")
	}
	//Get code from bc
	code, err := stub.GetState(guid)
	if err != nil {
		return shim.Error("[Management Chaincode][approveCode]Error getting chaincode")
	}
	var stored CodeStore
	err = json.Unmarshal(code, &stored)

	if stored.Name == "" {
		logger.Error("[Management Chaincode][approveCode] Code not exist")
		return shim.Error(err.Error())
	}
	logger.Debug(stored)
	if stored.Target[alias] != true {
		stored.Target[alias] = true
		stored.Approved++
	}
	if stored.Approved == len(stored.Target) {
		stored.Verified = true
	}
	//Convert to bytes to stored the new state
	//Index
	data, err := json.Marshal(stored)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(guid, data)
	if err != nil {
		logger.Error("[Management Chaincode][approveCode]Error updating the approval")
		return shim.Error("[Management Chaincode][approveCode]Error updating the approval")
	}
	return shim.Success([]byte("[Management Chaincode][approveCode]Approved chaincode from target " + alias))
}

/*
*  List all the available targets in the network
 */
func (t *ManagementChaincode) getAllTargets(stub shim.ChaincodeStubInterface) pb.Response {
	logger.Debug("[Management Chaincode][getAllTargets] Getting List of targets ..")
	state, err := stub.GetState("targetList")
	if err != nil {
		logger.Error("[Management Chaincode][getAllTargets]Problem getting new target..", err)
		return shim.Error(err.Error())
	}

	return shim.Success(state)
}

/*
* List all the chaincodes in the network , chaincodes are identifier by a ID
 */
func (t *ManagementChaincode) getListCC(stub shim.ChaincodeStubInterface) pb.Response {
	state, err := stub.GetState("codeList")
	if err != nil {
		logger.Error("[Management Chaincode][getListCC]Problem listing Codes..", err)
		return shim.Error(err.Error())
	}
	return shim.Success(state)
}

/************ UTILS************/
/*Get an alias from a caller */
func getAlias(stub shim.ChaincodeStubInterface) string {

	logger.Debug("[Management Chaincode][getAlias]Calling registrar...")
	//Get the transaction ID
	txID := stub.GetTxID()
	logger.Debug("[Management Chaincode][getAlias]Transaction ID", txID)
	// Get certs from transaction sender
	caller, err := stub.GetCreator()
	if err != nil {
		return ""
	}

	logger.Debug("[Management Chaincode][getAlias]Getting alias from caller---", string(caller[:]))
	state, err := stub.GetState(string(caller[:]))
	if err != nil {
		logger.Error("[Management Chaincode][getAlias]Problem adding new target..", err)
		return ""
	}
	if state == nil {
		logger.Error("[Management Chaincode][getAlias] Target not registered yet", err)
		return ""
	}
	logger.Debug("[Management Chaincode][getAlias]State", string(state[:]))
	return string(state[:])
}

/*
*  Add an available target to the network
 */
func addTarget(stub shim.ChaincodeStubInterface, newTarget string) bool {
	var targetList []string
	state, err := stub.GetState("targetList")
	if err != nil {
		logger.Error("[Management Chaincode][addTarget]Problem adding new target..", err)
		return false
	}
	json.Unmarshal(state, &targetList)
	logger.Debug("[Management Chaincode][addTarget] Old state of the target list ..", targetList)
	slice := append(targetList, newTarget)
	logger.Debug("[Management Chaincode][addTarget] Actual state of the target list ..", slice)
	toStore, err := json.Marshal(slice)
	if err != nil {
		return false
	}
	logger.Debug("[Management Chaincode][addTarget] Updating the state of the target list...")
	stub.PutState("targetList", toStore)
	logger.Debug("[Management Chaincode][addTarget] Updated the state of the target list")

	return true
}

func addToListCC(stub shim.ChaincodeStubInterface, newCode string) bool {
	var ccList []string
	state, err := stub.GetState("codeList")
	if err != nil {
		logger.Error("[Management Chaincode][addToListCC]Problem adding new Code..", err)
		return false
	}
	json.Unmarshal(state, &ccList)
	logger.Debug("[Management Chaincode][addToListCC] Old state of the listCC ..", state)

	slice := append(ccList, newCode)
	logger.Debug("[Management Chaincode][addToListCC] Actual state of the listCC ..", slice)
	toStore, err := json.Marshal(slice)
	if err != nil {
		return false
	}
	logger.Debug("[Management Chaincode][addToListCC] Updating the state...")
	stub.PutState("codeList", toStore)
	logger.Debug("[Management Chaincode][addToListCC] Updated the state")

	return true
}

func main() {
	err := shim.Start(new(ManagementChaincode))
	if err != nil {
		logger.Debugf("Error: %s", err)
	}
}
