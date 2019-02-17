/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/

// ====CHAINCODE EXECUTION SAMPLES (BCS REST API) ==================
/*
#TEST transaction / Init ledger

curl -H "Content-type:application/json" -X POST http://localhost:3100/bcsgw/rest/v1/transaction/invocation -d '{"channel":"channel1","chaincode":"vehiclenet","method":"initLedgerB","args":["ser1234"],"chaincodeVer":"v1"}'
curl -H "Content-type:application/json" -X POST http://localhost:3100/bcsgw/rest/v1/transaction/invocation -d '{"channel":"channel1","chaincode":"vehiclenet","method":"initLedgerC","args":["ser1234"],"chaincodeVer":"v1"}'
curl -H "Content-type:application/json" -X POST http://localhost:3100/bcsgw/rest/v1/transaction/invocation -d '{"channel":"channel1","chaincode":"vehiclenet","method":"initLedgerD","args":["ser1234"],"chaincodeVer":"v1"}'

# TEST transaction / Add Car Part

curl -H "Content-type:application/json" -X POST http://localhost:3100/bcsgw/rest/v1/transaction/invocation -d '{"channel":"channel1","chaincode":"vehiclenet","method":"initVehiclePart","args":["ser1234", "tata", "1502688979", "airbag 2020", "mazda", "false", "1502688979"],"chaincodeVer":"v1"}'
curl -H "Content-type:application/json" -X POST http://localhost:3100/bcsgw/rest/v1/transaction/invocation -d '{"channel":"channel1","chaincode":"vehiclenet","method":"initVehiclePart","args":["ser1235", "tata", "1502688979", "airbag 2020", "mercedes", "false", "1502688979"],"chaincodeVer":"v1"}'
curl -H "Content-type:application/json" -X POST http://localhost:3100/bcsgw/rest/v1/transaction/invocation -d '{"channel":"channel1","chaincode":"vehiclenet","method":"initVehiclePart","args":["ser1236", "tata", "1502688979", "airbag 2020", "toyota", "false", "15026889790"],"chaincodeVer":"v1"}'
curl -H "Content-type:application/json" -X POST http://localhost:3100/bcsgw/rest/v1/transaction/invocation -d '{"channel":"channel1","chaincode":"vehiclenet","method":"initVehiclePart","args":["ser1237", "tata", "1502688979", "airbag 5000", "mazda", "false", "1502688979"],"chaincodeVer":"v1"}'
curl -H "Content-type:application/json" -X POST http://localhost:3100/bcsgw/rest/v1/transaction/invocation -d '{"channel":"channel1","chaincode":"vehiclenet","method":"initVehiclePart","args":["ser1238", "tata", "1502688979", "airbag 5000", "mercedes", "false", "1502688979"],"chaincodeVer":"v1"}'
curl -H "Content-type:application/json" -X POST http://localhost:3100/bcsgw/rest/v1/transaction/invocation -d '{"channel":"channel1","chaincode":"vehiclenet","method":"initVehiclePart","args":["ser1239", "tata", "1502688979", "airbag 5000", "toyota", "false", "15026889790"],"chaincodeVer":"v1"}'

# TEST transaction / Add Car

curl -H "Content-type:application/json" -X POST http://localhost:3100/bcsgw/rest/v1/transaction/invocation -d '{"channel":"channel1","chaincode":"vehiclenet","method":"initVehicle","args":["mer1000001", "mercedes", "c class", "1502688979", "ser1234", "mercedes", "false", "1502688979"],"chaincodeVer":"v1"}'
curl -H "Content-type:application/json" -X POST http://localhost:3100/bcsgw/rest/v1/transaction/invocation -d '{"channel":"channel1","chaincode":"vehiclenet","method":"initVehicle","args":["maz1000001", "mazda", "mazda 6", "1502688979", "ser1235", "mazda", "false", "1502688979"],"chaincodeVer":"v1"}'
curl -H "Content-type:application/json" -X POST http://localhost:3100/bcsgw/rest/v1/transaction/invocation -d '{"channel":"channel1","chaincode":"vehiclenet","method":"initVehicle","args":["ren1000001", "renault", "megan", "1502688979", "ser1236", "renault", "false", "1502688979"],"chaincodeVer":"v1"}'
curl -H "Content-type:application/json" -X POST http://localhost:3100/bcsgw/rest/v1/transaction/invocation -d '{"channel":"channel1","chaincode":"vehiclenet","method":"initVehicle","args":["ford1000001", "ford", "mustang", "1502688979", "ser1237", "ford", "false", "1502688979"],"chaincodeVer":"v1"}'

# TEST query / Populated database

curl -H "Content-type:application/json" -X POST http://localhost:3100/bcsgw/rest/v1/transaction/query -d '{"channel":"channel1","chaincode":"vehiclenet","method":"readVehiclePart","args":["ser1234"],"chaincodeVer":"v1"}'
curl -H "Content-type:application/json" -X POST http://localhost:3100/bcsgw/rest/v1/transaction/query -d '{"channel":"channel1","chaincode":"vehiclenet","method":"readVehicle","args":["mer1000001"],"chaincodeVer":"v1"}'

# TEST transaction / Transfer ownership

curl -H "Content-type:application/json" -X POST http://localhost:3100/bcsgw/rest/v1/transaction/invocation -d '{"channel":"channel1","chaincode":"vehiclenet","method":"transferVehiclePart","args":["ser1234", "mercedes"],"chaincodeVer":"v1"}'
curl -H "Content-type:application/json" -X POST http://localhost:3100/bcsgw/rest/v1/transaction/invocation -d '{"channel":"channel1","chaincode":"vehiclenet","method":"transferVehicle","args":["mer1000001", "mercedes los angeles"],"chaincodeVer":"v1"}'

# TEST query / Get History

curl -H "Content-type:application/json" -X POST http://localhost:3100/bcsgw/rest/v1/transaction/query -d '{"channel":"channel1","chaincode":"vehiclenet","method":"getHistoryForRecord","args":["ser1234"],"chaincodeVer":"v1"}'
curl -H "Content-type:application/json" -X POST http://localhost:3100/bcsgw/rest/v1/transaction/query -d '{"channel":"channel1","chaincode":"vehiclenet","method":"getHistoryForRecord","args":["mer1000001"],"chaincodeVer":"v1"}'

# TEST transaction / delete records

curl -H "Content-type:application/json" -X POST http://localhost:3100/bcsgw/rest/v1/transaction/invocation -d '{"channel":"channel1","chaincode":"vehiclenet","method":"deleteVehiclePart","args":["ser1235"],"chaincodeVer":"v1"}'
curl -H "Content-type:application/json" -X POST http://localhost:3100/bcsgw/rest/v1/transaction/invocation -d '{"channel":"channel1","chaincode":"vehiclenet","method":"deleteVehicle","args":["maz1000001"],"chaincodeVer":"v1"}'

# TEST transaction / Recall Part

curl -H "Content-type:application/json" -X POST http://localhost:3100/bcsgw/{"channel":"channel1","chaincode":"vehiclenet","method":"setPartRecallState","args":["abg1234",true],"chaincodeVer":"v3"}'



CRYPTO
#Sign
go run cryptoHOL.go -s welcome

#Verify
go run cryptoHOL.go -v welcome 23465785510810132448457841429882907809251724155505686786147550387897 10848776947772665661803987914449872333300709981875993855742805426849
*/

