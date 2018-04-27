/*
 * Organization - Model
 */

package organization

import (
  "encoding/json"
  "errors"
  "fmt"
  "github.com/hyperledger/fabric/core/chaincode/shim"
  sc "github.com/hyperledger/fabric/protos/peer"
  "github.com/supply-chain-cc/chaincode/types"
  "github.com/supply-chain-cc/chaincode/utils"
  "strconv"
)

// CreateOrganization function creates organization.
func CreateOrganization(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
  /*
     Arguments:
       * ID
       * Name
       * ContactEmail
       * Contact Number
       * PictureURL
       * DomainName
       * OrganizationType
  */

  // Allow only 7 or 8 arguments. No less.  No more.
  if len(args) != 7 {
    return shim.Error("ERROR: #240YCD")
  }

  ID := "org-" + args[0]
  _, err := getOrganizationByDomainName(APIstub, args[4])
  if err != nil {
    var organization = types.Organization{Class: "Organization", ID: ID, Name: args[1], ContactEmail: args[2], ContactNumber: args[3], PictureURL: args[4], DomainName: args[5], OrganizationType: getOrganizationType(args[6])}
    OrganizationBytes, err := json.Marshal(organization)
    if err != nil {
      return shim.Error(err.Error())
    }
    APIstub.PutState(ID, OrganizationBytes)
    return shim.Success(OrganizationBytes)
  }
  return shim.Error("ERROR: #183GDC")
}

func getOrganizationType(ID string) string {
  RoleID, _ := strconv.Atoi(ID)
  // TODO HANDLE ERROR
  types := []string{"Manufacturer", "Supplier", "Consumer"}
  return types[RoleID]
}

func getOrganizationByDomainName(APIstub shim.ChaincodeStubInterface, domainName string) (types.Organization, error) {
  organization := types.Organization{}
  searchResultsBytes, err := utils.SearchQuery(APIstub, "{\"selector\": {\"domain_name\": \""+domainName+"\" }}")
  if err != nil {
    return organization, err
  }
  searchResults := []types.OrganizationSearchResult{}
  json.Unmarshal([]byte(searchResultsBytes), &searchResults)
  fmt.Println("SEARCH LEN: ", len(searchResults))
  if len(searchResults) > 0 {
    organization = searchResults[0].Data
    return organization, nil
  }
  return organization, errors.New("does not belongs to any organization")
}

// GetOrganizationByID function gets organization by ID
func GetOrganizationByID(APIstub shim.ChaincodeStubInterface, id string) (types.Organization, error) {
  organization := types.Organization{}
  organizationBytes, err := APIstub.GetState(id)
  json.Unmarshal(organizationBytes, &organization)
  return organization, err
}
