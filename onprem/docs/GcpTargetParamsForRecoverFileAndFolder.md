# GcpTargetParamsForRecoverFileAndFolder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetVm** | [**NullableRecoveryObjectIdentifier**](RecoveryObjectIdentifier.md) | Specifies the target VM to recover files and folders to. | 
**RecoverToOriginalPaths** | **NullableBool** | Specifies whether to recover files to original places. | 
**TargetVmCredentials** | [**NullableCredentials**](Credentials.md) | Specifies credentials to access the target VM. | 
**AlternateBaseDirectory** | Pointer to **NullableString** | Specifies a base directory under which all files and folders will be recovered. This is required if recoverToOriginalPaths is set to false. | [optional] 
**OverwriteOriginals** | Pointer to **NullableBool** | Specifies whether to override the existing files. Default is true. | [optional] 
**PreserveAttributes** | Pointer to **NullableBool** | Specifies whether to preserve original attributes. Default is true. | [optional] 
**ContinueOnError** | Pointer to **NullableBool** | Specifies whether to continue recovering other files if one of files or folders failed to recover. Default value is false. | [optional] 
**VlanConfig** | Pointer to [**NullableRecoveryVlanConfig**](RecoveryVlanConfig.md) | Specifies VLAN Params associated with the recovered files and folders. If this is not specified, then the VLAN settings will be automatically selected from one of the below options: a. If VLANs are configured on Cohesity, then the VLAN host/VIP will be automatically based on the client&#39;s (e.g. ESXI host) IP address. b. If VLANs are not configured on Cohesity, then the partition hostname or VIPs will be used for Recovery. | [optional] 

## Methods

### NewGcpTargetParamsForRecoverFileAndFolder

`func NewGcpTargetParamsForRecoverFileAndFolder(targetVm NullableRecoveryObjectIdentifier, recoverToOriginalPaths NullableBool, targetVmCredentials NullableCredentials, ) *GcpTargetParamsForRecoverFileAndFolder`

NewGcpTargetParamsForRecoverFileAndFolder instantiates a new GcpTargetParamsForRecoverFileAndFolder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpTargetParamsForRecoverFileAndFolderWithDefaults

`func NewGcpTargetParamsForRecoverFileAndFolderWithDefaults() *GcpTargetParamsForRecoverFileAndFolder`

NewGcpTargetParamsForRecoverFileAndFolderWithDefaults instantiates a new GcpTargetParamsForRecoverFileAndFolder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetVm

`func (o *GcpTargetParamsForRecoverFileAndFolder) GetTargetVm() RecoveryObjectIdentifier`

GetTargetVm returns the TargetVm field if non-nil, zero value otherwise.

### GetTargetVmOk

`func (o *GcpTargetParamsForRecoverFileAndFolder) GetTargetVmOk() (*RecoveryObjectIdentifier, bool)`

GetTargetVmOk returns a tuple with the TargetVm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetVm

`func (o *GcpTargetParamsForRecoverFileAndFolder) SetTargetVm(v RecoveryObjectIdentifier)`

SetTargetVm sets TargetVm field to given value.


### SetTargetVmNil

`func (o *GcpTargetParamsForRecoverFileAndFolder) SetTargetVmNil(b bool)`

 SetTargetVmNil sets the value for TargetVm to be an explicit nil

### UnsetTargetVm
`func (o *GcpTargetParamsForRecoverFileAndFolder) UnsetTargetVm()`

UnsetTargetVm ensures that no value is present for TargetVm, not even an explicit nil
### GetRecoverToOriginalPaths

`func (o *GcpTargetParamsForRecoverFileAndFolder) GetRecoverToOriginalPaths() bool`

GetRecoverToOriginalPaths returns the RecoverToOriginalPaths field if non-nil, zero value otherwise.

### GetRecoverToOriginalPathsOk

`func (o *GcpTargetParamsForRecoverFileAndFolder) GetRecoverToOriginalPathsOk() (*bool, bool)`

GetRecoverToOriginalPathsOk returns a tuple with the RecoverToOriginalPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverToOriginalPaths

`func (o *GcpTargetParamsForRecoverFileAndFolder) SetRecoverToOriginalPaths(v bool)`

