# RestoreADAppObjectParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdRestoreStatusVec** | Pointer to [**[]ADRestoreStatus**](ADRestoreStatus.md) | Status of the AD object/attribute restore operation. | [optional] 
**AdUpdateOptions** | Pointer to [**ADUpdateRestoreTaskOptions**](ADUpdateRestoreTaskOptions.md) |  | [optional] 
**Credentials** | Pointer to [**Credentials**](Credentials.md) |  | [optional] 
**LdapPort** | Pointer to **NullableInt32** | The ldap port on which the AD domain controller&#39;s NTDS database will be mounted. | [optional] 
**NumFailed** | Pointer to **NullableInt32** | Number of AD objects whose restore failed. Includes both AD object and attribute restored. | [optional] 
**NumRunning** | Pointer to **NullableInt32** | Number of AD objects whose restores are currently running. Includes both AD object and attribute recoveries. | [optional] 
**NumSuccessfull** | Pointer to **NullableInt32** | Number of AD objects restored successfully. Includes both AD object and attribute restored. | [optional] 
**ShouldMountAndRestore** | Pointer to **NullableBool** | The following field is set if user wants to mount AD, restore AD objects and destory AD mount in single task. | [optional] 

## Methods

### NewRestoreADAppObjectParams

`func NewRestoreADAppObjectParams() *RestoreADAppObjectParams`

NewRestoreADAppObjectParams instantiates a new RestoreADAppObjectParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreADAppObjectParamsWithDefaults

`func NewRestoreADAppObjectParamsWithDefaults() *RestoreADAppObjectParams`

NewRestoreADAppObjectParamsWithDefaults instantiates a new RestoreADAppObjectParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdRestoreStatusVec

`func (o *RestoreADAppObjectParams) GetAdRestoreStatusVec() []ADRestoreStatus`

GetAdRestoreStatusVec returns the AdRestoreStatusVec field if non-nil, zero value otherwise.

### GetAdRestoreStatusVecOk

`func (o *RestoreADAppObjectParams) GetAdRestoreStatusVecOk() (*[]ADRestoreStatus, bool)`

GetAdRestoreStatusVecOk returns a tuple with the AdRestoreStatusVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdRestoreStatusVec

`func (o *RestoreADAppObjectParams) SetAdRestoreStatusVec(v []ADRestoreStatus)`

SetAdRestoreStatusVec sets AdRestoreStatusVec field to given value.

### HasAdRestoreStatusVec

`func (o *RestoreADAppObjectParams) HasAdRestoreStatusVec() bool`

HasAdRestoreStatusVec returns a boolean if a field has been set.

### SetAdRestoreStatusVecNil

`func (o *RestoreADAppObjectParams) SetAdRestoreStatusVecNil(b bool)`

 SetAdRestoreStatusVecNil sets the value for AdRestoreStatusVec to be an explicit nil

### UnsetAdRestoreStatusVec
`func (o *RestoreADAppObjectParams) UnsetAdRestoreStatusVec()`

UnsetAdRestoreStatusVec ensures that no value is present for AdRestoreStatusVec, not even an explicit nil
### GetAdUpdateOptions

`func (o *RestoreADAppObjectParams) GetAdUpdateOptions() ADUpdateRestoreTaskOptions`

GetAdUpdateOptions returns the AdUpdateOptions field if non-nil, zero value otherwise.

### GetAdUpdateOptionsOk

`func (o *RestoreADAppObjectParams) GetAdUpdateOptionsOk() (*ADUpdateRestoreTaskOptions, bool)`

GetAdUpdateOptionsOk returns a tuple with the AdUpdateOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdUpdateOptions

`func (o *RestoreADAppObjectParams) SetAdUpdateOptions(v ADUpdateRestoreTaskOptions)`

SetAdUpdateOptions sets AdUpdateOptions field to given value.

### HasAdUpdateOptions

`func (o *RestoreADAppObjectParams) HasAdUpdateOptions() bool`

HasAdUpdateOptions returns a boolean if a field has been set.

### GetCredentials

