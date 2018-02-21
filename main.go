package main

import (
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/api/apitxn"

	"github.com/hyperledger/fabric-sdk-go/pkg/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

const (
	channelID = "channel"
	orgName   = "coreAdm"
	orgAdmin  = "Admin"
	ccID      = " "
)
const ExampleCCInitB = "200"

var initArgs = [][]byte{[]byte("init")}

func main() {

	// Creación del SDK
	sdk, err := fabsdk.New(config.FromFile("config.yaml"))
	if err != nil {
		fmt.Print("Failed to create new SDK: %s", err)
	}
	// coreResMgmt, err := sdk.NewClient(fabsdk.WithUser("Admin")).ResourceMgmt()
	// if err != nil {
	// 	fmt.Print("Failed to create new resource management client: %s", err)
	// }
	// fmt.Println("Creado el cliente")
	// //  Org2 resource management client
	// org2ResMgmt, err := sdk.NewClient(fabsdk.WithUser("Admin"), fabsdk.WithOrg("org2")).ResourceMgmt()
	// if err != nil {
	// 	fmt.Print("ERror")
	// }
	// org1ResMgmt, err := sdk.NewClient(fabsdk.WithUser("Admin"), fabsdk.WithOrg("org1")).ResourceMgmt()
	// if err != nil {
	// 	fmt.Print("ERror")
	// }

	// // // Crear el package para el chaincode
	// ccPkg, err := packager.NewCCPackage("mychainc", "chaincodes")
	// if err != nil {
	// 	fmt.Print(err)
	// }
	// installCCReq := resmgmt.InstallCCRequest{Name: "mycc1", Path: "mychainc", Version: "2", Package: ccPkg}
	// _, err = coreResMgmt.InstallCC(installCCReq)
	// if err != nil {
	// 	fmt.Println("Error al instalar")
	// }
	// _, err = org1ResMgmt.InstallCC(installCCReq)
	// if err != nil {
	// 	fmt.Println("Error al instalar")
	// }
	// _, err = org2ResMgmt.InstallCC(installCCReq)
	// if err != nil {
	// 	fmt.Println("Error al instalar")
	// }

	// ccPolicy := cauthdsl.SignedByAnyMember([]string{"org1MSP", "coreAdmMSP"})
	// instantciateCReq := resmgmt.InstantiateCCRequest{Name: "mycc1", Path: "mychainc", Version: "2", Args: initArgs, Policy: ccPolicy}
	// err = org1ResMgmt.InstantiateCC("channel", instantciateCReq)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// /** Invocation */
	chClientCoreAdminUser, err := sdk.NewClient(fabsdk.WithUser("Admin"), fabsdk.WithOrg("coreAdm")).Channel("channel")
	if err != nil {
		fmt.Println("Failed to create new channel client for Org1 user: %s", err)
	}
	// invokeArgs := [][]byte{[]byte(`{"Name": "mycc1", "Source": "ks20230", "Target":["org1", "org2"]}`)}
	// response, _, err := chClientCoreAdminUser.Execute(apitxn.Request{ChaincodeID: "mycc1", Fcn: "storeCode", Args: invokeArgs})
	// if err != nil {
	// 	fmt.Print("Failed to move funds: %s", err)
	// }
	// fmt.Printf(string(response[:]))
	// //guid := string(response[:])
	// getCodeArgs := [][]byte{[]byte("0")}

	// response2, err := chClientCoreAdminUser.Query(apitxn.Request{ChaincodeID: "mycc1", Fcn: "getCode", Args: getCodeArgs})
	// if err != nil {
	// 	fmt.Print("Failed to move funds: %s", err)
	// }
	// fmt.Printf("Response ----->" + string(response2[:]))

	response3, _, err := chClientCoreAdminUser.Execute(apitxn.Request{ChaincodeID: "mycc1", Fcn: "getListCC"})
	if err != nil {
		fmt.Print("Failed to move funds: %s", err)
	}
	fmt.Printf("Response ----->" + string(response3[:]))

	// chClientorg1User, err := sdk.NewClient(fabsdk.WithUser("Admin"), fabsdk.WithOrg("org1")).Channel("channel")
	// if err != nil {
	// 	fmt.Println("Failed to create new channel client for Org1 user: %s", err)
	// }
	// registrarOrg1 := [][]byte{[]byte(`org1`)}

	// response, _, err := chClientorg1User.Execute(apitxn.Request{ChaincodeID: "mycc1", Fcn: "registrar", Args: registrarOrg1})
	// if err != nil {
	// 	fmt.Print("Failed to move funds: %s", err)
	// }
	// fmt.Printf(string(response[:]))

	// chClientorg2User, err := sdk.NewClient(fabsdk.WithUser("Admin"), fabsdk.WithOrg("org2")).Channel("channel")
	// if err != nil {
	// 	fmt.Println("Failed to create new channel client for Org1 user: %s", err)
	// }
	// registrarOrg2 := [][]byte{[]byte(`org2`)}
	// response, _, err = chClientorg2User.Execute(apitxn.Request{ChaincodeID: "mycc1", Fcn: "registrar", Args: registrarOrg2})
	// if err != nil {
	// 	fmt.Print("Failed to move funds: %s", err)
	// }
	// fmt.Printf(string(response[:]))

	// aproval2 := [][]byte{[]byte(guid)}
	// response, _, err = chClientorg2User.Execute(apitxn.Request{ChaincodeID: "mycc1", Fcn: "approveCode", Args: aproval2})
	// if err != nil {
	// 	fmt.Print("Failed to move funds: %s", err)
	// }
	// fmt.Printf(string(response[:]))
	// response, _, err = chClientorg1User.Execute(apitxn.Request{ChaincodeID: "mycc1", Fcn: "getCode", Args: getCodeArgs})
	// if err != nil {
	// 	fmt.Print("Failed to move funds: %s", err)
	// }
	// fmt.Printf(string(response[:]))
}
