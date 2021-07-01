# CreatedApiKey

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
**Key** | Pointer to **NullableString** | Specifies the created key. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the API key name. | [optional] 
**OwnerUserSid** | Pointer to **NullableString** | Specifies the user sid who owns this API key. | [optional] 
**OwnerUsername** | Pointer to **NullableString** | Specifies the username who owns this API key. | [optional] 

## Methods

### NewCreatedApiKey

`func NewCreatedApiKey() *CreatedApiKey`

NewCreatedApiKey instantiates a new CreatedApiKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatedApiKeyWithDefaults

`func NewCreatedApiKeyWithDefaults() *CreatedApiKey`

NewCreatedApiKeyWithDefaults instantiates a new CreatedApiKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedTimeMsecs

`func (o *CreatedApiKey) GetCreatedTimeMsecs() int64`

GetCreatedTimeMsecs returns the CreatedTimeMsecs field if non-nil, zero value otherwise.

### GetCreatedTimeMsecsOk

`func (o *CreatedApiKey) GetCreatedTimeMsecsOk() (*int64, bool)`

GetCreatedTimeMsecsOk returns a tuple with the CreatedTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimeMsecs

`func (o *CreatedApiKey) SetCreatedTimeMsecs(v int64)`

SetCreatedTimeMsecs sets CreatedTimeMsecs field to given value.

### HasCreatedTimeMsecs

`func (o *CreatedApiKey) HasCreatedTimeMsecs() bool`

HasCreatedTimeMsecs returns a boolean if a field has been set.

### SetCreatedTimeMsecsNil

`func (o *CreatedApiKey) SetCreatedTimeMsecsNil(b bool)`

 SetCreatedTimeMsecsNil sets the value for CreatedTimeMsecs to be an explicit nil

### UnsetCreatedTimeMsecs
`func (o *CreatedApiKey) UnsetCreatedTimeMsecs()`

UnsetCreatedTimeMsecs ensures that no value is present for CreatedTimeMsecs, not even an explicit nil
### GetCreatedUserSid

`func (o *CreatedApiKey) GetCreatedUserSid() string`

GetCreatedUserSid returns the CreatedUserSid field if non-nil, zero value otherwise.

### GetCreatedUserSidOk

`func (o *CreatedApiKey) GetCreatedUserSidOk() (*string, bool)`

GetCreatedUserSidOk returns a tuple with the CreatedUserSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedUserSid

`func (o *CreatedApiKey) SetCreatedUserSid(v string)`

SetCreatedUserSid sets CreatedUserSid field to given value.

### HasCreatedUserSid

`func (o *CreatedApiKey) HasCreatedUserSid() bool`

HasCreatedUserSid returns a boolean if a field has been set.

### SetCreatedUserSidNil

`func (o *CreatedApiKey) SetCreatedUserSidNil(b bool)`

 SetCreatedUserSidNil sets the value for CreatedUserSid to be an explicit nil

### UnsetCreatedUserSid
`func (o *CreatedApiKey) UnsetCreatedUserSid()`

UnsetCreatedUserSid ensures that no value is present for CreatedUserSid, not even an explicit nil
### GetCreatedUsername

`func (o *CreatedApiKey) GetCreatedUsername() string`

GetCreatedUsername returns the CreatedUsername field if non-nil, zero value otherwise.

### GetCreatedUsernameOk

`func (o *CreatedApiKey) GetCreatedUsernameOk() (*string, bool)`

GetCreatedUsernameOk returns a tuple with the CreatedUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedUsername

`func (o *CreatedApiKey) SetCreatedUsername(v string)`

SetCreatedUsername sets CreatedUsername field to given value.

### HasCreatedUsername

`func (o *CreatedApiKey) HasCreatedUsername() bool`

HasCreatedUsername returns a boolean if a field has been set.

### SetCreatedUsernameNil

`func (o *CreatedApiKey) SetCreatedUsernameNil(b bool)`

 SetCreatedUsernameNil sets the value for CreatedUsername to be an explicit nil

### UnsetCreatedUsername
`func (o *CreatedApiKey) UnsetCreatedUsername()`

UnsetCreatedUsername ensures that no value is present for CreatedUsername, not even an explicit nil
### GetExpiringTimeMsecs

`func (o *CreatedApiKey) GetExpiringTimeMsecs() int64`

GetExpiringTimeMsecs returns the ExpiringTimeMsecs field if non-nil, zero value otherwise.

### GetExpiringTimeMsecsOk

`func (o *CreatedApiKey) GetExpiringTimeMsecsOk() (*int64, bool)`

GetExpiringTimeMsecsOk returns a tuple with the ExpiringTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiringTimeMsecs

`func (o *CreatedApiKey) SetExpiringTimeMsecs(v int64)`

