ORDERER_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/alastria.com/orderers/orderer.alastria.com/msp/tlscacerts/tlsca.alastria.com-cert.pem
CORE_PEER_LOCALMSPID="coreAdmMSP"
CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/coreAdm.alastria.com/peers/peer0.coreAdm.alastria.com/tls/ca.crt
CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/coreAdm.alastria.com/users/Admin@coreAdm.alastria.com/msp
CORE_PEER_ADDRESS=peer0.coreAdm.alastria.com:7051
CHANNEL_NAME=channel
CORE_PEER_TLS_ENABLED=true

peer channel create -o orderer.alastria.com:7050 -c $CHANNEL_NAME -f ./channel-artifacts/channel.tx --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA >&log.txt
cat log.txt
peer channel join -b $CHANNEL_NAME.block  >&log.txt
cat log.txt

CORE_PEER_LOCALMSPID="org1MSP"
CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.alastria.com/peers/peer0.org1.alastria.com/tls/ca.crt
CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.alastria.com/users/Admin@org1.alastria.com/msp
CORE_PEER_ADDRESS=peer0.org1.alastria.com:7051
peer channel join -b $CHANNEL_NAME.block  >&log.txt
cat log.txt

CORE_PEER_LOCALMSPID="org2MSP"
CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.alastria.com/peers/peer0.org2.alastria.com/tls/ca.crt
CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.alastria.com/users/Admin@org2.alastria.com/msp
CORE_PEER_ADDRESS=peer0.org2.alastria.com:7051
peer channel join -b $CHANNEL_NAME.block  >&log.txt
cat log.txt