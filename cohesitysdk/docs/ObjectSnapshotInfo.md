# ObjectSnapshotInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterPartitionId** | Pointer to **NullableInt64** | Specifies the Cohesity Cluster partition id where this object is stored. | [optional] 
**JobId** | Pointer to **NullableInt64** | Specifies the id for the Protection Job that is currently associated with the object. If the object was backed up on current Cohesity Cluster, this field contains the id for the Job that captured this backup object. If the object was backed up on a Primary Cluster and replicated to this Cohesity Cluster, a new Inactive Job is created, the object is now associated with new Inactive Job, and this field contains the id of the new Inactive Job. | [optional] 
**JobName** | Pointer to **NullableString** | Specifies the name of the Protection Job that captured the backup. | [optional] 
**JobUid** | Pointer to [**NullableUniversalId**](UniversalId.md) | Specifies the globally unique id of the Protection Job that backed up this object. This id is unique across Cohesity Clusters. Even if this object is replicated to a Remote Cohesity Cluster and the object is associated with a new Job, the value specified in this field does not change. | [optional] 
**ObjectName** | Pointer to **NullableString** | Specifies the primary name of the object. | [optional] 
**OsType** | Pointer to **NullableString** | Specifies the inferred OS type. | [optional] 
**RegisteredSource** | Pointer to [**ProtectionSource**](ProtectionSource.md) |  | [optional] 
**SnapshottedSource** | Pointer to [**ProtectionSource**](ProtectionSource.md) |  | [optional] 
**Versions** | Pointer to [**[]SnapshotVersion**](SnapshotVersion.md) | Array of Snapshots.  Specifies all snapshot versions of this object. Each time a Job Run of a Job executes, it may create a new snapshot version of an object. This array stores the different snapshots versions of the object. | [optional] 
**ViewBoxId** | Pointer to **NullableInt64** | Specifies the id of the Domain (View Box) where this object is stored. | [optional] 
**ViewName** | Pointer to **NullableString** | Specifies the View name where this object is stored. | [optional] 

## Methods

### NewObjectSnapshotInfo

`func NewObjectSnapshotInfo() *ObjectSnapshotInfo`

NewObjectSnapshotInfo instantiates a new ObjectSnapshotInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectSnapshotInfoWithDefaults

`func NewObjectSnapshotInfoWithDefaults() *ObjectSnapshotInfo`

NewObjectSnapshotInfoWithDefaults instantiates a new ObjectSnapshotInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterPartitionId

`func (o *ObjectSnapshotInfo) GetClusterPartitionId() int64`

GetClusterPartitionId returns the ClusterPartitionId field if non-nil, zero value otherwise.

### GetClusterPartitionIdOk

`func (o *ObjectSnapshotInfo) GetClusterPartitionIdOk() (*int64, bool)`

GetClusterPartitionIdOk returns a tuple with the ClusterPartitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterPartitionId

`func (o *ObjectSnapshotInfo) SetClusterPartitionId(v int64)`

SetClusterPartitionId sets ClusterPartitionId field to given value.

### HasClusterPartitionId

`func (o *ObjectSnapshotInfo) HasClusterPartitionId() bool`

HasClusterPartitionId returns a boolean if a field has been set.

### SetClusterPartitionIdNil

`func (o *ObjectSnapshotInfo) SetClusterPartitionIdNil(b bool)`

 SetClusterPartitionIdNil sets the value for ClusterPartitionId to be an explicit nil

### UnsetClusterPartitionId
`func (o *ObjectSnapshotInfo) UnsetClusterPartitionId()`

UnsetClusterPartitionId ensures that no value is present for ClusterPartitionId, not even an explicit nil
### GetJobId

`func (o *ObjectSnapshotInfo) GetJobId() int64`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *ObjectSnapshotInfo) GetJobIdOk() (*int64, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *ObjectSnapshotInfo) SetJobId(v int64)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *ObjectSnapshotInfo) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### SetJobIdNil

