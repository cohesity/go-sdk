# Disk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt64** | Specifies id to uniquely identify a disk. | [optional] 
**SerialNumber** | Pointer to **NullableString** | Specifies serial number of disk. | [optional] 
**CapacityInBytes** | Pointer to **NullableInt64** | Specifies capacity of disk in bytes. | [optional] 
**Model** | Pointer to **NullableString** | Specifies product model of disk. | [optional] 
**NodeId** | Pointer to **NullableInt64** | Specifies node id of the node that this disk belong to. | [optional] 
**Status** | Pointer to **string** | Specifies status of the disk. | [optional] 
**Type** | Pointer to **string** | Specifies type of the disk. | [optional] 
**Location** | Pointer to **NullableString** | Specifies location of the disk in node. | [optional] 

## Methods

### NewDisk

`func NewDisk() *Disk`

NewDisk instantiates a new Disk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskWithDefaults

`func NewDiskWithDefaults() *Disk`

NewDiskWithDefaults instantiates a new Disk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Disk) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Disk) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Disk) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Disk) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *Disk) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *Disk) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetSerialNumber

`func (o *Disk) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *Disk) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *Disk) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *Disk) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### SetSerialNumberNil

`func (o *Disk) SetSerialNumberNil(b bool)`

 SetSerialNumberNil sets the value for SerialNumber to be an explicit nil

### UnsetSerialNumber
`func (o *Disk) UnsetSerialNumber()`

UnsetSerialNumber ensures that no value is present for SerialNumber, not even an explicit nil
### GetCapacityInBytes

`func (o *Disk) GetCapacityInBytes() int64`

GetCapacityInBytes returns the CapacityInBytes field if non-nil, zero value otherwise.

### GetCapacityInBytesOk

`func (o *Disk) GetCapacityInBytesOk() (*int64, bool)`

GetCapacityInBytesOk returns a tuple with the CapacityInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityInBytes

`func (o *Disk) SetCapacityInBytes(v int64)`

SetCapacityInBytes sets CapacityInBytes field to given value.

### HasCapacityInBytes

`func (o *Disk) HasCapacityInBytes() bool`

HasCapacityInBytes returns a boolean if a field has been set.

### SetCapacityInBytesNil

`func (o *Disk) SetCapacityInBytesNil(b bool)`

 SetCapacityInBytesNil sets the value for CapacityInBytes to be an explicit nil

### UnsetCapacityInBytes
`func (o *Disk) UnsetCapacityInBytes()`

UnsetCapacityInBytes ensures that no value is present for CapacityInBytes, not even an explicit nil
### GetModel

`func (o *Disk) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *Disk) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *Disk) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *Disk) HasModel() bool`

HasModel returns a boolean if a field has been set.

### SetModelNil

`func (o *Disk) SetModelNil(b bool)`

 SetModelNil sets the value for Model to be an explicit nil

### UnsetModel
`func (o *Disk) UnsetModel()`

UnsetModel ensures that no value is present for Model, not even an explicit nil
### GetNodeId

`func (o *Disk) GetNodeId() int64`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *Disk) GetNodeIdOk() (*int64, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *Disk) SetNodeId(v int64)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *Disk) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### SetNodeIdNil

`func (o *Disk) SetNodeIdNil(b bool)`

 SetNodeIdNil sets the value for NodeId to be an explicit nil

### UnsetNodeId
`func (o *Disk) UnsetNodeId()`

UnsetNodeId ensures that no value is present for NodeId, not even an explicit nil
### GetStatus

`func (o *Disk) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Disk) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Disk) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Disk) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *Disk) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Disk) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Disk) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Disk) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLocation

`func (o *Disk) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Disk) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Disk) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *Disk) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *Disk) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *Disk) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