`func (o *RestoreADAppObjectParams) GetCredentials() Credentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *RestoreADAppObjectParams) GetCredentialsOk() (*Credentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *RestoreADAppObjectParams) SetCredentials(v Credentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *RestoreADAppObjectParams) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetLdapPort

`func (o *RestoreADAppObjectParams) GetLdapPort() int32`

GetLdapPort returns the LdapPort field if non-nil, zero value otherwise.

### GetLdapPortOk

`func (o *RestoreADAppObjectParams) GetLdapPortOk() (*int32, bool)`

GetLdapPortOk returns a tuple with the LdapPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapPort

`func (o *RestoreADAppObjectParams) SetLdapPort(v int32)`

SetLdapPort sets LdapPort field to given value.

### HasLdapPort

`func (o *RestoreADAppObjectParams) HasLdapPort() bool`

HasLdapPort returns a boolean if a field has been set.

### SetLdapPortNil

`func (o *RestoreADAppObjectParams) SetLdapPortNil(b bool)`

 SetLdapPortNil sets the value for LdapPort to be an explicit nil

### UnsetLdapPort
`func (o *RestoreADAppObjectParams) UnsetLdapPort()`

UnsetLdapPort ensures that no value is present for LdapPort, not even an explicit nil
### GetNumFailed

`func (o *RestoreADAppObjectParams) GetNumFailed() int32`

GetNumFailed returns the NumFailed field if non-nil, zero value otherwise.

### GetNumFailedOk

`func (o *RestoreADAppObjectParams) GetNumFailedOk() (*int32, bool)`

GetNumFailedOk returns a tuple with the NumFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumFailed

`func (o *RestoreADAppObjectParams) SetNumFailed(v int32)`

SetNumFailed sets NumFailed field to given value.

### HasNumFailed

`func (o *RestoreADAppObjectParams) HasNumFailed() bool`

HasNumFailed returns a boolean if a field has been set.

### SetNumFailedNil

`func (o *RestoreADAppObjectParams) SetNumFailedNil(b bool)`

 SetNumFailedNil sets the value for NumFailed to be an explicit nil

### UnsetNumFailed
`func (o *RestoreADAppObjectParams) UnsetNumFailed()`

UnsetNumFailed ensures that no value is present for NumFailed, not even an explicit nil
### GetNumRunning

`func (o *RestoreADAppObjectParams) GetNumRunning() int32`

GetNumRunning returns the NumRunning field if non-nil, zero value otherwise.

### GetNumRunningOk

`func (o *RestoreADAppObjectParams) GetNumRunningOk() (*int32, bool)`

GetNumRunningOk returns a tuple with the NumRunning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumRunning

`func (o *RestoreADAppObjectParams) SetNumRunning(v int32)`

SetNumRunning sets NumRunning field to given value.

### HasNumRunning

`func (o *RestoreADAppObjectParams) HasNumRunning() bool`

HasNumRunning returns a boolean if a field has been set.

### SetNumRunningNil

`func (o *RestoreADAppObjectParams) SetNumRunningNil(b bool)`

 SetNumRunningNil sets the value for NumRunning to be an explicit nil

### UnsetNumRunning
`func (o *RestoreADAppObjectParams) UnsetNumRunning()`

UnsetNumRunning ensures that no value is present for NumRunning, not even an explicit nil
### GetNumSuccessfull

`func (o *RestoreADAppObjectParams) GetNumSuccessfull() int32`

GetNumSuccessfull returns the NumSuccessfull field if non-nil, zero value otherwise.

### GetNumSuccessfullOk

`func (o *RestoreADAppObjectParams) GetNumSuccessfullOk() (*int32, bool)`

GetNumSuccessfullOk returns a tuple with the NumSuccessfull field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumSuccessfull

`func (o *RestoreADAppObjectParams) SetNumSuccessfull(v int32)`

SetNumSuccessfull sets NumSuccessfull field to given value.

### HasNumSuccessfull

`func (o *RestoreADAppObjectParams) HasNumSuccessfull() bool`

HasNumSuccessfull returns a boolean if a field has been set.

### SetNumSuccessfullNil

`func (o *RestoreADAppObjectParams) SetNumSuccessfullNil(b bool)`

 SetNumSuccessfullNil sets the value for NumSuccessfull to be an explicit nil

### UnsetNumSuccessfull
`func (o *RestoreADAppObjectParams) UnsetNumSuccessfull()`

UnsetNumSuccessfull ensures that no value is present for NumSuccessfull, not even an explicit nil
### GetShouldMountAndRestore

`func (o *RestoreADAppObjectParams) GetShouldMountAndRestore() bool`

GetShouldMountAndRestore returns the ShouldMountAndRestore field if non-nil, zero value otherwise.

### GetShouldMountAndRestoreOk

`func (o *RestoreADAppObjectParams) GetShouldMountAndRestoreOk() (*bool, bool)`

GetShouldMountAndRestoreOk returns a tuple with the ShouldMountAndRestore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldMountAndRestore

`func (o *RestoreADAppObjectParams) SetShouldMountAndRestore(v bool)`

SetShouldMountAndRestore sets ShouldMountAndRestore field to given value.

### HasShouldMountAndRestore

`func (o *RestoreADAppObjectParams) HasShouldMountAndRestore() bool`

HasShouldMountAndRestore returns a boolean if a field has been set.

### SetShouldMountAndRestoreNil

`func (o *RestoreADAppObjectParams) SetShouldMountAndRestoreNil(b bool)`

 SetShouldMountAndRestoreNil sets the value for ShouldMountAndRestore to be an explicit nil

### UnsetShouldMountAndRestore
`func (o *RestoreADAppObjectParams) UnsetShouldMountAndRestore()`

UnsetShouldMountAndRestore ensures that no value is present for ShouldMountAndRestore, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