`func (o *ObjectSnapshotInfo) SetJobIdNil(b bool)`

 SetJobIdNil sets the value for JobId to be an explicit nil

### UnsetJobId
`func (o *ObjectSnapshotInfo) UnsetJobId()`

UnsetJobId ensures that no value is present for JobId, not even an explicit nil
### GetJobName

`func (o *ObjectSnapshotInfo) GetJobName() string`

GetJobName returns the JobName field if non-nil, zero value otherwise.

### GetJobNameOk

`func (o *ObjectSnapshotInfo) GetJobNameOk() (*string, bool)`

GetJobNameOk returns a tuple with the JobName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobName

`func (o *ObjectSnapshotInfo) SetJobName(v string)`

SetJobName sets JobName field to given value.

### HasJobName

`func (o *ObjectSnapshotInfo) HasJobName() bool`

HasJobName returns a boolean if a field has been set.

### SetJobNameNil

`func (o *ObjectSnapshotInfo) SetJobNameNil(b bool)`

 SetJobNameNil sets the value for JobName to be an explicit nil

### UnsetJobName
`func (o *ObjectSnapshotInfo) UnsetJobName()`

UnsetJobName ensures that no value is present for JobName, not even an explicit nil
### GetJobUid

`func (o *ObjectSnapshotInfo) GetJobUid() UniversalId`

GetJobUid returns the JobUid field if non-nil, zero value otherwise.

### GetJobUidOk

`func (o *ObjectSnapshotInfo) GetJobUidOk() (*UniversalId, bool)`

GetJobUidOk returns a tuple with the JobUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobUid

`func (o *ObjectSnapshotInfo) SetJobUid(v UniversalId)`

SetJobUid sets JobUid field to given value.

### HasJobUid

`func (o *ObjectSnapshotInfo) HasJobUid() bool`

HasJobUid returns a boolean if a field has been set.

### SetJobUidNil

`func (o *ObjectSnapshotInfo) SetJobUidNil(b bool)`

 SetJobUidNil sets the value for JobUid to be an explicit nil

### UnsetJobUid
`func (o *ObjectSnapshotInfo) UnsetJobUid()`

UnsetJobUid ensures that no value is present for JobUid, not even an explicit nil
### GetObjectName

`func (o *ObjectSnapshotInfo) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *ObjectSnapshotInfo) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *ObjectSnapshotInfo) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *ObjectSnapshotInfo) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### SetObjectNameNil

`func (o *ObjectSnapshotInfo) SetObjectNameNil(b bool)`

 SetObjectNameNil sets the value for ObjectName to be an explicit nil

### UnsetObjectName
`func (o *ObjectSnapshotInfo) UnsetObjectName()`

UnsetObjectName ensures that no value is present for ObjectName, not even an explicit nil
### GetOsType

`func (o *ObjectSnapshotInfo) GetOsType() string`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *ObjectSnapshotInfo) GetOsTypeOk() (*string, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *ObjectSnapshotInfo) SetOsType(v string)`

SetOsType sets OsType field to given value.

### HasOsType

`func (o *ObjectSnapshotInfo) HasOsType() bool`

HasOsType returns a boolean if a field has been set.

### SetOsTypeNil

`func (o *ObjectSnapshotInfo) SetOsTypeNil(b bool)`

 SetOsTypeNil sets the value for OsType to be an explicit nil

### UnsetOsType
`func (o *ObjectSnapshotInfo) UnsetOsType()`

UnsetOsType ensures that no value is present for OsType, not even an explicit nil
### GetRegisteredSource

`func (o *ObjectSnapshotInfo) GetRegisteredSource() ProtectionSource`

GetRegisteredSource returns the RegisteredSource field if non-nil, zero value otherwise.

### GetRegisteredSourceOk

`func (o *ObjectSnapshotInfo) GetRegisteredSourceOk() (*ProtectionSource, bool)`

