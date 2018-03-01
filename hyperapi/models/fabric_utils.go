package models

import (
    "fmt"
	"encoding/json"
    "strings"
    "os"
	"github.com/astaxie/beego"
	"github.com/hyperledger/fabric-sdk-go/api/apitxn"
	"github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/common/cauthdsl"
    "github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"github.com/hyperledger/fabric-sdk-go/pkg/config"
    //"github.com/hyperledger/fabric-sdk-go/pkg/context/api/core"
    "github.com/hyperledger/fabric-sdk-go/pkg/fabric-client/peer"
    "github.com/hyperledger/fabric-sdk-go/api/apiconfig"
	resmgmt "github.com/hyperledger/fabric-sdk-go/api/apitxn/resmgmtclient"
	packager "github.com/hyperledger/fabric-sdk-go/pkg/fabric-client/ccpackager/gopackager"
)

// extracted from chaincode
type CodeStore struct {
	Name     string
	Source   string
	Target   map[string]bool
	Approved int  // autoincrement with the  approvement of a party
	Verified bool // If approved == map.length -> TRUE
}

var initArgs = [][]byte{[]byte("init")}
// initialize hyperledger sdk config
var channelID = beego.AppConfig.String("channelID")
var hyperledger_config_yaml = beego.AppConfig.String("hyperledgerConfigYamlPath")
var ccID = beego.AppConfig.String("ccID")
var chaincodePath = beego.AppConfig.String("chaincodePath")
var path = beego.AppConfig.String("path")
var version = beego.AppConfig.String("version")
var orgName = beego.AppConfig.String("orgName")
var sdk *fabsdk.FabricSDK
var chClient apitxn.ChannelClient
var luaExecutorPath = beego.AppConfig.String("luaExecutorPath")
var luaExecutorccID = beego.AppConfig.String("luaExecutorccID ")

func init(){
    configFile := os.Args[1]
    beego.LoadAppConfig("ini", configFile)
    // initialize hyperledger sdk config
    channelID = beego.AppConfig.String("channelID")
    hyperledger_config_yaml = beego.AppConfig.String("hyperledgerConfigYamlPath")
    ccID = beego.AppConfig.String("ccID")
    chaincodePath = beego.AppConfig.String("chaincodePath")
    path = beego.AppConfig.String("path")
    version = beego.AppConfig.String("version")
    orgName = beego.AppConfig.String("orgName")
    luaExecutorPath = beego.AppConfig.String("luaExecutorPath")
    luaExecutorccID = beego.AppConfig.String("luaExecutorccID")
    sdk, _ = fabsdk.New(config.FromFile(hyperledger_config_yaml))
    chClient, _ = sdk.NewClient(fabsdk.WithUser("Admin"), fabsdk.WithOrg(orgName)).Channel(channelID)
    createChaincodeFirstTime()
    createChaincodeLuaExecutorFirstTime()
}

func createChaincodeLuaExecutorFirstTime(){
    sdk, _ := fabsdk.New(config.FromFile(hyperledger_config_yaml))
    // Creación del usuario para interactuar.
    clientResMgmt, err := sdk.NewClient(fabsdk.WithUser("Admin")).ResourceMgmt()
	if err != nil {
		fmt.Print("Failed to create new resource management client: %s", err)
	}
	fmt.Println("Creado el cliente")

    // Crear el package para el chaincode
	ccPkg, err := packager.NewCCPackage(luaExecutorPath, chaincodePath)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println("Creado el paquete")


    fmt.Println(luaExecutorPath)
    fmt.Println(luaExecutorccID)
    installCCReq := resmgmt.InstallCCRequest{Name: luaExecutorccID, Path: luaExecutorPath, Version: version, Package: ccPkg}
	fmt.Println("Creada request")
	// Install example cc to Org1 peers
	_, err = clientResMgmt.InstallCC(installCCReq)
	if err != nil {
		fmt.Println("Error al instalar", err)
	}
	fmt.Println("Instalado chaincode en peer")
	fmt.Println("Instalado chaincode en peer")
	fmt.Println("Instalado chaincode en peer")
	fmt.Println("Instalado chaincode en peer")

	// Set up chaincode policy to 'any of two msps'
	ccPolicy:= cauthdsl.SignedByAnyMember([]string{"org1MSP", "coreAdmMSP", "org2MSP"})
	instantciateCReq := resmgmt.InstantiateCCRequest{Name: luaExecutorccID, Path: luaExecutorPath, Version: version, Args: initArgs, Policy: ccPolicy}
	// Instanciación del chaincode
	err = clientResMgmt.InstantiateCC(channelID, instantciateCReq)
	if err != nil {
		fmt.Println(err)
	}
}

