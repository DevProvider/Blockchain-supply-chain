package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
	//"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

//SmartContract is the data structure which represents this contract and on which  various contract lifecycle functions are attached
type SmartContract struct{}

//Model for Customer
type Customer struct {
	ObjectType		string    	`json:"Type"`
	CustomerID      string    	`json:"customerID"`
	Name			string    	`json:"name"`
	Address       	string   	`json:"address"`
	Location       	string    	`json:"location"`
	City         	string    	`json:"city"`
	Zip        		string   	`json:"zip"`
	BusinessPhone  	string    	`json:"businessPhone"`
}

type Order struct {
	ObjectType		string    	`json:"Type"`
	OrderID			string		`json:"orderID"`
	Fk_ProductID	string		`json:"fk_productID"`
	Fk_CustomerID	string		`json:"fk_customerID"`
	ShippingAddress	string		`json:"shippingAddress"`
	Quantity		string		`json:"quantity"`
	TotalPrice		string		`json:"totalPrice"`
	DateOfOrder		time.Time	`json:"dateOfOrder"`
	DateofReceiving	time.Time	`json:"dateofReceiving"`
}

type ShippingStatus struct {
	ObjectType		string    	`json:"Type"`
	ShippingID		string 		`json:"shippingID`
	Fk_OrderID		string		`json:"fk_orderID"`
	CurrentLocation	string 		`json:"currentLocation"`
	time_Stamp		time.Time	`json:"time_stamp"`
}

//Model for Product
type Product struct {
	ObjectType    	string    	`json:"Type"`
	ProductID		string		`json:"productID"`
	Name			string    	`json:"name"`
	ProductType		string 		`json:"productType"`
	Size			string 		`json:"size"`
	Price			string 		`json:"price"`
	Quantity		string 		`json:"quantity"`
	Company			string 		`json:"company"`
}

func (t *SmartContract) Init(stub shim.ChaincodeStubInterface) peer.Response {
	fmt.Println("Init Firing!")
	return shim.Success(nil)
}

func (t *SmartContract) Invoke(stub shim.ChaincodeStubInterface) peer.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := stub.GetFunctionAndParameters()
	fmt.Println("Chaincode Invoke Is Running " + function)

	if function == "addCustomer" { //Add new Customer
		return t.addCustomer(stub, args)
	}
	if function == "addOrder" { //Add new Order
		return t.addOrder(stub, args)
	}
	if function == "addShippingStatus" { //Add new ShippingStatus
		return t.addShippingStatus(stub, args)
	}
	if function == "addProduct" { //Add new Product
		return t.addProduct(stub, args)
	}
	if function == "queryCustomer" { //Add new Product
		return t.queryCustomer(stub, args)
	}
	if function == "queryOrder" { //Add new Product
		return t.queryOrder(stub, args)
	}
	if function == "queryShippingStatus" { //Add new Product
		return t.queryShippingStatus(stub, args)
	}
	if function == "queryProduct" { //Add new Product
		return t.queryProduct(stub, args)
	}

	fmt.Println("Invoke did not find specified function " + function)
	return shim.Error("Recieved unknown")
}

