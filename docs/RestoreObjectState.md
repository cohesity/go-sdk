# RestoreObjectState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to [**RequestError**](RequestError.md) |  | [optional] 
**ObjectStatus** | Pointer to **NullableString** | Specifies the status of an object during a Restore Task. &#39;kFilesCloned&#39; indicates that the cloning has completed. &#39;kFetchedEntityInfo&#39; indicates that information about the object was fetched from the primary source. &#39;kVMCreated&#39; indicates that the new VM was created. &#39;kRelocationStarted&#39; indicates that restoring to a different resource pool has started. &#39;kFinished&#39; indicates that the Restore Task has finished. Whether it was successful or not is indicated by the error code that that is stored with the Restore Task. &#39;kAborted&#39; indicates that the Restore Task was aborted before trying to restore this object. This can happen if the Restore Task encounters a global error. For example during a &#39;kCloneVMs&#39; Restore Task, the datastore could not be mounted. The entire Restore Task is aborted before trying to create VMs on the primary source. &#39;kDataCopyStarted&#39; indicates that the disk copy is started. &#39;kInProgress&#39; captures a generic in-progress state and can be used by restore operations that don&#39;t track individual states. | [optional] 
**ResourcePoolId** | Pointer to **NullableInt64** | Specifies the id of the Resource Pool that the restored object is attached to. For a &#39;kRecoverVMs&#39; Restore Task, an object can be recovered back to its original resource pool. This means while recovering a set of objects, this field can reference different resource pools. For a &#39;kCloneVMs&#39; Restore Task, all objects are attached to the same resource pool. However, this field will still be populated. NOTE: This field may not be populated if the restore of the object fails. | [optional] 
**RestoredObjectId** | Pointer to **NullableInt64** | Specifies the Id of the recovered or cloned object. NOTE: For a Restore Task that is recovering or cloning an object in the VMware environment, after the VM is created it is storage vMotioned to its primary datastore. If storage vMotion fails, the Cohesity Cluster marks the recovery task as failed. However, this field is still populated with the id of the recovered VM. This id can be used later to clean up the VM from primary environment (i.e, the vCenter Server). | [optional] 
**SourceObjectId** | Pointer to **NullableInt64** | Specifies the Protection Source id of the object to be recovered or cloned. | [optional] 

## Methods

### NewRestoreObjectState

`func NewRestoreObjectState() *RestoreObjectState`

NewRestoreObjectState instantiates a new RestoreObjectState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreObjectStateWithDefaults

`func NewRestoreObjectStateWithDefaults() *RestoreObjectState`

NewRestoreObjectStateWithDefaults instantiates a new RestoreObjectState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *RestoreObjectState) GetError() RequestError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *RestoreObjectState) GetErrorOk() (*RequestError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *RestoreObjectState) SetError(v RequestError)`

SetError sets Error field to given value.

### HasError

`func (o *RestoreObjectState) HasError() bool`

HasError returns a boolean if a field has been set.

### GetObjectStatus

`func (o *RestoreObjectState) GetObjectStatus() string`

GetObjectStatus returns the ObjectStatus field if non-nil, zero value otherwise.

### GetObjectStatusOk

`func (o *RestoreObjectState) GetObjectStatusOk() (*string, bool)`

GetObjectStatusOk returns a tuple with the ObjectStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectStatus

`func (o *RestoreObjectState) SetObjectStatus(v string)`

SetObjectStatus sets ObjectStatus field to given value.

### HasObjectStatus

`func (o *RestoreObjectState) HasObjectStatus() bool`

HasObjectStatus returns a boolean if a field has been set.

### SetObjectStatusNil

`func (o *RestoreObjectState) SetObjectStatusNil(b bool)`

 SetObjectStatusNil sets the value for ObjectStatus to be an explicit nil

### UnsetObjectStatus
`func (o *RestoreObjectState) UnsetObjectStatus()`

UnsetObjectStatus ensures that no value is present for ObjectStatus, not even an explicit nil
### GetResourcePoolId

`func (o *RestoreObjectState) GetResourcePoolId() int64`

GetResourcePoolId returns the ResourcePoolId field if non-nil, zero value otherwise.

### GetResourcePoolIdOk

`func (o *RestoreObjectState) GetResourcePoolIdOk() (*int64, bool)`

GetResourcePoolIdOk returns a tuple with the ResourcePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolId

`func (o *RestoreObjectState) SetResourcePoolId(v int64)`

SetResourcePoolId sets ResourcePoolId field to given value.

### HasResourcePoolId

`func (o *RestoreObjectState) HasResourcePoolId() bool`

HasResourcePoolId returns a boolean if a field has been set.

### SetResourcePoolIdNil

`func (o *RestoreObjectState) SetResourcePoolIdNil(b bool)`

 SetResourcePoolIdNil sets the value for ResourcePoolId to be an explicit nil

### UnsetResourcePoolId
`func (o *RestoreObjectState) UnsetResourcePoolId()`

UnsetResourcePoolId ensures that no value is present for ResourcePoolId, not even an explicit nil
### GetRestoredObjectId

`func (o *RestoreObjectState) GetRestoredObjectId() int64`

GetRestoredObjectId returns the RestoredObjectId field if non-nil, zero value otherwise.

### GetRestoredObjectIdOk

`func (o *RestoreObjectState) GetRestoredObjectIdOk() (*int64, bool)`

GetRestoredObjectIdOk returns a tuple with the RestoredObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoredObjectId

`func (o *RestoreObjectState) SetRestoredObjectId(v int64)`

SetRestoredObjectId sets RestoredObjectId field to given value.

### HasRestoredObjectId

`func (o *RestoreObjectState) HasRestoredObjectId() bool`

HasRestoredObjectId returns a boolean if a field has been set.

### SetRestoredObjectIdNil

`func (o *RestoreObjectState) SetRestoredObjectIdNil(b bool)`

 SetRestoredObjectIdNil sets the value for RestoredObjectId to be an explicit nil

### UnsetRestoredObjectId
`func (o *RestoreObjectState) UnsetRestoredObjectId()`

UnsetRestoredObjectId ensures that no value is present for RestoredObjectId, not even an explicit nil
### GetSourceObjectId

`func (o *RestoreObjectState) GetSourceObjectId() int64`

GetSourceObjectId returns the SourceObjectId field if non-nil, zero value otherwise.

### GetSourceObjectIdOk

`func (o *RestoreObjectState) GetSourceObjectIdOk() (*int64, bool)`

GetSourceObjectIdOk returns a tuple with the SourceObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceObjectId

`func (o *RestoreObjectState) SetSourceObjectId(v int64)`

SetSourceObjectId sets SourceObjectId field to given value.

### HasSourceObjectId

`func (o *RestoreObjectState) HasSourceObjectId() bool`

HasSourceObjectId returns a boolean if a field has been set.

### SetSourceObjectIdNil

`func (o *RestoreObjectState) SetSourceObjectIdNil(b bool)`

 SetSourceObjectIdNil sets the value for SourceObjectId to be an explicit nil

### UnsetSourceObjectId
`func (o *RestoreObjectState) UnsetSourceObjectId()`

UnsetSourceObjectId ensures that no value is present for SourceObjectId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


