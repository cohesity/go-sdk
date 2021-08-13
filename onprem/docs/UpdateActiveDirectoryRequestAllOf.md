# UpdateActiveDirectoryRequestAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveDirectoryAdminParams** | Pointer to [**NullableActiveDirectoryAdminParams**](ActiveDirectoryAdminParams.md) | Specifies the params of a user with administrative privilege of this Active Directory. This field is mandatory if machine accounts are updated. | [optional] 
**OverwriteMachineAccounts** | Pointer to **NullableBool** | Specifies if specified machine accounts should overwrite existing machine accounts. | [optional] 
**IdMappingParams** | Pointer to [**NullableIdMappingParams**](IdMappingParams.md) | Specifies the params of the user id mapping info of an Active Directory. | [optional] 
**TransitiveAdTrustLevelLimit** | Pointer to **NullableInt32** | Specifies level of transitive Active Directory trust domains to be used. | [optional] 

## Methods

### NewUpdateActiveDirectoryRequestAllOf

`func NewUpdateActiveDirectoryRequestAllOf() *UpdateActiveDirectoryRequestAllOf`

NewUpdateActiveDirectoryRequestAllOf instantiates a new UpdateActiveDirectoryRequestAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateActiveDirectoryRequestAllOfWithDefaults

`func NewUpdateActiveDirectoryRequestAllOfWithDefaults() *UpdateActiveDirectoryRequestAllOf`

NewUpdateActiveDirectoryRequestAllOfWithDefaults instantiates a new UpdateActiveDirectoryRequestAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveDirectoryAdminParams

`func (o *UpdateActiveDirectoryRequestAllOf) GetActiveDirectoryAdminParams() ActiveDirectoryAdminParams`

GetActiveDirectoryAdminParams returns the ActiveDirectoryAdminParams field if non-nil, zero value otherwise.

### GetActiveDirectoryAdminParamsOk

`func (o *UpdateActiveDirectoryRequestAllOf) GetActiveDirectoryAdminParamsOk() (*ActiveDirectoryAdminParams, bool)`

GetActiveDirectoryAdminParamsOk returns a tuple with the ActiveDirectoryAdminParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveDirectoryAdminParams

`func (o *UpdateActiveDirectoryRequestAllOf) SetActiveDirectoryAdminParams(v ActiveDirectoryAdminParams)`

SetActiveDirectoryAdminParams sets ActiveDirectoryAdminParams field to given value.

### HasActiveDirectoryAdminParams

`func (o *UpdateActiveDirectoryRequestAllOf) HasActiveDirectoryAdminParams() bool`

HasActiveDirectoryAdminParams returns a boolean if a field has been set.

### SetActiveDirectoryAdminParamsNil

`func (o *UpdateActiveDirectoryRequestAllOf) SetActiveDirectoryAdminParamsNil(b bool)`

 SetActiveDirectoryAdminParamsNil sets the value for ActiveDirectoryAdminParams to be an explicit nil

### UnsetActiveDirectoryAdminParams
`func (o *UpdateActiveDirectoryRequestAllOf) UnsetActiveDirectoryAdminParams()`

UnsetActiveDirectoryAdminParams ensures that no value is present for ActiveDirectoryAdminParams, not even an explicit nil
### GetOverwriteMachineAccounts

`func (o *UpdateActiveDirectoryRequestAllOf) GetOverwriteMachineAccounts() bool`

GetOverwriteMachineAccounts returns the OverwriteMachineAccounts field if non-nil, zero value otherwise.

### GetOverwriteMachineAccountsOk

`func (o *UpdateActiveDirectoryRequestAllOf) GetOverwriteMachineAccountsOk() (*bool, bool)`

GetOverwriteMachineAccountsOk returns a tuple with the OverwriteMachineAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwriteMachineAccounts

`func (o *UpdateActiveDirectoryRequestAllOf) SetOverwriteMachineAccounts(v bool)`

