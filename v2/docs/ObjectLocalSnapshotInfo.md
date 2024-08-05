# ObjectLocalSnapshotInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogicalSizeBytes** | Pointer to **NullableInt64** | Specifies the logical size of this snapshot in bytes. | [optional] 
**SnapshotId** | Pointer to **NullableString** | Specifies the id of the local snapshot for the object. | [optional] 

## Methods

### NewObjectLocalSnapshotInfo

`func NewObjectLocalSnapshotInfo() *ObjectLocalSnapshotInfo`

NewObjectLocalSnapshotInfo instantiates a new ObjectLocalSnapshotInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectLocalSnapshotInfoWithDefaults

`func NewObjectLocalSnapshotInfoWithDefaults() *ObjectLocalSnapshotInfo`

NewObjectLocalSnapshotInfoWithDefaults instantiates a new ObjectLocalSnapshotInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogicalSizeBytes

`func (o *ObjectLocalSnapshotInfo) GetLogicalSizeBytes() int64`

GetLogicalSizeBytes returns the LogicalSizeBytes field if non-nil, zero value otherwise.

### GetLogicalSizeBytesOk

`func (o *ObjectLocalSnapshotInfo) GetLogicalSizeBytesOk() (*int64, bool)`

GetLogicalSizeBytesOk returns a tuple with the LogicalSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalSizeBytes

`func (o *ObjectLocalSnapshotInfo) SetLogicalSizeBytes(v int64)`

SetLogicalSizeBytes sets LogicalSizeBytes field to given value.

### HasLogicalSizeBytes

`func (o *ObjectLocalSnapshotInfo) HasLogicalSizeBytes() bool`

HasLogicalSizeBytes returns a boolean if a field has been set.

### SetLogicalSizeBytesNil

`func (o *ObjectLocalSnapshotInfo) SetLogicalSizeBytesNil(b bool)`

 SetLogicalSizeBytesNil sets the value for LogicalSizeBytes to be an explicit nil

### UnsetLogicalSizeBytes
`func (o *ObjectLocalSnapshotInfo) UnsetLogicalSizeBytes()`

UnsetLogicalSizeBytes ensures that no value is present for LogicalSizeBytes, not even an explicit nil
### GetSnapshotId

`func (o *ObjectLocalSnapshotInfo) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *ObjectLocalSnapshotInfo) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *ObjectLocalSnapshotInfo) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.

### HasSnapshotId

`func (o *ObjectLocalSnapshotInfo) HasSnapshotId() bool`

HasSnapshotId returns a boolean if a field has been set.

### SetSnapshotIdNil

`func (o *ObjectLocalSnapshotInfo) SetSnapshotIdNil(b bool)`

 SetSnapshotIdNil sets the value for SnapshotId to be an explicit nil

### UnsetSnapshotId
`func (o *ObjectLocalSnapshotInfo) UnsetSnapshotId()`

UnsetSnapshotId ensures that no value is present for SnapshotId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


