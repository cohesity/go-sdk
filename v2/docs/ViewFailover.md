# ViewFailover

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsFailoverReady** | Pointer to **NullableBool** | Specifies if the view is ready for failover. | [optional] 
**RemoteClusterId** | Pointer to **NullableInt64** | Specifies the remote cluster id. | [optional] 
**RemoteClusterIncarnationId** | Pointer to **NullableInt64** | Specifies the remote cluster incarnation id. | [optional] 
**RemoteViewId** | Pointer to **NullableInt64** | Specifies the remote view id. | [optional] 
**ViewUid** | Pointer to **NullableString** | View uid. | [optional] 

## Methods

### NewViewFailover

`func NewViewFailover() *ViewFailover`

NewViewFailover instantiates a new ViewFailover object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewFailoverWithDefaults

`func NewViewFailoverWithDefaults() *ViewFailover`

NewViewFailoverWithDefaults instantiates a new ViewFailover object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsFailoverReady

`func (o *ViewFailover) GetIsFailoverReady() bool`

GetIsFailoverReady returns the IsFailoverReady field if non-nil, zero value otherwise.

### GetIsFailoverReadyOk

`func (o *ViewFailover) GetIsFailoverReadyOk() (*bool, bool)`

GetIsFailoverReadyOk returns a tuple with the IsFailoverReady field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFailoverReady

`func (o *ViewFailover) SetIsFailoverReady(v bool)`

SetIsFailoverReady sets IsFailoverReady field to given value.

### HasIsFailoverReady

`func (o *ViewFailover) HasIsFailoverReady() bool`

HasIsFailoverReady returns a boolean if a field has been set.

### SetIsFailoverReadyNil

`func (o *ViewFailover) SetIsFailoverReadyNil(b bool)`

 SetIsFailoverReadyNil sets the value for IsFailoverReady to be an explicit nil

### UnsetIsFailoverReady
`func (o *ViewFailover) UnsetIsFailoverReady()`

UnsetIsFailoverReady ensures that no value is present for IsFailoverReady, not even an explicit nil
### GetRemoteClusterId

`func (o *ViewFailover) GetRemoteClusterId() int64`

GetRemoteClusterId returns the RemoteClusterId field if non-nil, zero value otherwise.

### GetRemoteClusterIdOk

`func (o *ViewFailover) GetRemoteClusterIdOk() (*int64, bool)`

GetRemoteClusterIdOk returns a tuple with the RemoteClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteClusterId

`func (o *ViewFailover) SetRemoteClusterId(v int64)`

SetRemoteClusterId sets RemoteClusterId field to given value.

### HasRemoteClusterId

`func (o *ViewFailover) HasRemoteClusterId() bool`

HasRemoteClusterId returns a boolean if a field has been set.

### SetRemoteClusterIdNil

`func (o *ViewFailover) SetRemoteClusterIdNil(b bool)`

 SetRemoteClusterIdNil sets the value for RemoteClusterId to be an explicit nil

### UnsetRemoteClusterId
`func (o *ViewFailover) UnsetRemoteClusterId()`

UnsetRemoteClusterId ensures that no value is present for RemoteClusterId, not even an explicit nil
### GetRemoteClusterIncarnationId

`func (o *ViewFailover) GetRemoteClusterIncarnationId() int64`

GetRemoteClusterIncarnationId returns the RemoteClusterIncarnationId field if non-nil, zero value otherwise.

### GetRemoteClusterIncarnationIdOk

`func (o *ViewFailover) GetRemoteClusterIncarnationIdOk() (*int64, bool)`

GetRemoteClusterIncarnationIdOk returns a tuple with the RemoteClusterIncarnationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteClusterIncarnationId

`func (o *ViewFailover) SetRemoteClusterIncarnationId(v int64)`

SetRemoteClusterIncarnationId sets RemoteClusterIncarnationId field to given value.

### HasRemoteClusterIncarnationId

`func (o *ViewFailover) HasRemoteClusterIncarnationId() bool`

HasRemoteClusterIncarnationId returns a boolean if a field has been set.

### SetRemoteClusterIncarnationIdNil

`func (o *ViewFailover) SetRemoteClusterIncarnationIdNil(b bool)`

 SetRemoteClusterIncarnationIdNil sets the value for RemoteClusterIncarnationId to be an explicit nil

### UnsetRemoteClusterIncarnationId
`func (o *ViewFailover) UnsetRemoteClusterIncarnationId()`

UnsetRemoteClusterIncarnationId ensures that no value is present for RemoteClusterIncarnationId, not even an explicit nil
### GetRemoteViewId

`func (o *ViewFailover) GetRemoteViewId() int64`

GetRemoteViewId returns the RemoteViewId field if non-nil, zero value otherwise.

### GetRemoteViewIdOk

`func (o *ViewFailover) GetRemoteViewIdOk() (*int64, bool)`

GetRemoteViewIdOk returns a tuple with the RemoteViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteViewId

`func (o *ViewFailover) SetRemoteViewId(v int64)`

SetRemoteViewId sets RemoteViewId field to given value.

### HasRemoteViewId

`func (o *ViewFailover) HasRemoteViewId() bool`

HasRemoteViewId returns a boolean if a field has been set.

### SetRemoteViewIdNil

`func (o *ViewFailover) SetRemoteViewIdNil(b bool)`

 SetRemoteViewIdNil sets the value for RemoteViewId to be an explicit nil

### UnsetRemoteViewId
`func (o *ViewFailover) UnsetRemoteViewId()`

UnsetRemoteViewId ensures that no value is present for RemoteViewId, not even an explicit nil
### GetViewUid

`func (o *ViewFailover) GetViewUid() string`

GetViewUid returns the ViewUid field if non-nil, zero value otherwise.

### GetViewUidOk

`func (o *ViewFailover) GetViewUidOk() (*string, bool)`

GetViewUidOk returns a tuple with the ViewUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewUid

`func (o *ViewFailover) SetViewUid(v string)`

SetViewUid sets ViewUid field to given value.

### HasViewUid

`func (o *ViewFailover) HasViewUid() bool`

HasViewUid returns a boolean if a field has been set.

### SetViewUidNil

`func (o *ViewFailover) SetViewUidNil(b bool)`

 SetViewUidNil sets the value for ViewUid to be an explicit nil

### UnsetViewUid
`func (o *ViewFailover) UnsetViewUid()`

UnsetViewUid ensures that no value is present for ViewUid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


