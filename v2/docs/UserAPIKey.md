# UserAPIKey

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

## Methods

### NewUserAPIKey

`func NewUserAPIKey() *UserAPIKey`

NewUserAPIKey instantiates a new UserAPIKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserAPIKeyWithDefaults

`func NewUserAPIKeyWithDefaults() *UserAPIKey`

NewUserAPIKeyWithDefaults instantiates a new UserAPIKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedByUserSid

`func (o *UserAPIKey) GetCreatedByUserSid() string`

GetCreatedByUserSid returns the CreatedByUserSid field if non-nil, zero value otherwise.

### GetCreatedByUserSidOk

`func (o *UserAPIKey) GetCreatedByUserSidOk() (*string, bool)`

GetCreatedByUserSidOk returns a tuple with the CreatedByUserSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserSid

`func (o *UserAPIKey) SetCreatedByUserSid(v string)`

SetCreatedByUserSid sets CreatedByUserSid field to given value.

### HasCreatedByUserSid

`func (o *UserAPIKey) HasCreatedByUserSid() bool`

HasCreatedByUserSid returns a boolean if a field has been set.

### SetCreatedByUserSidNil

`func (o *UserAPIKey) SetCreatedByUserSidNil(b bool)`

 SetCreatedByUserSidNil sets the value for CreatedByUserSid to be an explicit nil

### UnsetCreatedByUserSid
`func (o *UserAPIKey) UnsetCreatedByUserSid()`

UnsetCreatedByUserSid ensures that no value is present for CreatedByUserSid, not even an explicit nil
### GetCreatedTimeMsecs

`func (o *UserAPIKey) GetCreatedTimeMsecs() int64`

GetCreatedTimeMsecs returns the CreatedTimeMsecs field if non-nil, zero value otherwise.

### GetCreatedTimeMsecsOk

`func (o *UserAPIKey) GetCreatedTimeMsecsOk() (*int64, bool)`

GetCreatedTimeMsecsOk returns a tuple with the CreatedTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimeMsecs

`func (o *UserAPIKey) SetCreatedTimeMsecs(v int64)`

SetCreatedTimeMsecs sets CreatedTimeMsecs field to given value.

### HasCreatedTimeMsecs

`func (o *UserAPIKey) HasCreatedTimeMsecs() bool`

HasCreatedTimeMsecs returns a boolean if a field has been set.

### SetCreatedTimeMsecsNil

`func (o *UserAPIKey) SetCreatedTimeMsecsNil(b bool)`

 SetCreatedTimeMsecsNil sets the value for CreatedTimeMsecs to be an explicit nil

### UnsetCreatedTimeMsecs
`func (o *UserAPIKey) UnsetCreatedTimeMsecs()`

UnsetCreatedTimeMsecs ensures that no value is present for CreatedTimeMsecs, not even an explicit nil
### GetExpiryTimeMsecs

`func (o *UserAPIKey) GetExpiryTimeMsecs() int64`

GetExpiryTimeMsecs returns the ExpiryTimeMsecs field if non-nil, zero value otherwise.

### GetExpiryTimeMsecsOk

`func (o *UserAPIKey) GetExpiryTimeMsecsOk() (*int64, bool)`

GetExpiryTimeMsecsOk returns a tuple with the ExpiryTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryTimeMsecs

`func (o *UserAPIKey) SetExpiryTimeMsecs(v int64)`

SetExpiryTimeMsecs sets ExpiryTimeMsecs field to given value.

### HasExpiryTimeMsecs

`func (o *UserAPIKey) HasExpiryTimeMsecs() bool`

HasExpiryTimeMsecs returns a boolean if a field has been set.

### SetExpiryTimeMsecsNil

`func (o *UserAPIKey) SetExpiryTimeMsecsNil(b bool)`

 SetExpiryTimeMsecsNil sets the value for ExpiryTimeMsecs to be an explicit nil

### UnsetExpiryTimeMsecs
`func (o *UserAPIKey) UnsetExpiryTimeMsecs()`

