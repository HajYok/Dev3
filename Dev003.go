package main

import (
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

// Init resets all the things
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}
	err := stub.PutState("TRY!", []byte(args[0]))
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// Invoke isur entry point to invoke a chaincode function
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "init" {
		return t.Init(stub, "init", args)
	} else if function == "write" {
		return t.write(stub, args)
	} else if function == "write2" {
		return t.write2(stub, args)
	} else if function == "write3" {
		return t.write3(stub, args)
	} else if function == "write4" {
		return t.write4(stub, args)
	}
	fmt.Println("invoke did not find func: " + function)

	return nil, errors.New("Received unknown function invocation: " + function)
}

// Query is our entry point for queries
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)

	// Handle different functions
	if function == "read" { //read a variable
		return t.read(stub, args)
	} else if function == "read2" {
		return t.read2(stub, args)
	} else if function == "read3" {
		return t.read3(stub, args)
	} else if function == "read4" {
		return t.read4(stub, args)
	}
	fmt.Println("query did not find func: " + function)

	return nil, errors.New("Received unknown function query: " + function)
}

// ============================================================================================================================
// Write - write variable into chaincode state
// ============================================================================================================================
func (t *SimpleChaincode) write(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var name, date, stim, etim string // Entities
	var err error
	fmt.Println("running write()")

	if len(args) != 4 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2. name of the variable and value to set")
	}

	name = args[0]										//rename for funsies
	date = args[1]
	stim = args[2]
	etim = args[3]
	str := `{"name": "` + name + `", "date": "` + date + `", "start_time": "` + stim + `", "end_time": "` + etim + `"}`
	err = stub.PutState(name, []byte(str))						//write the variable into the chaincode state
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// ============================================================================================================================
// Read - read a variable from chaincode state
// ============================================================================================================================
func (t *SimpleChaincode) read(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var name, jsonResp string
	var err error

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting name of the var to query")
	}

	name = args[0]
	valAsbytes, err := stub.GetState(name)					//get the var from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + name + "\"}"
		return nil, errors.New(jsonResp)
	}

	return valAsbytes, nil							//send it onward
}

// ============================================================================================================================
// Write2 - write variable into chaincode state
// ============================================================================================================================
func (t *SimpleChaincode) write2(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var key, name, subj, vote string // Entities
	var err error
	fmt.Println("running write2()")

	if len(args) != 3 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2. name of the variable and value to set")
	}

	key  = args[0] + args[1]										//rename for funsies
	name = args[0]
	subj = args[1]
	vote = args[2]
	str := `{"key": "` + key + `", "name": "` + name + `", "subj": "` + subj + `", "vote": "` + vote + `"}`
	err = stub.PutState(key, []byte(str))						//write the variable into the chaincode state
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// ============================================================================================================================
// Read2 - read a variable from chaincode state
// ============================================================================================================================
func (t *SimpleChaincode) read2(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var key, jsonResp string
	var err error

	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting name of the var to query")
	}

	key = args[0] + args[1]
	valAsbytes, err := stub.GetState(key)					//get the var from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
		return nil, errors.New(jsonResp)
	}

	return valAsbytes, nil							//send it onward
}

// ============================================================================================================================
// Write3 - write variable into chaincode state
// ============================================================================================================================
func (t *SimpleChaincode) write3(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var subject__number, subject__title, subject__content, subject__type string // Entities
	var err error
	fmt.Println("running write3()")

	if len(args) != 4 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2. name of the variable and value to set")
	}

	subject__number = args[0]									//rename for funsies
	subject__title = args[1]
	subject__content = args[2]
	subject__type = args[3]
	str := `{"subject__number": "` + subject__number + `", "subject__title": "` + subject__title + `", "subject__content": "` + subject__content + `", "subject__type": "` + subject__type + `"}`
	err = stub.PutState(subject__number, []byte(str))						//write the variable into the chaincode state
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// ============================================================================================================================
// Read3 - read a variable from chaincode state
// ============================================================================================================================
func (t *SimpleChaincode) read3(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var subject__number, jsonResp string
	var err error

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting name of the var to query")
	}

	subject__number = args[0]
	valAsbytes, err := stub.GetState(subject__number)					//get the var from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + subject__number + "\"}"
		return nil, errors.New(jsonResp)
	}

	return valAsbytes, nil							//send it onward
}

// ============================================================================================================================
// Write4 - write variable into chaincode state
// ============================================================================================================================
func (t *SimpleChaincode) write4(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var key_shares_held, bill_number, user, number_of_shares_held string // Entities
	var err error
	fmt.Println("running write4()")

	if len(args) != 4 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2. name of the variable and value to set")
	}

	key_shares_held = args[0]										//rename for funsies
	bill_number = args[1]
	user = args[2]
	number_of_shares_held = args[3]
	str := `{"key_shares_held": "` + key_shares_held + `", "bill_number": "` + bill_number + `", "user": "` + user + `", "number_of_shares_held": "` + number_of_shares_held + `"}`
	err = stub.PutState(key_shares_held, []byte(str))						//write the variable into the chaincode state
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// ============================================================================================================================
// Read4 - read a variable from chaincode state
// ============================================================================================================================
func (t *SimpleChaincode) read4(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var key_shares_held, jsonResp string
	var err error

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting name of the var to query")
	}

	key_shares_held = args[0]
	valAsbytes, err := stub.GetState(key_shares_held)					//get the var from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + key_shares_held + "\"}"
		return nil, errors.New(jsonResp)
	}

	return valAsbytes, nil							//send it onward
}
