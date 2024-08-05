# EncryptionSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableAdditionalSecurity** | Pointer to **NullableBool** | Enable Additional security by managing key manually | [optional] 
**EncryptionKeyFileDownloaded** | Pointer to **NullableBool** | Specifies if the encryption key file has been downloaded using the Cohesity Dashboard (Cohesity UI). If true, the encryption key has been downloaded using the Cohesity Dashboard. An encryption key can only be downloaded once using the Cohesity Dashboard. After setting it to true once, subsequent updates to this field will be ignored. | [optional] 
**EncryptionLevel** | **NullableString** | Specifies the type of encryption for the Setting. | 
**KeyFileDownloadTimeUsecs** | Pointer to **NullableInt64** | Specifies the time (in microseconds) when the encryption key file was downloaded from the Cohesity Dashboard (Cohesity UI). An encryption key can only be downloaded once using Cohesity Dashboard. Can be set only once when the key is downloaded. | [optional] [readonly] 
**KeyFileDownloadUser** | Pointer to **NullableString** | Specifies the user who downloaded the encryption key from the Cohesity Dashboard (Cohesity UI). This field is only populated if encryption is enabled for the Vault and customerManagingEncryptionKeys is true. Can be set only once when the key is downloaded. | [optional] [readonly] 
**KmsServerId** | Pointer to **NullableInt64** | Specifies the Key Management Service Server ID for the Encryption Setting. | [optional] 

## Methods

### NewEncryptionSettings

`func NewEncryptionSettings(encryptionLevel NullableString, ) *EncryptionSettings`

NewEncryptionSettings instantiates a new EncryptionSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEncryptionSettingsWithDefaults

`func NewEncryptionSettingsWithDefaults() *EncryptionSettings`

NewEncryptionSettingsWithDefaults instantiates a new EncryptionSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableAdditionalSecurity

`func (o *EncryptionSettings) GetEnableAdditionalSecurity() bool`

GetEnableAdditionalSecurity returns the EnableAdditionalSecurity field if non-nil, zero value otherwise.

### GetEnableAdditionalSecurityOk

`func (o *EncryptionSettings) GetEnableAdditionalSecurityOk() (*bool, bool)`

GetEnableAdditionalSecurityOk returns a tuple with the EnableAdditionalSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAdditionalSecurity

`func (o *EncryptionSettings) SetEnableAdditionalSecurity(v bool)`

SetEnableAdditionalSecurity sets EnableAdditionalSecurity field to given value.

### HasEnableAdditionalSecurity

`func (o *EncryptionSettings) HasEnableAdditionalSecurity() bool`

HasEnableAdditionalSecurity returns a boolean if a field has been set.

### SetEnableAdditionalSecurityNil

`func (o *EncryptionSettings) SetEnableAdditionalSecurityNil(b bool)`

 SetEnableAdditionalSecurityNil sets the value for EnableAdditionalSecurity to be an explicit nil

### UnsetEnableAdditionalSecurity
`func (o *EncryptionSettings) UnsetEnableAdditionalSecurity()`

UnsetEnableAdditionalSecurity ensures that no value is present for EnableAdditionalSecurity, not even an explicit nil
### GetEncryptionKeyFileDownloaded

`func (o *EncryptionSettings) GetEncryptionKeyFileDownloaded() bool`

GetEncryptionKeyFileDownloaded returns the EncryptionKeyFileDownloaded field if non-nil, zero value otherwise.

### GetEncryptionKeyFileDownloadedOk

`func (o *EncryptionSettings) GetEncryptionKeyFileDownloadedOk() (*bool, bool)`

GetEncryptionKeyFileDownloadedOk returns a tuple with the EncryptionKeyFileDownloaded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionKeyFileDownloaded

`func (o *EncryptionSettings) SetEncryptionKeyFileDownloaded(v bool)`

SetEncryptionKeyFileDownloaded sets EncryptionKeyFileDownloaded field to given value.

### HasEncryptionKeyFileDownloaded

`func (o *EncryptionSettings) HasEncryptionKeyFileDownloaded() bool`

HasEncryptionKeyFileDownloaded returns a boolean if a field has been set.

### SetEncryptionKeyFileDownloadedNil

`func (o *EncryptionSettings) SetEncryptionKeyFileDownloadedNil(b bool)`

 SetEncryptionKeyFileDownloadedNil sets the value for EncryptionKeyFileDownloaded to be an explicit nil

### UnsetEncryptionKeyFileDownloaded
`func (o *EncryptionSettings) UnsetEncryptionKeyFileDownloaded()`

UnsetEncryptionKeyFileDownloaded ensures that no value is present for EncryptionKeyFileDownloaded, not even an explicit nil
### GetEncryptionLevel

`func (o *EncryptionSettings) GetEncryptionLevel() string`

