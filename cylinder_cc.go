package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type CylinderTraceChaincode struct {
}

type cylinder struct {
	ObjectType         string `json:"docType"`       
	DocumentNumber       string `json:"documentNumber"` 
	BatchNumber       string `json:"batchNumber"`
	SerialNumber              string `json:"serialNumber"`
	ValveBrand       string `json:"valveBrand"`
	ValveCode       string    `json:"valveCode"`
	CylinderPartNumber string `json:"cylinderPartNumber"`
	TagId              string `json:"tagId"`  
	Lot                string   `json:"lot"`     
	ManufacturingDate    string    `json:"manufacturingDate"`
}

func main() {
	err := shim.Start(new(CylinderTraceChaincode))
	if err != nil {
		fmt.Printf("Error starting Cylinder Trace chaincode: %s", err)
	}
}

func (t *CylinderTraceChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

func (t *CylinderTraceChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	fmt.Println("invoke is running " + function)

	
	if function == "initCylinder" { 
		return t.initCylinder(stub, args)	
	} else if function == "readCylinder" { 
		return t.readCylinder(stub, args)
	} else if function == "queryCylinderBySerialNumber" { 
		return t.queryCylinderBySerialNumber(stub, args)
	} else if function == "queryCylinder" { 
		return t.queryCylinder(stub, args)
	} else if function == "getHistoryForCylinderRecord" { 
		return t.getHistoryForCylinderRecord(stub, args)
	} 

	fmt.Println("invoke did not find func: " + function) 
	return shim.Error("Received unknown function invocation")
}

func (t *CylinderTraceChaincode) initCylinder(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	// data entity with medical device atributes (JSON)
	//   0       	    1      		    2     	 3				4				5						6			7		8
	// "documentNumber", "batchNumber", "serialNumber", "valveBrand", "valveCode", "cylinderPartNumber",  "tagId",  "lot",  "manufacturingDate"

	// if len(args) != 8 {
	// 	return shim.Error("Incorrect number of arguments. Expecting 9")
	// }	

	
	fmt.Println("- start init Cylinder part")
	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}
	

	documentNumber := args[0]
	batchNumber := args[1]
	serialNumber := args[2]
	valveBrand := args[3]
	valveCode := args[4]
	cylinderPartNumber := args[5]
	tagId := args[6]
	lot := args[7]
	manufacturingDate := args[8]


		
	// ==== Check if Cylinder part already exists ====
	// CylinderAsBytes, err := stub.GetState(serialNumber)
	// if err != nil {
	// 	return shim.Error("Failed to get the device: " + err.Error())
	// } else if cylinderAsBytes != nil {
	// 	fmt.Println("This cylinder already exists: " + serialNumber)
	// 	return shim.Error("This cylinder already exists: " + serialNumber)
	// }

		objectType := "cylinder"
	cylinder := &cylinder{objectType, documentNumber, batchNumber, serialNumber, valveBrand, valveCode, cylinderPartNumber, tagId, lot, manufacturingDate}
	cylinderJSONasBytes, err := json.Marshal(cylinder)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(serialNumber, cylinderJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end init cylinder")
	return shim.Success(nil)
}

func (t *CylinderTraceChaincode) readCylinder(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var serialNumber, jsonResp string
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting device number of the cylinder to query")
	}

	serialNumber = args[0]
	valAsbytes, err := stub.GetState(serialNumber) 
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + serialNumber + "\"}"
		fmt.Println(jsonResp)
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"Cylinder does not exist: " + serialNumber + "\"}"
		fmt.Println(jsonResp)
		return shim.Error(jsonResp)
	}

	return shim.Success(valAsbytes)
}
func (t *CylinderTraceChaincode) queryCylinderBySerialNumber(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	serialNumber := strings.ToLower(args[0])

	queryString := fmt.Sprintf("SELECT valueJson FROM <STATE> WHERE json_extract(valueJson, '$.docType', '$.serialNumber') = '[\"cylinder\",\"%s\"]'", serialNumber)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

func (t *CylinderTraceChaincode) queryCylinder(stub shim.ChaincodeStubInterface, args []string) pb.Response {

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

func getQueryResultForQueryString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {

	fmt.Printf("- getQueryResultForQueryString queryString:\n%s\n", queryString)

	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString(string(queryResponse.Value))
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())

	return buffer.Bytes(), nil
}

func (t *CylinderTraceChaincode) getHistoryForCylinderRecord(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	recordKey := args[0]

	fmt.Printf("- start getHistoryForCylinderRecord: %s\n", recordKey)

	resultsIterator, err := stub.GetHistoryForKey(recordKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
	
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"TxId\":")
		buffer.WriteString("\"")
		buffer.WriteString(response.TxId)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Value\":")
	
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

	fmt.Printf("- getHistoryForCylinderRecord returning:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}
