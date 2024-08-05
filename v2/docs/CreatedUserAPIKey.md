# CreatedUserAPIKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedByUserSid** | Pointer to **NullableString** | Specifies the user SID who created the API key. | [optional] [readonly] 
**CreatedTimeMsecs** | Pointer to **NullableInt64** | Specifies the time in milliseconds when the API key was created. | [optional] [readonly] 
**ExpiryTimeMsecs** | Pointer to **NullableInt64** | Specifies the time in milliseconds when the API key will expire. null signifies no-expiry. | [optional] [readonly] 
**Id** | Pointer to **NullableString** | Specifies the unique id of the API key. | [optional] [readonly] 
**IsActive** | Pointer to **NullableBool** | Specifies if the API key is active. | [optional] [readonly] [default to true]
**IsExpired** | Pointer to **NullableBool** | Specifies if the API key has expired. | [optional] [readonly] 
**LastRotatedTimeMsecs** | Pointer to **NullableInt64** | Specifies the time in milliseconds when the API key was last rotated. | [optional] [readonly] 
**Name** | Pointer to **NullableString** | Specifies the API key name. | [optional] [readonly] 
**UserSid** | Pointer to **NullableString** | Specifies the user who owns the API key. | [optional] [readonly] 
**ApiKey** | Pointer to **NullableString** | Specifies the API key. | [optional] [readonly] 

## Methods

### NewCreatedUserAPIKey

`func NewCreatedUserAPIKey() *CreatedUserAPIKey`

NewCreatedUserAPIKey instantiates a new CreatedUserAPIKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatedUserAPIKeyWithDefaults

`func NewCreatedUserAPIKeyWithDefaults() *CreatedUserAPIKey`

NewCreatedUserAPIKeyWithDefaults instantiates a new CreatedUserAPIKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedByUserSid

`func (o *CreatedUserAPIKey) GetCreatedByUserSid() string`

GetCreatedByUserSid returns the CreatedByUserSid field if non-nil, zero value otherwise.

### GetCreatedByUserSidOk

`func (o *CreatedUserAPIKey) GetCreatedByUserSidOk() (*string, bool)`

GetCreatedByUserSidOk returns a tuple with the CreatedByUserSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserSid

`func (o *CreatedUserAPIKey) SetCreatedByUserSid(v string)`

SetCreatedByUserSid sets CreatedByUserSid field to given value.

### HasCreatedByUserSid

`func (o *CreatedUserAPIKey) HasCreatedByUserSid() bool`

HasCreatedByUserSid returns a boolean if a field has been set.

### SetCreatedByUserSidNil

`func (o *CreatedUserAPIKey) SetCreatedByUserSidNil(b bool)`

 SetCreatedByUserSidNil sets the value for CreatedByUserSid to be an explicit nil

### UnsetCreatedByUserSid
`func (o *CreatedUserAPIKey) UnsetCreatedByUserSid()`

UnsetCreatedByUserSid ensures that no value is present for CreatedByUserSid, not even an explicit nil
### GetCreatedTimeMsecs

`func (o *CreatedUserAPIKey) GetCreatedTimeMsecs() int64`

GetCreatedTimeMsecs returns the CreatedTimeMsecs field if non-nil, zero value otherwise.

### GetCreatedTimeMsecsOk

`func (o *CreatedUserAPIKey) GetCreatedTimeMsecsOk() (*int64, bool)`

GetCreatedTimeMsecsOk returns a tuple with the CreatedTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimeMsecs

`func (o *CreatedUserAPIKey) SetCreatedTimeMsecs(v int64)`

SetCreatedTimeMsecs sets CreatedTimeMsecs field to given value.

### HasCreatedTimeMsecs

`func (o *CreatedUserAPIKey) HasCreatedTimeMsecs() bool`

HasCreatedTimeMsecs returns a boolean if a field has been set.

### SetCreatedTimeMsecsNil

`func (o *CreatedUserAPIKey) SetCreatedTimeMsecsNil(b bool)`

 SetCreatedTimeMsecsNil sets the value for CreatedTimeMsecs to be an explicit nil

### UnsetCreatedTimeMsecs
`func (o *CreatedUserAPIKey) UnsetCreatedTimeMsecs()`

UnsetCreatedTimeMsecs ensures that no value is present for CreatedTimeMsecs, not even an explicit nil
### GetExpiryTimeMsecs

`func (o *CreatedUserAPIKey) GetExpiryTimeMsecs() int64`

GetExpiryTimeMsecs returns the ExpiryTimeMsecs field if non-nil, zero value otherwise.

### GetExpiryTimeMsecsOk

`func (o *CreatedUserAPIKey) GetExpiryTimeMsecsOk() (*int64, bool)`

GetExpiryTimeMsecsOk returns a tuple with the ExpiryTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryTimeMsecs

`func (o *CreatedUserAPIKey) SetExpiryTimeMsecs(v int64)`

SetExpiryTimeMsecs sets ExpiryTimeMsecs field to given value.

### HasExpiryTimeMsecs

`func (o *CreatedUserAPIKey) HasExpiryTimeMsecs() bool`

HasExpiryTimeMsecs returns a boolean if a field has been set.

### SetExpiryTimeMsecsNil