// Index for chaincodeid, docType, owner, size (descending order).
// Note that docType, owner and size fields must be prefixed with the "data" wrapper
// chaincodeid must be added for all queries
//
// Definition for use with Fauxton interface
// {"index":{"fields":[{"data.size":"desc"},{"chaincodeid":"desc"},{"data.docType":"desc"},{"data.owner":"desc"}]},"ddoc":"indexSizeSortDoc", "name":"indexSizeSortDesc","type":"json"}
//
// example curl definition for use with command line
// curl -i -X POST -H "Content-Type: application/json" -d "{\"index\":{\"fields\":[{\"data.size\":\"desc\"},{\"chaincodeid\":\"desc\"},{\"data.docType\":\"desc\"},{\"data.owner\":\"desc\"}]},\"ddoc\":\"indexSizeSortDoc\", \"name\":\"indexSizeSortDesc\",\"type\":\"json\"}" http://hostname:port/channelNameGoesHere/_index

// Rich Query with index design doc and index name specified (Only supported if CouchDB is used as state database):
//   peer chaincode query -C channelNameGoesHere -n vehicleParts -c '{"Args":["queryVehiclePart","{\"selector\":{\"docType\":\"vehiclePart\",\"owner\":\"mercedes\"}, \"use_index\":[\"_design/indexOwnerDoc\", \"indexOwner\"]}"]}'

// Rich Query with index design doc specified only (Only supported if CouchDB is used as state database):
//   peer chaincode query -C channelNameGoesHere -n vehicleParts -c '{"Args":["queryVehiclePart","{\"selector\":{\"docType\":{\"$eq\":\"vehiclePart\"},\"owner\":{\"$eq\":\"mercedes\"},\"assemblyDate\":{\"$gt\":1502688979}},\"fields\":[\"docType\",\"owner\",\"assemblyDate\"],\"sort\":[{\"assemblyDate\":\"desc\"}],\"use_index\":\"_design/indexSizeSortDoc\"}"]}'

package main

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// AutoTraceChaincode example simple Chaincode implementation
type AutoTraceChaincode struct {
}

// @MODIFY_HERE add recall fields to vehiclePart JSON object
type vehiclePart struct {
	ObjectType   string `json:"docType"` //docType is used to distinguish the various types of objects in state database
	SerialNumber string `json:"serialNumber"`
	Assembler    string `json:"assembler"` //the fieldtags are needed to keep case from bouncing around
	AssemblyDate int    `json:"assemblyDate"`
	Name         string `json:"name"`
	Owner        string `json:"owner"`
	Recall       bool   `json:"recall"`     // to be added at workshop
	RecallDate   int    `json:"recallDate"` // to be added at workshop
}

