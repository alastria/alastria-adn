
#  Alastria DNA
Decentraliced network administration for Alastria platform.

Install prerequisites
 - Docker and Docker Compose
 - Go version 1.9.x [install](https://golang.org/doc/install)
 - Npm and Node.js [install](https://docs.npmjs.com/getting-started/installing-node)
 - npm install -g gulp-cli bower (It must be installed like normal user, not like root user)

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
For more information about Admin console please click [here](https://github.com/alastria/alastria-dna/blob/develop/dna-frontAdmin/README.md).

## Run Dna Alastria FrontEnd for the entity

You need to clone the folder if you need more entities

``` bash
cd ./dna-frontEntity
npm install
bower install
gulp serve
```
For more information about Entity console please click [here](https://github.com/alastria/alastria-dna/blob/develop/dna-frontEntity/README.md).

## Run the API Dna Alastria [Backend](hyperapi/README.md)

API Backend is a middleware to exchange information from front-end to hyperledger fabric chaincode and vice versa. To run the api backend for the whole infrastructure, we need to:


* terminal 1$ ```cd hyperapi && go run main.go conf/coreAdm.conf```
* terminal 2$ ```cd hyperapi && go run main.go conf/entity1.conf```
* terminal 3$ ```cd hyperapi && go run main.go conf/entity2.conf```

For the first time running the application you have to run the server with ```createChaincodeFirstTime()``` and ```createChaincodeLuaExecutorFirstTime()``` uncommented in order to install all the chaincodes we need to run this demo

