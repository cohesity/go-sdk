# DiskIdentify

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiskId** | Pointer to **NullableInt64** | Specifies the disk id of the disk. This parameter is incompatible with &#39;nodeId&#39; and &#39;serialNumber&#39;. | [optional] 
**Identify** | **NullableBool** | Turn on/off led light if it is set to true/false | 
**NodeId** | Pointer to **NullableInt64** | Specifies the node id of node that disk belongs to. This parameter is incompatible with &#39;diskId&#39;. Must be used together with &#39;serialNumber&#39;. | [optional] 
**SerialNumber** | Pointer to **NullableString** | Specifies serial number of disk. This parameter is incompatible with &#39;diskId&#39;. Must be used together with &#39;nodeId&#39;. | [optional] 

## Methods

### NewDiskIdentify

`func NewDiskIdentify(identify NullableBool, ) *DiskIdentify`

NewDiskIdentify instantiates a new DiskIdentify object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskIdentifyWithDefaults

`func NewDiskIdentifyWithDefaults() *DiskIdentify`

NewDiskIdentifyWithDefaults instantiates a new DiskIdentify object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiskId

`func (o *DiskIdentify) GetDiskId() int64`

GetDiskId returns the DiskId field if non-nil, zero value otherwise.

### GetDiskIdOk

`func (o *DiskIdentify) GetDiskIdOk() (*int64, bool)`

GetDiskIdOk returns a tuple with the DiskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskId

`func (o *DiskIdentify) SetDiskId(v int64)`

SetDiskId sets DiskId field to given value.

### HasDiskId

`func (o *DiskIdentify) HasDiskId() bool`

HasDiskId returns a boolean if a field has been set.

### SetDiskIdNil

`func (o *DiskIdentify) SetDiskIdNil(b bool)`

 SetDiskIdNil sets the value for DiskId to be an explicit nil

### UnsetDiskId
`func (o *DiskIdentify) UnsetDiskId()`

UnsetDiskId ensures that no value is present for DiskId, not even an explicit nil
### GetIdentify

`func (o *DiskIdentify) GetIdentify() bool`

GetIdentify returns the Identify field if non-nil, zero value otherwise.

### GetIdentifyOk

`func (o *DiskIdentify) GetIdentifyOk() (*bool, bool)`

GetIdentifyOk returns a tuple with the Identify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentify

`func (o *DiskIdentify) SetIdentify(v bool)`

SetIdentify sets Identify field to given value.


### SetIdentifyNil

`func (o *DiskIdentify) SetIdentifyNil(b bool)`

 SetIdentifyNil sets the value for Identify to be an explicit nil

### UnsetIdentify
`func (o *DiskIdentify) UnsetIdentify()`

UnsetIdentify ensures that no value is present for Identify, not even an explicit nil
### GetNodeId

`func (o *DiskIdentify) GetNodeId() int64`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *DiskIdentify) GetNodeIdOk() (*int64, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *DiskIdentify) SetNodeId(v int64)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *DiskIdentify) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### SetNodeIdNil

`func (o *DiskIdentify) SetNodeIdNil(b bool)`

 SetNodeIdNil sets the value for NodeId to be an explicit nil

### UnsetNodeId
`func (o *DiskIdentify) UnsetNodeId()`

UnsetNodeId ensures that no value is present for NodeId, not even an explicit nil
### GetSerialNumber

`func (o *DiskIdentify) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *DiskIdentify) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *DiskIdentify) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *DiskIdentify) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### SetSerialNumberNil

`func (o *DiskIdentify) SetSerialNumberNil(b bool)`

 SetSerialNumberNil sets the value for SerialNumber to be an explicit nil

### UnsetSerialNumber
`func (o *DiskIdentify) UnsetSerialNumber()`

UnsetSerialNumber ensures that no value is present for SerialNumber, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


