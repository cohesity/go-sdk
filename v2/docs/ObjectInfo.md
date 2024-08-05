# ObjectInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseSnapshotHandle** | Pointer to [**NullableSnapshotHandle**](SnapshotHandle.md) |  | [optional] 
**GlobalId** | Pointer to **NullableString** | Specifies the object id. | [optional] 
**ObjectName** | Pointer to **NullableString** | Specifies the object name | [optional] 
**ObjectType** | Pointer to [**ObjectType**](ObjectType.md) |  | [optional] 
**Snapshots** | Pointer to [**[]GaiaSnapshotInfo**](GaiaSnapshotInfo.md) | Snapshots of the object. | [optional] 

## Methods

### NewObjectInfo

`func NewObjectInfo() *ObjectInfo`

NewObjectInfo instantiates a new ObjectInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectInfoWithDefaults

`func NewObjectInfoWithDefaults() *ObjectInfo`

NewObjectInfoWithDefaults instantiates a new ObjectInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseSnapshotHandle

`func (o *ObjectInfo) GetBaseSnapshotHandle() SnapshotHandle`

GetBaseSnapshotHandle returns the BaseSnapshotHandle field if non-nil, zero value otherwise.

### GetBaseSnapshotHandleOk

`func (o *ObjectInfo) GetBaseSnapshotHandleOk() (*SnapshotHandle, bool)`

GetBaseSnapshotHandleOk returns a tuple with the BaseSnapshotHandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseSnapshotHandle

`func (o *ObjectInfo) SetBaseSnapshotHandle(v SnapshotHandle)`

SetBaseSnapshotHandle sets BaseSnapshotHandle field to given value.

### HasBaseSnapshotHandle

`func (o *ObjectInfo) HasBaseSnapshotHandle() bool`

HasBaseSnapshotHandle returns a boolean if a field has been set.

### SetBaseSnapshotHandleNil

`func (o *ObjectInfo) SetBaseSnapshotHandleNil(b bool)`

 SetBaseSnapshotHandleNil sets the value for BaseSnapshotHandle to be an explicit nil

### UnsetBaseSnapshotHandle
`func (o *ObjectInfo) UnsetBaseSnapshotHandle()`

UnsetBaseSnapshotHandle ensures that no value is present for BaseSnapshotHandle, not even an explicit nil
### GetGlobalId

`func (o *ObjectInfo) GetGlobalId() string`

GetGlobalId returns the GlobalId field if non-nil, zero value otherwise.

### GetGlobalIdOk

`func (o *ObjectInfo) GetGlobalIdOk() (*string, bool)`

GetGlobalIdOk returns a tuple with the GlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalId

`func (o *ObjectInfo) SetGlobalId(v string)`

SetGlobalId sets GlobalId field to given value.

### HasGlobalId

`func (o *ObjectInfo) HasGlobalId() bool`

HasGlobalId returns a boolean if a field has been set.

### SetGlobalIdNil

`func (o *ObjectInfo) SetGlobalIdNil(b bool)`

 SetGlobalIdNil sets the value for GlobalId to be an explicit nil

### UnsetGlobalId
`func (o *ObjectInfo) UnsetGlobalId()`

UnsetGlobalId ensures that no value is present for GlobalId, not even an explicit nil
### GetObjectName

`func (o *ObjectInfo) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *ObjectInfo) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *ObjectInfo) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *ObjectInfo) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### SetObjectNameNil

`func (o *ObjectInfo) SetObjectNameNil(b bool)`

 SetObjectNameNil sets the value for ObjectName to be an explicit nil

### UnsetObjectName
`func (o *ObjectInfo) UnsetObjectName()`

UnsetObjectName ensures that no value is present for ObjectName, not even an explicit nil
### GetObjectType

`func (o *ObjectInfo) GetObjectType() ObjectType`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ObjectInfo) GetObjectTypeOk() (*ObjectType, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ObjectInfo) SetObjectType(v ObjectType)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *ObjectInfo) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetSnapshots

`func (o *ObjectInfo) GetSnapshots() []GaiaSnapshotInfo`

GetSnapshots returns the Snapshots field if non-nil, zero value otherwise.

### GetSnapshotsOk

`func (o *ObjectInfo) GetSnapshotsOk() (*[]GaiaSnapshotInfo, bool)`

GetSnapshotsOk returns a tuple with the Snapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshots

`func (o *ObjectInfo) SetSnapshots(v []GaiaSnapshotInfo)`

SetSnapshots sets Snapshots field to given value.

### HasSnapshots

`func (o *ObjectInfo) HasSnapshots() bool`

HasSnapshots returns a boolean if a field has been set.

### SetSnapshotsNil

`func (o *ObjectInfo) SetSnapshotsNil(b bool)`

 SetSnapshotsNil sets the value for Snapshots to be an explicit nil

### UnsetSnapshots
`func (o *ObjectInfo) UnsetSnapshots()`

UnsetSnapshots ensures that no value is present for Snapshots, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


