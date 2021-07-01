# VMwareBackupSourceParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceAppParams** | Pointer to [**SourceAppParams**](SourceAppParams.md) |  | [optional] 
**VmCredentials** | Pointer to [**Credentials**](Credentials.md) |  | [optional] 
**VmwareDiskExclusionInfo** | Pointer to [**[]VMwareDiskExclusionProto**](VMwareDiskExclusionProto.md) | List of Virtual Disk(s) to be excluded from the backup job for the source. Overrides the exclusion list requested (if any) through EnvBackupParams.VMwareBackupEnvParams. | [optional] 

## Methods

### NewVMwareBackupSourceParams

`func NewVMwareBackupSourceParams() *VMwareBackupSourceParams`

NewVMwareBackupSourceParams instantiates a new VMwareBackupSourceParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVMwareBackupSourceParamsWithDefaults

`func NewVMwareBackupSourceParamsWithDefaults() *VMwareBackupSourceParams`

NewVMwareBackupSourceParamsWithDefaults instantiates a new VMwareBackupSourceParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceAppParams

`func (o *VMwareBackupSourceParams) GetSourceAppParams() SourceAppParams`

GetSourceAppParams returns the SourceAppParams field if non-nil, zero value otherwise.

### GetSourceAppParamsOk

`func (o *VMwareBackupSourceParams) GetSourceAppParamsOk() (*SourceAppParams, bool)`

GetSourceAppParamsOk returns a tuple with the SourceAppParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAppParams

`func (o *VMwareBackupSourceParams) SetSourceAppParams(v SourceAppParams)`

SetSourceAppParams sets SourceAppParams field to given value.

### HasSourceAppParams

`func (o *VMwareBackupSourceParams) HasSourceAppParams() bool`

HasSourceAppParams returns a boolean if a field has been set.

### GetVmCredentials

`func (o *VMwareBackupSourceParams) GetVmCredentials() Credentials`

GetVmCredentials returns the VmCredentials field if non-nil, zero value otherwise.

### GetVmCredentialsOk

`func (o *VMwareBackupSourceParams) GetVmCredentialsOk() (*Credentials, bool)`

GetVmCredentialsOk returns a tuple with the VmCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmCredentials

`func (o *VMwareBackupSourceParams) SetVmCredentials(v Credentials)`

SetVmCredentials sets VmCredentials field to given value.

### HasVmCredentials

`func (o *VMwareBackupSourceParams) HasVmCredentials() bool`

HasVmCredentials returns a boolean if a field has been set.

### GetVmwareDiskExclusionInfo

`func (o *VMwareBackupSourceParams) GetVmwareDiskExclusionInfo() []VMwareDiskExclusionProto`

GetVmwareDiskExclusionInfo returns the VmwareDiskExclusionInfo field if non-nil, zero value otherwise.

### GetVmwareDiskExclusionInfoOk

`func (o *VMwareBackupSourceParams) GetVmwareDiskExclusionInfoOk() (*[]VMwareDiskExclusionProto, bool)`

GetVmwareDiskExclusionInfoOk returns a tuple with the VmwareDiskExclusionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmwareDiskExclusionInfo

`func (o *VMwareBackupSourceParams) SetVmwareDiskExclusionInfo(v []VMwareDiskExclusionProto)`

SetVmwareDiskExclusionInfo sets VmwareDiskExclusionInfo field to given value.

### HasVmwareDiskExclusionInfo

`func (o *VMwareBackupSourceParams) HasVmwareDiskExclusionInfo() bool`

HasVmwareDiskExclusionInfo returns a boolean if a field has been set.

### SetVmwareDiskExclusionInfoNil

`func (o *VMwareBackupSourceParams) SetVmwareDiskExclusionInfoNil(b bool)`

 SetVmwareDiskExclusionInfoNil sets the value for VmwareDiskExclusionInfo to be an explicit nil

### UnsetVmwareDiskExclusionInfo
`func (o *VMwareBackupSourceParams) UnsetVmwareDiskExclusionInfo()`

UnsetVmwareDiskExclusionInfo ensures that no value is present for VmwareDiskExclusionInfo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


