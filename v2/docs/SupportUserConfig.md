# SupportUserConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableSudoAccess** | Pointer to **NullableBool** | Specifies if the support user has sudo access. | [optional] 
**PasswordSet** | Pointer to **NullableBool** | Specifies if the password for the support user has been set. | [optional] 
**SudoAccessEndTimestampMsecs** | Pointer to **NullableInt32** | Specifies the sudo access end time stamp in milliseconds since unix epoch. | [optional] 
**SudoAccessMode** | Pointer to **NullableInt64** | Specifies whether the sudo access mode is enabled or not | [optional] 

## Methods

### NewSupportUserConfig

`func NewSupportUserConfig() *SupportUserConfig`

NewSupportUserConfig instantiates a new SupportUserConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportUserConfigWithDefaults

`func NewSupportUserConfigWithDefaults() *SupportUserConfig`

NewSupportUserConfigWithDefaults instantiates a new SupportUserConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableSudoAccess

`func (o *SupportUserConfig) GetEnableSudoAccess() bool`

GetEnableSudoAccess returns the EnableSudoAccess field if non-nil, zero value otherwise.

### GetEnableSudoAccessOk

`func (o *SupportUserConfig) GetEnableSudoAccessOk() (*bool, bool)`

GetEnableSudoAccessOk returns a tuple with the EnableSudoAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSudoAccess

`func (o *SupportUserConfig) SetEnableSudoAccess(v bool)`

SetEnableSudoAccess sets EnableSudoAccess field to given value.

### HasEnableSudoAccess

`func (o *SupportUserConfig) HasEnableSudoAccess() bool`

HasEnableSudoAccess returns a boolean if a field has been set.

### SetEnableSudoAccessNil

`func (o *SupportUserConfig) SetEnableSudoAccessNil(b bool)`

 SetEnableSudoAccessNil sets the value for EnableSudoAccess to be an explicit nil

### UnsetEnableSudoAccess
`func (o *SupportUserConfig) UnsetEnableSudoAccess()`

UnsetEnableSudoAccess ensures that no value is present for EnableSudoAccess, not even an explicit nil
### GetPasswordSet

`func (o *SupportUserConfig) GetPasswordSet() bool`

GetPasswordSet returns the PasswordSet field if non-nil, zero value otherwise.

### GetPasswordSetOk

`func (o *SupportUserConfig) GetPasswordSetOk() (*bool, bool)`

GetPasswordSetOk returns a tuple with the PasswordSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordSet

`func (o *SupportUserConfig) SetPasswordSet(v bool)`

SetPasswordSet sets PasswordSet field to given value.

### HasPasswordSet

`func (o *SupportUserConfig) HasPasswordSet() bool`

HasPasswordSet returns a boolean if a field has been set.

### SetPasswordSetNil

`func (o *SupportUserConfig) SetPasswordSetNil(b bool)`

 SetPasswordSetNil sets the value for PasswordSet to be an explicit nil

### UnsetPasswordSet
`func (o *SupportUserConfig) UnsetPasswordSet()`

UnsetPasswordSet ensures that no value is present for PasswordSet, not even an explicit nil
### GetSudoAccessEndTimestampMsecs

`func (o *SupportUserConfig) GetSudoAccessEndTimestampMsecs() int32`

GetSudoAccessEndTimestampMsecs returns the SudoAccessEndTimestampMsecs field if non-nil, zero value otherwise.

### GetSudoAccessEndTimestampMsecsOk

`func (o *SupportUserConfig) GetSudoAccessEndTimestampMsecsOk() (*int32, bool)`

GetSudoAccessEndTimestampMsecsOk returns a tuple with the SudoAccessEndTimestampMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSudoAccessEndTimestampMsecs

`func (o *SupportUserConfig) SetSudoAccessEndTimestampMsecs(v int32)`

SetSudoAccessEndTimestampMsecs sets SudoAccessEndTimestampMsecs field to given value.

### HasSudoAccessEndTimestampMsecs

`func (o *SupportUserConfig) HasSudoAccessEndTimestampMsecs() bool`

HasSudoAccessEndTimestampMsecs returns a boolean if a field has been set.

### SetSudoAccessEndTimestampMsecsNil

`func (o *SupportUserConfig) SetSudoAccessEndTimestampMsecsNil(b bool)`

 SetSudoAccessEndTimestampMsecsNil sets the value for SudoAccessEndTimestampMsecs to be an explicit nil

### UnsetSudoAccessEndTimestampMsecs
`func (o *SupportUserConfig) UnsetSudoAccessEndTimestampMsecs()`

UnsetSudoAccessEndTimestampMsecs ensures that no value is present for SudoAccessEndTimestampMsecs, not even an explicit nil
### GetSudoAccessMode

`func (o *SupportUserConfig) GetSudoAccessMode() int64`

GetSudoAccessMode returns the SudoAccessMode field if non-nil, zero value otherwise.

### GetSudoAccessModeOk

`func (o *SupportUserConfig) GetSudoAccessModeOk() (*int64, bool)`

GetSudoAccessModeOk returns a tuple with the SudoAccessMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSudoAccessMode

`func (o *SupportUserConfig) SetSudoAccessMode(v int64)`

SetSudoAccessMode sets SudoAccessMode field to given value.

### HasSudoAccessMode

`func (o *SupportUserConfig) HasSudoAccessMode() bool`

HasSudoAccessMode returns a boolean if a field has been set.

### SetSudoAccessModeNil

`func (o *SupportUserConfig) SetSudoAccessModeNil(b bool)`

 SetSudoAccessModeNil sets the value for SudoAccessMode to be an explicit nil

### UnsetSudoAccessMode
`func (o *SupportUserConfig) UnsetSudoAccessMode()`

UnsetSudoAccessMode ensures that no value is present for SudoAccessMode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


