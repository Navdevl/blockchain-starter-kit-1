/*
 * Product - Model
 */

package product

import (
  "encoding/json"
  "github.com/hyperledger/fabric/core/chaincode/shim"
  sc "github.com/hyperledger/fabric/protos/peer"
  "github.com/supply-chain-cc/chaincode/types"
  "strconv"
  "time"
)

// CreateProduct function creates/transfers products
func CreateProduct(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
  /*
     Arguments:
       * ID
       * Name
       * HeatId
       * Quantity
       * Description
       * Organization
       * Unit
       * CreatedAt
  */

  ID := "product-" + args[0]
  quantity, err := strconv.ParseFloat(args[4], 64)
  if err != nil {
    return shim.Error(err.Error())
  }
  createdAt := time.Now().Format("2006-01-02 15:04:05 +0000 UTC")
  if quantity < 0 {
    return shim.Error("Quantity should be positive")
  }

  var product = types.Product{Class: "Product", ID: ID, Name: args[1], HeatID: args[2], Quantity: quantity, Description: args[4], Owner: args[5], Unit: args[6], CreatedAt: createdAt, ModifiedAt: createdAt}

  CreatedProductBytes, _ := json.Marshal(product)
  APIstub.PutState(ID, CreatedProductBytes)

  return shim.Success(CreatedProductBytes)
}
