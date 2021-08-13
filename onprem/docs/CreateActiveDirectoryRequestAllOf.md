# CreateActiveDirectoryRequestAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainName** | **NullableString** | Specifies the domain name of the Active Directory. | 
**ActiveDirectoryAdminParams** | [**NullableActiveDirectoryAdminParams**](ActiveDirectoryAdminParams.md) | Specifies the params of a user with administrative privilege of this Active Directory. | 
**OverwriteMachineAccounts** | Pointer to **NullableBool** | Specifies if specified machine accounts should overwrite existing machine accounts. | [optional] 

## Methods

### NewCreateActiveDirectoryRequestAllOf

`func NewCreateActiveDirectoryRequestAllOf(domainName NullableString, activeDirectoryAdminParams NullableActiveDirectoryAdminParams, ) *CreateActiveDirectoryRequestAllOf`

NewCreateActiveDirectoryRequestAllOf instantiates a new CreateActiveDirectoryRequestAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateActiveDirectoryRequestAllOfWithDefaults

`func NewCreateActiveDirectoryRequestAllOfWithDefaults() *CreateActiveDirectoryRequestAllOf`

NewCreateActiveDirectoryRequestAllOfWithDefaults instantiates a new CreateActiveDirectoryRequestAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainName

`func (o *CreateActiveDirectoryRequestAllOf) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *CreateActiveDirectoryRequestAllOf) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *CreateActiveDirectoryRequestAllOf) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.


### SetDomainNameNil

`func (o *CreateActiveDirectoryRequestAllOf) SetDomainNameNil(b bool)`

 SetDomainNameNil sets the value for DomainName to be an explicit nil

### UnsetDomainName
`func (o *CreateActiveDirectoryRequestAllOf) UnsetDomainName()`

UnsetDomainName ensures that no value is present for DomainName, not even an explicit nil
### GetActiveDirectoryAdminParams

`func (o *CreateActiveDirectoryRequestAllOf) GetActiveDirectoryAdminParams() ActiveDirectoryAdminParams`

GetActiveDirectoryAdminParams returns the ActiveDirectoryAdminParams field if non-nil, zero value otherwise.

### GetActiveDirectoryAdminParamsOk

`func (o *CreateActiveDirectoryRequestAllOf) GetActiveDirectoryAdminParamsOk() (*ActiveDirectoryAdminParams, bool)`

GetActiveDirectoryAdminParamsOk returns a tuple with the ActiveDirectoryAdminParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveDirectoryAdminParams

`func (o *CreateActiveDirectoryRequestAllOf) SetActiveDirectoryAdminParams(v ActiveDirectoryAdminParams)`

SetActiveDirectoryAdminParams sets ActiveDirectoryAdminParams field to given value.


### SetActiveDirectoryAdminParamsNil

`func (o *CreateActiveDirectoryRequestAllOf) SetActiveDirectoryAdminParamsNil(b bool)`

 SetActiveDirectoryAdminParamsNil sets the value for ActiveDirectoryAdminParams to be an explicit nil

### UnsetActiveDirectoryAdminParams
`func (o *CreateActiveDirectoryRequestAllOf) UnsetActiveDirectoryAdminParams()`

UnsetActiveDirectoryAdminParams ensures that no value is present for ActiveDirectoryAdminParams, not even an explicit nil
### GetOverwriteMachineAccounts

`func (o *CreateActiveDirectoryRequestAllOf) GetOverwriteMachineAccounts() bool`

GetOverwriteMachineAccounts returns the OverwriteMachineAccounts field if non-nil, zero value otherwise.

### GetOverwriteMachineAccountsOk

`func (o *CreateActiveDirectoryRequestAllOf) GetOverwriteMachineAccountsOk() (*bool, bool)`

GetOverwriteMachineAccountsOk returns a tuple with the OverwriteMachineAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwriteMachineAccounts

`func (o *CreateActiveDirectoryRequestAllOf) SetOverwriteMachineAccounts(v bool)`

SetOverwriteMachineAccounts sets OverwriteMachineAccounts field to given value.

### HasOverwriteMachineAccounts

`func (o *CreateActiveDirectoryRequestAllOf) HasOverwriteMachineAccounts() bool`

HasOverwriteMachineAccounts returns a boolean if a field has been set.

### SetOverwriteMachineAccountsNil

`func (o *CreateActiveDirectoryRequestAllOf) SetOverwriteMachineAccountsNil(b bool)`

 SetOverwriteMachineAccountsNil sets the value for OverwriteMachineAccounts to be an explicit nil

### UnsetOverwriteMachineAccounts
`func (o *CreateActiveDirectoryRequestAllOf) UnsetOverwriteMachineAccounts()`

UnsetOverwriteMachineAccounts ensures that no value is present for OverwriteMachineAccounts, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


