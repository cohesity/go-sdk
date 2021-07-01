# ClusterNetworkingResourceInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Endpoints** | Pointer to [**[]ClusterNetworkingEndpoint**](ClusterNetworkingEndpoint.md) | The endpoints by which the resource is accessible. | [optional] 
**Type** | Pointer to **NullableString** | The type of the resource. | [optional] 

## Methods

### NewClusterNetworkingResourceInformation

`func NewClusterNetworkingResourceInformation() *ClusterNetworkingResourceInformation`

NewClusterNetworkingResourceInformation instantiates a new ClusterNetworkingResourceInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterNetworkingResourceInformationWithDefaults

`func NewClusterNetworkingResourceInformationWithDefaults() *ClusterNetworkingResourceInformation`

NewClusterNetworkingResourceInformationWithDefaults instantiates a new ClusterNetworkingResourceInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpoints

`func (o *ClusterNetworkingResourceInformation) GetEndpoints() []ClusterNetworkingEndpoint`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *ClusterNetworkingResourceInformation) GetEndpointsOk() (*[]ClusterNetworkingEndpoint, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *ClusterNetworkingResourceInformation) SetEndpoints(v []ClusterNetworkingEndpoint)`

SetEndpoints sets Endpoints field to given value.

### HasEndpoints

`func (o *ClusterNetworkingResourceInformation) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.

### SetEndpointsNil

`func (o *ClusterNetworkingResourceInformation) SetEndpointsNil(b bool)`

 SetEndpointsNil sets the value for Endpoints to be an explicit nil

### UnsetEndpoints
`func (o *ClusterNetworkingResourceInformation) UnsetEndpoints()`

UnsetEndpoints ensures that no value is present for Endpoints, not even an explicit nil
### GetType

`func (o *ClusterNetworkingResourceInformation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ClusterNetworkingResourceInformation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ClusterNetworkingResourceInformation) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ClusterNetworkingResourceInformation) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *ClusterNetworkingResourceInformation) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ClusterNetworkingResourceInformation) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


