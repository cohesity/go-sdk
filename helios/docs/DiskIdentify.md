# DiskIdentify

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeId** | **NullableInt64** | Specifies the node id of node that disk belongs to. | 
**SerialNumber** | **NullableString** | Specifies serial number of disk. | 
**Identify** | **NullableBool** | Turn on/off led light if it is set to true/false | 

## Methods

### NewDiskIdentify

`func NewDiskIdentify(nodeId NullableInt64, serialNumber NullableString, identify NullableBool, ) *DiskIdentify`

NewDiskIdentify instantiates a new DiskIdentify object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskIdentifyWithDefaults

`func NewDiskIdentifyWithDefaults() *DiskIdentify`

NewDiskIdentifyWithDefaults instantiates a new DiskIdentify object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


### SetSerialNumberNil

`func (o *DiskIdentify) SetSerialNumberNil(b bool)`

 SetSerialNumberNil sets the value for SerialNumber to be an explicit nil

### UnsetSerialNumber
`func (o *DiskIdentify) UnsetSerialNumber()`

UnsetSerialNumber ensures that no value is present for SerialNumber, not even an explicit nil
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

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


