package utils

import (
  "bytes"
  "github.com/hyperledger/fabric/core/chaincode/shim"
)

// SearchQuery function searches using the query
func SearchQuery(APIstub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {
  resultsIterator, err := APIstub.GetQueryResult(queryString)
  defer resultsIterator.Close()
  if err != nil {
    return nil, err
  }
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
    buffer.WriteString("{\"id\":")
    buffer.WriteString("\"")
    buffer.WriteString(queryResponse.Key)
    buffer.WriteString("\"")
    buffer.WriteString(", \"data\":")
    // Record is a JSON object, so we write as-is
    buffer.WriteString(string(queryResponse.Value))
    buffer.WriteString("}")
    bArrayMemberAlreadyWritten = true
  }
  buffer.WriteString("]")
  return buffer.Bytes(), nil
}

// PushToTransactedFromX Function
func PushToTransactedFromX(ReceiverTransactedFrom []string, SenderTransactedFrom []string, SenderID string) []string {
  if len(SenderTransactedFrom) < 1 {
    transactionString := SenderID
    if !contains(ReceiverTransactedFrom, transactionString) {
      ReceiverTransactedFrom = append(ReceiverTransactedFrom, transactionString)
    }
  } else {
    for _, transaction := range SenderTransactedFrom {
      transactionString := transaction + ", " + SenderID
      if !contains(ReceiverTransactedFrom, transactionString) {
        ReceiverTransactedFrom = append(ReceiverTransactedFrom, transactionString)
      } else {
        if !contains(ReceiverTransactedFrom, SenderID) {
          ReceiverTransactedFrom = append(ReceiverTransactedFrom, SenderID)
        }
      }
    }
  }
  return ReceiverTransactedFrom
}

func contains(arr []string, str string) bool {
  for _, a := range arr {
    if a == str {
      return true
    }
  }
  return false
}