UnsetExpiryTimeMsecs ensures that no value is present for ExpiryTimeMsecs, not even an explicit nil
### GetId

`func (o *UserAPIKey) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserAPIKey) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserAPIKey) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserAPIKey) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *UserAPIKey) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *UserAPIKey) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIsActive

`func (o *UserAPIKey) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *UserAPIKey) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *UserAPIKey) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *UserAPIKey) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### SetIsActiveNil

`func (o *UserAPIKey) SetIsActiveNil(b bool)`

 SetIsActiveNil sets the value for IsActive to be an explicit nil

### UnsetIsActive
`func (o *UserAPIKey) UnsetIsActive()`

UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
### GetIsExpired

`func (o *UserAPIKey) GetIsExpired() bool`

GetIsExpired returns the IsExpired field if non-nil, zero value otherwise.

### GetIsExpiredOk

`func (o *UserAPIKey) GetIsExpiredOk() (*bool, bool)`

GetIsExpiredOk returns a tuple with the IsExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExpired

`func (o *UserAPIKey) SetIsExpired(v bool)`

SetIsExpired sets IsExpired field to given value.

### HasIsExpired

`func (o *UserAPIKey) HasIsExpired() bool`

HasIsExpired returns a boolean if a field has been set.

### SetIsExpiredNil

`func (o *UserAPIKey) SetIsExpiredNil(b bool)`

 SetIsExpiredNil sets the value for IsExpired to be an explicit nil

### UnsetIsExpired
`func (o *UserAPIKey) UnsetIsExpired()`

UnsetIsExpired ensures that no value is present for IsExpired, not even an explicit nil
### GetLastRotatedTimeMsecs

`func (o *UserAPIKey) GetLastRotatedTimeMsecs() int64`

GetLastRotatedTimeMsecs returns the LastRotatedTimeMsecs field if non-nil, zero value otherwise.

### GetLastRotatedTimeMsecsOk

`func (o *UserAPIKey) GetLastRotatedTimeMsecsOk() (*int64, bool)`

GetLastRotatedTimeMsecsOk returns a tuple with the LastRotatedTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRotatedTimeMsecs

`func (o *UserAPIKey) SetLastRotatedTimeMsecs(v int64)`

SetLastRotatedTimeMsecs sets LastRotatedTimeMsecs field to given value.

### HasLastRotatedTimeMsecs

`func (o *UserAPIKey) HasLastRotatedTimeMsecs() bool`

HasLastRotatedTimeMsecs returns a boolean if a field has been set.

### SetLastRotatedTimeMsecsNil

`func (o *UserAPIKey) SetLastRotatedTimeMsecsNil(b bool)`

 SetLastRotatedTimeMsecsNil sets the value for LastRotatedTimeMsecs to be an explicit nil

### UnsetLastRotatedTimeMsecs
`func (o *UserAPIKey) UnsetLastRotatedTimeMsecs()`

UnsetLastRotatedTimeMsecs ensures that no value is present for LastRotatedTimeMsecs, not even an explicit nil
### GetName

`func (o *UserAPIKey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserAPIKey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserAPIKey) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserAPIKey) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UserAPIKey) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UserAPIKey) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetUserSid

`func (o *UserAPIKey) GetUserSid() string`

GetUserSid returns the UserSid field if non-nil, zero value otherwise.

### GetUserSidOk

`func (o *UserAPIKey) GetUserSidOk() (*string, bool)`

GetUserSidOk returns a tuple with the UserSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSid

`func (o *UserAPIKey) SetUserSid(v string)`

SetUserSid sets UserSid field to given value.

### HasUserSid

`func (o *UserAPIKey) HasUserSid() bool`

HasUserSid returns a boolean if a field has been set.

### SetUserSidNil

`func (o *UserAPIKey) SetUserSidNil(b bool)`

 SetUserSidNil sets the value for UserSid to be an explicit nil

### UnsetUserSid
`func (o *UserAPIKey) UnsetUserSid()`

UnsetUserSid ensures that no value is present for UserSid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