SetRecoverToOriginalPaths sets RecoverToOriginalPaths field to given value.


### SetRecoverToOriginalPathsNil

`func (o *GcpTargetParamsForRecoverFileAndFolder) SetRecoverToOriginalPathsNil(b bool)`

 SetRecoverToOriginalPathsNil sets the value for RecoverToOriginalPaths to be an explicit nil

### UnsetRecoverToOriginalPaths
`func (o *GcpTargetParamsForRecoverFileAndFolder) UnsetRecoverToOriginalPaths()`

UnsetRecoverToOriginalPaths ensures that no value is present for RecoverToOriginalPaths, not even an explicit nil
### GetTargetVmCredentials

`func (o *GcpTargetParamsForRecoverFileAndFolder) GetTargetVmCredentials() Credentials`

GetTargetVmCredentials returns the TargetVmCredentials field if non-nil, zero value otherwise.

### GetTargetVmCredentialsOk

`func (o *GcpTargetParamsForRecoverFileAndFolder) GetTargetVmCredentialsOk() (*Credentials, bool)`

GetTargetVmCredentialsOk returns a tuple with the TargetVmCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetVmCredentials

`func (o *GcpTargetParamsForRecoverFileAndFolder) SetTargetVmCredentials(v Credentials)`

SetTargetVmCredentials sets TargetVmCredentials field to given value.


### SetTargetVmCredentialsNil

`func (o *GcpTargetParamsForRecoverFileAndFolder) SetTargetVmCredentialsNil(b bool)`

 SetTargetVmCredentialsNil sets the value for TargetVmCredentials to be an explicit nil

### UnsetTargetVmCredentials
`func (o *GcpTargetParamsForRecoverFileAndFolder) UnsetTargetVmCredentials()`

UnsetTargetVmCredentials ensures that no value is present for TargetVmCredentials, not even an explicit nil
### GetAlternateBaseDirectory

`func (o *GcpTargetParamsForRecoverFileAndFolder) GetAlternateBaseDirectory() string`

GetAlternateBaseDirectory returns the AlternateBaseDirectory field if non-nil, zero value otherwise.

### GetAlternateBaseDirectoryOk

`func (o *GcpTargetParamsForRecoverFileAndFolder) GetAlternateBaseDirectoryOk() (*string, bool)`

GetAlternateBaseDirectoryOk returns a tuple with the AlternateBaseDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateBaseDirectory

`func (o *GcpTargetParamsForRecoverFileAndFolder) SetAlternateBaseDirectory(v string)`

SetAlternateBaseDirectory sets AlternateBaseDirectory field to given value.

### HasAlternateBaseDirectory

`func (o *GcpTargetParamsForRecoverFileAndFolder) HasAlternateBaseDirectory() bool`

HasAlternateBaseDirectory returns a boolean if a field has been set.

### SetAlternateBaseDirectoryNil

`func (o *GcpTargetParamsForRecoverFileAndFolder) SetAlternateBaseDirectoryNil(b bool)`

 SetAlternateBaseDirectoryNil sets the value for AlternateBaseDirectory to be an explicit nil

### UnsetAlternateBaseDirectory
`func (o *GcpTargetParamsForRecoverFileAndFolder) UnsetAlternateBaseDirectory()`

UnsetAlternateBaseDirectory ensures that no value is present for AlternateBaseDirectory, not even an explicit nil
### GetOverwriteOriginals

`func (o *GcpTargetParamsForRecoverFileAndFolder) GetOverwriteOriginals() bool`

GetOverwriteOriginals returns the OverwriteOriginals field if non-nil, zero value otherwise.

### GetOverwriteOriginalsOk

`func (o *GcpTargetParamsForRecoverFileAndFolder) GetOverwriteOriginalsOk() (*bool, bool)`

GetOverwriteOriginalsOk returns a tuple with the OverwriteOriginals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwriteOriginals

`func (o *GcpTargetParamsForRecoverFileAndFolder) SetOverwriteOriginals(v bool)`

SetOverwriteOriginals sets OverwriteOriginals field to given value.

### HasOverwriteOriginals

`func (o *GcpTargetParamsForRecoverFileAndFolder) HasOverwriteOriginals() bool`

HasOverwriteOriginals returns a boolean if a field has been set.