// @MODIFY_HERE add recall fields to vehicle JSON object
type vehicle struct {
	ObjectType         string `json:"docType"`       //docType is used to distinguish the various types of objects in state database
	ChassisNumber      string `json:"chassisNumber"` //the fieldtags are needed to keep case from bouncing around
	Manufacturer       string `json:"manufacturer"`
	Model              string `json:"model"`
	AssemblyDate       int    `json:"assemblyDate"`
	AirbagSerialNumber string `json:"airbagSerialNumber"`
	Owner              string `json:"owner"`
	Recall             bool   `json:"recall"`     // to be added at workshop
	RecallDate         int    `json:"recallDate"` // to be added at workshop
}

// ===================================================================================
// Main
// ===================================================================================
func main() {
	err := shim.Start(new(AutoTraceChaincode))
	if err != nil {
		fmt.Printf("Error starting Parts Trace chaincode: %s", err)
	}
}

// Init initializes chaincode
// ===========================
func (t *AutoTraceChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

// Invoke - Our entry point for Invocations
// ========================================
func (t *AutoTraceChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "initVehiclePart" { //create a new vehiclePart
		return t.initVehiclePart(stub, args)
	} else if function == "transferVehiclePart" { //change owner of a specific vehicle part
		return t.transferVehiclePart(stub, args)
	} else if function == "deleteVehiclePart" { //delete a vehicle part
		return t.deleteVehiclePart(stub, args)
	} else if function == "readVehiclePart" { //read a vehiclePart
		return t.readVehiclePart(stub, args)
	} else if function == "queryVehiclePartByOwner" { //find vehicle part for owner X using rich query
		return t.queryVehiclePartByOwner(stub, args)
	} else if function == "queryVehiclePart" { //find vehicle part based on an ad hoc rich query
		return t.queryVehiclePart(stub, args)
	} else if function == "getHistoryForRecord" { //get history of values for a record
		return t.getHistoryForRecord(stub, args)
	} else if function == "getVehiclePartByRange" { //get vehicle part based on range query
		return t.getVehiclePartByRange(stub, args)
	} else if function == "initVehicle" { //create a new vehicle
		return t.initVehicle(stub, args)
	} else if function == "transferVehicle" { //change owner of a specific vehicle
		return t.transferVehicle(stub, args)
	} else if function == "readVehicle" { //read a vehicle
		return t.readVehicle(stub, args)
	} else if function == "deleteVehicle" { //delete a vehicle
		return t.deleteVehicle(stub, args)
	} else if function == "transferPartToVehicle" { // transfer airbag to vehicle
		return t.transferPartToVehicle(stub, args)
	} else if function == "setPartRecallState" { // set recall state of vehicle part
		return t.setPartRecallState(stub, args)
	}

	// @MODIFY_HERE
	// ==== Write a sub-routine to mark a vehicle part as recalled by ".Name"

	// @MODIFY_HERE
	// ==== Write a sub-routine to mark a vehicle as recalled by ".Manufacturer" & ".Model"

	fmt.Println("invoke did not find func: " + function) //error
	return shim.Error("Received unknown function invocation")
}

// @MODIFY_HERE un-comment
// ============================================================
// setPartRecallState - sets recall field of a vehicle
// ============================================================
func (t *AutoTraceChaincode) setPartRecallState(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	// dexpects following arguements
	//   	0       		1
	// "serialNumber", "status (boolean)"
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}

	serialNumber := args[0]
	recall, err := strconv.ParseBool(args[1])
	if err != nil {
		return shim.Error("2nd argument must be a boolean string")
	}

	// ==== Check if vehicle part already exists ====
	vehiclePartAsBytes, err := stub.GetState(serialNumber)
	if err != nil {
		return shim.Error("Failed to get vehicle part: " + err.Error())
	} else if vehiclePartAsBytes == nil {
		fmt.Println("This vehicle part does not exist: " + serialNumber)
		return shim.Error("This vehicle part does not exist:: " + serialNumber)
	}

	vehiclePartJSON := vehiclePart{}
	err = json.Unmarshal(vehiclePartAsBytes, &vehiclePartJSON) //unmarshal it aka JSON.parse()
	if err != nil {
		fmt.Println("Unable to unmarshall vehicle part from byte to JSON object: " + serialNumber)
		return shim.Error("Unable to unmarshall vehicle part from byte to JSON object: " + serialNumber)
	}

	// ==== Create vehiclePart object and marshal to JSON ====
	objectType := "vehiclePart"
	vehiclePart := &vehiclePart{objectType, serialNumber, vehiclePartJSON.Assembler, vehiclePartJSON.AssemblyDate, vehiclePartJSON.Name, vehiclePartJSON.Owner, recall, 1502688979}
	vehiclePartJSONasBytes, err := json.Marshal(vehiclePart)
	if err != nil {
		return shim.Error(err.Error())
	}

	// === Save vehiclePart to state ===
	err = stub.PutState(serialNumber, vehiclePartJSONasBytes)

	// ==== Vehicle part saved. Return success ====
	fmt.Println("- end setPartRecallState")
	return shim.Success(nil)
}

