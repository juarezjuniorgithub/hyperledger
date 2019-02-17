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

type ContainerTraceChaincode struct {
}

type container struct {
	ObjectType         string `json:"docType"`       
	DocumentNumber       string `json:"documentNumber"` 
	BatchNumber       string `json:"batchNumber"`	
	TagId              string `json:"tagId"`  
}

func main() {
	err := shim.Start(new(ContainerTraceChaincode))
	if err != nil {
		fmt.Printf("Error starting Container Trace chaincode: %s", err)
	}
}

func (t *ContainerTraceChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

func (t *ContainerTraceChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	fmt.Println("invoke is running " + function)

	
	if function == "initContainer" { 
		return t.initContainer(stub, args)	
	} else if function == "readContainer" { 
		return t.readContainer(stub, args)
	} else if function == "queryContainerByDocumentNumber" { 
		return t.queryContainerByDocumentNumber(stub, args)
	} else if function == "queryContainer" { 
		return t.queryContainer(stub, args)
	} else if function == "getHistoryForContainerRecord" { 
		return t.getHistoryForContainerRecord(stub, args)
	} 

	fmt.Println("invoke did not find func: " + function) 
	return shim.Error("Received unknown function invocation")
}

func (t *ContainerTraceChaincode) initContainer(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	// data entity with medical device atributes (JSON)
	//   0       	    1      		    2     	 
	// "documentNumber", "batchNumber", "tagId"

	// if len(args) != 3 {
	// 	return shim.Error("Incorrect number of arguments. Expecting 3")
	// }
	
	
	fmt.Println("- start init Container part")
	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}
	
	documentNumber := args[0]
	batchNumber := args[1]
	tagId := args[2]
	
		
	// ==== Check if Container part already exists ====
	// ContainerAsBytes, err := stub.GetState(documentNumber)
	// if err != nil {
	// 	return shim.Error("Failed to get the device: " + err.Error())
	// } else if containerAsBytes != nil {
	// 	fmt.Println("This container already exists: " + documentNumber)
	// 	return shim.Error("This container already exists: " + documentNumber)
	// }

	objectType := "container"
	container := &container{objectType, documentNumber, batchNumber, tagId}
	containerJSONasBytes, err := json.Marshal(container)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(documentNumber, containerJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end init container")
	return shim.Success(nil)
}

func (t *ContainerTraceChaincode) readContainer(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var documentNumber, jsonResp string
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting device number of the container to query")
	}

	documentNumber = args[0]
	valAsbytes, err := stub.GetState(documentNumber) 
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + documentNumber + "\"}"
		fmt.Println(jsonResp)
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"Container does not exist: " + documentNumber + "\"}"
		fmt.Println(jsonResp)
		return shim.Error(jsonResp)
	}

	return shim.Success(valAsbytes)
}
func (t *ContainerTraceChaincode) queryContainerByDocumentNumber(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	documentNumber := strings.ToLower(args[0])

	queryString := fmt.Sprintf("SELECT valueJson FROM <STATE> WHERE json_extract(valueJson, '$.docType', '$.documentNumber') = '[\"container\",\"%s\"]'", documentNumber)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

func (t *ContainerTraceChaincode) queryContainer(stub shim.ChaincodeStubInterface, args []string) pb.Response {

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

func (t *ContainerTraceChaincode) getHistoryForContainerRecord(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	recordKey := args[0]

	fmt.Printf("- start getHistoryForContainerRecord: %s\n", recordKey)

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

	fmt.Printf("- getHistoryForContainerRecord returning:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}
