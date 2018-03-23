# Hyperledger Fabric TestNet intallation guide

Pasos para levantar una red de Hyperledger Fabric.

1. Primero vamos a asegurarnos que no tenemos ningún contenedor de *Hyperledger Fabric* levantado o parado. Para ello ejecutamos el script _fabrickOps.sh_ con el argumento **clean**

```shell
./fabricOps.sh clean
```

2. Por último volvemos a ejecutar el script _fabricOps.sh_ con el argumento **start**

```shell
./fabricOps.sh start

...

=================================================
---------- Starting the network -----------------
=================================================

Starting peer0.org1.alastria.com
Recreating ca_peerCoreAdm
Recreating ca_peerOrg1
Recreating ca_peerOrg2
Starting orderer.alastria.com
Starting peer0.coreAdm.alastria.com
Starting peer0.org2.alastria.com
Starting cli
```

Aquí podremos ejecutar el comando *docker ps* para comprobar que se nos han creado todos los contenedores necesarios de nuestra red.

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