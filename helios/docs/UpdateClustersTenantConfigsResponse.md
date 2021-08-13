# UpdateClustersTenantConfigsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clusters** | Pointer to [**[]ClusterTenantConfig**](ClusterTenantConfig.md) | The list of clusters updated, with errors if any. | [optional] 

## Methods

### NewUpdateClustersTenantConfigsResponse

`func NewUpdateClustersTenantConfigsResponse() *UpdateClustersTenantConfigsResponse`

NewUpdateClustersTenantConfigsResponse instantiates a new UpdateClustersTenantConfigsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateClustersTenantConfigsResponseWithDefaults

`func NewUpdateClustersTenantConfigsResponseWithDefaults() *UpdateClustersTenantConfigsResponse`

NewUpdateClustersTenantConfigsResponseWithDefaults instantiates a new UpdateClustersTenantConfigsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusters

`func (o *UpdateClustersTenantConfigsResponse) GetClusters() []ClusterTenantConfig`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *UpdateClustersTenantConfigsResponse) GetClustersOk() (*[]ClusterTenantConfig, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *UpdateClustersTenantConfigsResponse) SetClusters(v []ClusterTenantConfig)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *UpdateClustersTenantConfigsResponse) HasClusters() bool`

HasClusters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