// ============================================================
// initVehiclePart - create a new vehicle part, store into chaincode state
// ============================================================
func (t *AutoTraceChaincode) initVehiclePart(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	// data model without recall fields
	//   0       	1      		2     		3				4
	// "ser1234", "tata", "1502688979", "airbag 2020", "aaimler ag / mercedes"

	// data model with recall fields
	//   0       	1      		2     		3				4						5	  6
	// "ser1234", "tata", "1502688979", "airbag 2020", "aaimler ag / mercedes", "false", "0"

	// @MODIFY_HERE extend to expect 7 arguements, up from 5
	if len(args) != 7 {
		return shim.Error("Incorrect number of arguments. Expecting 7")
	}

	// ==== Input sanitation ====
	fmt.Println("- start init vehicle part")
	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}
	if len(args[1]) <= 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}
	if len(args[3]) <= 0 {
		return shim.Error("4th argument must be a non-empty string")
	}
	if len(args[4]) <= 0 {
		return shim.Error("5th argument must be a non-empty string")
	}

	serialNumber := args[0]
	assembler := strings.ToLower(args[1])
	assemblyDate, err := strconv.Atoi(args[2])
	if err != nil {
		return shim.Error("3rd argument must be a numeric string")
	}
	name := strings.ToLower(args[3])
	owner := strings.ToLower(args[4])

	// @MODIFY_HERE parts recall fields
	recall, err := strconv.ParseBool(args[5])
	if err != nil {
		return shim.Error("6th argument must be a boolean string")
	}
	recallDate, err := strconv.Atoi(args[6])
	if err != nil {
		return shim.Error("7th argument must be a numeric string")
	}

	// ==== Check if vehicle part already exists ====
	vehiclePartAsBytes, err := stub.GetState(serialNumber)
	if err != nil {
		return shim.Error("Failed to get vehicle part: " + err.Error())
	} else if vehiclePartAsBytes != nil {
		fmt.Println("This vehicle part already exists: " + serialNumber)
		return shim.Error("This vehicle part already exists: " + serialNumber)
	}

	// @MODIFY_HERE parts recall fields
	// ==== Create vehiclePart object and marshal to JSON ====
	objectType := "vehiclePart"
	//vehiclePart := &vehiclePart{objectType, serialNumber, assembler, assemblyDate, name, owner}
	vehiclePart := &vehiclePart{objectType, serialNumber, assembler, assemblyDate, name, owner, recall, recallDate}
	vehiclePartJSONasBytes, err := json.Marshal(vehiclePart)
	if err != nil {
		return shim.Error(err.Error())
	}

	// === Save vehiclePart to state ===
	err = stub.PutState(serialNumber, vehiclePartJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	//  ==== Index the vehicle parts to enable assember & owner-based range queries, e.g. return all tata parts ====
	//  An 'index' is a normal key/value entry in state.
	//  The key is a composite key, with the elements that you want to range query on listed first.
	//  In our case, the composite key is based on indexName~assember~serialNumber.
	//  This will enable very efficißent state range queries based on composite keys matching indexName~color~*
	indexName := "assembler~serialNumber"
	ownersIndex := "owner~identifier"
	err = t.createIndex(stub, indexName, []string{vehiclePart.Assembler, vehiclePart.SerialNumber})
	if err != nil {
		return shim.Error(err.Error())
	}
	err = t.createIndex(stub, ownersIndex, []string{vehiclePart.Owner, vehiclePart.SerialNumber})
	if err != nil {
		return shim.Error(err.Error())
	}

	// ==== Vehicle part saved and indexed. Return success ====
	fmt.Println("- end init vehicle part")
	return shim.Success(nil)
}

