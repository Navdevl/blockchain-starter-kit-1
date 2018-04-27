package types

// Organization definition
type Organization struct {
  ID                string  `json:"id"`
  MspID             string  `json:"msp_id"`
  Class             string  `json:"class"`
  Name              string  `json:"name"`
  ContactEmail      string  `json:"contact_email"`
  ContactNumber     string  `json:"contact_number"`
  PictureURL        string  `json:"picture_url"`
  DomainName        string  `json:"domain_name"`
  OverallRatings    float64 `json:"overall_ratings"`
  TotalRatingsCount int     `json:"total_ratings_count"`
  OrganizationType  string  `json:"organization_type"`
  ReferredBy        string  `json:"referred_by"`
}

// Referral definition
type Referral struct {
  ID           string `json:"id"`
  Class        string `json:"class"`
  Code         string `json:"code"`
  Name         string `json:"name"`
  ContactName  string `json:"contact_name"`
  ContactEmail string `json:"contact_email"`
  ReferredBy   string `json:"referred_by"`
}

// User definition
type User struct {
  ID           string `json:"id"`
  Class        string `json:"class"`
  Name         string `json:"name"`
  Email        string `json:"email"`
  Organization string `json:"organization"`
  UserRole     string `json:"user_role"`
}

// IdentityRecord definition
type IdentityRecord struct {
  ID             string `json:"id"`
  Class          string `json:"class"`
  EnrollmentID   string `json:"enrollment_id"`
  UserID         string `json:"user_id"`
  Status         bool   `json:"status"`
  EnrollmentTime string `json:"enrollment_time"`
}

// Product definition
type Product struct {
  ID             string   `json:"id"`
  Class          string   `json:"class"`
  Name           string   `json:"name"`
  HeatID         string   `json:"heat_id"`
  Quantity       float64  `json:"quantity"`
  Description    string   `json:"description"`
  Owner          string   `json:"owner"`
  Categories     []string `json:"categories"`
  TransactedFrom []string `json:"transacted_from"`
  Unit           string   `json:"unit"`
  CreatedAt      string   `json:"created_at"`
  ModifiedAt     string   `json:"modified_at"`
  Status         string   `json:"status"`
}

// ProductSearchResult definition
type ProductSearchResult struct {
  ID   string  `json:"id"`
  Data Product `json:"data"`
}

// UserSearchResult definition
type UserSearchResult struct {
  ID   string `json:"id"`
  Data User   `json:"data"`
}

// IdentityRecordSearchResult definition
type IdentityRecordSearchResult struct {
  ID   string         `json:"id"`
  Data IdentityRecord `json:"data"`
}

// OrganizationSearchResult definition
type OrganizationSearchResult struct {
  ID   string       `json:"id"`
  Data Organization `json:"data"`
}

// TransactionRecordSearchResult definition
type TransactionRecordSearchResult struct {
  ID   string            `json:"id"`
  Data TransactionRecord `json:"data"`
}

// OrganizationTransactedWithSearchResult definition
type OrganizationTransactedWithSearchResult struct {
  ID   string                     `json:"id"`
  Data OrganizationTransactedWith `json:"data"`
}

// TransactionRecord definition
type TransactionRecord struct {
  ID                 string  `json:"id"`
  Class              string  `json:"class"`
  TransactionID      string  `json:"transactionId"`
  TransactionType    string  `json:"transactionType"`
  Product            string  `json:"product"`
  Quantity           float64 `json:"quantity"`
  User               string  `json:"user"`
  Organization       string  `json:"organization"`
  SecondOrganization string  `json:"second_organization"`
  Returnable         bool    `json:"returnable"`
  CreatedAt          string  `json:"created_at"`
}

// ProductWithOrganization definition
type ProductWithOrganization struct {
  ID             string       `json:"id"`
  Class          string       `json:"class"`
  Name           string       `json:"name"`
  HeatID         string       `json:"heat_id"`
  CertificateURL string       `json:"certificate_url"`
  Quantity       float64      `json:"quantity"`
  Description    string       `json:"description"`
  Owner          Organization `json:"owner"`
  TransactedFrom string       `json:"transacted_from"`
  CreatedAt      string       `json:"created_at"`
}

// UserWithOrganization definition
type UserWithOrganization struct {
  ID           string       `json:"id"`
  Class        string       `json:"class"`
  Name         string       `json:"name"`
  Email        string       `json:"email"`
  Organization Organization `json:"organization"`
  UserRole     string       `json:"user_role"`
}

// OrganizationTransactedWith definition
type OrganizationTransactedWith struct {
  ID                 string  `json:"id"`
  Class              string  `json:"class"`
  Count              int     `json:"count"`
  Rating             float64 `json:"rating"`
  IsRating           bool    `json:"is_rating"`
  LastTransactionAt  string  `json:"last_transaction_at"`
  Organization       string  `json:"organization"`
  SecondOrganization string  `json:"second_organization"`
}

// OrganizationTransactedWithOrganization definition
type OrganizationTransactedWithOrganization struct {
  ID                 string       `json:"id"`
  Class              string       `json:"class"`
  Count              int          `json:"count"`
  Rating             float64      `json:"rating"`
  IsRating           bool         `json:"is_rating"`
  LastTransactionAt  string       `json:"last_transaction_at"`
  Organization       string       `json:"organization"`
  SecondOrganization Organization `json:"second_organization"`
}

// TransactionRecordWithProductOrg definition
type TransactionRecordWithProductOrg struct {
  ID                 string       `json:"id"`
  Class              string       `json:"class"`
  TransactionID      string       `json:"transactionId"`
  TransactionType    string       `json:"transactionType"`
  Product            Product      `json:"product"`
  Quantity           float64      `json:"quantity"`
  User               string       `json:"user"`
  Organization       Organization `json:"organization"`
  SecondOrganization Organization `json:"second_organization"`
  Returnable         bool         `json:"returnable"`
  CreatedAt          string       `json:"created_at"`
}
