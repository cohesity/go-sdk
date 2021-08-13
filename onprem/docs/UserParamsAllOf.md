# UserParamsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **NullableString** | Specifies the sid of the User. | [optional] 
**CreatedTimeMsecs** | Pointer to **NullableInt64** | Specifies the epoch time in milliseconds when the user account was created. | [optional] [readonly] 
**LastUpdatedTimeMsecs** | Pointer to **NullableInt64** | Specifies the epoch time in milliseconds when the user account was last modified. | [optional] [readonly] 
**TenantId** | Pointer to **NullableString** | Specifies the tenant id of the User. | [optional] 
**S3AccountParams** | Pointer to [**S3AccountParams**](S3AccountParams.md) | Specifies the S3 Account parameters of the User. | [optional] 
**LockedReason** | Pointer to **NullableString** | Specifies the reason for locking the User. | [optional] [readonly] 

## Methods

### NewUserParamsAllOf

`func NewUserParamsAllOf() *UserParamsAllOf`

NewUserParamsAllOf instantiates a new UserParamsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserParamsAllOfWithDefaults

`func NewUserParamsAllOfWithDefaults() *UserParamsAllOf`

NewUserParamsAllOfWithDefaults instantiates a new UserParamsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSid

`func (o *UserParamsAllOf) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *UserParamsAllOf) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *UserParamsAllOf) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *UserParamsAllOf) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *UserParamsAllOf) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *UserParamsAllOf) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetCreatedTimeMsecs

`func (o *UserParamsAllOf) GetCreatedTimeMsecs() int64`

GetCreatedTimeMsecs returns the CreatedTimeMsecs field if non-nil, zero value otherwise.

### GetCreatedTimeMsecsOk

`func (o *UserParamsAllOf) GetCreatedTimeMsecsOk() (*int64, bool)`

GetCreatedTimeMsecsOk returns a tuple with the CreatedTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimeMsecs

`func (o *UserParamsAllOf) SetCreatedTimeMsecs(v int64)`

SetCreatedTimeMsecs sets CreatedTimeMsecs field to given value.

### HasCreatedTimeMsecs

`func (o *UserParamsAllOf) HasCreatedTimeMsecs() bool`

HasCreatedTimeMsecs returns a boolean if a field has been set.

### SetCreatedTimeMsecsNil

`func (o *UserParamsAllOf) SetCreatedTimeMsecsNil(b bool)`

 SetCreatedTimeMsecsNil sets the value for CreatedTimeMsecs to be an explicit nil

### UnsetCreatedTimeMsecs
`func (o *UserParamsAllOf) UnsetCreatedTimeMsecs()`

UnsetCreatedTimeMsecs ensures that no value is present for CreatedTimeMsecs, not even an explicit nil
### GetLastUpdatedTimeMsecs

`func (o *UserParamsAllOf) GetLastUpdatedTimeMsecs() int64`

GetLastUpdatedTimeMsecs returns the LastUpdatedTimeMsecs field if non-nil, zero value otherwise.

### GetLastUpdatedTimeMsecsOk

`func (o *UserParamsAllOf) GetLastUpdatedTimeMsecsOk() (*int64, bool)`

GetLastUpdatedTimeMsecsOk returns a tuple with the LastUpdatedTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTimeMsecs

`func (o *UserParamsAllOf) SetLastUpdatedTimeMsecs(v int64)`

SetLastUpdatedTimeMsecs sets LastUpdatedTimeMsecs field to given value.

### HasLastUpdatedTimeMsecs

`func (o *UserParamsAllOf) HasLastUpdatedTimeMsecs() bool`

HasLastUpdatedTimeMsecs returns a boolean if a field has been set.

### SetLastUpdatedTimeMsecsNil

`func (o *UserParamsAllOf) SetLastUpdatedTimeMsecsNil(b bool)`

 SetLastUpdatedTimeMsecsNil sets the value for LastUpdatedTimeMsecs to be an explicit nil

### UnsetLastUpdatedTimeMsecs
`func (o *UserParamsAllOf) UnsetLastUpdatedTimeMsecs()`

UnsetLastUpdatedTimeMsecs ensures that no value is present for LastUpdatedTimeMsecs, not even an explicit nil
### GetTenantId

`func (o *UserParamsAllOf) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *UserParamsAllOf) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *UserParamsAllOf) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *UserParamsAllOf) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *UserParamsAllOf) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *UserParamsAllOf) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetS3AccountParams

`func (o *UserParamsAllOf) GetS3AccountParams() S3AccountParams`

GetS3AccountParams returns the S3AccountParams field if non-nil, zero value otherwise.

### GetS3AccountParamsOk

`func (o *UserParamsAllOf) GetS3AccountParamsOk() (*S3AccountParams, bool)`

GetS3AccountParamsOk returns a tuple with the S3AccountParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3AccountParams

`func (o *UserParamsAllOf) SetS3AccountParams(v S3AccountParams)`

SetS3AccountParams sets S3AccountParams field to given value.

### HasS3AccountParams

`func (o *UserParamsAllOf) HasS3AccountParams() bool`

HasS3AccountParams returns a boolean if a field has been set.

### GetLockedReason

`func (o *UserParamsAllOf) GetLockedReason() string`

GetLockedReason returns the LockedReason field if non-nil, zero value otherwise.

### GetLockedReasonOk

`func (o *UserParamsAllOf) GetLockedReasonOk() (*string, bool)`

GetLockedReasonOk returns a tuple with the LockedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedReason

`func (o *UserParamsAllOf) SetLockedReason(v string)`

SetLockedReason sets LockedReason field to given value.

### HasLockedReason

`func (o *UserParamsAllOf) HasLockedReason() bool`

HasLockedReason returns a boolean if a field has been set.

### SetLockedReasonNil

`func (o *UserParamsAllOf) SetLockedReasonNil(b bool)`

 SetLockedReasonNil sets the value for LockedReason to be an explicit nil

### UnsetLockedReason
`func (o *UserParamsAllOf) UnsetLockedReason()`

UnsetLockedReason ensures that no value is present for LockedReason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


