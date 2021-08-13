# ObjectArchivalSnapshotInfoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SnapshotId** | Pointer to **NullableString** | Specifies the id of the archival snapshot for the object. | [optional] 
**LogicalSizeBytes** | Pointer to **NullableInt64** | Specifies the logical size of this snapshot in bytes. | [optional] 

## Methods

### NewObjectArchivalSnapshotInfoAllOf

`func NewObjectArchivalSnapshotInfoAllOf() *ObjectArchivalSnapshotInfoAllOf`

NewObjectArchivalSnapshotInfoAllOf instantiates a new ObjectArchivalSnapshotInfoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectArchivalSnapshotInfoAllOfWithDefaults

`func NewObjectArchivalSnapshotInfoAllOfWithDefaults() *ObjectArchivalSnapshotInfoAllOf`

NewObjectArchivalSnapshotInfoAllOfWithDefaults instantiates a new ObjectArchivalSnapshotInfoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSnapshotId

`func (o *ObjectArchivalSnapshotInfoAllOf) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *ObjectArchivalSnapshotInfoAllOf) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *ObjectArchivalSnapshotInfoAllOf) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.

### HasSnapshotId

`func (o *ObjectArchivalSnapshotInfoAllOf) HasSnapshotId() bool`

HasSnapshotId returns a boolean if a field has been set.

### SetSnapshotIdNil

`func (o *ObjectArchivalSnapshotInfoAllOf) SetSnapshotIdNil(b bool)`

 SetSnapshotIdNil sets the value for SnapshotId to be an explicit nil

### UnsetSnapshotId
`func (o *ObjectArchivalSnapshotInfoAllOf) UnsetSnapshotId()`

UnsetSnapshotId ensures that no value is present for SnapshotId, not even an explicit nil
### GetLogicalSizeBytes

`func (o *ObjectArchivalSnapshotInfoAllOf) GetLogicalSizeBytes() int64`

GetLogicalSizeBytes returns the LogicalSizeBytes field if non-nil, zero value otherwise.

### GetLogicalSizeBytesOk

`func (o *ObjectArchivalSnapshotInfoAllOf) GetLogicalSizeBytesOk() (*int64, bool)`

GetLogicalSizeBytesOk returns a tuple with the LogicalSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalSizeBytes

`func (o *ObjectArchivalSnapshotInfoAllOf) SetLogicalSizeBytes(v int64)`

SetLogicalSizeBytes sets LogicalSizeBytes field to given value.

### HasLogicalSizeBytes

`func (o *ObjectArchivalSnapshotInfoAllOf) HasLogicalSizeBytes() bool`

HasLogicalSizeBytes returns a boolean if a field has been set.

### SetLogicalSizeBytesNil

`func (o *ObjectArchivalSnapshotInfoAllOf) SetLogicalSizeBytesNil(b bool)`

 SetLogicalSizeBytesNil sets the value for LogicalSizeBytes to be an explicit nil

### UnsetLogicalSizeBytes
`func (o *ObjectArchivalSnapshotInfoAllOf) UnsetLogicalSizeBytes()`

UnsetLogicalSizeBytes ensures that no value is present for LogicalSizeBytes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


