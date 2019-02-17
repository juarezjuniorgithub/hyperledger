// WINNER - 2nd runner - Irish Government Blockchain Hackathon 2019 - #BlockPirates Team
// https://www.youtube.com/watch?v=MRTzLKNPq14&
// https://twitter.com/BlockAthonIRE
// https://twitter.com/juarezjunior/status/1091778229487177728
// https://twitter.com/hashtag/BlockAthonIRE?src=hash

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

type MedDevTraceChaincode struct {
}

type medicalDevice struct {
	ObjectType         string `json:"docType"`       
	DeviceNumber       string `json:"deviceNumber"` 
	Manufacturer       string `json:"manufacturer"`
	Model              string `json:"model"`
	LocationData       string `json:"locationData"`
	AssemblyDate       string    `json:"assemblyDate"`
	DeviceSerialNumber string `json:"deviceSerialNumber"`
	Owner              string `json:"owner"`  
	Active             string   `json:"active"`     
	DecomissionDate    string    `json:"decomissionDate"`
}

func main() {
	err := shim.Start(new(MedDevTraceChaincode))
	if err != nil {
		fmt.Printf("Error starting Flat Trace chaincode: %s", err)
	}
}

func (t *MedDevTraceChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

func (t *MedDevTraceChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	fmt.Println("invoke is running " + function)

	
	if function == "initMedicalDevice" { 
		return t.initMedicalDevice(stub, args)	
	} else if function == "readMedicalDevice" { 
		return t.readMedicalDevice(stub, args)
	} else if function == "queryDeviceByDeviceNumber" { 
		return t.queryDeviceByDeviceNumber(stub, args)
	} else if function == "queryDevice" { 
		return t.queryDevice(stub, args)
	} else if function == "getHistoryForDeviceRecord" { 
		return t.getHistoryForDeviceRecord(stub, args)
	} 

	fmt.Println("invoke did not find func: " + function) 
	return shim.Error("Received unknown function invocation")
}

func (t *MedDevTraceChaincode) initMedicalDevice(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	// data entity with medical device atributes (JSON)
	//   0       	    1      		    2     	 3				4				5						6			7		8
	// "deviceNumber", "manufacturer", "model", "locationData", "assemblyDate", "deviceSerialNumber",  "owner",  "active",  "decomissionDate"

	// if len(args) != 8 {
	// 	return shim.Error("Incorrect number of arguments. Expecting 6")
	// }

	// type medicalDevice struct {
	// 	ObjectType         string `json:"docType"`       
	// 	DeviceNumber       string `json:"deviceNumber"` 
	// 	Manufacturer       string `json:"manufacturer"`
	// 	Model              string `json:"model"`
	// 	LocationData       string `json:"locationData"`
	// 	AssemblyDate       string    `json:"assemblyDate"`
	// 	DeviceSerialNumber string `json:"deviceSerialNumber"`
	// 	Owner              string `json:"owner"`  
	// 	Active             string   `json:"active"`     
	// 	DecomissionDate    string    `json:"decomissionDate"`
	// }

	
	fmt.Println("- start init medicalDevice part")
	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}
	

	deviceNumber := args[0]
	manufacturer := args[1]
	model := args[2]
	locationData := args[3]
	assemblyDate := args[4]
	deviceSerialNumber := args[5]
	owner := args[6]
	active := args[7]
	decomissionDate := args[8]
		
	// ==== Check if medicalDevice part already exists ====
	// medicalDeviceAsBytes, err := stub.GetState(deviceNumber)
	// if err != nil {
	// 	return shim.Error("Failed to get the device: " + err.Error())
	// } else if medicalDeviceAsBytes != nil {
	// 	fmt.Println("This device already exists: " + deviceNumber)
	// 	return shim.Error("This device part already exists: " + deviceNumber)
	// }

		objectType := "medicalDevice"
	medicalDevice := &medicalDevice{objectType, deviceNumber, manufacturer, model, locationData, assemblyDate, deviceSerialNumber, owner, active, decomissionDate}
	medicalDeviceJSONasBytes, err := json.Marshal(medicalDevice)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(deviceNumber, medicalDeviceJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end init medicalDevice")
	return shim.Success(nil)
}

func (t *MedDevTraceChaincode) readMedicalDevice(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var deviceNumber, jsonResp string
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting device number of the medicalDevice to query")
	}

	deviceNumber = args[0]
	valAsbytes, err := stub.GetState(deviceNumber) 
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + deviceNumber + "\"}"
		fmt.Println(jsonResp)
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"Vehicle part does not exist: " + deviceNumber + "\"}"
		fmt.Println(jsonResp)
		return shim.Error(jsonResp)
	}

	return shim.Success(valAsbytes)
}
func (t *MedDevTraceChaincode) queryDeviceByDeviceNumber(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	deviceNumber := strings.ToLower(args[0])

	queryString := fmt.Sprintf("SELECT valueJson FROM <STATE> WHERE json_extract(valueJson, '$.docType', '$.deviceNumber') = '[\"medicalDevice\",\"%s\"]'", deviceNumber)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

func (t *MedDevTraceChaincode) queryDevice(stub shim.ChaincodeStubInterface, args []string) pb.Response {

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

func (t *MedDevTraceChaincode) getHistoryForDeviceRecord(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	recordKey := args[0]

	fmt.Printf("- start getHistoryForDeviceRecord: %s\n", recordKey)

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

	fmt.Printf("- getHistoryForDeviceRecord returning:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}
