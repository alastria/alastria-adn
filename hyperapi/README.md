# API Backend

Is a middleware to exchange information from front-end to hyperledger fabric chaincode and vice versa.

Let's see the *configuration files* used by the **API** to start the services that is in the directory *conf/*.

One of them is the configuration file of the *Regulator* (coreAdm.conf).

```code
// coreAdm.conf

httpport = 8083
runmode = dev
autorender = false
copyrequestbody = true
EnableDocs = true
channelID = channel
hyperledgerConfigYamlPath = conf/coreadmin_config.yaml
ccID = thecc38
chaincodePath = ../chaincodes
path = managementCC
version = 1
orgName = coreAdm
luaExecutorPath = luaMonitorAlastria
luaExecutorccID = invokeLua17
```

In this file we find configuration parameters to start the service of the Regulator, type the name of the service, the port that exposes or the environment in which it will be deployed (dev). In addition to the configuration of the service we have the file *coreadmin_config.yaml* commissioned with it to be able to use the **SDK of Go**.

Each entity will have a file *entity.conf* and a file *entity_config.yaml*. In this case we have 3 entities, the Regulator, Entity1 and Entity2.

First of all we will check that the Hyperledger network is up and running. For this, from a terminal we launch the command *docker ps* (to know how to start a Hyperledger Fabric network look at the [README.md](../hf-testnet/README.md) on hf-testnet)

```shell
docker ps

CONTAINER ID        IMAGE                        COMMAND                  CREATED             STATUS              PORTS                                                                    NAMES
217f7f98dde2        hyperledger/fabric-ca        "sh -c 'fabric-ca-se…"   35 seconds ago      Up 34 seconds       0.0.0.0:9054->7054/tcp                                                   ca_peerOrg2
6450f30931f3        hyperledger/fabric-ca        "sh -c 'fabric-ca-se…"   35 seconds ago      Up 33 seconds       0.0.0.0:8054->7054/tcp                                                   ca_peerOrg1
8f3aeae65189        hyperledger/fabric-ca        "sh -c 'fabric-ca-se…"   35 seconds ago      Up 34 seconds       0.0.0.0:7054->7054/tcp                                                   ca_peerCoreAdm
6eace4e42339        hyperledger/fabric-peer      "peer node start"        23 hours ago        Up 34 seconds       0.0.0.0:9051->7051/tcp, 0.0.0.0:9052->7052/tcp, 0.0.0.0:9053->7053/tcp   peer0.org2.alastria.com
360c864d71cd        hyperledger/fabric-orderer   "orderer"                23 hours ago        Up 34 seconds       0.0.0.0:7050->7050/tcp                                                   orderer.alastria.com
372faddc72de        hyperledger/fabric-peer      "peer node start"        23 hours ago        Up 33 seconds       0.0.0.0:7051-7053->7051-7053/tcp                                         peer0.coreAdm.alastria.com
8f6b454cf416        hyperledger/fabric-peer      "peer node start"        23 hours ago        Up 34 seconds       0.0.0.0:8051->7051/tcp, 0.0.0.0:8052->7052/tcp, 0.0.0.0:8053->7053/tcp   peer0.org1.alastria.com
```

Now let's start each one of the services. For this, we open a terminal and go to the directory *hyperapi/* of our project.

```shell
pwd
~/workdir/smartreg
cd hyperapi/
```

We are going to start the *Regulator* service first. For this we execute the following command in a terminal.

```shell
go run main.go conf/coreAdm.conf
```

This command, what it does is to up the beego service on the port *8083* with the configuration that has been indicated, in addition to installing the chaincode and registering the new organization within the network and creating a channel where the rest of entities will be added.

Once the Regulator service is started, we are ready to start the service of the Entity1. Again, in another terminal we launched the previous command but changing the configuration file with that of Entity1

```shell
go run main.go conf/entity1.conf
```

Again, this command deploy the service of Entity1 on port *8081* ,registers the entity as **org1** and puts it into a network, adds Entity1 to the channel created by the Regulator. If we look at the logs at the beginning of the application, we will see that an error appears when installing the chaincode. This is normal since the chaincode was already installed by the Regulator and can not be installed with two identical chaincodes.

We will now do the same with the Entity2.

```shell
go run main.go conf/entity2.conf
```

As in the previous entity, this command will start the service of Entity2 on port *8082*, register the entity in the network as **org2** and adds it to the channel together with Entity1 and the Regulator. Again in the logs we can see how it has given an error when installing the chaincode, but as in the previous case, the error is normal because that chaincode is already installed in the network.

Now the three services are started and we can connect it to the Fronts-end.