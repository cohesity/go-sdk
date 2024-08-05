# LockFileParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiryTimestampMsecs** | **NullableInt32** | Specifies a definite timestamp in milliseconds for retaining the file or to extend it&#39;s expiry timestamp. | 
**FilePath** | **NullableString** | Specifies the file path that needs to be locked. | 

## Methods

### NewLockFileParams

`func NewLockFileParams(expiryTimestampMsecs NullableInt32, filePath NullableString, ) *LockFileParams`

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


### SetExpiryTimestampMsecsNil

`func (o *LockFileParams) SetExpiryTimestampMsecsNil(b bool)`

 SetExpiryTimestampMsecsNil sets the value for ExpiryTimestampMsecs to be an explicit nil

### UnsetExpiryTimestampMsecs
`func (o *LockFileParams) UnsetExpiryTimestampMsecs()`

UnsetExpiryTimestampMsecs ensures that no value is present for ExpiryTimestampMsecs, not even an explicit nil
### GetFilePath

`func (o *LockFileParams) GetFilePath() string`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *LockFileParams) GetFilePathOk() (*string, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *LockFileParams) SetFilePath(v string)`

SetFilePath sets FilePath field to given value.


### SetFilePathNil

`func (o *LockFileParams) SetFilePathNil(b bool)`

 SetFilePathNil sets the value for FilePath to be an explicit nil

### UnsetFilePath
`func (o *LockFileParams) UnsetFilePath()`

UnsetFilePath ensures that no value is present for FilePath, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


