# RemoteAdapterProtectionGroupParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hosts** | [**[]RemoteAdapterHost**](RemoteAdapterHost.md) | Specifies a list of hosts to protected in this protection group. | 
**ViewId** | **NullableInt64** | Specifies the id of the view where we put the script result data. | 
**RemoteViewParams** | Pointer to [**RemoteAdapterProtectionGroupReplicationParams**](RemoteAdapterProtectionGroupReplicationParams.md) |  | [optional] 
**IndexingPolicy** | Pointer to [**IndexingPolicy**](IndexingPolicy.md) |  | [optional] 
**AppConsistentSnapshot** | Pointer to **NullableBool** | Specifies whether or not to quiesce apps and the file system in order to take app consistent snapshots. | [optional] 

## Methods

### NewRemoteAdapterProtectionGroupParams

`func NewRemoteAdapterProtectionGroupParams(hosts []RemoteAdapterHost, viewId NullableInt64, ) *RemoteAdapterProtectionGroupParams`

NewRemoteAdapterProtectionGroupParams instantiates a new RemoteAdapterProtectionGroupParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteAdapterProtectionGroupParamsWithDefaults

`func NewRemoteAdapterProtectionGroupParamsWithDefaults() *RemoteAdapterProtectionGroupParams`

NewRemoteAdapterProtectionGroupParamsWithDefaults instantiates a new RemoteAdapterProtectionGroupParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHosts

`func (o *RemoteAdapterProtectionGroupParams) GetHosts() []RemoteAdapterHost`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *RemoteAdapterProtectionGroupParams) GetHostsOk() (*[]RemoteAdapterHost, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *RemoteAdapterProtectionGroupParams) SetHosts(v []RemoteAdapterHost)`

SetHosts sets Hosts field to given value.


### SetHostsNil

`func (o *RemoteAdapterProtectionGroupParams) SetHostsNil(b bool)`

 SetHostsNil sets the value for Hosts to be an explicit nil

### UnsetHosts
`func (o *RemoteAdapterProtectionGroupParams) UnsetHosts()`

UnsetHosts ensures that no value is present for Hosts, not even an explicit nil
### GetViewId

`func (o *RemoteAdapterProtectionGroupParams) GetViewId() int64`

GetViewId returns the ViewId field if non-nil, zero value otherwise.

### GetViewIdOk

`func (o *RemoteAdapterProtectionGroupParams) GetViewIdOk() (*int64, bool)`

GetViewIdOk returns a tuple with the ViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewId

`func (o *RemoteAdapterProtectionGroupParams) SetViewId(v int64)`

SetViewId sets ViewId field to given value.


### SetViewIdNil

`func (o *RemoteAdapterProtectionGroupParams) SetViewIdNil(b bool)`

 SetViewIdNil sets the value for ViewId to be an explicit nil

### UnsetViewId
`func (o *RemoteAdapterProtectionGroupParams) UnsetViewId()`

UnsetViewId ensures that no value is present for ViewId, not even an explicit nil
### GetRemoteViewParams

`func (o *RemoteAdapterProtectionGroupParams) GetRemoteViewParams() RemoteAdapterProtectionGroupReplicationParams`

GetRemoteViewParams returns the RemoteViewParams field if non-nil, zero value otherwise.

### GetRemoteViewParamsOk

`func (o *RemoteAdapterProtectionGroupParams) GetRemoteViewParamsOk() (*RemoteAdapterProtectionGroupReplicationParams, bool)`

GetRemoteViewParamsOk returns a tuple with the RemoteViewParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteViewParams

`func (o *RemoteAdapterProtectionGroupParams) SetRemoteViewParams(v RemoteAdapterProtectionGroupReplicationParams)`

SetRemoteViewParams sets RemoteViewParams field to given value.

### HasRemoteViewParams

`func (o *RemoteAdapterProtectionGroupParams) HasRemoteViewParams() bool`

HasRemoteViewParams returns a boolean if a field has been set.

### GetIndexingPolicy

`func (o *RemoteAdapterProtectionGroupParams) GetIndexingPolicy() IndexingPolicy`

GetIndexingPolicy returns the IndexingPolicy field if non-nil, zero value otherwise.

### GetIndexingPolicyOk

`func (o *RemoteAdapterProtectionGroupParams) GetIndexingPolicyOk() (*IndexingPolicy, bool)`

GetIndexingPolicyOk returns a tuple with the IndexingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexingPolicy

`func (o *RemoteAdapterProtectionGroupParams) SetIndexingPolicy(v IndexingPolicy)`

SetIndexingPolicy sets IndexingPolicy field to given value.

### HasIndexingPolicy

`func (o *RemoteAdapterProtectionGroupParams) HasIndexingPolicy() bool`

HasIndexingPolicy returns a boolean if a field has been set.

### GetAppConsistentSnapshot

`func (o *RemoteAdapterProtectionGroupParams) GetAppConsistentSnapshot() bool`

GetAppConsistentSnapshot returns the AppConsistentSnapshot field if non-nil, zero value otherwise.

### GetAppConsistentSnapshotOk

`func (o *RemoteAdapterProtectionGroupParams) GetAppConsistentSnapshotOk() (*bool, bool)`

GetAppConsistentSnapshotOk returns a tuple with the AppConsistentSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppConsistentSnapshot

`func (o *RemoteAdapterProtectionGroupParams) SetAppConsistentSnapshot(v bool)`

SetAppConsistentSnapshot sets AppConsistentSnapshot field to given value.

### HasAppConsistentSnapshot

`func (o *RemoteAdapterProtectionGroupParams) HasAppConsistentSnapshot() bool`

HasAppConsistentSnapshot returns a boolean if a field has been set.

### SetAppConsistentSnapshotNil

`func (o *RemoteAdapterProtectionGroupParams) SetAppConsistentSnapshotNil(b bool)`

 SetAppConsistentSnapshotNil sets the value for AppConsistentSnapshot to be an explicit nil

### UnsetAppConsistentSnapshot
`func (o *RemoteAdapterProtectionGroupParams) UnsetAppConsistentSnapshot()`

UnsetAppConsistentSnapshot ensures that no value is present for AppConsistentSnapshot, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