`func (o *CreatedUserAPIKey) SetExpiryTimeMsecsNil(b bool)`

 SetExpiryTimeMsecsNil sets the value for ExpiryTimeMsecs to be an explicit nil

### UnsetExpiryTimeMsecs
`func (o *CreatedUserAPIKey) UnsetExpiryTimeMsecs()`

UnsetExpiryTimeMsecs ensures that no value is present for ExpiryTimeMsecs, not even an explicit nil
### GetId

`func (o *CreatedUserAPIKey) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreatedUserAPIKey) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreatedUserAPIKey) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreatedUserAPIKey) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CreatedUserAPIKey) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CreatedUserAPIKey) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIsActive

`func (o *CreatedUserAPIKey) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *CreatedUserAPIKey) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *CreatedUserAPIKey) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *CreatedUserAPIKey) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### SetIsActiveNil

`func (o *CreatedUserAPIKey) SetIsActiveNil(b bool)`

 SetIsActiveNil sets the value for IsActive to be an explicit nil

### UnsetIsActive
`func (o *CreatedUserAPIKey) UnsetIsActive()`

UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
### GetIsExpired

`func (o *CreatedUserAPIKey) GetIsExpired() bool`

GetIsExpired returns the IsExpired field if non-nil, zero value otherwise.

### GetIsExpiredOk

`func (o *CreatedUserAPIKey) GetIsExpiredOk() (*bool, bool)`

GetIsExpiredOk returns a tuple with the IsExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExpired

`func (o *CreatedUserAPIKey) SetIsExpired(v bool)`

SetIsExpired sets IsExpired field to given value.

### HasIsExpired

`func (o *CreatedUserAPIKey) HasIsExpired() bool`

HasIsExpired returns a boolean if a field has been set.

### SetIsExpiredNil

`func (o *CreatedUserAPIKey) SetIsExpiredNil(b bool)`

 SetIsExpiredNil sets the value for IsExpired to be an explicit nil

### UnsetIsExpired
`func (o *CreatedUserAPIKey) UnsetIsExpired()`

UnsetIsExpired ensures that no value is present for IsExpired, not even an explicit nil
### GetLastRotatedTimeMsecs

`func (o *CreatedUserAPIKey) GetLastRotatedTimeMsecs() int64`

GetLastRotatedTimeMsecs returns the LastRotatedTimeMsecs field if non-nil, zero value otherwise.

### GetLastRotatedTimeMsecsOk

`func (o *CreatedUserAPIKey) GetLastRotatedTimeMsecsOk() (*int64, bool)`

GetLastRotatedTimeMsecsOk returns a tuple with the LastRotatedTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRotatedTimeMsecs

`func (o *CreatedUserAPIKey) SetLastRotatedTimeMsecs(v int64)`

SetLastRotatedTimeMsecs sets LastRotatedTimeMsecs field to given value.

### HasLastRotatedTimeMsecs

`func (o *CreatedUserAPIKey) HasLastRotatedTimeMsecs() bool`

HasLastRotatedTimeMsecs returns a boolean if a field has been set.

### SetLastRotatedTimeMsecsNil

`func (o *CreatedUserAPIKey) SetLastRotatedTimeMsecsNil(b bool)`

 SetLastRotatedTimeMsecsNil sets the value for LastRotatedTimeMsecs to be an explicit nil

### UnsetLastRotatedTimeMsecs
`func (o *CreatedUserAPIKey) UnsetLastRotatedTimeMsecs()`

UnsetLastRotatedTimeMsecs ensures that no value is present for LastRotatedTimeMsecs, not even an explicit nil
### GetName

`func (o *CreatedUserAPIKey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreatedUserAPIKey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreatedUserAPIKey) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreatedUserAPIKey) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CreatedUserAPIKey) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreatedUserAPIKey) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetUserSid

`func (o *CreatedUserAPIKey) GetUserSid() string`

GetUserSid returns the UserSid field if non-nil, zero value otherwise.

### GetUserSidOk

`func (o *CreatedUserAPIKey) GetUserSidOk() (*string, bool)`

GetUserSidOk returns a tuple with the UserSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSid

`func (o *CreatedUserAPIKey) SetUserSid(v string)`

SetUserSid sets UserSid field to given value.

### HasUserSid

`func (o *CreatedUserAPIKey) HasUserSid() bool`

HasUserSid returns a boolean if a field has been set.

### SetUserSidNil

`func (o *CreatedUserAPIKey) SetUserSidNil(b bool)`

 SetUserSidNil sets the value for UserSid to be an explicit nil

### UnsetUserSid
`func (o *CreatedUserAPIKey) UnsetUserSid()`

UnsetUserSid ensures that no value is present for UserSid, not even an explicit nil
### GetApiKey

`func (o *CreatedUserAPIKey) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *CreatedUserAPIKey) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *CreatedUserAPIKey) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *CreatedUserAPIKey) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### SetApiKeyNil

`func (o *CreatedUserAPIKey) SetApiKeyNil(b bool)`

 SetApiKeyNil sets the value for ApiKey to be an explicit nil

### UnsetApiKey
`func (o *CreatedUserAPIKey) UnsetApiKey()`

UnsetApiKey ensures that no value is present for ApiKey, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