GetEncryptionLevel returns the EncryptionLevel field if non-nil, zero value otherwise.

### GetEncryptionLevelOk

`func (o *EncryptionSettings) GetEncryptionLevelOk() (*string, bool)`

GetEncryptionLevelOk returns a tuple with the EncryptionLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionLevel

`func (o *EncryptionSettings) SetEncryptionLevel(v string)`

SetEncryptionLevel sets EncryptionLevel field to given value.


### SetEncryptionLevelNil

`func (o *EncryptionSettings) SetEncryptionLevelNil(b bool)`

 SetEncryptionLevelNil sets the value for EncryptionLevel to be an explicit nil

### UnsetEncryptionLevel
`func (o *EncryptionSettings) UnsetEncryptionLevel()`

UnsetEncryptionLevel ensures that no value is present for EncryptionLevel, not even an explicit nil
### GetKeyFileDownloadTimeUsecs

`func (o *EncryptionSettings) GetKeyFileDownloadTimeUsecs() int64`

GetKeyFileDownloadTimeUsecs returns the KeyFileDownloadTimeUsecs field if non-nil, zero value otherwise.

### GetKeyFileDownloadTimeUsecsOk

`func (o *EncryptionSettings) GetKeyFileDownloadTimeUsecsOk() (*int64, bool)`

GetKeyFileDownloadTimeUsecsOk returns a tuple with the KeyFileDownloadTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyFileDownloadTimeUsecs

`func (o *EncryptionSettings) SetKeyFileDownloadTimeUsecs(v int64)`

SetKeyFileDownloadTimeUsecs sets KeyFileDownloadTimeUsecs field to given value.

### HasKeyFileDownloadTimeUsecs

`func (o *EncryptionSettings) HasKeyFileDownloadTimeUsecs() bool`

HasKeyFileDownloadTimeUsecs returns a boolean if a field has been set.

### SetKeyFileDownloadTimeUsecsNil

`func (o *EncryptionSettings) SetKeyFileDownloadTimeUsecsNil(b bool)`

 SetKeyFileDownloadTimeUsecsNil sets the value for KeyFileDownloadTimeUsecs to be an explicit nil

### UnsetKeyFileDownloadTimeUsecs
`func (o *EncryptionSettings) UnsetKeyFileDownloadTimeUsecs()`

UnsetKeyFileDownloadTimeUsecs ensures that no value is present for KeyFileDownloadTimeUsecs, not even an explicit nil
### GetKeyFileDownloadUser

`func (o *EncryptionSettings) GetKeyFileDownloadUser() string`

GetKeyFileDownloadUser returns the KeyFileDownloadUser field if non-nil, zero value otherwise.

### GetKeyFileDownloadUserOk

`func (o *EncryptionSettings) GetKeyFileDownloadUserOk() (*string, bool)`

GetKeyFileDownloadUserOk returns a tuple with the KeyFileDownloadUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyFileDownloadUser

`func (o *EncryptionSettings) SetKeyFileDownloadUser(v string)`

SetKeyFileDownloadUser sets KeyFileDownloadUser field to given value.

### HasKeyFileDownloadUser

`func (o *EncryptionSettings) HasKeyFileDownloadUser() bool`

HasKeyFileDownloadUser returns a boolean if a field has been set.

### SetKeyFileDownloadUserNil

`func (o *EncryptionSettings) SetKeyFileDownloadUserNil(b bool)`

 SetKeyFileDownloadUserNil sets the value for KeyFileDownloadUser to be an explicit nil

### UnsetKeyFileDownloadUser
`func (o *EncryptionSettings) UnsetKeyFileDownloadUser()`

UnsetKeyFileDownloadUser ensures that no value is present for KeyFileDownloadUser, not even an explicit nil
### GetKmsServerId

`func (o *EncryptionSettings) GetKmsServerId() int64`

GetKmsServerId returns the KmsServerId field if non-nil, zero value otherwise.

### GetKmsServerIdOk

`func (o *EncryptionSettings) GetKmsServerIdOk() (*int64, bool)`

GetKmsServerIdOk returns a tuple with the KmsServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKmsServerId

`func (o *EncryptionSettings) SetKmsServerId(v int64)`

SetKmsServerId sets KmsServerId field to given value.

### HasKmsServerId

`func (o *EncryptionSettings) HasKmsServerId() bool`

HasKmsServerId returns a boolean if a field has been set.

### SetKmsServerIdNil

`func (o *EncryptionSettings) SetKmsServerIdNil(b bool)`

 SetKmsServerIdNil sets the value for KmsServerId to be an explicit nil

### UnsetKmsServerId
`func (o *EncryptionSettings) UnsetKmsServerId()`

UnsetKmsServerId ensures that no value is present for KmsServerId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