func createChaincodeFirstTime(){
    sdk, _ := fabsdk.New(config.FromFile(hyperledger_config_yaml))
    // Creación del usuario para interactuar.
    clientResMgmt, err := sdk.NewClient(fabsdk.WithUser("Admin")).ResourceMgmt()
	if err != nil {
		fmt.Print("Failed to create new resource management client: %s", err)
	}
	fmt.Println("Creado el cliente")

    // Crear el package para el chaincode
	ccPkg, err := packager.NewCCPackage(path, chaincodePath)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println("Creado el paquete")


    installCCReq := resmgmt.InstallCCRequest{Name: ccID, Path: path, Version: version, Package: ccPkg}
	fmt.Println("Creada request")
	// Install example cc to Org1 peers
	_, err = clientResMgmt.InstallCC(installCCReq)
	if err != nil {
		fmt.Println("Error al instalar", err)
	}
	fmt.Println("Instalado chaincode en peer")
	fmt.Println("Instalado chaincode en peer")
	fmt.Println("Instalado chaincode en peer")
	fmt.Println("Instalado chaincode en peer")
	fmt.Println("Instalado chaincode en peer")

	// Set up chaincode policy to 'any of two msps'
	ccPolicy:= cauthdsl.SignedByAnyMember([]string{"org1MSP", "coreAdmMSP", "org2MSP"})

    exeucuteLuaInitArgs := [][]byte{[]byte("init"), []byte(luaExecutorccID)}
    fmt.Println(exeucuteLuaInitArgs)
	instantciateCReq := resmgmt.InstantiateCCRequest{Name: ccID, Path: path, Version: version, Args: exeucuteLuaInitArgs, Policy: ccPolicy}
	// Instanciación del chaincode
	err = clientResMgmt.InstantiateCC(channelID, instantciateCReq)
	if err != nil {
		fmt.Println(err)
	}

}

func execute(){
    // //invoke
	invokeArgs := [][]byte{[]byte(`{"Name": "` + ccID + `", Source": "ks20230", "Target":["bank1", "bank2"]}`)}

	// Invoke chaincode
    res, _, err := chClient.Execute(apitxn.Request{ChaincodeID: "thecc", Fcn: "storeCode", Args: invokeArgs})
	if err != nil {
		fmt.Print("Failed to send code: %s", err)
	}

    fmt.Print("res: %s", res)
}

func fabric_add_code(name string, source string, targets []string) string{
    invokeArgs := [][]byte{[]byte(`{"Name": "` + name + `", "Source": "` + source + `", "Target": ["` + strings.Join(targets[:],"\",\"") + `"]}`)}

    fmt.Println(`{"Name": "` + name + `", "Source": "` + source + `", "Target": ["` + strings.Join(targets[:],"\",\"") + `"]}`)
    fmt.Println(`{"Name": "` + name + `", "Source": "` + source + `", "Target": ["` + strings.Join(targets[:],"\",\"") + `"]}`)
    fmt.Println(`{"Name": "` + name + `", "Source": "` + source + `", "Target": ["` + strings.Join(targets[:],"\",\"") + `"]}`)
    value, _, err := chClient.Execute(apitxn.Request{ChaincodeID: ccID, Fcn: "storeCode", Args: invokeArgs})
	if err != nil {
		fmt.Println("Failed to query values: %s", err)
	}
	fmt.Println("response value: ", string(value))
    return string(value)
}

