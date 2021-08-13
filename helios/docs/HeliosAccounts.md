# HeliosAccounts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | Pointer to [**[]HeliosAccount**](HeliosAccount.md) | Specifies a list of Helios Accounts. | [optional] 

## Methods

### NewHeliosAccounts

`func NewHeliosAccounts() *HeliosAccounts`

NewHeliosAccounts instantiates a new HeliosAccounts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeliosAccountsWithDefaults

`func NewHeliosAccountsWithDefaults() *HeliosAccounts`

NewHeliosAccountsWithDefaults instantiates a new HeliosAccounts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *HeliosAccounts) GetAccounts() []HeliosAccount`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *HeliosAccounts) GetAccountsOk() (*[]HeliosAccount, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *HeliosAccounts) SetAccounts(v []HeliosAccount)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *HeliosAccounts) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### SetAccountsNil

`func (o *HeliosAccounts) SetAccountsNil(b bool)`

 SetAccountsNil sets the value for Accounts to be an explicit nil

### UnsetAccounts
`func (o *HeliosAccounts) UnsetAccounts()`

UnsetAccounts ensures that no value is present for Accounts, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


