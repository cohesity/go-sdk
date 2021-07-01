# DestroyClonedVMTaskInfoProto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatastoreNotUnmountedReason** | Pointer to **NullableString** | If datastore was not unmounted, this field contains the reason for the same. | [optional] 
**DatastoreUnmounted** | Pointer to **NullableBool** | Whether the datastore corresponding to the clone view was unmounted from primary environment. | [optional] 
**DestroyClonedEntityInfoVec** | Pointer to [**[]DestroyClonedEntityInfoProto**](DestroyClonedEntityInfoProto.md) | Vector of all cloned entities that this destroy task will teardown. | [optional] 
**Type** | Pointer to **NullableInt32** | The type of environment this destroy clone task info pertains to. | [optional] 
**ViewDeleted** | Pointer to **NullableBool** | Whether the clone view was deleted by the destroy task. | [optional] 

## Methods

### NewDestroyClonedVMTaskInfoProto

`func NewDestroyClonedVMTaskInfoProto() *DestroyClonedVMTaskInfoProto`

NewDestroyClonedVMTaskInfoProto instantiates a new DestroyClonedVMTaskInfoProto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestroyClonedVMTaskInfoProtoWithDefaults

`func NewDestroyClonedVMTaskInfoProtoWithDefaults() *DestroyClonedVMTaskInfoProto`

NewDestroyClonedVMTaskInfoProtoWithDefaults instantiates a new DestroyClonedVMTaskInfoProto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatastoreNotUnmountedReason

`func (o *DestroyClonedVMTaskInfoProto) GetDatastoreNotUnmountedReason() string`

GetDatastoreNotUnmountedReason returns the DatastoreNotUnmountedReason field if non-nil, zero value otherwise.

### GetDatastoreNotUnmountedReasonOk

`func (o *DestroyClonedVMTaskInfoProto) GetDatastoreNotUnmountedReasonOk() (*string, bool)`

GetDatastoreNotUnmountedReasonOk returns a tuple with the DatastoreNotUnmountedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreNotUnmountedReason

`func (o *DestroyClonedVMTaskInfoProto) SetDatastoreNotUnmountedReason(v string)`

SetDatastoreNotUnmountedReason sets DatastoreNotUnmountedReason field to given value.

### HasDatastoreNotUnmountedReason

`func (o *DestroyClonedVMTaskInfoProto) HasDatastoreNotUnmountedReason() bool`

HasDatastoreNotUnmountedReason returns a boolean if a field has been set.

### SetDatastoreNotUnmountedReasonNil

`func (o *DestroyClonedVMTaskInfoProto) SetDatastoreNotUnmountedReasonNil(b bool)`

 SetDatastoreNotUnmountedReasonNil sets the value for DatastoreNotUnmountedReason to be an explicit nil

### UnsetDatastoreNotUnmountedReason
`func (o *DestroyClonedVMTaskInfoProto) UnsetDatastoreNotUnmountedReason()`

UnsetDatastoreNotUnmountedReason ensures that no value is present for DatastoreNotUnmountedReason, not even an explicit nil
### GetDatastoreUnmounted

`func (o *DestroyClonedVMTaskInfoProto) GetDatastoreUnmounted() bool`

GetDatastoreUnmounted returns the DatastoreUnmounted field if non-nil, zero value otherwise.

### GetDatastoreUnmountedOk

`func (o *DestroyClonedVMTaskInfoProto) GetDatastoreUnmountedOk() (*bool, bool)`

GetDatastoreUnmountedOk returns a tuple with the DatastoreUnmounted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreUnmounted

`func (o *DestroyClonedVMTaskInfoProto) SetDatastoreUnmounted(v bool)`

SetDatastoreUnmounted sets DatastoreUnmounted field to given value.

### HasDatastoreUnmounted

`func (o *DestroyClonedVMTaskInfoProto) HasDatastoreUnmounted() bool`

HasDatastoreUnmounted returns a boolean if a field has been set.

### SetDatastoreUnmountedNil

`func (o *DestroyClonedVMTaskInfoProto) SetDatastoreUnmountedNil(b bool)`

 SetDatastoreUnmountedNil sets the value for DatastoreUnmounted to be an explicit nil

### UnsetDatastoreUnmounted
`func (o *DestroyClonedVMTaskInfoProto) UnsetDatastoreUnmounted()`

UnsetDatastoreUnmounted ensures that no value is present for DatastoreUnmounted, not even an explicit nil
### GetDestroyClonedEntityInfoVec

`func (o *DestroyClonedVMTaskInfoProto) GetDestroyClonedEntityInfoVec() []DestroyClonedEntityInfoProto`

GetDestroyClonedEntityInfoVec returns the DestroyClonedEntityInfoVec field if non-nil, zero value otherwise.

### GetDestroyClonedEntityInfoVecOk

`func (o *DestroyClonedVMTaskInfoProto) GetDestroyClonedEntityInfoVecOk() (*[]DestroyClonedEntityInfoProto, bool)`

GetDestroyClonedEntityInfoVecOk returns a tuple with the DestroyClonedEntityInfoVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestroyClonedEntityInfoVec

`func (o *DestroyClonedVMTaskInfoProto) SetDestroyClonedEntityInfoVec(v []DestroyClonedEntityInfoProto)`

SetDestroyClonedEntityInfoVec sets DestroyClonedEntityInfoVec field to given value.

### HasDestroyClonedEntityInfoVec

`func (o *DestroyClonedVMTaskInfoProto) HasDestroyClonedEntityInfoVec() bool`

HasDestroyClonedEntityInfoVec returns a boolean if a field has been set.

### SetDestroyClonedEntityInfoVecNil

`func (o *DestroyClonedVMTaskInfoProto) SetDestroyClonedEntityInfoVecNil(b bool)`

 SetDestroyClonedEntityInfoVecNil sets the value for DestroyClonedEntityInfoVec to be an explicit nil

### UnsetDestroyClonedEntityInfoVec
`func (o *DestroyClonedVMTaskInfoProto) UnsetDestroyClonedEntityInfoVec()`

UnsetDestroyClonedEntityInfoVec ensures that no value is present for DestroyClonedEntityInfoVec, not even an explicit nil
### GetType

`func (o *DestroyClonedVMTaskInfoProto) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DestroyClonedVMTaskInfoProto) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DestroyClonedVMTaskInfoProto) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *DestroyClonedVMTaskInfoProto) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *DestroyClonedVMTaskInfoProto) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *DestroyClonedVMTaskInfoProto) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetViewDeleted

`func (o *DestroyClonedVMTaskInfoProto) GetViewDeleted() bool`

GetViewDeleted returns the ViewDeleted field if non-nil, zero value otherwise.

### GetViewDeletedOk

`func (o *DestroyClonedVMTaskInfoProto) GetViewDeletedOk() (*bool, bool)`

GetViewDeletedOk returns a tuple with the ViewDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewDeleted

`func (o *DestroyClonedVMTaskInfoProto) SetViewDeleted(v bool)`

SetViewDeleted sets ViewDeleted field to given value.

### HasViewDeleted

`func (o *DestroyClonedVMTaskInfoProto) HasViewDeleted() bool`

HasViewDeleted returns a boolean if a field has been set.

### SetViewDeletedNil

`func (o *DestroyClonedVMTaskInfoProto) SetViewDeletedNil(b bool)`

 SetViewDeletedNil sets the value for ViewDeleted to be an explicit nil

### UnsetViewDeleted
`func (o *DestroyClonedVMTaskInfoProto) UnsetViewDeleted()`

UnsetViewDeleted ensures that no value is present for ViewDeleted, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