SetOverwriteMachineAccounts sets OverwriteMachineAccounts field to given value.

### HasOverwriteMachineAccounts

`func (o *UpdateActiveDirectoryRequestAllOf) HasOverwriteMachineAccounts() bool`

HasOverwriteMachineAccounts returns a boolean if a field has been set.

### SetOverwriteMachineAccountsNil

`func (o *UpdateActiveDirectoryRequestAllOf) SetOverwriteMachineAccountsNil(b bool)`

 SetOverwriteMachineAccountsNil sets the value for OverwriteMachineAccounts to be an explicit nil

### UnsetOverwriteMachineAccounts
`func (o *UpdateActiveDirectoryRequestAllOf) UnsetOverwriteMachineAccounts()`

UnsetOverwriteMachineAccounts ensures that no value is present for OverwriteMachineAccounts, not even an explicit nil
### GetIdMappingParams

`func (o *UpdateActiveDirectoryRequestAllOf) GetIdMappingParams() IdMappingParams`

GetIdMappingParams returns the IdMappingParams field if non-nil, zero value otherwise.

### GetIdMappingParamsOk

`func (o *UpdateActiveDirectoryRequestAllOf) GetIdMappingParamsOk() (*IdMappingParams, bool)`

GetIdMappingParamsOk returns a tuple with the IdMappingParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdMappingParams

`func (o *UpdateActiveDirectoryRequestAllOf) SetIdMappingParams(v IdMappingParams)`

SetIdMappingParams sets IdMappingParams field to given value.

### HasIdMappingParams

`func (o *UpdateActiveDirectoryRequestAllOf) HasIdMappingParams() bool`

HasIdMappingParams returns a boolean if a field has been set.

### SetIdMappingParamsNil

`func (o *UpdateActiveDirectoryRequestAllOf) SetIdMappingParamsNil(b bool)`

 SetIdMappingParamsNil sets the value for IdMappingParams to be an explicit nil

### UnsetIdMappingParams
`func (o *UpdateActiveDirectoryRequestAllOf) UnsetIdMappingParams()`

UnsetIdMappingParams ensures that no value is present for IdMappingParams, not even an explicit nil
### GetTransitiveAdTrustLevelLimit

`func (o *UpdateActiveDirectoryRequestAllOf) GetTransitiveAdTrustLevelLimit() int32`

GetTransitiveAdTrustLevelLimit returns the TransitiveAdTrustLevelLimit field if non-nil, zero value otherwise.

### GetTransitiveAdTrustLevelLimitOk

`func (o *UpdateActiveDirectoryRequestAllOf) GetTransitiveAdTrustLevelLimitOk() (*int32, bool)`

GetTransitiveAdTrustLevelLimitOk returns a tuple with the TransitiveAdTrustLevelLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitiveAdTrustLevelLimit

`func (o *UpdateActiveDirectoryRequestAllOf) SetTransitiveAdTrustLevelLimit(v int32)`

SetTransitiveAdTrustLevelLimit sets TransitiveAdTrustLevelLimit field to given value.

### HasTransitiveAdTrustLevelLimit

`func (o *UpdateActiveDirectoryRequestAllOf) HasTransitiveAdTrustLevelLimit() bool`

HasTransitiveAdTrustLevelLimit returns a boolean if a field has been set.

### SetTransitiveAdTrustLevelLimitNil

`func (o *UpdateActiveDirectoryRequestAllOf) SetTransitiveAdTrustLevelLimitNil(b bool)`

 SetTransitiveAdTrustLevelLimitNil sets the value for TransitiveAdTrustLevelLimit to be an explicit nil

### UnsetTransitiveAdTrustLevelLimit
`func (o *UpdateActiveDirectoryRequestAllOf) UnsetTransitiveAdTrustLevelLimit()`

UnsetTransitiveAdTrustLevelLimit ensures that no value is present for TransitiveAdTrustLevelLimit, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