### SetOverwriteOriginalsNil

`func (o *GcpTargetParamsForRecoverFileAndFolder) SetOverwriteOriginalsNil(b bool)`

 SetOverwriteOriginalsNil sets the value for OverwriteOriginals to be an explicit nil

### UnsetOverwriteOriginals
`func (o *GcpTargetParamsForRecoverFileAndFolder) UnsetOverwriteOriginals()`

UnsetOverwriteOriginals ensures that no value is present for OverwriteOriginals, not even an explicit nil
### GetPreserveAttributes

`func (o *GcpTargetParamsForRecoverFileAndFolder) GetPreserveAttributes() bool`

GetPreserveAttributes returns the PreserveAttributes field if non-nil, zero value otherwise.

### GetPreserveAttributesOk

`func (o *GcpTargetParamsForRecoverFileAndFolder) GetPreserveAttributesOk() (*bool, bool)`

GetPreserveAttributesOk returns a tuple with the PreserveAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveAttributes

`func (o *GcpTargetParamsForRecoverFileAndFolder) SetPreserveAttributes(v bool)`

SetPreserveAttributes sets PreserveAttributes field to given value.

### HasPreserveAttributes

`func (o *GcpTargetParamsForRecoverFileAndFolder) HasPreserveAttributes() bool`

HasPreserveAttributes returns a boolean if a field has been set.

### SetPreserveAttributesNil

`func (o *GcpTargetParamsForRecoverFileAndFolder) SetPreserveAttributesNil(b bool)`

 SetPreserveAttributesNil sets the value for PreserveAttributes to be an explicit nil

### UnsetPreserveAttributes
`func (o *GcpTargetParamsForRecoverFileAndFolder) UnsetPreserveAttributes()`

UnsetPreserveAttributes ensures that no value is present for PreserveAttributes, not even an explicit nil
### GetContinueOnError

`func (o *GcpTargetParamsForRecoverFileAndFolder) GetContinueOnError() bool`

GetContinueOnError returns the ContinueOnError field if non-nil, zero value otherwise.

### GetContinueOnErrorOk

`func (o *GcpTargetParamsForRecoverFileAndFolder) GetContinueOnErrorOk() (*bool, bool)`

GetContinueOnErrorOk returns a tuple with the ContinueOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnError

`func (o *GcpTargetParamsForRecoverFileAndFolder) SetContinueOnError(v bool)`

SetContinueOnError sets ContinueOnError field to given value.

### HasContinueOnError

`func (o *GcpTargetParamsForRecoverFileAndFolder) HasContinueOnError() bool`

HasContinueOnError returns a boolean if a field has been set.

### SetContinueOnErrorNil

`func (o *GcpTargetParamsForRecoverFileAndFolder) SetContinueOnErrorNil(b bool)`

 SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil

### UnsetContinueOnError
`func (o *GcpTargetParamsForRecoverFileAndFolder) UnsetContinueOnError()`

UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
### GetVlanConfig

`func (o *GcpTargetParamsForRecoverFileAndFolder) GetVlanConfig() RecoveryVlanConfig`

GetVlanConfig returns the VlanConfig field if non-nil, zero value otherwise.

### GetVlanConfigOk

`func (o *GcpTargetParamsForRecoverFileAndFolder) GetVlanConfigOk() (*RecoveryVlanConfig, bool)`

GetVlanConfigOk returns a tuple with the VlanConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanConfig

`func (o *GcpTargetParamsForRecoverFileAndFolder) SetVlanConfig(v RecoveryVlanConfig)`

SetVlanConfig sets VlanConfig field to given value.

### HasVlanConfig

`func (o *GcpTargetParamsForRecoverFileAndFolder) HasVlanConfig() bool`

HasVlanConfig returns a boolean if a field has been set.

### SetVlanConfigNil

`func (o *GcpTargetParamsForRecoverFileAndFolder) SetVlanConfigNil(b bool)`

 SetVlanConfigNil sets the value for VlanConfig to be an explicit nil

### UnsetVlanConfig
`func (o *GcpTargetParamsForRecoverFileAndFolder) UnsetVlanConfig()`

UnsetVlanConfig ensures that no value is present for VlanConfig, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


