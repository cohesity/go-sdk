# UpdateApiKeyParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiringTimeMsecs** | Pointer to **NullableInt64** | Specifies a time stamp when the API key will expire in milli seconds. | [optional] 
**IsActive** | Pointer to **NullableBool** | Specifies if the API key is active. Only an active and not expired API key can be used for authentication. | [optional] 
**Name** | **NullableString** | Specifies the name of API key. | 

## Methods

### NewUpdateApiKeyParams

`func NewUpdateApiKeyParams(name NullableString, ) *UpdateApiKeyParams`

NewUpdateApiKeyParams instantiates a new UpdateApiKeyParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateApiKeyParamsWithDefaults

`func NewUpdateApiKeyParamsWithDefaults() *UpdateApiKeyParams`

NewUpdateApiKeyParamsWithDefaults instantiates a new UpdateApiKeyParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiringTimeMsecs

`func (o *UpdateApiKeyParams) GetExpiringTimeMsecs() int64`

GetExpiringTimeMsecs returns the ExpiringTimeMsecs field if non-nil, zero value otherwise.

### GetExpiringTimeMsecsOk

`func (o *UpdateApiKeyParams) GetExpiringTimeMsecsOk() (*int64, bool)`

GetExpiringTimeMsecsOk returns a tuple with the ExpiringTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiringTimeMsecs

`func (o *UpdateApiKeyParams) SetExpiringTimeMsecs(v int64)`

SetExpiringTimeMsecs sets ExpiringTimeMsecs field to given value.

### HasExpiringTimeMsecs

`func (o *UpdateApiKeyParams) HasExpiringTimeMsecs() bool`

HasExpiringTimeMsecs returns a boolean if a field has been set.

### SetExpiringTimeMsecsNil

`func (o *UpdateApiKeyParams) SetExpiringTimeMsecsNil(b bool)`

 SetExpiringTimeMsecsNil sets the value for ExpiringTimeMsecs to be an explicit nil

### UnsetExpiringTimeMsecs
`func (o *UpdateApiKeyParams) UnsetExpiringTimeMsecs()`

UnsetExpiringTimeMsecs ensures that no value is present for ExpiringTimeMsecs, not even an explicit nil
### GetIsActive

`func (o *UpdateApiKeyParams) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *UpdateApiKeyParams) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *UpdateApiKeyParams) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *UpdateApiKeyParams) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### SetIsActiveNil

`func (o *UpdateApiKeyParams) SetIsActiveNil(b bool)`

 SetIsActiveNil sets the value for IsActive to be an explicit nil

### UnsetIsActive
`func (o *UpdateApiKeyParams) UnsetIsActive()`

UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
### GetName

`func (o *UpdateApiKeyParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateApiKeyParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateApiKeyParams) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *UpdateApiKeyParams) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateApiKeyParams) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


