# Chassis

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChassisNodeBase** | Pointer to **int32** | This field is initialized as sum of maximum slots of all the chassis added to the cluster so far plus one. This is required to assign unique node index for the nodes when they are added to the cluster. Please refer to cluster_node_index in Node below. | [optional] 
**HardwareModel** | Pointer to **NullableString** | Specifies the hardware model of the chassis.Like ivybridge, haswell. | [optional] 
**Id** | Pointer to **NullableInt64** | Each chassis in a cluster is assigned a unique id when the chassis is added to the cluster first time. The index starts from 1. The use of an integer id helps speed up internal computations involving chassis. This integer will not change during the lifetime of the chassis in the cluster. | [optional] 
**Location** | Pointer to **NullableString** | Location of the chassis within the rack. | [optional] 
**Name** | Pointer to **NullableString** | Unique name assigned to this chassis. This is set to the serial number of the chassis by one of the following two ways. 1) by the chassis manufacturer for non-cohesity systems, and cohesity systems built before jira ticket ECO-2 was approved. 2) by a cohesity contract manufacturer for cohesity systems built after jira ticket ECO-2 was approved. | [optional] 
**NodeIds** | Pointer to **[]int64** | Specifies list of ids of all the nodes in chassis. | [optional] 
**RackId** | Pointer to **NullableInt64** | Rack Id that this chassis belong to | [optional] 
**SerialNumber** | Pointer to **NullableString** | Specifies the serial number of the chassis. | [optional] 

## Methods

### NewChassis

`func NewChassis() *Chassis`

NewChassis instantiates a new Chassis object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChassisWithDefaults

`func NewChassisWithDefaults() *Chassis`

NewChassisWithDefaults instantiates a new Chassis object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChassisNodeBase

`func (o *Chassis) GetChassisNodeBase() int32`

GetChassisNodeBase returns the ChassisNodeBase field if non-nil, zero value otherwise.

### GetChassisNodeBaseOk

`func (o *Chassis) GetChassisNodeBaseOk() (*int32, bool)`

GetChassisNodeBaseOk returns a tuple with the ChassisNodeBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisNodeBase

`func (o *Chassis) SetChassisNodeBase(v int32)`

SetChassisNodeBase sets ChassisNodeBase field to given value.

### HasChassisNodeBase

`func (o *Chassis) HasChassisNodeBase() bool`

HasChassisNodeBase returns a boolean if a field has been set.

### GetHardwareModel

`func (o *Chassis) GetHardwareModel() string`

GetHardwareModel returns the HardwareModel field if non-nil, zero value otherwise.

### GetHardwareModelOk

`func (o *Chassis) GetHardwareModelOk() (*string, bool)`

GetHardwareModelOk returns a tuple with the HardwareModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareModel

`func (o *Chassis) SetHardwareModel(v string)`

SetHardwareModel sets HardwareModel field to given value.

### HasHardwareModel

`func (o *Chassis) HasHardwareModel() bool`

HasHardwareModel returns a boolean if a field has been set.

### SetHardwareModelNil

`func (o *Chassis) SetHardwareModelNil(b bool)`

 SetHardwareModelNil sets the value for HardwareModel to be an explicit nil

### UnsetHardwareModel
`func (o *Chassis) UnsetHardwareModel()`

UnsetHardwareModel ensures that no value is present for HardwareModel, not even an explicit nil
### GetId

`func (o *Chassis) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Chassis) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Chassis) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Chassis) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *Chassis) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *Chassis) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetLocation

`func (o *Chassis) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Chassis) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Chassis) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *Chassis) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *Chassis) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *Chassis) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetName

`func (o *Chassis) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Chassis) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Chassis) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Chassis) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Chassis) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Chassis) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetNodeIds

`func (o *Chassis) GetNodeIds() []int64`

GetNodeIds returns the NodeIds field if non-nil, zero value otherwise.

### GetNodeIdsOk

`func (o *Chassis) GetNodeIdsOk() (*[]int64, bool)`

GetNodeIdsOk returns a tuple with the NodeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeIds

`func (o *Chassis) SetNodeIds(v []int64)`

SetNodeIds sets NodeIds field to given value.

### HasNodeIds

`func (o *Chassis) HasNodeIds() bool`

HasNodeIds returns a boolean if a field has been set.

### SetNodeIdsNil

`func (o *Chassis) SetNodeIdsNil(b bool)`

 SetNodeIdsNil sets the value for NodeIds to be an explicit nil

### UnsetNodeIds
`func (o *Chassis) UnsetNodeIds()`

UnsetNodeIds ensures that no value is present for NodeIds, not even an explicit nil
### GetRackId

`func (o *Chassis) GetRackId() int64`

GetRackId returns the RackId field if non-nil, zero value otherwise.

### GetRackIdOk

`func (o *Chassis) GetRackIdOk() (*int64, bool)`

GetRackIdOk returns a tuple with the RackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRackId

`func (o *Chassis) SetRackId(v int64)`

SetRackId sets RackId field to given value.

### HasRackId

`func (o *Chassis) HasRackId() bool`

HasRackId returns a boolean if a field has been set.

### SetRackIdNil

`func (o *Chassis) SetRackIdNil(b bool)`

 SetRackIdNil sets the value for RackId to be an explicit nil

### UnsetRackId
`func (o *Chassis) UnsetRackId()`

UnsetRackId ensures that no value is present for RackId, not even an explicit nil
### GetSerialNumber

`func (o *Chassis) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *Chassis) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *Chassis) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *Chassis) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### SetSerialNumberNil

`func (o *Chassis) SetSerialNumberNil(b bool)`

 SetSerialNumberNil sets the value for SerialNumber to be an explicit nil

### UnsetSerialNumber
`func (o *Chassis) UnsetSerialNumber()`

UnsetSerialNumber ensures that no value is present for SerialNumber, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