SetExpiringTimeMsecs sets ExpiringTimeMsecs field to given value.

### HasExpiringTimeMsecs

`func (o *CreatedApiKey) HasExpiringTimeMsecs() bool`

HasExpiringTimeMsecs returns a boolean if a field has been set.

### SetExpiringTimeMsecsNil

`func (o *CreatedApiKey) SetExpiringTimeMsecsNil(b bool)`

 SetExpiringTimeMsecsNil sets the value for ExpiringTimeMsecs to be an explicit nil

### UnsetExpiringTimeMsecs
`func (o *CreatedApiKey) UnsetExpiringTimeMsecs()`

UnsetExpiringTimeMsecs ensures that no value is present for ExpiringTimeMsecs, not even an explicit nil
### GetId

`func (o *CreatedApiKey) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreatedApiKey) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreatedApiKey) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreatedApiKey) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CreatedApiKey) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CreatedApiKey) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIsActive

`func (o *CreatedApiKey) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *CreatedApiKey) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *CreatedApiKey) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *CreatedApiKey) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### SetIsActiveNil

`func (o *CreatedApiKey) SetIsActiveNil(b bool)`

 SetIsActiveNil sets the value for IsActive to be an explicit nil

### UnsetIsActive
`func (o *CreatedApiKey) UnsetIsActive()`

UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
### GetIsExpired

`func (o *CreatedApiKey) GetIsExpired() bool`

GetIsExpired returns the IsExpired field if non-nil, zero value otherwise.

### GetIsExpiredOk

`func (o *CreatedApiKey) GetIsExpiredOk() (*bool, bool)`

GetIsExpiredOk returns a tuple with the IsExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExpired

`func (o *CreatedApiKey) SetIsExpired(v bool)`

SetIsExpired sets IsExpired field to given value.

### HasIsExpired

`func (o *CreatedApiKey) HasIsExpired() bool`

HasIsExpired returns a boolean if a field has been set.

### SetIsExpiredNil

`func (o *CreatedApiKey) SetIsExpiredNil(b bool)`

 SetIsExpiredNil sets the value for IsExpired to be an explicit nil

### UnsetIsExpired
`func (o *CreatedApiKey) UnsetIsExpired()`

UnsetIsExpired ensures that no value is present for IsExpired, not even an explicit nil
### GetKey

`func (o *CreatedApiKey) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CreatedApiKey) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CreatedApiKey) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *CreatedApiKey) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *CreatedApiKey) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *CreatedApiKey) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetName

`func (o *CreatedApiKey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreatedApiKey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreatedApiKey) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreatedApiKey) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CreatedApiKey) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreatedApiKey) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOwnerUserSid

`func (o *CreatedApiKey) GetOwnerUserSid() string`

GetOwnerUserSid returns the OwnerUserSid field if non-nil, zero value otherwise.

### GetOwnerUserSidOk

`func (o *CreatedApiKey) GetOwnerUserSidOk() (*string, bool)`

GetOwnerUserSidOk returns a tuple with the OwnerUserSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerUserSid

`func (o *CreatedApiKey) SetOwnerUserSid(v string)`

SetOwnerUserSid sets OwnerUserSid field to given value.

### HasOwnerUserSid

`func (o *CreatedApiKey) HasOwnerUserSid() bool`

HasOwnerUserSid returns a boolean if a field has been set.

### SetOwnerUserSidNil

`func (o *CreatedApiKey) SetOwnerUserSidNil(b bool)`

 SetOwnerUserSidNil sets the value for OwnerUserSid to be an explicit nil

### UnsetOwnerUserSid
`func (o *CreatedApiKey) UnsetOwnerUserSid()`

UnsetOwnerUserSid ensures that no value is present for OwnerUserSid, not even an explicit nil
### GetOwnerUsername

`func (o *CreatedApiKey) GetOwnerUsername() string`

GetOwnerUsername returns the OwnerUsername field if non-nil, zero value otherwise.

### GetOwnerUsernameOk

`func (o *CreatedApiKey) GetOwnerUsernameOk() (*string, bool)`

GetOwnerUsernameOk returns a tuple with the OwnerUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerUsername

`func (o *CreatedApiKey) SetOwnerUsername(v string)`

SetOwnerUsername sets OwnerUsername field to given value.

### HasOwnerUsername

`func (o *CreatedApiKey) HasOwnerUsername() bool`

HasOwnerUsername returns a boolean if a field has been set.

### SetOwnerUsernameNil

`func (o *CreatedApiKey) SetOwnerUsernameNil(b bool)`

 SetOwnerUsernameNil sets the value for OwnerUsername to be an explicit nil

### UnsetOwnerUsername
`func (o *CreatedApiKey) UnsetOwnerUsername()`

UnsetOwnerUsername ensures that no value is present for OwnerUsername, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


