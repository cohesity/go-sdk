# RemoteAdapterProtectionGroupReplicationParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateView** | **NullableBool** | Specifies whether or not to create a remote view on replication cluster. | 
**ViewName** | Pointer to **NullableString** | Specifies the name of the remote view. By default the name will be same as the protected view. If a view with the specified name already exists on the remote cluster, the created name will have a suffix &#39;-1&#39;. | [optional] 

## Methods

### NewRemoteAdapterProtectionGroupReplicationParams

`func NewRemoteAdapterProtectionGroupReplicationParams(createView NullableBool, ) *RemoteAdapterProtectionGroupReplicationParams`

NewRemoteAdapterProtectionGroupReplicationParams instantiates a new RemoteAdapterProtectionGroupReplicationParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteAdapterProtectionGroupReplicationParamsWithDefaults

`func NewRemoteAdapterProtectionGroupReplicationParamsWithDefaults() *RemoteAdapterProtectionGroupReplicationParams`

NewRemoteAdapterProtectionGroupReplicationParamsWithDefaults instantiates a new RemoteAdapterProtectionGroupReplicationParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateView

`func (o *RemoteAdapterProtectionGroupReplicationParams) GetCreateView() bool`

GetCreateView returns the CreateView field if non-nil, zero value otherwise.

### GetCreateViewOk

`func (o *RemoteAdapterProtectionGroupReplicationParams) GetCreateViewOk() (*bool, bool)`

GetCreateViewOk returns a tuple with the CreateView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateView

`func (o *RemoteAdapterProtectionGroupReplicationParams) SetCreateView(v bool)`

SetCreateView sets CreateView field to given value.


### SetCreateViewNil

`func (o *RemoteAdapterProtectionGroupReplicationParams) SetCreateViewNil(b bool)`

 SetCreateViewNil sets the value for CreateView to be an explicit nil

### UnsetCreateView
`func (o *RemoteAdapterProtectionGroupReplicationParams) UnsetCreateView()`

UnsetCreateView ensures that no value is present for CreateView, not even an explicit nil
### GetViewName

`func (o *RemoteAdapterProtectionGroupReplicationParams) GetViewName() string`

GetViewName returns the ViewName field if non-nil, zero value otherwise.

### GetViewNameOk

`func (o *RemoteAdapterProtectionGroupReplicationParams) GetViewNameOk() (*string, bool)`

GetViewNameOk returns a tuple with the ViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewName

`func (o *RemoteAdapterProtectionGroupReplicationParams) SetViewName(v string)`

SetViewName sets ViewName field to given value.

### HasViewName

`func (o *RemoteAdapterProtectionGroupReplicationParams) HasViewName() bool`

HasViewName returns a boolean if a field has been set.

### SetViewNameNil

`func (o *RemoteAdapterProtectionGroupReplicationParams) SetViewNameNil(b bool)`

 SetViewNameNil sets the value for ViewName to be an explicit nil

### UnsetViewName
`func (o *RemoteAdapterProtectionGroupReplicationParams) UnsetViewName()`

UnsetViewName ensures that no value is present for ViewName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