// ============================================================
// initVehicle - create a new vehicle , store into chaincode state
// ============================================================
func (t *AutoTraceChaincode) initVehicle(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	// data model without recall fields
	//   0       		1      		2     		3			   4		5
	// "mer1000001", "mercedes", "c class", "1502688979", "ser1234", "mercedes"

	// data model with recall fields
	//   0       		1      		2     		3			   4		5	       6			7
	// "mer1000001", "mercedes", "c class", "1502688979", "ser1234", "mercedes", "false", "1502688979"

	// @MODIFY_HERE extend to expect 8 arguements, up from 6
	if len(args) != 8 {
		return shim.Error("Incorrect number of arguments. Expecting 7")
	}

	// ==== Input sanitation ====
	fmt.Println("- start init vehicle")
	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}
	if len(args[1]) <= 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}
	if len(args[2]) <= 0 {
		return shim.Error("3rd argument must be a non-empty string")
	}
	if len(args[4]) <= 0 {
		return shim.Error("5th argument must be a non-empty string")
	}
	if len(args[5]) <= 0 {
		return shim.Error("6th argument must be a non-empty string")
	}

	chassisNumber := args[0]
	manufacturer := strings.ToLower(args[1])
	model := strings.ToLower(args[2])
	assemblyDate, err := strconv.Atoi(args[3])
	if err != nil {
		return shim.Error("3rd argument must be a numeric string")
	}
	airbagSerialNumber := strings.ToLower(args[4])
	owner := strings.ToLower(args[5])

	// @MODIFY_HERE vehicle recall fields
	recall, err := strconv.ParseBool(args[6])
	if err != nil {
		return shim.Error("7th argument must be a boolean string")
	}
	recallDate, err := strconv.Atoi(args[7])
	if err != nil {
		return shim.Error("8th argument must be a numeric string")
	}

	// ==== Check if vehicle already exists ====
	vehicleAsBytes, err := stub.GetState(chassisNumber)
	if err != nil {
		return shim.Error("Failed to get vehicle: " + err.Error())
	} else if vehicleAsBytes != nil {
		return shim.Error("This vehicle already exists: " + chassisNumber)
	}

	// @MODIFY_HERE parts recall fields
	// ==== Create vehicle object and marshal to JSON ====
	objectType := "vehicle"
	//vehicle := &vehicle{objectType, chassisNumber, manufacturer, model, assemblyDate, airbagSerialNumber, owner}
	vehicle := &vehicle{objectType, chassisNumber, manufacturer, model, assemblyDate, airbagSerialNumber, owner, recall, recallDate}
	vehicleJSONasBytes, err := json.Marshal(vehicle)
	if err != nil {
		return shim.Error(err.Error())
	}

	// === Save vehicle to state ===
	err = stub.PutState(chassisNumber, vehicleJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	//  ==== Index the vehicle parts to enable assember & owner-based range queries, e.g. return all tata parts ====
	//  An 'index' is a normal key/value entry in state.
	//  The key is a composite key, with the elements that you want to range query on listed first.
	//  In our case, the composite key is based on indexName~assember~chassisNumber.
	//  This will enable very efficient state range queries based on composite keys matching indexName~color~*
	indexName := "manufacturer~chassisNumber"
	ownersIndex := "owner~identifier"
	err = t.createIndex(stub, indexName, []string{vehicle.Manufacturer, vehicle.ChassisNumber})
	if err != nil {
		return shim.Error(err.Error())
	}
	err = t.createIndex(stub, ownersIndex, []string{vehicle.Owner, vehicle.ChassisNumber})
	if err != nil {
		return shim.Error(err.Error())
	}

	// ==== Vehicle part saved and indexed. Return success ====
	fmt.Println("- end init vehicle")
	return shim.Success(nil)
}

// ===============================================
// createIndex - create search index for ledger
// ===============================================
func (t *AutoTraceChaincode) createIndex(stub shim.ChaincodeStubInterface, indexName string, attributes []string) error {
	fmt.Println("- start create index")
	var err error
	//  ==== Index the object to enable range queries, e.g. return all parts made by supplier b ====
	//  An 'index' is a normal key/value entry in state.
	//  The key is a composite key, with the elements that you want to range query on listed first.
	//  This will enable very efficient state range queries based on composite keys matching indexName~color~*
	indexKey, err := stub.CreateCompositeKey(indexName, attributes)
	if err != nil {
		return err
	}
	//  Save index entry to state. Only the key name is needed, no need to store a duplicate copy of object.
	//  Note - passing a 'nil' value will effectively delete the key from state, therefore we pass null character as value
	value := []byte{0x00}
	stub.PutState(indexKey, value)

	fmt.Println("- end create index")
	return nil
}

// ===============================================
// deleteIndex - remove search index for ledger
// ===============================================
func (t *AutoTraceChaincode) deleteIndex(stub shim.ChaincodeStubInterface, indexName string, attributes []string) error {
	fmt.Println("- start delete index")
	var err error
	//  ==== Index the object to enable range queries, e.g. return all parts made by supplier b ====
	//  An 'index' is a normal key/value entry in state.
	//  The key is a composite key, with the elements that you want to range query on listed first.
	//  This will enable very efficient state range queries based on composite keys matching indexName~color~*
	indexKey, err := stub.CreateCompositeKey(indexName, attributes)
	if err != nil {
		return err
	}
	//  Delete index by key
	stub.DelState(indexKey)

	fmt.Println("- end delete index")
	return nil
}

// ===============================================
// readVehiclePart - read a vehicle part from chaincode state
// ===============================================
func (t *AutoTraceChaincode) readVehiclePart(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var serialNumber, jsonResp string
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting serial number of the vehicle part to query")
	}

	serialNumber = args[0]
	valAsbytes, err := stub.GetState(serialNumber) //get the vehiclePart from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + serialNumber + "\"}"
		fmt.Println(jsonResp)
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"Vehicle part does not exist: " + serialNumber + "\"}"
		fmt.Println(jsonResp)
		return shim.Error(jsonResp)
	}

	return shim.Success(valAsbytes)
}

// ===============================================
// readVehicle - read a vehicle from chaincode state
// ===============================================
func (t *AutoTraceChaincode) readVehicle(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var chassisNumber, jsonResp string
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting chassis number of the vehicle to query")
	}

	chassisNumber = args[0]
	valAsbytes, err := stub.GetState(chassisNumber) //get the vehicle from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + chassisNumber + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"Vehicle does not exist: " + chassisNumber + "\"}"
		return shim.Error(jsonResp)
	}

	return shim.Success(valAsbytes)
}