GetRegisteredSourceOk returns a tuple with the RegisteredSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredSource

`func (o *ObjectSnapshotInfo) SetRegisteredSource(v ProtectionSource)`

SetRegisteredSource sets RegisteredSource field to given value.

### HasRegisteredSource

`func (o *ObjectSnapshotInfo) HasRegisteredSource() bool`

HasRegisteredSource returns a boolean if a field has been set.

### GetSnapshottedSource

`func (o *ObjectSnapshotInfo) GetSnapshottedSource() ProtectionSource`

GetSnapshottedSource returns the SnapshottedSource field if non-nil, zero value otherwise.

### GetSnapshottedSourceOk

`func (o *ObjectSnapshotInfo) GetSnapshottedSourceOk() (*ProtectionSource, bool)`

GetSnapshottedSourceOk returns a tuple with the SnapshottedSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshottedSource

`func (o *ObjectSnapshotInfo) SetSnapshottedSource(v ProtectionSource)`

SetSnapshottedSource sets SnapshottedSource field to given value.

### HasSnapshottedSource

`func (o *ObjectSnapshotInfo) HasSnapshottedSource() bool`

HasSnapshottedSource returns a boolean if a field has been set.

### GetVersions

`func (o *ObjectSnapshotInfo) GetVersions() []SnapshotVersion`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *ObjectSnapshotInfo) GetVersionsOk() (*[]SnapshotVersion, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *ObjectSnapshotInfo) SetVersions(v []SnapshotVersion)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *ObjectSnapshotInfo) HasVersions() bool`

HasVersions returns a boolean if a field has been set.

### SetVersionsNil

`func (o *ObjectSnapshotInfo) SetVersionsNil(b bool)`

 SetVersionsNil sets the value for Versions to be an explicit nil

### UnsetVersions
`func (o *ObjectSnapshotInfo) UnsetVersions()`

UnsetVersions ensures that no value is present for Versions, not even an explicit nil
### GetViewBoxId

`func (o *ObjectSnapshotInfo) GetViewBoxId() int64`

GetViewBoxId returns the ViewBoxId field if non-nil, zero value otherwise.

### GetViewBoxIdOk

`func (o *ObjectSnapshotInfo) GetViewBoxIdOk() (*int64, bool)`

GetViewBoxIdOk returns a tuple with the ViewBoxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewBoxId

`func (o *ObjectSnapshotInfo) SetViewBoxId(v int64)`

SetViewBoxId sets ViewBoxId field to given value.

### HasViewBoxId

`func (o *ObjectSnapshotInfo) HasViewBoxId() bool`

HasViewBoxId returns a boolean if a field has been set.

### SetViewBoxIdNil

`func (o *ObjectSnapshotInfo) SetViewBoxIdNil(b bool)`

 SetViewBoxIdNil sets the value for ViewBoxId to be an explicit nil

### UnsetViewBoxId
`func (o *ObjectSnapshotInfo) UnsetViewBoxId()`

UnsetViewBoxId ensures that no value is present for ViewBoxId, not even an explicit nil
### GetViewName

`func (o *ObjectSnapshotInfo) GetViewName() string`

GetViewName returns the ViewName field if non-nil, zero value otherwise.

### GetViewNameOk

`func (o *ObjectSnapshotInfo) GetViewNameOk() (*string, bool)`

GetViewNameOk returns a tuple with the ViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewName

`func (o *ObjectSnapshotInfo) SetViewName(v string)`

SetViewName sets ViewName field to given value.

### HasViewName

`func (o *ObjectSnapshotInfo) HasViewName() bool`

HasViewName returns a boolean if a field has been set.

### SetViewNameNil

`func (o *ObjectSnapshotInfo) SetViewNameNil(b bool)`

 SetViewNameNil sets the value for ViewName to be an explicit nil

### UnsetViewName
`func (o *ObjectSnapshotInfo) UnsetViewName()`

UnsetViewName ensures that no value is present for ViewName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


