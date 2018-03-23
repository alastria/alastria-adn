API Backend is a middleware to exchange information from front-end to hyperledger fabric chaincode and vice versa.
To run the api backend for the whole infrastructure, we need to:

* terminal 1: cd hyperapi && go run main.go conf/coreAdm.conf
* terminal 2: cd hyperapi && go run main.go conf/entity1.conf
* terminal 3: cd hyperapi && go run main.go conf/entity2.conf

For the first time running the application you have to run the server with createChaincodeFirstTime() and createChaincodeLuaExecutorFirstTime() uncommented in order to install all the chaincodes we need to run this demo


DISTRIBUTED CERTIFICATIONS

According to our sdk confs we need following certifications for each entity

