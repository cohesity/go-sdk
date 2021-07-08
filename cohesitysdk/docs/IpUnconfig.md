# IpUnconfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpFamily** | Pointer to **NullableInt32** | IpFamily of this config. | [optional] 
**InterfaceName** | Pointer to **NullableString** | The interface name. | [optional] 
**NodeIds** | Pointer to **[]int64** | Node ids. | [optional] 

## Methods

### NewIpUnconfig

`func NewIpUnconfig() *IpUnconfig`

NewIpUnconfig instantiates a new IpUnconfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpUnconfigWithDefaults

`func NewIpUnconfigWithDefaults() *IpUnconfig`

NewIpUnconfigWithDefaults instantiates a new IpUnconfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpFamily

`func (o *IpUnconfig) GetIpFamily() int32`

GetIpFamily returns the IpFamily field if non-nil, zero value otherwise.

### GetIpFamilyOk

`func (o *IpUnconfig) GetIpFamilyOk() (*int32, bool)`

GetIpFamilyOk returns a tuple with the IpFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpFamily

`func (o *IpUnconfig) SetIpFamily(v int32)`

SetIpFamily sets IpFamily field to given value.

### HasIpFamily

`func (o *IpUnconfig) HasIpFamily() bool`

HasIpFamily returns a boolean if a field has been set.

### SetIpFamilyNil

`func (o *IpUnconfig) SetIpFamilyNil(b bool)`

 SetIpFamilyNil sets the value for IpFamily to be an explicit nil

### UnsetIpFamily
`func (o *IpUnconfig) UnsetIpFamily()`

UnsetIpFamily ensures that no value is present for IpFamily, not even an explicit nil
### GetInterfaceName

`func (o *IpUnconfig) GetInterfaceName() string`

GetInterfaceName returns the InterfaceName field if non-nil, zero value otherwise.

### GetInterfaceNameOk

`func (o *IpUnconfig) GetInterfaceNameOk() (*string, bool)`

GetInterfaceNameOk returns a tuple with the InterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceName

`func (o *IpUnconfig) SetInterfaceName(v string)`

SetInterfaceName sets InterfaceName field to given value.

### HasInterfaceName

`func (o *IpUnconfig) HasInterfaceName() bool`

HasInterfaceName returns a boolean if a field has been set.

### SetInterfaceNameNil

`func (o *IpUnconfig) SetInterfaceNameNil(b bool)`

 SetInterfaceNameNil sets the value for InterfaceName to be an explicit nil

### UnsetInterfaceName
`func (o *IpUnconfig) UnsetInterfaceName()`

UnsetInterfaceName ensures that no value is present for InterfaceName, not even an explicit nil
### GetNodeIds

`func (o *IpUnconfig) GetNodeIds() []int64`

GetNodeIds returns the NodeIds field if non-nil, zero value otherwise.

### GetNodeIdsOk

`func (o *IpUnconfig) GetNodeIdsOk() (*[]int64, bool)`

GetNodeIdsOk returns a tuple with the NodeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeIds

`func (o *IpUnconfig) SetNodeIds(v []int64)`

SetNodeIds sets NodeIds field to given value.

### HasNodeIds

`func (o *IpUnconfig) HasNodeIds() bool`

HasNodeIds returns a boolean if a field has been set.

### SetNodeIdsNil

`func (o *IpUnconfig) SetNodeIdsNil(b bool)`

 SetNodeIdsNil sets the value for NodeIds to be an explicit nil

### UnsetNodeIds
`func (o *IpUnconfig) UnsetNodeIds()`

UnsetNodeIds ensures that no value is present for NodeIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