// ==================================================
// deleteVehiclePart - remove a vehiclePart key/value pair from state
// ==================================================
func (t *AutoTraceChaincode) deleteVehiclePart(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var jsonResp string
	var vehiclePartJSON vehiclePart
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	serialNumber := args[0]

	// to maintain the assember~serialNumber index, we need to read the vehiclePart first and get its assembler
	valAsbytes, err := stub.GetState(serialNumber) //get the vehiclePart from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + serialNumber + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"VehiclePart does not exist: " + serialNumber + "\"}"
		return shim.Error(jsonResp)
	}

	err = json.Unmarshal([]byte(valAsbytes), &vehiclePartJSON)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to decode JSON of: " + serialNumber + "\"}"
		return shim.Error(jsonResp)
	}

	err = stub.DelState(serialNumber) //remove the vehiclePart from chaincode state
	if err != nil {
		return shim.Error("Failed to delete state:" + err.Error())
	}

	// maintain the index
	indexName := "assembler~serialNumber"
	ownersIndex := "owner~identifier"
	// remove previous indexes
	err = t.deleteIndex(stub, indexName, []string{vehiclePartJSON.Assembler, vehiclePartJSON.SerialNumber})
	if err != nil {
		return shim.Error(err.Error())
	}
	err = t.deleteIndex(stub, ownersIndex, []string{vehiclePartJSON.Owner, vehiclePartJSON.SerialNumber})
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

// ==================================================
// deleteVehicle - remove a vehicle key/value pair from state
// ==================================================
func (t *AutoTraceChaincode) deleteVehicle(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var jsonResp string
	var vehicleJSON vehicle
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	chassisNumber := args[0]

	// to maintain the manufacturer~chassisNumber index, we need to read the vehicle first and get its assembler
	valAsbytes, err := stub.GetState(chassisNumber) //get the vehicle from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + chassisNumber + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"Vehicle does not exist: " + chassisNumber + "\"}"
		return shim.Error(jsonResp)
	}

	err = json.Unmarshal([]byte(valAsbytes), &vehicleJSON)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to decode JSON of: " + chassisNumber + "\"}"
		return shim.Error(jsonResp)
	}

	err = stub.DelState(chassisNumber) //remove the vehicle from chaincode state
	if err != nil {
		return shim.Error("Failed to delete state:" + err.Error())
	}

	// maintain the index
	indexName := "manufacturer~chassisNumber"
	ownersIndex := "owner~identifier"
	// remove previous indexes
	err = t.deleteIndex(stub, indexName, []string{vehicleJSON.Manufacturer, vehicleJSON.ChassisNumber})
	if err != nil {
		return shim.Error(err.Error())
	}
	err = t.deleteIndex(stub, ownersIndex, []string{vehicleJSON.Owner, vehicleJSON.ChassisNumber})
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

// ===========================================================
// transfer a vehicle part by setting a new owner name on the vehiclePart
// ===========================================================
func (t *AutoTraceChaincode) transferVehiclePart(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	//   0       1       3
	// "name", "from", "to"
	if len(args) < 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	serialNumber := args[0]
	currentOwner := strings.ToLower(args[1])
	newOwner := strings.ToLower(args[2])
	fmt.Println("- start transferVehiclePart ", serialNumber, currentOwner, newOwner)

	message, err := t.transferPartHelper(stub, serialNumber, currentOwner, newOwner)
	if err != nil {
		return shim.Error(message + err.Error())
	} else if message != "" {
		return shim.Error(message)
	}

	fmt.Println("- end transferVehiclePart (success)")
	return shim.Success(nil)
}

// ===========================================================
// transferParts : helper method for transferVehiclePart
// ===========================================================
func (t *AutoTraceChaincode) transferPartHelper(stub shim.ChaincodeStubInterface, serialNumber string, currentOwner string, newOwner string) (string, error) {
	// attempt to get the current vehiclePart object by serial number.
	// if sucessful, returns us a byte array we can then us JSON.parse to unmarshal
	fmt.Println("Transfering part with serial number: " + serialNumber + " To: " + newOwner)
	vehiclePartAsBytes, err := stub.GetState(serialNumber)
	if err != nil {
		return "Failed to get vehicle part: " + serialNumber, err
	} else if vehiclePartAsBytes == nil {
		return "Vehicle part does not exist: " + serialNumber, nil
	}

	vehiclePartToTransfer := vehiclePart{}
	err = json.Unmarshal(vehiclePartAsBytes, &vehiclePartToTransfer) //unmarshal it aka JSON.parse()
	if err != nil {
		return "", err
	}

	if currentOwner != vehiclePartToTransfer.Owner {
		return "This asset is currently owned by another entity.", err
	}

	vehiclePartToTransfer.Owner = newOwner //change the owner

	vehiclePartJSONBytes, _ := json.Marshal(vehiclePartToTransfer)
	err = stub.PutState(serialNumber, vehiclePartJSONBytes) //rewrite the vehiclePart
	if err != nil {
		return "", err
	}

	// maintain indexes
	ownersIndex := "owner~identifier"
	// remove previous index
	err = t.deleteIndex(stub, ownersIndex, []string{currentOwner, serialNumber})
	if err != nil {
		return "", err
	}
	// create new index
	err = t.createIndex(stub, ownersIndex, []string{newOwner, serialNumber})
	if err != nil {
		return "", err
	}

	return "", nil
}

