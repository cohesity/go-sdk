# VMwareMountVolumesOriginalTargetConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BringDisksOnline** | **NullableBool** | Specifies whether the volumes need to be online within the target environment after attaching the disks. For linux VMs, this should always be set to true. For Windows, this is optional. If this is set to true, VMware tools must be installed on the VM. If this is set to false, useExistingAgent and targetCredentials are not needed. | 
**TargetVmCredentials** | Pointer to [**NullableVMwareMountVolumesOriginalTargetConfigTargetVmCredentials**](VMwareMountVolumesOriginalTargetConfigTargetVmCredentials.md) |  | [optional] 
**UseExistingAgent** | Pointer to **NullableBool** | Specifies whether this will use an existing agent on the target vm or will deploy a new agent. This is required if bringDisksOnline is set to true. | [optional] 

## Methods

### NewVMwareMountVolumesOriginalTargetConfig

`func NewVMwareMountVolumesOriginalTargetConfig(bringDisksOnline NullableBool, ) *VMwareMountVolumesOriginalTargetConfig`

NewVMwareMountVolumesOriginalTargetConfig instantiates a new VMwareMountVolumesOriginalTargetConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVMwareMountVolumesOriginalTargetConfigWithDefaults

`func NewVMwareMountVolumesOriginalTargetConfigWithDefaults() *VMwareMountVolumesOriginalTargetConfig`

NewVMwareMountVolumesOriginalTargetConfigWithDefaults instantiates a new VMwareMountVolumesOriginalTargetConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBringDisksOnline

`func (o *VMwareMountVolumesOriginalTargetConfig) GetBringDisksOnline() bool`

GetBringDisksOnline returns the BringDisksOnline field if non-nil, zero value otherwise.

### GetBringDisksOnlineOk

`func (o *VMwareMountVolumesOriginalTargetConfig) GetBringDisksOnlineOk() (*bool, bool)`

GetBringDisksOnlineOk returns a tuple with the BringDisksOnline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBringDisksOnline

`func (o *VMwareMountVolumesOriginalTargetConfig) SetBringDisksOnline(v bool)`

SetBringDisksOnline sets BringDisksOnline field to given value.


### SetBringDisksOnlineNil

`func (o *VMwareMountVolumesOriginalTargetConfig) SetBringDisksOnlineNil(b bool)`

 SetBringDisksOnlineNil sets the value for BringDisksOnline to be an explicit nil

### UnsetBringDisksOnline
`func (o *VMwareMountVolumesOriginalTargetConfig) UnsetBringDisksOnline()`

UnsetBringDisksOnline ensures that no value is present for BringDisksOnline, not even an explicit nil
### GetTargetVmCredentials

`func (o *VMwareMountVolumesOriginalTargetConfig) GetTargetVmCredentials() VMwareMountVolumesOriginalTargetConfigTargetVmCredentials`

GetTargetVmCredentials returns the TargetVmCredentials field if non-nil, zero value otherwise.

### GetTargetVmCredentialsOk

`func (o *VMwareMountVolumesOriginalTargetConfig) GetTargetVmCredentialsOk() (*VMwareMountVolumesOriginalTargetConfigTargetVmCredentials, bool)`

GetTargetVmCredentialsOk returns a tuple with the TargetVmCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetVmCredentials

`func (o *VMwareMountVolumesOriginalTargetConfig) SetTargetVmCredentials(v VMwareMountVolumesOriginalTargetConfigTargetVmCredentials)`

SetTargetVmCredentials sets TargetVmCredentials field to given value.

### HasTargetVmCredentials

`func (o *VMwareMountVolumesOriginalTargetConfig) HasTargetVmCredentials() bool`

HasTargetVmCredentials returns a boolean if a field has been set.

### SetTargetVmCredentialsNil

`func (o *VMwareMountVolumesOriginalTargetConfig) SetTargetVmCredentialsNil(b bool)`

 SetTargetVmCredentialsNil sets the value for TargetVmCredentials to be an explicit nil

### UnsetTargetVmCredentials
`func (o *VMwareMountVolumesOriginalTargetConfig) UnsetTargetVmCredentials()`

UnsetTargetVmCredentials ensures that no value is present for TargetVmCredentials, not even an explicit nil
### GetUseExistingAgent

`func (o *VMwareMountVolumesOriginalTargetConfig) GetUseExistingAgent() bool`

GetUseExistingAgent returns the UseExistingAgent field if non-nil, zero value otherwise.

### GetUseExistingAgentOk

`func (o *VMwareMountVolumesOriginalTargetConfig) GetUseExistingAgentOk() (*bool, bool)`

GetUseExistingAgentOk returns a tuple with the UseExistingAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseExistingAgent

`func (o *VMwareMountVolumesOriginalTargetConfig) SetUseExistingAgent(v bool)`

SetUseExistingAgent sets UseExistingAgent field to given value.

### HasUseExistingAgent

`func (o *VMwareMountVolumesOriginalTargetConfig) HasUseExistingAgent() bool`

HasUseExistingAgent returns a boolean if a field has been set.

### SetUseExistingAgentNil

`func (o *VMwareMountVolumesOriginalTargetConfig) SetUseExistingAgentNil(b bool)`

 SetUseExistingAgentNil sets the value for UseExistingAgent to be an explicit nil

### UnsetUseExistingAgent
`func (o *VMwareMountVolumesOriginalTargetConfig) UnsetUseExistingAgent()`

UnsetUseExistingAgent ensures that no value is present for UseExistingAgent, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


