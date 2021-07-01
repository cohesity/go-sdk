# LockFileParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiryTimestampMsecs** | Pointer to **NullableInt32** | Specifies a definite timestamp in milliseconds for retaining the file, or to extend it&#39;s expiry timestamp. | [optional] 
**Path** | Pointer to **NullableString** | Specifies the file path that needs to be locked. | [optional] 

## Methods

### NewLockFileParams

`func NewLockFileParams() *LockFileParams`

NewLockFileParams instantiates a new LockFileParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLockFileParamsWithDefaults

`func NewLockFileParamsWithDefaults() *LockFileParams`

NewLockFileParamsWithDefaults instantiates a new LockFileParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiryTimestampMsecs

`func (o *LockFileParams) GetExpiryTimestampMsecs() int32`

GetExpiryTimestampMsecs returns the ExpiryTimestampMsecs field if non-nil, zero value otherwise.

### GetExpiryTimestampMsecsOk

`func (o *LockFileParams) GetExpiryTimestampMsecsOk() (*int32, bool)`

GetExpiryTimestampMsecsOk returns a tuple with the ExpiryTimestampMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryTimestampMsecs

`func (o *LockFileParams) SetExpiryTimestampMsecs(v int32)`

SetExpiryTimestampMsecs sets ExpiryTimestampMsecs field to given value.

### HasExpiryTimestampMsecs

`func (o *LockFileParams) HasExpiryTimestampMsecs() bool`

HasExpiryTimestampMsecs returns a boolean if a field has been set.

### SetExpiryTimestampMsecsNil

`func (o *LockFileParams) SetExpiryTimestampMsecsNil(b bool)`

 SetExpiryTimestampMsecsNil sets the value for ExpiryTimestampMsecs to be an explicit nil

### UnsetExpiryTimestampMsecs
`func (o *LockFileParams) UnsetExpiryTimestampMsecs()`

UnsetExpiryTimestampMsecs ensures that no value is present for ExpiryTimestampMsecs, not even an explicit nil
### GetPath

`func (o *LockFileParams) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *LockFileParams) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *LockFileParams) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *LockFileParams) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *LockFileParams) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *LockFileParams) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


