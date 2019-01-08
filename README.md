# Supplychain

This is a simple Supplychain application to start working with Hyperledger Fabric blockchain.

To run this example and invoke the chaincode. Pull this repository to any location you want. 

Go to the directory fabric-sample/scripts and run

$ ./bootstrap.sh

this will install Hyperledger Fabric 1.3 for you and all its required files.

Now its time to install and invoke your chaincode with Fabric Basic Network.

Go to the directory fabric-samples/SupplyChain and run 

$ ./startFabric.sh

This command will install chaincode on the hyperledger fabric and start fabric as well. 

To insert any data into the blockchain use the following terminal commands.

$ docker exec -e "CORE_PEER_LOCALMSPID=Org1MSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp" cli peer chaincode invoke -o orderer.example.com:7050 -C mychannel -n supplychain -c '{"Args":["addCustomer","customer_1","data1","data2","data3","data4","data5","data6"]}'

Use the command below to query that record.
$ docker exec -e "CORE_PEER_LOCALMSPID=Org1MSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp" cli peer chaincode invoke -o orderer.example.com:7050 -C mychannel -n supplychain -c '{"Args":["queryCustomer","customer_1"]}'


NOTE: To see the chaincode. go to director fabric-samples/chaincodes/supplychain. The chaincode file is there name mycode.go