func fabric_query_users() []string{
    var targetList []string
    invokeArgs := [][]byte{[]byte("")}
    value, _, err := chClient.Execute(apitxn.Request{ChaincodeID: ccID, Fcn: "getAllTargets", Args: invokeArgs})
	if err != nil {
		fmt.Println("Failed to query values: %s", err)
	}
	fmt.Println("response value: ", string(value))
    if string(value) != "null"{
	    json.Unmarshal(value, &targetList)
        return targetList 
    }else{
        targetList = make([]string, 0)
        return targetList  
    }
}

func fabric_query_user(uid string) string{
	fmt.Println("response value: NOT IMPLEMENTED")
    return "1"
}

func fabric_query_codes() []LuaChaincode{
    var codeIds []string
    var luaCodes []LuaChaincode
    var luaCode LuaChaincode
    invokeArgs := [][]byte{[]byte("")}
    value, _, err := chClient.Execute(apitxn.Request{ChaincodeID: ccID, Fcn: "getListCC", Args: invokeArgs})
	if err != nil {
		fmt.Println("Failed to query values: %s", err)
	}
	fmt.Println("response value: ", string(value))
    if string(value) != "null"{
	    json.Unmarshal(value, &codeIds)
        var i = 0 
        for i < len(codeIds) {
            luaCode = fabric_query_code(codeIds[i])
            luaCodes = append(luaCodes, luaCode)
            i++
        }
        return luaCodes 
    }else{
        luaCodes = make([]LuaChaincode, 0)
        return luaCodes 
    }
}

func fabric_query_code(cid string) LuaChaincode{
    var tempCode CodeStore
    var code LuaChaincode
    invokeArgs := [][]byte{[]byte(cid)}
    value, _, err := chClient.Execute(apitxn.Request{ChaincodeID: ccID, Fcn: "getCode", Args: invokeArgs})
	if err != nil {
		fmt.Println("Failed to query values: %s", err)
	}
	fmt.Println("response value: ", string(value))
	json.Unmarshal(value, &tempCode)
    code.LuaChaincodeId = cid  
    code.Name = tempCode.Name
    code.SourceCode = tempCode.Source
    for key, value := range tempCode.Target {
        code.Targets = append(code.Targets, key)
        code.Validations = append(code.Validations, value)
    }
    return code
}

func fabric_add_user(uid string) string{
    invokeArgs := [][]byte{[]byte(uid)}
    value, _, err := chClient.Execute(apitxn.Request{ChaincodeID: ccID, Fcn: "registrar", Args: invokeArgs})
	if err != nil {
		fmt.Println("Failed to query values: %s", err)
	}
	fmt.Println("response value: ", string(value))
    return string(value)
}

func fabric_validate_code(LuaChaincodeId string) string{
    invokeArgs := [][]byte{[]byte(LuaChaincodeId)}
    value, _, err := chClient.Execute(apitxn.Request{ChaincodeID: ccID, Fcn: "approveCode", Args: invokeArgs})
	if err != nil {
		fmt.Println("Failed to query values: %s", err)
	}
	fmt.Println("response value: ", string(value))

    return string(value)
}

func fabric_execute_code(LuaChaincodeId string) []ExecutionResponse{
    code := fabric_query_code(LuaChaincodeId)
    var executionResponse []ExecutionResponse
    var i = 0
    invokeArgs := [][]byte{[]byte(LuaChaincodeId)}
    for i < len(code.Targets) {
        var execution ExecutionResponse
        peers, err := sdk.Config().PeersConfig(code.Targets[i])
        if err != nil {
            fmt.Println("Failed to get organization peers: %s", err)
        }
        peer0, err := peer.New(sdk.Config(), peer.FromPeerConfig(&apiconfig.NetworkPeer{PeerConfig: peers[0]}))
        if err != nil {
            fmt.Println("Failed to get peer: %s", err)
        }
        value, _, err := chClient.Execute(apitxn.Request{ChaincodeID: ccID, Fcn: "exectuteCC", Args: invokeArgs}, apitxn.WithProposalProcessor(peer0))
        if err != nil {
            fmt.Println("Failed to execute lua chaincode: %s", err)
        }
        fmt.Println("response value: ", string(value))
        execution.orgName = code.Targets[i]
        execution.executionResult = string(value)
        executionResponse = append(executionResponse , execution)
    }
	return executionResponse 
}