// ===========================================================
// transfer a vehicle part by setting a new owner name on the vehiclePart
// ===========================================================
func (t *AutoTraceChaincode) transferPartToVehicle(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("- start transferPartToVehicle")
	//   	0      			 1
	// "serialNumber", "chassisNumber"
	if len(args) < 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	serialNumber := args[0]
	chassisNumber := args[1]

	message, err := t.transferPartToVehicleHelper(stub, serialNumber, chassisNumber)
	if err != nil {
		return shim.Error(message)
	}

	fmt.Println("- end transferPartToVehicle (success)")
	return shim.Success(nil)
}

// ===========================================================
// transferPartToVehicleHelper : helper for transferPartToVehicle
// ===========================================================
func (t *AutoTraceChaincode) transferPartToVehicleHelper(stub shim.ChaincodeStubInterface, serialNumber string, chassisNumber string) (string, error) {
	vehiclePartAsBytes, err := stub.GetState(serialNumber)
	if err != nil {
		return "Failed to get vehicle part: " + serialNumber, err
	} else if vehiclePartAsBytes == nil {
		return "Vehicle part does not exist: " + serialNumber, nil
	}

	vehicleAsBytes, err := stub.GetState(chassisNumber)
	if err != nil {
		return "Failed to get vehicle: " + chassisNumber, err
	} else if vehicleAsBytes == nil {
		return "Vehicle does not exist: " + chassisNumber, err
	}

	part := vehiclePart{}
	err = json.Unmarshal(vehiclePartAsBytes, &part) //unmarshal it aka JSON.parse()
	if err != nil {
		return "", err
	}

	car := vehicle{}
	err = json.Unmarshal(vehicleAsBytes, &car) //unmarshal it aka JSON.parse()
	if err != nil {
		return "", err
	}

	if car.Owner != part.Owner {
		return "Illegal Transfer.", err
	}

	vehicleToModify := vehicle{}
	err = json.Unmarshal(vehicleAsBytes, &vehicleToModify) //unmarshal it aka JSON.parse()
	if err != nil {
		return "", err
	}
	vehicleToModify.AirbagSerialNumber = serialNumber //change the serialnumber of the vehicle

	vehicleJSONBytes, _ := json.Marshal(vehicleToModify)
	err = stub.PutState(chassisNumber, vehicleJSONBytes) //rewrite the vehicle
	if err != nil {
		return "", err
	}
	return "", nil
}

// ===========================================================
// transferVehicleHelper: transfer a vehicle  by setting a new owner name on the vehicle
// ===========================================================
func (t *AutoTraceChaincode) transferVehicle(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	//   0       1       3
	// "name", "from", "to"
	if len(args) < 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	chassisNumber := args[0]
	currentOnwer := strings.ToLower(args[1])
	newOwner := strings.ToLower(args[2])
	fmt.Println("- start transferVehicle ", chassisNumber, currentOnwer, newOwner)

	// attempt to get the current vehicle object by serial number.
	// if sucessful, returns us a byte array we can then us JSON.parse to unmarshal
	message, err := t.trannsferVehicleHelper(stub, chassisNumber, currentOnwer, newOwner)
	if err != nil {
		return shim.Error(message + err.Error())
	} else if message != "" {
		return shim.Error(message)
	}

	fmt.Println("- end transferVehicle (success)")
	return shim.Success(nil)
}

// ===========================================================
// trannsferVehicleHelper : helper method for transferVehicle
// ===========================================================
func (t *AutoTraceChaincode) trannsferVehicleHelper(stub shim.ChaincodeStubInterface, chassisNumber string, currentOwner string, newOwner string) (string, error) {
	// attempt to get the current vehicle object by serial number.
	// if sucessful, returns us a byte array we can then us JSON.parse to unmarshal
	fmt.Println("Transfering vehicle with chassis number: " + chassisNumber + " To: " + newOwner)
	vehicleAsBytes, err := stub.GetState(chassisNumber)
	if err != nil {
		return "Failed to get vehicle:", err
	} else if vehicleAsBytes == nil {
		return "Vehicle does not exist", err
	}

	vehicleToTransfer := vehicle{}
	err = json.Unmarshal(vehicleAsBytes, &vehicleToTransfer) //unmarshal it aka JSON.parse()
	if err != nil {
		return "", err
	}

	if currentOwner != vehicleToTransfer.Owner {
		return "This asset is currently owned by another entity.", err
	}

	vehicleToTransfer.Owner = newOwner //change the owner

	vehicleJSONBytes, _ := json.Marshal(vehicleToTransfer)
	err = stub.PutState(chassisNumber, vehicleJSONBytes) //rewrite the vehicle
	if err != nil {
		return "", err
	}

	// maintain indexes
	ownersIndex := "owner~identifier"
	// remove previous index
	err = t.deleteIndex(stub, ownersIndex, []string{currentOwner, chassisNumber})
	if err != nil {
		return "", err
	}
	// create new index
	err = t.createIndex(stub, ownersIndex, []string{newOwner, chassisNumber})
	if err != nil {
		return "", err
	}

	return "", nil
}

