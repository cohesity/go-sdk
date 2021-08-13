# HeliosAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyName** | Pointer to **NullableString** | Specifies the name of the company to which this account belongs. | [optional] 
**SalesforceAccountId** | Pointer to **NullableString** | Specifies the Salesforce account ID of this account. | [optional] 

## Methods

### NewHeliosAccount

`func NewHeliosAccount() *HeliosAccount`

NewHeliosAccount instantiates a new HeliosAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeliosAccountWithDefaults

`func NewHeliosAccountWithDefaults() *HeliosAccount`

NewHeliosAccountWithDefaults instantiates a new HeliosAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyName

`func (o *HeliosAccount) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *HeliosAccount) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *HeliosAccount) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *HeliosAccount) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### SetCompanyNameNil

`func (o *HeliosAccount) SetCompanyNameNil(b bool)`

 SetCompanyNameNil sets the value for CompanyName to be an explicit nil

### UnsetCompanyName
`func (o *HeliosAccount) UnsetCompanyName()`

UnsetCompanyName ensures that no value is present for CompanyName, not even an explicit nil
### GetSalesforceAccountId

`func (o *HeliosAccount) GetSalesforceAccountId() string`

GetSalesforceAccountId returns the SalesforceAccountId field if non-nil, zero value otherwise.

### GetSalesforceAccountIdOk

`func (o *HeliosAccount) GetSalesforceAccountIdOk() (*string, bool)`

GetSalesforceAccountIdOk returns a tuple with the SalesforceAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesforceAccountId

`func (o *HeliosAccount) SetSalesforceAccountId(v string)`

SetSalesforceAccountId sets SalesforceAccountId field to given value.

### HasSalesforceAccountId

`func (o *HeliosAccount) HasSalesforceAccountId() bool`

HasSalesforceAccountId returns a boolean if a field has been set.

### SetSalesforceAccountIdNil

`func (o *HeliosAccount) SetSalesforceAccountIdNil(b bool)`

 SetSalesforceAccountIdNil sets the value for SalesforceAccountId to be an explicit nil

### UnsetSalesforceAccountId
`func (o *HeliosAccount) UnsetSalesforceAccountId()`

UnsetSalesforceAccountId ensures that no value is present for SalesforceAccountId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


