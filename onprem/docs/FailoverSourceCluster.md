# FailoverSourceCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableInt64** | Specifies the source cluster Id involved in failover operation. | 
**IncarnationId** | Pointer to **NullableInt64** | Specifies the source cluster incarnation Id involved in failover operation. | [optional] [readonly] 
**ProtectionGroupId** | Pointer to **NullableString** | Specifies the protection group Id involved in failover operation. | [optional] [readonly] 
**ViewId** | Pointer to **NullableInt64** | If failover is initiated by view based orchastrator, then this field specifies the local view id of source cluster which is being failed over. | [optional] [readonly] 

## Methods

### NewFailoverSourceCluster

`func NewFailoverSourceCluster(id NullableInt64, ) *FailoverSourceCluster`

NewFailoverSourceCluster instantiates a new FailoverSourceCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFailoverSourceClusterWithDefaults

`func NewFailoverSourceClusterWithDefaults() *FailoverSourceCluster`

NewFailoverSourceClusterWithDefaults instantiates a new FailoverSourceCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FailoverSourceCluster) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FailoverSourceCluster) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FailoverSourceCluster) SetId(v int64)`

SetId sets Id field to given value.


### SetIdNil

`func (o *FailoverSourceCluster) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *FailoverSourceCluster) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIncarnationId

`func (o *FailoverSourceCluster) GetIncarnationId() int64`

GetIncarnationId returns the IncarnationId field if non-nil, zero value otherwise.

### GetIncarnationIdOk

`func (o *FailoverSourceCluster) GetIncarnationIdOk() (*int64, bool)`

GetIncarnationIdOk returns a tuple with the IncarnationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncarnationId

`func (o *FailoverSourceCluster) SetIncarnationId(v int64)`

SetIncarnationId sets IncarnationId field to given value.

### HasIncarnationId

`func (o *FailoverSourceCluster) HasIncarnationId() bool`

HasIncarnationId returns a boolean if a field has been set.

### SetIncarnationIdNil

`func (o *FailoverSourceCluster) SetIncarnationIdNil(b bool)`

 SetIncarnationIdNil sets the value for IncarnationId to be an explicit nil

### UnsetIncarnationId
`func (o *FailoverSourceCluster) UnsetIncarnationId()`

UnsetIncarnationId ensures that no value is present for IncarnationId, not even an explicit nil
### GetProtectionGroupId

`func (o *FailoverSourceCluster) GetProtectionGroupId() string`

GetProtectionGroupId returns the ProtectionGroupId field if non-nil, zero value otherwise.

### GetProtectionGroupIdOk

`func (o *FailoverSourceCluster) GetProtectionGroupIdOk() (*string, bool)`

GetProtectionGroupIdOk returns a tuple with the ProtectionGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupId

`func (o *FailoverSourceCluster) SetProtectionGroupId(v string)`

SetProtectionGroupId sets ProtectionGroupId field to given value.

### HasProtectionGroupId

`func (o *FailoverSourceCluster) HasProtectionGroupId() bool`

HasProtectionGroupId returns a boolean if a field has been set.

### SetProtectionGroupIdNil

`func (o *FailoverSourceCluster) SetProtectionGroupIdNil(b bool)`

 SetProtectionGroupIdNil sets the value for ProtectionGroupId to be an explicit nil

### UnsetProtectionGroupId
`func (o *FailoverSourceCluster) UnsetProtectionGroupId()`

UnsetProtectionGroupId ensures that no value is present for ProtectionGroupId, not even an explicit nil
### GetViewId

`func (o *FailoverSourceCluster) GetViewId() int64`

GetViewId returns the ViewId field if non-nil, zero value otherwise.

### GetViewIdOk

`func (o *FailoverSourceCluster) GetViewIdOk() (*int64, bool)`

GetViewIdOk returns a tuple with the ViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewId

`func (o *FailoverSourceCluster) SetViewId(v int64)`

SetViewId sets ViewId field to given value.

### HasViewId

`func (o *FailoverSourceCluster) HasViewId() bool`

HasViewId returns a boolean if a field has been set.

### SetViewIdNil

`func (o *FailoverSourceCluster) SetViewIdNil(b bool)`

 SetViewIdNil sets the value for ViewId to be an explicit nil

### UnsetViewId
`func (o *FailoverSourceCluster) UnsetViewId()`

UnsetViewId ensures that no value is present for ViewId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


