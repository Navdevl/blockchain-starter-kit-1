/*
 * Supply Chain -  Chaincode
 */

package main

import (
  "fmt"
  "github.com/hyperledger/fabric/core/chaincode/shim"
  sc "github.com/hyperledger/fabric/protos/peer"
  "github.com/supply-chain-cc/chaincode/organization"
  "github.com/supply-chain-cc/chaincode/product"
)

// SmartContract defines the chaincode
type SmartContract struct {
}

// Init function
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
  fmt.Println("Name: Supply Chain")
  fmt.Println("Author: Skcript")
  return shim.Success(nil)
}

// Invoke function
func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {
  function, args := APIstub.GetFunctionAndParameters()

  fmt.Println("Function Invoked: %s", function)
  fmt.Println("Args Involved: %s", args)

  if function == "createOrganization" {
    return s.createOrganization(APIstub, args)
  } else if function == "createProduct" {
    return s.createProduct(APIstub, args)
  }

  return shim.Error("Invalid Action.")
}

func (s *SmartContract) createOrganization(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
  return organization.CreateOrganization(APIstub, args)
}

func (s *SmartContract) createProduct(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
  return product.CreateProduct(APIstub, args)
}

func main() {

  // Create a new Smart Contract
  err := shim.Start(new(SmartContract))
  if err != nil {
    fmt.Printf("Error creating new Smart Contract: %s", err)
  }

}
