package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// Asset chaincode implementation
type Asset struct {
}

// Init initializes chaincode
func (t *Asset) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

// Invoke entry point for Invocations
func (t *Asset) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fn, args := stub.GetFunctionAndParameters()
	fmt.Println("Invoke function: " + fn)

	switch fn {
	case "set":
		return set(stub, args)
	case "get":
		return get(stub, args)
	default:
		return shim.Error("Received unknown function invocation")
	}
}

func set(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 3 {
		return shim.Error("Incorrect arguments, expecting a key and a value")
	}

	if err := stub.PutState(args[0], []byte(args[1])); err != nil {
		return shim.Error("Failed to set asset")
	}

	txID := []byte(stub.GetTxID())

	return shim.Success(txID)
}

func get(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect arguments, expecting a key")
	}

	value, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error("Failed to get asset")
	}

	return shim.Success(value)
}

func main() {
	err := shim.Start(new(Asset))
	if err != nil {
		fmt.Printf("Error starting Asset chaincode: %s", err)
	}
}