//Contract for Adding new Customer
func (t *SmartContract) addCustomer(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	var err error

	if len(args) != 7 {
		return shim.Error("Incorrect Number of Aruments. Expecting 7")
	}

	fmt.Println("Adding New Customer")

	// ==== Input sanitation ====
	if len(args[0]) <= 0 {
		return shim.Error("1st Argument Must be a Non-Empty String")
	}
	if len(args[1]) <= 0 {
		return shim.Error("2nd Argument Must be a Non-Empty String")
	}
	if len(args[2]) <= 0 {
		return shim.Error("3rd Argument Must be a Non-Empty String")
	}
	if len(args[3]) <= 0 {
		return shim.Error("4th Argument Must be a Non-Empty String")
	}
	if len(args[4]) <= 0 {
		return shim.Error("5th Argument Must be a Non-Empty String")
	}
	if len(args[5]) <= 0 {
		return shim.Error("6th Argument Must be a Non-Empty String")
	}
	if len(args[6]) <= 0 {
		return shim.Error("7th Argument Must be a Non-Empty String")
	}

	customerID := args[0]
	name := args[1]
	address := args[2]
	location := args[3]
	city := args[4]
	zip := args[5]
	businessPhone := args[6]

	// ======Check if Customer Already exists

	customerAsBytes, err := stub.GetState(customerID)
	if err != nil {
		return shim.Error("Transaction Failed with Error: " + err.Error())
	} else if customerAsBytes != nil {
		return shim.Error("The Inserted customer ID already Exists: " + customerID)
	}

	// ===== Create customer Object and Marshal to JSON

	objectType := "customer"
	customer := &Customer{objectType, customerID, name, address, location,	city, zip, businessPhone}
	customerJSONasBytes, err := json.Marshal(customer)

	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Save customer to State

	err = stub.PutState(customerID, customerJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Return Success

	fmt.Println("Successfully Saved customer")
	return shim.Success(nil)
}

//Contract for Adding new Order
func (t *SmartContract) addOrder(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	var err error

	if len(args) != 8 {
		return shim.Error("Incorrect Number of Aruments. Expecting 8")
	}

	fmt.Println("Adding New Order")

	// ==== Input sanitation ====
	if len(args[0]) <= 0 {
		return shim.Error("1st Argument Must be a Non-Empty String")
	}
	if len(args[1]) <= 0 {
		return shim.Error("2nd Argument Must be a Non-Empty String")
	}
	if len(args[2]) <= 0 {
		return shim.Error("3rd Argument Must be a Non-Empty String")
	}
	if len(args[3]) <= 0 {
		return shim.Error("4th Argument Must be a Non-Empty String")
	}
	if len(args[4]) <= 0 {
		return shim.Error("5th Argument Must be a Non-Empty String")
	}
	if len(args[5]) <= 0 {
		return shim.Error("6th Argument Must be a Non-Empty String")
	}
	if len(args[6]) <= 0 {
		return shim.Error("7th Argument Must be a Non-Empty String")
	}
	if len(args[7]) <= 0 {
		return shim.Error("8th Argument Must be a Non-Empty String")
	}

	orderID := args[0]
	fk_productID := args[1]
	fk_customerID := args[2]
	shippingAddress := args[3]
	quantity := args[4]
	totalPrice := args[5]
	dateOfOrder, err1 := time.Parse(time.RFC3339, args[6])
	if err1 != nil {
		return shim.Error(err.Error())
	}
	dateofReceiving, err1 := time.Parse(time.RFC3339, args[7])
	if err1 != nil {
		return shim.Error(err.Error())
	}
	// ======Check if Order Already exists

	orderAsBytes, err := stub.GetState(orderID)
	if err != nil {
		return shim.Error("Transaction Failed with Error: " + err.Error())
	} else if orderAsBytes != nil {
		return shim.Error("The Inserted order ID already Exists: " + orderID)
	}

	// ===== Create order Object and Marshal to JSON

	objectType := "order"
	order := &Order{objectType, orderID, fk_productID, fk_customerID, shippingAddress,	quantity, totalPrice, dateOfOrder, dateofReceiving}
	orderJSONasBytes, err := json.Marshal(order)

	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Save order to State

	err = stub.PutState(orderID, orderJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Return Success

	fmt.Println("Successfully Saved order")
	return shim.Success(nil)
}

//Contract for Adding new ShippingStatus
func (t *SmartContract) addShippingStatus(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	var err error

	if len(args) != 4 {
		return shim.Error("Incorrect Number of Aruments. Expecting 4")
	}

	fmt.Println("Adding New ShippingStatus")

	// ==== Input sanitation ====
	if len(args[0]) <= 0 {
		return shim.Error("1st Argument Must be a Non-Empty String")
	}
	if len(args[1]) <= 0 {
		return shim.Error("2nd Argument Must be a Non-Empty String")
	}
	if len(args[2]) <= 0 {
		return shim.Error("3rd Argument Must be a Non-Empty String")
	}
	if len(args[3]) <= 0 {
		return shim.Error("4th Argument Must be a Non-Empty String")
	}

	shippingID := args[0]
	fk_orderID := args[1]
	currentLocation := args[2]
	time_stamp, err1 := time.Parse(time.RFC3339, args[3])
	if err1 != nil {
		return shim.Error(err.Error())
	}
	// ======Check if ShippingStatus Already exists

	shippingAsBytes, err := stub.GetState(shippingID)
	if err != nil {
		return shim.Error("Transaction Failed with Error: " + err.Error())
	} else if shippingAsBytes != nil {
		return shim.Error("The Inserted shipping ID already Exists: " + shippingID)
	}

	// ===== Create order Object and Marshal to JSON

	objectType := "shipping"
	shipping := &ShippingStatus{objectType, shippingID, fk_orderID, currentLocation, time_stamp}
	shippingJSONasBytes, err := json.Marshal(shipping)

	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Save shipping to State

	err = stub.PutState(shippingID, shippingJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Return Success

	fmt.Println("Successfully Saved shippingStatus")
	return shim.Success(nil)
}

//Contract for Adding new Product
func (t *SmartContract) addProduct(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	var err error

	if len(args) != 7 {
		return shim.Error("Incorrect Number of Aruments. Expecting 7")
	}

	fmt.Println("Adding New Product")

	// ==== Input sanitation ====
	if len(args[0]) <= 0 {
		return shim.Error("1st Argument Must be a Non-Empty String")
	}
	if len(args[1]) <= 0 {
		return shim.Error("2nd Argument Must be a Non-Empty String")
	}
	if len(args[2]) <= 0 {
		return shim.Error("3rd Argument Must be a Non-Empty String")
	}
	if len(args[3]) <= 0 {
		return shim.Error("4th Argument Must be a Non-Empty String")
	}
	if len(args[4]) <= 0 {
		return shim.Error("4th Argument Must be a Non-Empty String")
	}
	if len(args[5]) <= 0 {
		return shim.Error("4th Argument Must be a Non-Empty String")
	}
	if len(args[6]) <= 0 {
		return shim.Error("4th Argument Must be a Non-Empty String")
	}

	productID := args[0]
	name := args[1]	
	productType := args[2]
	size := args[1]
	price := args[1]
	quantity := args[1]
	company := args[1]

	// ======Check if ShippingStatus Already exists

	productAsBytes, err := stub.GetState(productID)
	if err != nil {
		return shim.Error("Transaction Failed with Error: " + err.Error())
	} else if productAsBytes != nil {
		return shim.Error("The Inserted product ID already Exists: " + productID)
	}

	// ===== Create order Object and Marshal to JSON

	objectType := "product"
	product := &Product{objectType, productID, name, productType, size, price, quantity, company}
	productJSONasBytes, err := json.Marshal(product)

	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Save product to State

	err = stub.PutState(productID, productJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Return Success

	fmt.Println("Successfully Saved product")
	return shim.Success(nil)
}

func (t *SmartContract) queryCustomer(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	customerID := args[0]

	queryString := fmt.Sprintf("{\"selector\":{\"Type\":\"customer\",\"customerID\":\"%s\"}}", customerID)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
	
}

func (t *SmartContract) queryOrder(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	orderID := args[0]

	queryString := fmt.Sprintf("{\"selector\":{\"Type\":\"order\",\"orderID\":\"%s\"}}", orderID)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
	
}

func (t *SmartContract) queryShippingStatus(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	shippingID := args[0]

	queryString := fmt.Sprintf("{\"selector\":{\"Type\":\"shippingstatus\",\"shippingID\":\"%s\"}}", shippingID)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
	
}

func (t *SmartContract) queryProduct(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	productID := args[0]

	queryString := fmt.Sprintf("{\"selector\":{\"Type\":\"product\",\"productID\":\"%s\"}}", productID)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
	
}















// =========================================================================================
// getQueryResultForQueryString executes the passed in query string.
// Result set is built and returned as a byte array containing the JSON results.
// =========================================================================================
func getQueryResultForQueryString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {

	fmt.Printf("- getQueryResultForQueryString queryString:\n%s\n", queryString)

	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryRecords
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())

	return buffer.Bytes(), nil
}

//Main Function starts up the Chaincode
func main() {
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Smart Contract could not be run. Error Occured: %s", err)
	} else {
		fmt.Println("Smart Contract successfully Initiated")
	}
}
