# UpdateMachineAccountsParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MachineAccounts** | Pointer to **[]string** | Array of Machine Accounts.  Specifies an array of computer names used to identify the Cohesity Cluster on the domain. | [optional] 
**OverwriteExistingAccounts** | Pointer to **NullableBool** | Specifies whether the specified machine accounts should overwrite the existing machine accounts in this domain. | [optional] 
**Password** | Pointer to **NullableString** | Specifies the password for the specified userName. | [optional] 
**UserName** | Pointer to **NullableString** | Specifies a userName that has administrative privileges in the domain. | [optional] 

## Methods

### NewUpdateMachineAccountsParams

`func NewUpdateMachineAccountsParams() *UpdateMachineAccountsParams`

NewUpdateMachineAccountsParams instantiates a new UpdateMachineAccountsParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMachineAccountsParamsWithDefaults

`func NewUpdateMachineAccountsParamsWithDefaults() *UpdateMachineAccountsParams`

NewUpdateMachineAccountsParamsWithDefaults instantiates a new UpdateMachineAccountsParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMachineAccounts

`func (o *UpdateMachineAccountsParams) GetMachineAccounts() []string`

GetMachineAccounts returns the MachineAccounts field if non-nil, zero value otherwise.

### GetMachineAccountsOk

`func (o *UpdateMachineAccountsParams) GetMachineAccountsOk() (*[]string, bool)`

GetMachineAccountsOk returns a tuple with the MachineAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineAccounts

`func (o *UpdateMachineAccountsParams) SetMachineAccounts(v []string)`

SetMachineAccounts sets MachineAccounts field to given value.

### HasMachineAccounts

`func (o *UpdateMachineAccountsParams) HasMachineAccounts() bool`

HasMachineAccounts returns a boolean if a field has been set.

### SetMachineAccountsNil

`func (o *UpdateMachineAccountsParams) SetMachineAccountsNil(b bool)`

 SetMachineAccountsNil sets the value for MachineAccounts to be an explicit nil

### UnsetMachineAccounts
`func (o *UpdateMachineAccountsParams) UnsetMachineAccounts()`

UnsetMachineAccounts ensures that no value is present for MachineAccounts, not even an explicit nil
### GetOverwriteExistingAccounts

`func (o *UpdateMachineAccountsParams) GetOverwriteExistingAccounts() bool`

GetOverwriteExistingAccounts returns the OverwriteExistingAccounts field if non-nil, zero value otherwise.

### GetOverwriteExistingAccountsOk

`func (o *UpdateMachineAccountsParams) GetOverwriteExistingAccountsOk() (*bool, bool)`

GetOverwriteExistingAccountsOk returns a tuple with the OverwriteExistingAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwriteExistingAccounts

`func (o *UpdateMachineAccountsParams) SetOverwriteExistingAccounts(v bool)`

SetOverwriteExistingAccounts sets OverwriteExistingAccounts field to given value.

### HasOverwriteExistingAccounts

`func (o *UpdateMachineAccountsParams) HasOverwriteExistingAccounts() bool`

HasOverwriteExistingAccounts returns a boolean if a field has been set.

### SetOverwriteExistingAccountsNil

`func (o *UpdateMachineAccountsParams) SetOverwriteExistingAccountsNil(b bool)`

 SetOverwriteExistingAccountsNil sets the value for OverwriteExistingAccounts to be an explicit nil

### UnsetOverwriteExistingAccounts
`func (o *UpdateMachineAccountsParams) UnsetOverwriteExistingAccounts()`

UnsetOverwriteExistingAccounts ensures that no value is present for OverwriteExistingAccounts, not even an explicit nil
### GetPassword

`func (o *UpdateMachineAccountsParams) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UpdateMachineAccountsParams) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UpdateMachineAccountsParams) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UpdateMachineAccountsParams) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *UpdateMachineAccountsParams) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *UpdateMachineAccountsParams) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetUserName

`func (o *UpdateMachineAccountsParams) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *UpdateMachineAccountsParams) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *UpdateMachineAccountsParams) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *UpdateMachineAccountsParams) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### SetUserNameNil

`func (o *UpdateMachineAccountsParams) SetUserNameNil(b bool)`

 SetUserNameNil sets the value for UserName to be an explicit nil

### UnsetUserName
`func (o *UpdateMachineAccountsParams) UnsetUserName()`

UnsetUserName ensures that no value is present for UserName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


