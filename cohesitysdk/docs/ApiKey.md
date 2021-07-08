# ApiKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedTimeMsecs** | Pointer to **NullableInt64** | Specifies the API key created time in milli seconds. | [optional] 
**CreatedUserSid** | Pointer to **NullableString** | Specifies the user sid who created this API key. | [optional] 
**CreatedUsername** | Pointer to **NullableString** | Specifies the username who created this API key. | [optional] 
**ExpiringTimeMsecs** | Pointer to **NullableInt64** | Specifies a time stamp when the API key will expire in milli seconds. | [optional] 
**Id** | Pointer to **NullableString** | Specifies the API key id. | [optional] 
**IsActive** | Pointer to **NullableBool** | Specifies if the API key is active. Only an active and not expired API key can be used for authentication. | [optional] 
**IsExpired** | Pointer to **NullableBool** | Specifies if the API key is expired. Expired API keys cannot be used for authentication. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the API key name. | [optional] 
**OwnerUserSid** | Pointer to **NullableString** | Specifies the user sid who owns this API key. | [optional] 
**OwnerUsername** | Pointer to **NullableString** | Specifies the username who owns this API key. | [optional] 

## Methods

### NewApiKey

`func NewApiKey() *ApiKey`

NewApiKey instantiates a new ApiKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiKeyWithDefaults

`func NewApiKeyWithDefaults() *ApiKey`

NewApiKeyWithDefaults instantiates a new ApiKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedTimeMsecs

`func (o *ApiKey) GetCreatedTimeMsecs() int64`

GetCreatedTimeMsecs returns the CreatedTimeMsecs field if non-nil, zero value otherwise.

### GetCreatedTimeMsecsOk

`func (o *ApiKey) GetCreatedTimeMsecsOk() (*int64, bool)`

GetCreatedTimeMsecsOk returns a tuple with the CreatedTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimeMsecs

`func (o *ApiKey) SetCreatedTimeMsecs(v int64)`

SetCreatedTimeMsecs sets CreatedTimeMsecs field to given value.

### HasCreatedTimeMsecs

`func (o *ApiKey) HasCreatedTimeMsecs() bool`

HasCreatedTimeMsecs returns a boolean if a field has been set.

### SetCreatedTimeMsecsNil

`func (o *ApiKey) SetCreatedTimeMsecsNil(b bool)`

 SetCreatedTimeMsecsNil sets the value for CreatedTimeMsecs to be an explicit nil

### UnsetCreatedTimeMsecs
`func (o *ApiKey) UnsetCreatedTimeMsecs()`

UnsetCreatedTimeMsecs ensures that no value is present for CreatedTimeMsecs, not even an explicit nil
### GetCreatedUserSid

`func (o *ApiKey) GetCreatedUserSid() string`

GetCreatedUserSid returns the CreatedUserSid field if non-nil, zero value otherwise.

### GetCreatedUserSidOk

`func (o *ApiKey) GetCreatedUserSidOk() (*string, bool)`

GetCreatedUserSidOk returns a tuple with the CreatedUserSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedUserSid

`func (o *ApiKey) SetCreatedUserSid(v string)`

SetCreatedUserSid sets CreatedUserSid field to given value.

### HasCreatedUserSid

`func (o *ApiKey) HasCreatedUserSid() bool`

HasCreatedUserSid returns a boolean if a field has been set.

### SetCreatedUserSidNil

`func (o *ApiKey) SetCreatedUserSidNil(b bool)`

 SetCreatedUserSidNil sets the value for CreatedUserSid to be an explicit nil

### UnsetCreatedUserSid
`func (o *ApiKey) UnsetCreatedUserSid()`

UnsetCreatedUserSid ensures that no value is present for CreatedUserSid, not even an explicit nil
### GetCreatedUsername

`func (o *ApiKey) GetCreatedUsername() string`

GetCreatedUsername returns the CreatedUsername field if non-nil, zero value otherwise.

### GetCreatedUsernameOk

`func (o *ApiKey) GetCreatedUsernameOk() (*string, bool)`

GetCreatedUsernameOk returns a tuple with the CreatedUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedUsername

`func (o *ApiKey) SetCreatedUsername(v string)`

SetCreatedUsername sets CreatedUsername field to given value.

### HasCreatedUsername

`func (o *ApiKey) HasCreatedUsername() bool`

HasCreatedUsername returns a boolean if a field has been set.

### SetCreatedUsernameNil

`func (o *ApiKey) SetCreatedUsernameNil(b bool)`

 SetCreatedUsernameNil sets the value for CreatedUsername to be an explicit nil

### UnsetCreatedUsername
`func (o *ApiKey) UnsetCreatedUsername()`

UnsetCreatedUsername ensures that no value is present for CreatedUsername, not even an explicit nil
### GetExpiringTimeMsecs

