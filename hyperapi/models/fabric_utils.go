package models

import (
    "fmt"
    "strings"
    "os"
	"github.com/astaxie/beego"
	"github.com/hyperledger/fabric-sdk-go/api/apitxn"
	"github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/common/cauthdsl"
    "github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"github.com/hyperledger/fabric-sdk-go/pkg/config"
	resmgmt "github.com/hyperledger/fabric-sdk-go/api/apitxn/resmgmtclient"
	packager "github.com/hyperledger/fabric-sdk-go/pkg/fabric-client/ccpackager/gopackager"
)


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
	fmt.Println("Instalado chaincode en peer")

	// Set up chaincode policy to 'any of two msps'
	ccPolicy:= cauthdsl.SignedByAnyMember([]string{"org1MSP", "coreAdmMSP"})

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
	ccPolicy:= cauthdsl.SignedByAnyMember([]string{"org1MSP", "coreAdmMSP"})

	instantciateCReq := resmgmt.InstantiateCCRequest{Name: ccID, Path: path, Version: version, Args: initArgs, Policy: ccPolicy}
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

func fabric_add_code(name string, source string, targets []string) []byte{
    invokeArgs := [][]byte{[]byte(`{"Name": "` + name + `", Source": "` + source + `", "Target": [` + strings.Join(targets[:],",") + `]}`)}
    value, _, err := chClient.Execute(apitxn.Request{ChaincodeID: ccID, Fcn: "storeCode", Args: invokeArgs})
	if err != nil {
		fmt.Println("Failed to query values: %s", err)
	}
	fmt.Println("response value: ", string(value))
    return value
}

func fabric_query_users() []byte{
    invokeArgs := [][]byte{[]byte("")}
    value, _, err := chClient.Execute(apitxn.Request{ChaincodeID: ccID, Fcn: "getAlias", Args: invokeArgs})
	if err != nil {
		fmt.Println("Failed to query values: %s", err)
	}
	fmt.Println("response value: ", string(value))
    return value
}

func fabric_query_user(uid string) string{
	fmt.Println("response value: NOT IMPLEMENTED")
    return "1"
}

func fabric_query_codes() []byte{
    invokeArgs := [][]byte{[]byte("")}
    value, _, err := chClient.Execute(apitxn.Request{ChaincodeID: ccID, Fcn: "getListCC", Args: invokeArgs})
	if err != nil {
		fmt.Println("Failed to query values: %s", err)
	}
	fmt.Println("response value: ", string(value))
    return value
}
