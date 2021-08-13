# AssignClusterToTenantParamsBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterIdentifier** | **NullableString** | Specifies the list of cluster identifiers. The format is clusterId:clusterIncarnationId. | 
**Network** | [**TenantNetwork**](TenantNetwork.md) |  | 

## Methods

### NewAssignClusterToTenantParamsBody

`func NewAssignClusterToTenantParamsBody(clusterIdentifier NullableString, network TenantNetwork, ) *AssignClusterToTenantParamsBody`

NewAssignClusterToTenantParamsBody instantiates a new AssignClusterToTenantParamsBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssignClusterToTenantParamsBodyWithDefaults

`func NewAssignClusterToTenantParamsBodyWithDefaults() *AssignClusterToTenantParamsBody`

NewAssignClusterToTenantParamsBodyWithDefaults instantiates a new AssignClusterToTenantParamsBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterIdentifier

`func (o *AssignClusterToTenantParamsBody) GetClusterIdentifier() string`

GetClusterIdentifier returns the ClusterIdentifier field if non-nil, zero value otherwise.

### GetClusterIdentifierOk

`func (o *AssignClusterToTenantParamsBody) GetClusterIdentifierOk() (*string, bool)`

GetClusterIdentifierOk returns a tuple with the ClusterIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIdentifier

`func (o *AssignClusterToTenantParamsBody) SetClusterIdentifier(v string)`

SetClusterIdentifier sets ClusterIdentifier field to given value.


### SetClusterIdentifierNil

`func (o *AssignClusterToTenantParamsBody) SetClusterIdentifierNil(b bool)`

 SetClusterIdentifierNil sets the value for ClusterIdentifier to be an explicit nil

### UnsetClusterIdentifier
`func (o *AssignClusterToTenantParamsBody) UnsetClusterIdentifier()`

UnsetClusterIdentifier ensures that no value is present for ClusterIdentifier, not even an explicit nil
### GetNetwork

`func (o *AssignClusterToTenantParamsBody) GetNetwork() TenantNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *AssignClusterToTenantParamsBody) GetNetworkOk() (*TenantNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *AssignClusterToTenantParamsBody) SetNetwork(v TenantNetwork)`

SetNetwork sets Network field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


