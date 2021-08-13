# ViewProtectionGroupReplicationParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ViewNameConfigList** | Pointer to [**[]ReplicatedViewNameConfig**](ReplicatedViewNameConfig.md) | Specifies the list of remote view names for the protected views in the Protection Group. By default the names will be the same as the name of the protected view. | [optional] 

## Methods

### NewViewProtectionGroupReplicationParams

`func NewViewProtectionGroupReplicationParams() *ViewProtectionGroupReplicationParams`

NewViewProtectionGroupReplicationParams instantiates a new ViewProtectionGroupReplicationParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewProtectionGroupReplicationParamsWithDefaults

`func NewViewProtectionGroupReplicationParamsWithDefaults() *ViewProtectionGroupReplicationParams`

NewViewProtectionGroupReplicationParamsWithDefaults instantiates a new ViewProtectionGroupReplicationParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetViewNameConfigList

`func (o *ViewProtectionGroupReplicationParams) GetViewNameConfigList() []ReplicatedViewNameConfig`

GetViewNameConfigList returns the ViewNameConfigList field if non-nil, zero value otherwise.

### GetViewNameConfigListOk

`func (o *ViewProtectionGroupReplicationParams) GetViewNameConfigListOk() (*[]ReplicatedViewNameConfig, bool)`

GetViewNameConfigListOk returns a tuple with the ViewNameConfigList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewNameConfigList

`func (o *ViewProtectionGroupReplicationParams) SetViewNameConfigList(v []ReplicatedViewNameConfig)`

SetViewNameConfigList sets ViewNameConfigList field to given value.

### HasViewNameConfigList

`func (o *ViewProtectionGroupReplicationParams) HasViewNameConfigList() bool`

HasViewNameConfigList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