// ===========================================================================================
// getVehiclePartByRange performs a range query based on the start and end keys provided.

// Read-only function results are not typically submitted to ordering. If the read-only
// results are submitted to ordering, or if the query is used in an update transaction
// and submitted to ordering, then the committing peers will re-execute to guarantee that
// result sets are stable between endorsement time and commit time. The transaction is
// invalidated by the committing peers if the result set has changed between endorsement
// time and commit time.
// Therefore, range queries are a safe option for performing update transactions based on query results.
// ===========================================================================================
func (t *AutoTraceChaincode) getVehiclePartByRange(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) < 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	startKey := args[0]
	endKey := args[1]

	resultsIterator, err := stub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
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

	fmt.Printf("- getVehiclePartByRange queryResult:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

// =======Rich queries =========================================================================
// Two examples of rich queries are provided below (parameterized query and ad hoc query).

// Rich queries pass a query string to the state database.
// Rich queries are only supported by state database implementations
//  that support rich query (e.g. CouchDB).
// The query string is in the syntax of the underlying state database.
// With rich queries there is no guarantee that the result set hasn't changed between
//  endorsement time and commit time, aka 'phantom reads'.
// Therefore, rich queries should not be used in update transactions, unless the
// application handles the possibility of result set changes between endorsement and commit time.
// Rich queries can be used for point-in-time queries against a peer.
// ============================================================================================

// ===== Example: Parameterized rich query =================================================
// queryVehiclePartByOwner queries for vehicle part based on a passed in owner.
// This is an example of a parameterized query where the query logic is baked into the chaincode,
// and accepting a single query parameter (owner).
// Only available on state databases that support rich query (e.g. CouchDB)
// =========================================================================================
func (t *AutoTraceChaincode) queryVehiclePartByOwner(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	//   0
	// "bob"
	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	owner := strings.ToLower(args[0])

	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"vehiclePart\",\"owner\":\"%s\"}}", owner)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

// ===== Example: Ad hoc rich query ========================================================
// queryVehiclePart uses a query string to perform a query for vehiclePart.
// Query string matching state database syntax is passed in and executed as is.
// Supports ad hoc queries that can be defined at runtime by the client.
// If this is not desired, follow the queryVehiclePartByOwner example for parameterized queries.
// Only available on state databases that support rich query (e.g. CouchDB)
// =========================================================================================
func (t *AutoTraceChaincode) queryVehiclePart(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	//   0
	// "queryString"
	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	queryString := args[0]

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

// ===========================================================================================
// getHistoryForRecord returns the histotical state transitions for a given key of a record
// ===========================================================================================
func (t *AutoTraceChaincode) getHistoryForRecord(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	recordKey := args[0]

	fmt.Printf("- start getHistoryForRecord: %s\n", recordKey)

	resultsIterator, err := stub.GetHistoryForKey(recordKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing historic values for the key/value pair
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"TxId\":")
		buffer.WriteString("\"")
		buffer.WriteString(response.TxId)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Value\":")
		// if it was a delete operation on given key, then we need to set the
		//corresponding value null. Else, we will write the response.Value
		//as-is (as the Value itself a JSON vehiclePart)
		if response.IsDelete {
			buffer.WriteString("null")
		} else {
			buffer.WriteString(string(response.Value))
		}

		buffer.WriteString(", \"Timestamp\":")
		buffer.WriteString("\"")
		buffer.WriteString(time.Unix(response.Timestamp.Seconds, int64(response.Timestamp.Nanos)).String())
		buffer.WriteString("\"")

		buffer.WriteString(", \"IsDelete\":")
		buffer.WriteString("\"")
		buffer.WriteString(strconv.FormatBool(response.IsDelete))
		buffer.WriteString("\"")

		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getHistoryForRecord returning:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

// ===========================================================================================
// cryptoVerify : Verifies signed message against public key
// Public Key of Authority:
// [48 78 48 16 6 7 42 134 72 206 61 2 1 6 5 43 129 4 0 33 3 58 0 4 21 162 242 84 40 78 13 26 160 33 97 191 210 22 152 134 162 66 12 77 221 129 138 60 74 243 198 34 102 209 14 48 16 2 98 96 172 47 170 216 228 169 103 121 153 100 84 111 33 13 106 42 46 227 52 91]
// ===========================================================================================
func cryptoVerify(hash []byte, publicKeyBytes []byte, r *big.Int, s *big.Int) (result bool) {
	fmt.Println("- Verifying ECDSA signature")
	fmt.Println("Message")
	fmt.Println(hash)
	fmt.Println("Public Key")
	fmt.Println(publicKeyBytes)
	fmt.Println("r")
	fmt.Println(r)
	fmt.Println("s")
	fmt.Println(s)

	publicKey, err := x509.ParsePKIXPublicKey(publicKeyBytes)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	switch publicKey := publicKey.(type) {
	case *ecdsa.PublicKey:
		return ecdsa.Verify(publicKey, hash, r, s)
	default:
		return false
	}
}