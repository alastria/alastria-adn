package main


import (
    "os"
	_ "hyperapi/routers"
    //"fmt"

	"github.com/astaxie/beego"
    
	//"github.com/hyperledger/fabric-sdk-go/api/apitxn"
	//"github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/common/cauthdsl"
    //"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	//"github.com/hyperledger/fabric-sdk-go/pkg/config"
	//resmgmt "github.com/hyperledger/fabric-sdk-go/api/apitxn/resmgmtclient"
	//packager "github.com/hyperledger/fabric-sdk-go/pkg/fabric-client/ccpackager/gopackager"
)
/*
var initArgs = [][]byte{[]byte("init")}

func createChaincodeFirstTime(){
    // initialize hyperledger sdk config
    channelID := beego.AppConfig.String("channelID")
    hyperledger_config_yaml := beego.AppConfig.String("hyperledgerConfigYamlPath")
    ccID := beego.AppConfig.String("ccID")
    chaincodePath := beego.AppConfig.String("chaincodePath")
    path := beego.AppConfig.String("path")
    version := beego.AppConfig.String("version")


    //SDK creation
	sdk, err := fabsdk.New(config.FromFile(hyperledger_config_yaml))
	if err != nil {
		fmt.Print("Failed to create new SDK: %s", err)
	}
	fmt.Println("Creado el SDK")

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

	fmt.Println("Instanciado")
	fmt.Println("Instanciado")
	fmt.Println("Instanciado")
	fmt.Println("Instanciado")
	fmt.Println("Instanciado")
}

func execute(){
    // initialize hyperledger sdk config
    channelID := beego.AppConfig.String("channelID")
    hyperledger_config_yaml := beego.AppConfig.String("hyperledgerConfigYamlPath")
    ccID := beego.AppConfig.String("ccID")
    //chaincodePath := beego.AppConfig.String("chaincodePath")
    //path := beego.AppConfig.String("path")
    //version := beego.AppConfig.String("version")

    //SDK creation
	sdk, err := fabsdk.New(config.FromFile(hyperledger_config_yaml))
	if err != nil {
		fmt.Print("Failed to create new SDK: %s", err)
	}
	fmt.Println("Creado el SDK")

	// admin client is used to query and execute transactions (Org1 is default org)
	chClient, err := sdk.NewClient(fabsdk.WithUser("Admin"), fabsdk.WithOrg("coreAdm")).Channel(channelID)
	if err != nil {
		fmt.Errorf("Failed to create new channel client: %s", err)
	}

    // //invoke
	invokeArgs := [][]byte{[]byte(`{"Name": "` + ccID + `", Source": "ks20230", "Target":["bank1", "bank2"]}`)}

	// Invoke chaincode
    res, _, err := chClient.Execute(apitxn.Request{ChaincodeID: "thecc", Fcn: "storeCode", Args: invokeArgs})
	if err != nil {
		fmt.Print("Failed to send code: %s", err)
	}

    fmt.Print("res: %s", res)
}

func query(){
    // initialize hyperledger sdk config
    channelID := beego.AppConfig.String("channelID")
    hyperledger_config_yaml := beego.AppConfig.String("hyperledgerConfigYamlPath")
    ccID := beego.AppConfig.String("ccID")
    //chaincodePath := beego.AppConfig.String("chaincodePath")
    //path := beego.AppConfig.String("path")
    //version := beego.AppConfig.String("version")

    //SDK creation
	sdk, err := fabsdk.New(config.FromFile(hyperledger_config_yaml))
	if err != nil {
		fmt.Print("Failed to create new SDK: %s", err)
	}
	fmt.Println("Creado el SDK")

	// admin client is used to query and execute transactions (Org1 is default org)
	chClient, err := sdk.NewClient(fabsdk.WithUser("Admin"), fabsdk.WithOrg("coreAdm")).Channel(channelID)
	if err != nil {
		fmt.Errorf("Failed to create new channel client: %s", err)
	}

    invokeArgs := [][]byte{[]byte("")}
    initialValue, _, err := chClient.Execute(apitxn.Request{ChaincodeID: ccID, Fcn: "getListCC", Args: invokeArgs})
	if err != nil {
		fmt.Println("Failed to query funds: %s", err)
	}
	fmt.Println("B value: ", string(initialValue))
}
*/

func main() {
    configFile := os.Args[1]
    beego.LoadAppConfig("ini", configFile)
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
    //createChaincodeFirstTime()
    //execute()
    //query()
	beego.Run()
}

