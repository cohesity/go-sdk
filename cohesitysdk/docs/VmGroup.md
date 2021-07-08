# VmGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Specifies name of the VM group. | [optional] 
**Vms** | Pointer to [**[]VmInfo**](VmInfo.md) | Specifies VMs in the group. | [optional] 

## Methods

### NewVmGroup

`func NewVmGroup() *VmGroup`

NewVmGroup instantiates a new VmGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmGroupWithDefaults

`func NewVmGroupWithDefaults() *VmGroup`

NewVmGroupWithDefaults instantiates a new VmGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VmGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VmGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VmGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VmGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *VmGroup) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *VmGroup) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetVms

`func (o *VmGroup) GetVms() []VmInfo`

GetVms returns the Vms field if non-nil, zero value otherwise.

### GetVmsOk

`func (o *VmGroup) GetVmsOk() (*[]VmInfo, bool)`

GetVmsOk returns a tuple with the Vms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVms

`func (o *VmGroup) SetVms(v []VmInfo)`

SetVms sets Vms field to given value.

### HasVms

`func (o *VmGroup) HasVms() bool`

HasVms returns a boolean if a field has been set.

### SetVmsNil

`func (o *VmGroup) SetVmsNil(b bool)`

 SetVmsNil sets the value for Vms to be an explicit nil

### UnsetVms
`func (o *VmGroup) UnsetVms()`

UnsetVms ensures that no value is present for Vms, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


