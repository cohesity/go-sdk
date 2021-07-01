# VmInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HealthDetail** | Pointer to **NullableString** | Specifies the reason if vm is unhealthy. | [optional] 
**HealthStatus** | Pointer to **NullableInt32** | Specifies the current health status of the app instance. | [optional] 
**Name** | Pointer to **NullableString** | Specifies name of the VM. | [optional] 
**NodePorts** | Pointer to [**[]NodePort**](NodePort.md) | Specifies nodeports assigned to the vm. | [optional] 

## Methods

### NewVmInfo

`func NewVmInfo() *VmInfo`

NewVmInfo instantiates a new VmInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmInfoWithDefaults

`func NewVmInfoWithDefaults() *VmInfo`

NewVmInfoWithDefaults instantiates a new VmInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHealthDetail

`func (o *VmInfo) GetHealthDetail() string`

GetHealthDetail returns the HealthDetail field if non-nil, zero value otherwise.

### GetHealthDetailOk

`func (o *VmInfo) GetHealthDetailOk() (*string, bool)`

GetHealthDetailOk returns a tuple with the HealthDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthDetail

`func (o *VmInfo) SetHealthDetail(v string)`

SetHealthDetail sets HealthDetail field to given value.

### HasHealthDetail

`func (o *VmInfo) HasHealthDetail() bool`

HasHealthDetail returns a boolean if a field has been set.

### SetHealthDetailNil

`func (o *VmInfo) SetHealthDetailNil(b bool)`

 SetHealthDetailNil sets the value for HealthDetail to be an explicit nil

### UnsetHealthDetail
`func (o *VmInfo) UnsetHealthDetail()`

UnsetHealthDetail ensures that no value is present for HealthDetail, not even an explicit nil
### GetHealthStatus

`func (o *VmInfo) GetHealthStatus() int32`

GetHealthStatus returns the HealthStatus field if non-nil, zero value otherwise.

### GetHealthStatusOk

`func (o *VmInfo) GetHealthStatusOk() (*int32, bool)`

GetHealthStatusOk returns a tuple with the HealthStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthStatus

`func (o *VmInfo) SetHealthStatus(v int32)`

SetHealthStatus sets HealthStatus field to given value.

### HasHealthStatus

`func (o *VmInfo) HasHealthStatus() bool`

HasHealthStatus returns a boolean if a field has been set.

### SetHealthStatusNil

`func (o *VmInfo) SetHealthStatusNil(b bool)`

 SetHealthStatusNil sets the value for HealthStatus to be an explicit nil

### UnsetHealthStatus
`func (o *VmInfo) UnsetHealthStatus()`

UnsetHealthStatus ensures that no value is present for HealthStatus, not even an explicit nil
### GetName

`func (o *VmInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VmInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VmInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VmInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *VmInfo) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *VmInfo) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetNodePorts

`func (o *VmInfo) GetNodePorts() []NodePort`

GetNodePorts returns the NodePorts field if non-nil, zero value otherwise.

### GetNodePortsOk

`func (o *VmInfo) GetNodePortsOk() (*[]NodePort, bool)`

GetNodePortsOk returns a tuple with the NodePorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodePorts

`func (o *VmInfo) SetNodePorts(v []NodePort)`

SetNodePorts sets NodePorts field to given value.

### HasNodePorts

`func (o *VmInfo) HasNodePorts() bool`

HasNodePorts returns a boolean if a field has been set.

### SetNodePortsNil

`func (o *VmInfo) SetNodePortsNil(b bool)`

 SetNodePortsNil sets the value for NodePorts to be an explicit nil

### UnsetNodePorts
`func (o *VmInfo) UnsetNodePorts()`

UnsetNodePorts ensures that no value is present for NodePorts, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


