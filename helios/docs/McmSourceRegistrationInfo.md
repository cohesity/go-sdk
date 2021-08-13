# McmSourceRegistrationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionId** | Pointer to **NullableInt64** | Specifies the id of the connection from where this source is reachable. | [optional] 
**RegistrationTimeUsecs** | Pointer to **NullableInt64** | Specifies the registration time of the Source in Unix time epoch in microseconds. | [optional] 
**RegistrationError** | Pointer to **NullableString** | Specifies the error detail occured during the protection source registration. | [optional] 
**LastRefreshTimeUsecs** | Pointer to **NullableInt64** | Specifies the registration time of the Source in Unix time epoch in microseconds. | [optional] 
**RefreshError** | Pointer to **NullableString** | Specifies the error detail occured during the refersh of a protection source. | [optional] 

## Methods

### NewMcmSourceRegistrationInfo

`func NewMcmSourceRegistrationInfo() *McmSourceRegistrationInfo`

NewMcmSourceRegistrationInfo instantiates a new McmSourceRegistrationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcmSourceRegistrationInfoWithDefaults

`func NewMcmSourceRegistrationInfoWithDefaults() *McmSourceRegistrationInfo`

NewMcmSourceRegistrationInfoWithDefaults instantiates a new McmSourceRegistrationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionId

`func (o *McmSourceRegistrationInfo) GetConnectionId() int64`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *McmSourceRegistrationInfo) GetConnectionIdOk() (*int64, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *McmSourceRegistrationInfo) SetConnectionId(v int64)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *McmSourceRegistrationInfo) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### SetConnectionIdNil

`func (o *McmSourceRegistrationInfo) SetConnectionIdNil(b bool)`

 SetConnectionIdNil sets the value for ConnectionId to be an explicit nil

### UnsetConnectionId
`func (o *McmSourceRegistrationInfo) UnsetConnectionId()`

UnsetConnectionId ensures that no value is present for ConnectionId, not even an explicit nil
### GetRegistrationTimeUsecs

`func (o *McmSourceRegistrationInfo) GetRegistrationTimeUsecs() int64`

GetRegistrationTimeUsecs returns the RegistrationTimeUsecs field if non-nil, zero value otherwise.

### GetRegistrationTimeUsecsOk

`func (o *McmSourceRegistrationInfo) GetRegistrationTimeUsecsOk() (*int64, bool)`

GetRegistrationTimeUsecsOk returns a tuple with the RegistrationTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationTimeUsecs

`func (o *McmSourceRegistrationInfo) SetRegistrationTimeUsecs(v int64)`

SetRegistrationTimeUsecs sets RegistrationTimeUsecs field to given value.

### HasRegistrationTimeUsecs

`func (o *McmSourceRegistrationInfo) HasRegistrationTimeUsecs() bool`

HasRegistrationTimeUsecs returns a boolean if a field has been set.

### SetRegistrationTimeUsecsNil

`func (o *McmSourceRegistrationInfo) SetRegistrationTimeUsecsNil(b bool)`

 SetRegistrationTimeUsecsNil sets the value for RegistrationTimeUsecs to be an explicit nil

### UnsetRegistrationTimeUsecs
`func (o *McmSourceRegistrationInfo) UnsetRegistrationTimeUsecs()`

UnsetRegistrationTimeUsecs ensures that no value is present for RegistrationTimeUsecs, not even an explicit nil
### GetRegistrationError

`func (o *McmSourceRegistrationInfo) GetRegistrationError() string`

GetRegistrationError returns the RegistrationError field if non-nil, zero value otherwise.

### GetRegistrationErrorOk

`func (o *McmSourceRegistrationInfo) GetRegistrationErrorOk() (*string, bool)`

GetRegistrationErrorOk returns a tuple with the RegistrationError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationError

`func (o *McmSourceRegistrationInfo) SetRegistrationError(v string)`

SetRegistrationError sets RegistrationError field to given value.

### HasRegistrationError

`func (o *McmSourceRegistrationInfo) HasRegistrationError() bool`

HasRegistrationError returns a boolean if a field has been set.

### SetRegistrationErrorNil

`func (o *McmSourceRegistrationInfo) SetRegistrationErrorNil(b bool)`

 SetRegistrationErrorNil sets the value for RegistrationError to be an explicit nil

### UnsetRegistrationError
`func (o *McmSourceRegistrationInfo) UnsetRegistrationError()`

UnsetRegistrationError ensures that no value is present for RegistrationError, not even an explicit nil
### GetLastRefreshTimeUsecs

`func (o *McmSourceRegistrationInfo) GetLastRefreshTimeUsecs() int64`

GetLastRefreshTimeUsecs returns the LastRefreshTimeUsecs field if non-nil, zero value otherwise.

### GetLastRefreshTimeUsecsOk

`func (o *McmSourceRegistrationInfo) GetLastRefreshTimeUsecsOk() (*int64, bool)`

GetLastRefreshTimeUsecsOk returns a tuple with the LastRefreshTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRefreshTimeUsecs

`func (o *McmSourceRegistrationInfo) SetLastRefreshTimeUsecs(v int64)`

SetLastRefreshTimeUsecs sets LastRefreshTimeUsecs field to given value.

### HasLastRefreshTimeUsecs

`func (o *McmSourceRegistrationInfo) HasLastRefreshTimeUsecs() bool`

HasLastRefreshTimeUsecs returns a boolean if a field has been set.

### SetLastRefreshTimeUsecsNil

`func (o *McmSourceRegistrationInfo) SetLastRefreshTimeUsecsNil(b bool)`

 SetLastRefreshTimeUsecsNil sets the value for LastRefreshTimeUsecs to be an explicit nil

### UnsetLastRefreshTimeUsecs
`func (o *McmSourceRegistrationInfo) UnsetLastRefreshTimeUsecs()`

UnsetLastRefreshTimeUsecs ensures that no value is present for LastRefreshTimeUsecs, not even an explicit nil
### GetRefreshError

`func (o *McmSourceRegistrationInfo) GetRefreshError() string`

GetRefreshError returns the RefreshError field if non-nil, zero value otherwise.

### GetRefreshErrorOk

`func (o *McmSourceRegistrationInfo) GetRefreshErrorOk() (*string, bool)`

GetRefreshErrorOk returns a tuple with the RefreshError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshError

`func (o *McmSourceRegistrationInfo) SetRefreshError(v string)`

SetRefreshError sets RefreshError field to given value.

### HasRefreshError

`func (o *McmSourceRegistrationInfo) HasRefreshError() bool`

HasRefreshError returns a boolean if a field has been set.

### SetRefreshErrorNil

`func (o *McmSourceRegistrationInfo) SetRefreshErrorNil(b bool)`

 SetRefreshErrorNil sets the value for RefreshError to be an explicit nil

### UnsetRefreshError
`func (o *McmSourceRegistrationInfo) UnsetRefreshError()`

UnsetRefreshError ensures that no value is present for RefreshError, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


