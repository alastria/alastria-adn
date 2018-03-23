#  Alastria DNA
Decentraliced network administration for Alastria platform.

Install prerequisites
 - Docker and Docker Compose
 - Go version 1.9.x [install](https://golang.org/doc/install)
 - Npm and Node.js [install](https://docs.npmjs.com/getting-started/installing-node)
 - npm install -g gulp-cli bower

``` bash
go get github.com/hyperledger/fabric
go get github.com/hyperledger/fabric-sdk-go
````


Hyperledger fabric is an _Active_ Hyperledger project. Information on what _Active_ entails can be found in
the [Hyperledger Project Lifecycle document](https://wiki.hyperledger.org/community/project-lifecycle).

So Make sure you're working on the same release/commits
* Release Hyperledger Fabric v1.0.6
* Hyperledger Fabric Client SDK for Go v1.0.0-alpha2

``` bash
cd $GOPATH/src/github.com/hyperledger/fabric
git checkout 70f3f2
cd $GOPATH/src/github.com/hyperledger/fabric-sdk-go
git checkout b0efb7
```

If you have Alastria-Dna in other directory, then you shoud link it into $GOPATH/src directory:
``` bash
ln -s [Your path]/alastria-dna $GOPATH/src/github/
ln -s [Your path]/alastria-dna/hyperapi $GOPATH/src/
```

## Run the Test Network

The testnet archicture is made up by 3 organizations with one peer each organization. All the peers belongs to the same channel called "channel"
To run the testnet in development mode execute the following command:

``` bash
cd ./hf-testnet
./fabricOps.sh start
```

## Run Dna Alastria FrontEnd for the admin
``` bash
cd ./dna-frontAdmin
npm install
bower install
gulp serve
```
## Run Dna Alastria FrontEnd for the entity
You need to clone the folder if you need more entities
``` bash
cd ./dna-frontEntity
npm install
bower install
gulp serve
```

## Run the API Dna Alastria [Backend](hyperapi/README.md)

API Backend is a middleware to exchange information from front-end to hyperledger fabric chaincode and vice versa. To run the api backend for the whole infrastructure, we need to:


* terminal 1$ ```cd hyperapi && go run main.go conf/coreAdm.conf```
* terminal 2$ ```cd hyperapi && go run main.go conf/entity1.conf```
* terminal 3$ ```cd hyperapi && go run main.go conf/entity2.conf```

For the first time running the application you have to run the server with ```createChaincodeFirstTime()``` and ```createChaincodeLuaExecutorFirstTime()``` uncommented in order to install all the chaincodes we need to run this demo

# Alastria DNA Distributed (example)

To deploy the network in a distributed way we will need 3 EC2 machines. 
- admin
- node 1 (Org1)
- node 2 (Org2)

**hyperapi / conf-distributed /**
In the folder "hyperapi / conf-distributed /" we have an example of configuration files for the machines, in which the IP of each machine has been modified.
**hf-testnet / base-distributed**
In the base-distributed folder you can find the files docker-compose-base.yml and peer.yml with the extra_host and the IPs assigned to the peers.
**hf-testnet/crypto-config-distributed**
Here lies the Crypto Material generated for this example.
**hf-testnet/docker-compose-distributed.yaml**
This file is what we use to generate the containers.
**hf-testnet/fabricOp-distributed.sh**
In this script the following functions of *start* and *clean* have been commented for the example:
- #generateCerts
- #generateChannelArtifacts
- #replacePrivateKey
- #startNetwork
- And the *clean* function now only removes the containers, leaving the docker images intact.
In this way we do not rebuild our crypto materials.

With all this and assuming that we have the machines and the source code and all the dependencies installed in each machine, we proceed with the creation of the 
Hyperledger network.

### 1. In the admin ec2 machine:


    docker-compose -f docker-compose.yaml up -d caCoreAdm orderer.alastria.com peer0.coreAdm.alastria.com

  
### 2. In the node1 ec2 machine

    docker-compose -f docker-compose.yaml up -d caOrg1 peer0.org1.alastria.com

### 3. In the node2 ec2 machine

    docker-compose -f docker-compose.yaml up -d caOrg2 peer0.org2.alastria.com

### 4. Run the cli container in the admin ec2

    docker-compose -f docker-compose.yaml up cli

This 4 steps create the whole hf-testnet distributed.
