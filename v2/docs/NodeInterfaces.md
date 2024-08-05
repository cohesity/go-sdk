# NodeInterfaces

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChassisSerial** | Pointer to **NullableString** | Specifies the ip of the node. | [optional] 
**Id** | Pointer to **NullableInt64** | Specifies the id of the node. | [optional] 
**Interfaces** | Pointer to [**[]NetworkInterface**](NetworkInterface.md) | Specifies the list of network interfaces present on this Node. | [optional] 
**Ip** | Pointer to **NullableString** | Specifies the ip of the node. | [optional] 
**SlotNumber** | Pointer to **NullableInt32** | Specifies the slot number. | [optional] 

## Methods

### NewNodeInterfaces

`func NewNodeInterfaces() *NodeInterfaces`

NewNodeInterfaces instantiates a new NodeInterfaces object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeInterfacesWithDefaults

`func NewNodeInterfacesWithDefaults() *NodeInterfaces`

NewNodeInterfacesWithDefaults instantiates a new NodeInterfaces object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChassisSerial

`func (o *NodeInterfaces) GetChassisSerial() string`

GetChassisSerial returns the ChassisSerial field if non-nil, zero value otherwise.

### GetChassisSerialOk

`func (o *NodeInterfaces) GetChassisSerialOk() (*string, bool)`

GetChassisSerialOk returns a tuple with the ChassisSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisSerial

`func (o *NodeInterfaces) SetChassisSerial(v string)`

SetChassisSerial sets ChassisSerial field to given value.

### HasChassisSerial

`func (o *NodeInterfaces) HasChassisSerial() bool`

HasChassisSerial returns a boolean if a field has been set.

### SetChassisSerialNil

`func (o *NodeInterfaces) SetChassisSerialNil(b bool)`

 SetChassisSerialNil sets the value for ChassisSerial to be an explicit nil

### UnsetChassisSerial
`func (o *NodeInterfaces) UnsetChassisSerial()`

UnsetChassisSerial ensures that no value is present for ChassisSerial, not even an explicit nil
### GetId

`func (o *NodeInterfaces) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NodeInterfaces) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NodeInterfaces) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *NodeInterfaces) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *NodeInterfaces) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *NodeInterfaces) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetInterfaces

`func (o *NodeInterfaces) GetInterfaces() []NetworkInterface`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *NodeInterfaces) GetInterfacesOk() (*[]NetworkInterface, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *NodeInterfaces) SetInterfaces(v []NetworkInterface)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *NodeInterfaces) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.

### GetIp

`func (o *NodeInterfaces) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *NodeInterfaces) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *NodeInterfaces) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *NodeInterfaces) HasIp() bool`

HasIp returns a boolean if a field has been set.

### SetIpNil

`func (o *NodeInterfaces) SetIpNil(b bool)`

 SetIpNil sets the value for Ip to be an explicit nil

### UnsetIp
`func (o *NodeInterfaces) UnsetIp()`

UnsetIp ensures that no value is present for Ip, not even an explicit nil
### GetSlotNumber

`func (o *NodeInterfaces) GetSlotNumber() int32`

GetSlotNumber returns the SlotNumber field if non-nil, zero value otherwise.

### GetSlotNumberOk

`func (o *NodeInterfaces) GetSlotNumberOk() (*int32, bool)`

GetSlotNumberOk returns a tuple with the SlotNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotNumber

`func (o *NodeInterfaces) SetSlotNumber(v int32)`

SetSlotNumber sets SlotNumber field to given value.

### HasSlotNumber

`func (o *NodeInterfaces) HasSlotNumber() bool`

HasSlotNumber returns a boolean if a field has been set.

### SetSlotNumberNil

`func (o *NodeInterfaces) SetSlotNumberNil(b bool)`

 SetSlotNumberNil sets the value for SlotNumber to be an explicit nil

### UnsetSlotNumber
`func (o *NodeInterfaces) UnsetSlotNumber()`

UnsetSlotNumber ensures that no value is present for SlotNumber, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


