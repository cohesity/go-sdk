# IdMappingParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SidMappedToUnixRootUser** | **NullableString** | Specifies the sid of an Active Directory domain user mapping to unix root user. | 
**UserIdMappingParams** | [**NullableIdMappingParamsUserIdMappingParams**](IdMappingParamsUserIdMappingParams.md) |  | 

## Methods

### NewIdMappingParams

`func NewIdMappingParams(sidMappedToUnixRootUser NullableString, userIdMappingParams NullableIdMappingParamsUserIdMappingParams, ) *IdMappingParams`

NewIdMappingParams instantiates a new IdMappingParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdMappingParamsWithDefaults

`func NewIdMappingParamsWithDefaults() *IdMappingParams`

NewIdMappingParamsWithDefaults instantiates a new IdMappingParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSidMappedToUnixRootUser

`func (o *IdMappingParams) GetSidMappedToUnixRootUser() string`

GetSidMappedToUnixRootUser returns the SidMappedToUnixRootUser field if non-nil, zero value otherwise.

### GetSidMappedToUnixRootUserOk

`func (o *IdMappingParams) GetSidMappedToUnixRootUserOk() (*string, bool)`

GetSidMappedToUnixRootUserOk returns a tuple with the SidMappedToUnixRootUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSidMappedToUnixRootUser

`func (o *IdMappingParams) SetSidMappedToUnixRootUser(v string)`

SetSidMappedToUnixRootUser sets SidMappedToUnixRootUser field to given value.


### SetSidMappedToUnixRootUserNil

`func (o *IdMappingParams) SetSidMappedToUnixRootUserNil(b bool)`

 SetSidMappedToUnixRootUserNil sets the value for SidMappedToUnixRootUser to be an explicit nil

### UnsetSidMappedToUnixRootUser
`func (o *IdMappingParams) UnsetSidMappedToUnixRootUser()`

UnsetSidMappedToUnixRootUser ensures that no value is present for SidMappedToUnixRootUser, not even an explicit nil
### GetUserIdMappingParams

`func (o *IdMappingParams) GetUserIdMappingParams() IdMappingParamsUserIdMappingParams`

GetUserIdMappingParams returns the UserIdMappingParams field if non-nil, zero value otherwise.

### GetUserIdMappingParamsOk

`func (o *IdMappingParams) GetUserIdMappingParamsOk() (*IdMappingParamsUserIdMappingParams, bool)`

GetUserIdMappingParamsOk returns a tuple with the UserIdMappingParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdMappingParams

`func (o *IdMappingParams) SetUserIdMappingParams(v IdMappingParamsUserIdMappingParams)`

SetUserIdMappingParams sets UserIdMappingParams field to given value.


### SetUserIdMappingParamsNil

`func (o *IdMappingParams) SetUserIdMappingParamsNil(b bool)`

 SetUserIdMappingParamsNil sets the value for UserIdMappingParams to be an explicit nil

### UnsetUserIdMappingParams
`func (o *IdMappingParams) UnsetUserIdMappingParams()`

UnsetUserIdMappingParams ensures that no value is present for UserIdMappingParams, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