`func (o *ApiKey) GetExpiringTimeMsecs() int64`

GetExpiringTimeMsecs returns the ExpiringTimeMsecs field if non-nil, zero value otherwise.

### GetExpiringTimeMsecsOk

`func (o *ApiKey) GetExpiringTimeMsecsOk() (*int64, bool)`

GetExpiringTimeMsecsOk returns a tuple with the ExpiringTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiringTimeMsecs

`func (o *ApiKey) SetExpiringTimeMsecs(v int64)`

SetExpiringTimeMsecs sets ExpiringTimeMsecs field to given value.

### HasExpiringTimeMsecs

`func (o *ApiKey) HasExpiringTimeMsecs() bool`

HasExpiringTimeMsecs returns a boolean if a field has been set.

### SetExpiringTimeMsecsNil

`func (o *ApiKey) SetExpiringTimeMsecsNil(b bool)`

 SetExpiringTimeMsecsNil sets the value for ExpiringTimeMsecs to be an explicit nil

### UnsetExpiringTimeMsecs
`func (o *ApiKey) UnsetExpiringTimeMsecs()`

UnsetExpiringTimeMsecs ensures that no value is present for ExpiringTimeMsecs, not even an explicit nil
### GetId

`func (o *ApiKey) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiKey) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiKey) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiKey) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ApiKey) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ApiKey) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIsActive

`func (o *ApiKey) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *ApiKey) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *ApiKey) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *ApiKey) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### SetIsActiveNil

`func (o *ApiKey) SetIsActiveNil(b bool)`

 SetIsActiveNil sets the value for IsActive to be an explicit nil

### UnsetIsActive
`func (o *ApiKey) UnsetIsActive()`

UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
### GetIsExpired

`func (o *ApiKey) GetIsExpired() bool`

GetIsExpired returns the IsExpired field if non-nil, zero value otherwise.

### GetIsExpiredOk

`func (o *ApiKey) GetIsExpiredOk() (*bool, bool)`

GetIsExpiredOk returns a tuple with the IsExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExpired

`func (o *ApiKey) SetIsExpired(v bool)`

SetIsExpired sets IsExpired field to given value.

### HasIsExpired

`func (o *ApiKey) HasIsExpired() bool`

HasIsExpired returns a boolean if a field has been set.

### SetIsExpiredNil

`func (o *ApiKey) SetIsExpiredNil(b bool)`

 SetIsExpiredNil sets the value for IsExpired to be an explicit nil

### UnsetIsExpired
`func (o *ApiKey) UnsetIsExpired()`

UnsetIsExpired ensures that no value is present for IsExpired, not even an explicit nil
### GetName

`func (o *ApiKey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiKey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiKey) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiKey) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ApiKey) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ApiKey) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOwnerUserSid

`func (o *ApiKey) GetOwnerUserSid() string`

GetOwnerUserSid returns the OwnerUserSid field if non-nil, zero value otherwise.

### GetOwnerUserSidOk

`func (o *ApiKey) GetOwnerUserSidOk() (*string, bool)`

GetOwnerUserSidOk returns a tuple with the OwnerUserSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerUserSid

`func (o *ApiKey) SetOwnerUserSid(v string)`

SetOwnerUserSid sets OwnerUserSid field to given value.

### HasOwnerUserSid

`func (o *ApiKey) HasOwnerUserSid() bool`

HasOwnerUserSid returns a boolean if a field has been set.

### SetOwnerUserSidNil

`func (o *ApiKey) SetOwnerUserSidNil(b bool)`

 SetOwnerUserSidNil sets the value for OwnerUserSid to be an explicit nil

### UnsetOwnerUserSid
`func (o *ApiKey) UnsetOwnerUserSid()`

UnsetOwnerUserSid ensures that no value is present for OwnerUserSid, not even an explicit nil
### GetOwnerUsername

`func (o *ApiKey) GetOwnerUsername() string`

GetOwnerUsername returns the OwnerUsername field if non-nil, zero value otherwise.

### GetOwnerUsernameOk

`func (o *ApiKey) GetOwnerUsernameOk() (*string, bool)`

GetOwnerUsernameOk returns a tuple with the OwnerUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerUsername

`func (o *ApiKey) SetOwnerUsername(v string)`

SetOwnerUsername sets OwnerUsername field to given value.

### HasOwnerUsername

`func (o *ApiKey) HasOwnerUsername() bool`

HasOwnerUsername returns a boolean if a field has been set.

### SetOwnerUsernameNil

`func (o *ApiKey) SetOwnerUsernameNil(b bool)`

 SetOwnerUsernameNil sets the value for OwnerUsername to be an explicit nil

### UnsetOwnerUsername
`func (o *ApiKey) UnsetOwnerUsername()`

UnsetOwnerUsername ensures that no value is present for OwnerUsername, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


