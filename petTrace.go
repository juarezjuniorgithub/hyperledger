package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type PetTraceChaincode struct {
}

type pet struct {
	ObjectType         string `json:"docType"`      
	PedigreeRegistrationNumber      string `json:"pedigreeRegistrationNumber"` 
	Breed       string `json:"breed"`
	BirthDate       int    `json:"birthDate"`
	Owner              string `json:"owner"`	
}

func main() {
	err := shim.Start(new(PetTraceChaincode))
	if err != nil {
		fmt.Printf("Error starting Pets Trace chaincode: %s", err)
	}
}

func (t *PetTraceChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

func (t *PetTraceChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	function, args := stub.GetFunctionAndParameters()
	fmt.Println("invoke is running " + function)
	
	if function == "initPet" { //create a new pet
		return t.initPet(stub, args)
	} else if function == "transferPet" { //change owner of a specific pet
		return t.transferPet(stub, args)
	} else if function == "readPet" { //read a pet
		return t.readPet(stub, args)
	} else if function == "deletePet" { //delete a pet
		return t.deletePet(stub, args)
	} 

	fmt.Println("invoke did not find func: " + function) //error
	return shim.Error("Received unknown function invocation")
}

func (t *PetTraceChaincode) initPet(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	var err error

	// data model 
	//   0       		1      		2     		3
	// "pet1000001", "bulldog", "1502688979", "juarez"
	
	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}

	// ==== Input sanitation ====
	fmt.Println("- start init pet")
	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}
	if len(args[1]) <= 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}
	if len(args[2]) <= 0 {
		return shim.Error("3rd argument must be a non-empty string")
	}
	if len(args[3]) <= 0 {
		return shim.Error("5th argument must be a non-empty string")
	}
	

	pedigreeRegistrationNumber := args[0]
	breed := strings.ToLower(args[1])	
	birthDate, err := strconv.Atoi(args[2])
	if err != nil {
		return shim.Error("2nd argument must be a numeric string")
	}	
	owner := strings.ToLower(args[3])

	// ==== Check if pet already exists ====
	petAsBytes, err := stub.GetState(pedigreeRegistrationNumber)
	if err != nil {
		return shim.Error("Failed to get pet: " + err.Error())
	} else if petAsBytes != nil {
		return shim.Error("This pet already exists: " + pedigreeRegistrationNumber)
	}

	// ==== Create pet object and marshal to JSON ====
	objectType := "pet"
	pet := &pet{objectType, pedigreeRegistrationNumber, breed, birthDate, owner}
	petJSONasBytes, err := json.Marshal(pet)
	if err != nil {
		return shim.Error(err.Error())
	}

	// === Save pet to state ===
	err = stub.PutState(pedigreeRegistrationNumber, petJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// ==== Pet saved and indexed. Return success ====
	fmt.Println("- end init pet")
	return shim.Success(nil)
}

// ===============================================
// createIndex - create search index for ledger
// ===============================================
func (t *PetTraceChaincode) createIndex(stub shim.ChaincodeStubInterface, indexName string, attributes []string) error {
	fmt.Println("- start create index")
	var err error
	
	indexKey, err := stub.CreateCompositeKey(indexName, attributes)
	if err != nil {
		return err
	}
	
	value := []byte{0x00}
	stub.PutState(indexKey, value)

	fmt.Println("- end create index")
	return nil
}


func (t *PetTraceChaincode) deleteIndex(stub shim.ChaincodeStubInterface, indexName string, attributes []string) error {
	fmt.Println("- start delete index")
	var err error
	
	indexKey, err := stub.CreateCompositeKey(indexName, attributes)
	if err != nil {
		return err
	}
	//  Delete index by key
	stub.DelState(indexKey)

	fmt.Println("- end delete index")
	return nil
}

func (t *PetTraceChaincode) readPet(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	var pedigreeRegistrationNumber, jsonResp string
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting pedigree number of the pet to query")
	}

	pedigreeRegistrationNumber = args[0]
	valAsbytes, err := stub.GetState(pedigreeRegistrationNumber) //get the pet from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + pedigreeRegistrationNumber + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"Pet does not exist: " + pedigreeRegistrationNumber + "\"}"
		return shim.Error(jsonResp)
	}

	return shim.Success(valAsbytes)
}


func (t *PetTraceChaincode) deletePet(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	var jsonResp string
	var petJSON pet
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	pedigreeRegistrationNumber := args[0]
	
	valAsbytes, err := stub.GetState(pedigreeRegistrationNumber) //get the pet from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + pedigreeRegistrationNumber + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"Pet does not exist: " + pedigreeRegistrationNumber + "\"}"
		return shim.Error(jsonResp)
	}

	err = json.Unmarshal([]byte(valAsbytes), &petJSON)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to decode JSON of: " + pedigreeRegistrationNumber + "\"}"
		return shim.Error(jsonResp)
	}

	err = stub.DelState(pedigreeRegistrationNumber) //remove the pet from chaincode state
	if err != nil {
		return shim.Error("Failed to delete state:" + err.Error())
	}

	return shim.Success(nil)
}

func (t *PetTraceChaincode) transferPet(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	//   0       1       3
	// "name", "from", "to"
	if len(args) < 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	pedigreeRegistrationNumber := args[0]
	currentOnwer := strings.ToLower(args[1])
	newOwner := strings.ToLower(args[2])
	fmt.Println("- start transferPet ", pedigreeRegistrationNumber, currentOnwer, newOwner)

	message, err := t.transferPetHelper(stub, pedigreeRegistrationNumber, currentOnwer, newOwner)
	if err != nil {
		return shim.Error(message + err.Error())
	} else if message != "" {
		return shim.Error(message)
	}

	fmt.Println("- end transferPet (success)")
	return shim.Success(nil)
}

func (t *PetTraceChaincode) transferPetHelper(stub shim.ChaincodeStubInterface, pedigreeRegistrationNumber string, currentOwner string, newOwner string) (string, error) {
	
	fmt.Println("Transfering pet with pedigree registration number: " + pedigreeRegistrationNumber + " To: " + newOwner)
	petAsBytes, err := stub.GetState(pedigreeRegistrationNumber)
	if err != nil {
		return "Failed to get pet:", err
	} else if petAsBytes == nil {
		return "Pet does not exist", err
	}

	petToTransfer := pet{}
	err = json.Unmarshal(petAsBytes, &petToTransfer) //unmarshal it aka JSON.parse()
	if err != nil {
		return "", err
	}

	if currentOwner != petToTransfer.Owner {
		return "This pet is currently owned by another entity.", err
	}

	petToTransfer.Owner = newOwner //change the owner

	petJSONBytes, _ := json.Marshal(petToTransfer)
	err = stub.PutState(pedigreeRegistrationNumber, petJSONBytes) //rewrite the pet
	if err != nil {
		return "", err
	}

	return "", nil
}

func (t *PetTraceChaincode) getHistoryForRecord(stub shim.ChaincodeStubInterface, args []string) peer.Response {

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
