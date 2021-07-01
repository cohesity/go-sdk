# ExpandPhysicalClusterParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChassisSerialToRackIdMap** | Pointer to **map[string]int64** | ChassisSerialToRackId map. | [optional] 
**NodeConfigs** | [**[]PhysicalNodeConfiguration**](PhysicalNodeConfiguration.md) | Specifies the configuration details of the Nodes in the Cluster. | 
**Vips** | Pointer to **[]string** | Specifies the virtual IPs to add to the Cluster. | [optional] 

## Methods

### NewExpandPhysicalClusterParameters

`func NewExpandPhysicalClusterParameters(nodeConfigs []PhysicalNodeConfiguration, ) *ExpandPhysicalClusterParameters`

NewExpandPhysicalClusterParameters instantiates a new ExpandPhysicalClusterParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpandPhysicalClusterParametersWithDefaults

`func NewExpandPhysicalClusterParametersWithDefaults() *ExpandPhysicalClusterParameters`

NewExpandPhysicalClusterParametersWithDefaults instantiates a new ExpandPhysicalClusterParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChassisSerialToRackIdMap

`func (o *ExpandPhysicalClusterParameters) GetChassisSerialToRackIdMap() map[string]int64`

GetChassisSerialToRackIdMap returns the ChassisSerialToRackIdMap field if non-nil, zero value otherwise.

### GetChassisSerialToRackIdMapOk

`func (o *ExpandPhysicalClusterParameters) GetChassisSerialToRackIdMapOk() (*map[string]int64, bool)`

GetChassisSerialToRackIdMapOk returns a tuple with the ChassisSerialToRackIdMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisSerialToRackIdMap

`func (o *ExpandPhysicalClusterParameters) SetChassisSerialToRackIdMap(v map[string]int64)`

SetChassisSerialToRackIdMap sets ChassisSerialToRackIdMap field to given value.

### HasChassisSerialToRackIdMap

`func (o *ExpandPhysicalClusterParameters) HasChassisSerialToRackIdMap() bool`

HasChassisSerialToRackIdMap returns a boolean if a field has been set.

### SetChassisSerialToRackIdMapNil

`func (o *ExpandPhysicalClusterParameters) SetChassisSerialToRackIdMapNil(b bool)`

 SetChassisSerialToRackIdMapNil sets the value for ChassisSerialToRackIdMap to be an explicit nil

### UnsetChassisSerialToRackIdMap
`func (o *ExpandPhysicalClusterParameters) UnsetChassisSerialToRackIdMap()`

UnsetChassisSerialToRackIdMap ensures that no value is present for ChassisSerialToRackIdMap, not even an explicit nil
### GetNodeConfigs

`func (o *ExpandPhysicalClusterParameters) GetNodeConfigs() []PhysicalNodeConfiguration`

GetNodeConfigs returns the NodeConfigs field if non-nil, zero value otherwise.

### GetNodeConfigsOk

`func (o *ExpandPhysicalClusterParameters) GetNodeConfigsOk() (*[]PhysicalNodeConfiguration, bool)`

GetNodeConfigsOk returns a tuple with the NodeConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeConfigs

`func (o *ExpandPhysicalClusterParameters) SetNodeConfigs(v []PhysicalNodeConfiguration)`

SetNodeConfigs sets NodeConfigs field to given value.


### SetNodeConfigsNil

`func (o *ExpandPhysicalClusterParameters) SetNodeConfigsNil(b bool)`

 SetNodeConfigsNil sets the value for NodeConfigs to be an explicit nil

### UnsetNodeConfigs
`func (o *ExpandPhysicalClusterParameters) UnsetNodeConfigs()`

UnsetNodeConfigs ensures that no value is present for NodeConfigs, not even an explicit nil
### GetVips

`func (o *ExpandPhysicalClusterParameters) GetVips() []string`

GetVips returns the Vips field if non-nil, zero value otherwise.

### GetVipsOk

`func (o *ExpandPhysicalClusterParameters) GetVipsOk() (*[]string, bool)`

GetVipsOk returns a tuple with the Vips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVips

`func (o *ExpandPhysicalClusterParameters) SetVips(v []string)`

SetVips sets Vips field to given value.

### HasVips

`func (o *ExpandPhysicalClusterParameters) HasVips() bool`

HasVips returns a boolean if a field has been set.

### SetVipsNil

`func (o *ExpandPhysicalClusterParameters) SetVipsNil(b bool)`

 SetVipsNil sets the value for Vips to be an explicit nil

### UnsetVips
`func (o *ExpandPhysicalClusterParameters) UnsetVips()`

UnsetVips ensures that no value is present for Vips, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


