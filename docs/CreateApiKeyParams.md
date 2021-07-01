# CreateApiKeyParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiringTimeMsecs** | Pointer to **NullableInt64** | Specifies a time stamp when the API key will expire in milli seconds. | [optional] 
**IsActive** | Pointer to **NullableBool** | Specifies if the API key is active. Only an active and not expired API key can be used for authentication. | [optional] 
**Name** | **NullableString** | Specifies the name of API key. | 

## Methods

### NewCreateApiKeyParams

`func NewCreateApiKeyParams(name NullableString, ) *CreateApiKeyParams`

NewCreateApiKeyParams instantiates a new CreateApiKeyParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateApiKeyParamsWithDefaults

`func NewCreateApiKeyParamsWithDefaults() *CreateApiKeyParams`

NewCreateApiKeyParamsWithDefaults instantiates a new CreateApiKeyParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiringTimeMsecs

`func (o *CreateApiKeyParams) GetExpiringTimeMsecs() int64`

GetExpiringTimeMsecs returns the ExpiringTimeMsecs field if non-nil, zero value otherwise.

### GetExpiringTimeMsecsOk

`func (o *CreateApiKeyParams) GetExpiringTimeMsecsOk() (*int64, bool)`

GetExpiringTimeMsecsOk returns a tuple with the ExpiringTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiringTimeMsecs

`func (o *CreateApiKeyParams) SetExpiringTimeMsecs(v int64)`

SetExpiringTimeMsecs sets ExpiringTimeMsecs field to given value.

### HasExpiringTimeMsecs

`func (o *CreateApiKeyParams) HasExpiringTimeMsecs() bool`

HasExpiringTimeMsecs returns a boolean if a field has been set.

### SetExpiringTimeMsecsNil

`func (o *CreateApiKeyParams) SetExpiringTimeMsecsNil(b bool)`

 SetExpiringTimeMsecsNil sets the value for ExpiringTimeMsecs to be an explicit nil

### UnsetExpiringTimeMsecs
`func (o *CreateApiKeyParams) UnsetExpiringTimeMsecs()`

UnsetExpiringTimeMsecs ensures that no value is present for ExpiringTimeMsecs, not even an explicit nil
### GetIsActive

`func (o *CreateApiKeyParams) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *CreateApiKeyParams) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *CreateApiKeyParams) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *CreateApiKeyParams) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### SetIsActiveNil

`func (o *CreateApiKeyParams) SetIsActiveNil(b bool)`

 SetIsActiveNil sets the value for IsActive to be an explicit nil

### UnsetIsActive
`func (o *CreateApiKeyParams) UnsetIsActive()`

UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
### GetName

`func (o *CreateApiKeyParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateApiKeyParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateApiKeyParams) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *CreateApiKeyParams) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateApiKeyParams) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


