# CreateOrUpdateAPIKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiryTimeMsecs** | Pointer to **NullableInt64** | Specifies the time in milliseconds when the API key will expire. Set null for no expiry. | [optional] 
**IsActive** | Pointer to **NullableBool** | Specifies if the API key is active. | [optional] [default to true]
**Name** | **NullableString** | Specifies the API key name. | 

## Methods

### NewCreateOrUpdateAPIKeyRequest

`func NewCreateOrUpdateAPIKeyRequest(name NullableString, ) *CreateOrUpdateAPIKeyRequest`

NewCreateOrUpdateAPIKeyRequest instantiates a new CreateOrUpdateAPIKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrUpdateAPIKeyRequestWithDefaults

`func NewCreateOrUpdateAPIKeyRequestWithDefaults() *CreateOrUpdateAPIKeyRequest`

NewCreateOrUpdateAPIKeyRequestWithDefaults instantiates a new CreateOrUpdateAPIKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiryTimeMsecs

`func (o *CreateOrUpdateAPIKeyRequest) GetExpiryTimeMsecs() int64`

GetExpiryTimeMsecs returns the ExpiryTimeMsecs field if non-nil, zero value otherwise.

### GetExpiryTimeMsecsOk

`func (o *CreateOrUpdateAPIKeyRequest) GetExpiryTimeMsecsOk() (*int64, bool)`

GetExpiryTimeMsecsOk returns a tuple with the ExpiryTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryTimeMsecs

`func (o *CreateOrUpdateAPIKeyRequest) SetExpiryTimeMsecs(v int64)`

SetExpiryTimeMsecs sets ExpiryTimeMsecs field to given value.

### HasExpiryTimeMsecs

`func (o *CreateOrUpdateAPIKeyRequest) HasExpiryTimeMsecs() bool`

HasExpiryTimeMsecs returns a boolean if a field has been set.

### SetExpiryTimeMsecsNil

`func (o *CreateOrUpdateAPIKeyRequest) SetExpiryTimeMsecsNil(b bool)`

 SetExpiryTimeMsecsNil sets the value for ExpiryTimeMsecs to be an explicit nil

### UnsetExpiryTimeMsecs
`func (o *CreateOrUpdateAPIKeyRequest) UnsetExpiryTimeMsecs()`

UnsetExpiryTimeMsecs ensures that no value is present for ExpiryTimeMsecs, not even an explicit nil
### GetIsActive

`func (o *CreateOrUpdateAPIKeyRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *CreateOrUpdateAPIKeyRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *CreateOrUpdateAPIKeyRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *CreateOrUpdateAPIKeyRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### SetIsActiveNil

`func (o *CreateOrUpdateAPIKeyRequest) SetIsActiveNil(b bool)`

 SetIsActiveNil sets the value for IsActive to be an explicit nil

### UnsetIsActive
`func (o *CreateOrUpdateAPIKeyRequest) UnsetIsActive()`

UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
### GetName

`func (o *CreateOrUpdateAPIKeyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateOrUpdateAPIKeyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateOrUpdateAPIKeyRequest) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *CreateOrUpdateAPIKeyRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateOrUpdateAPIKeyRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


