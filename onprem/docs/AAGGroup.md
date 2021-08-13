# AAGGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Specifies the unique identifier of the AGGGroup. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the AAG Group. | [optional] 
**ResourceInfo** | Pointer to [**AppResource**](AppResource.md) |  | [optional] 
**Servers** | Pointer to [**[]SQLServer**](SQLServer.md) | Specifies the list of SQL servers which belongs to the given AAG Group. | [optional] 
**FciClusters** | Pointer to [**[]FCICluster**](FCICluster.md) | Specifies the list of FCI clusters which belongs to the given AAG Group. | [optional] 

## Methods

### NewAAGGroup

`func NewAAGGroup() *AAGGroup`

NewAAGGroup instantiates a new AAGGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAAGGroupWithDefaults

`func NewAAGGroupWithDefaults() *AAGGroup`

NewAAGGroupWithDefaults instantiates a new AAGGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AAGGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AAGGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AAGGroup) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AAGGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *AAGGroup) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *AAGGroup) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *AAGGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AAGGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AAGGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AAGGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AAGGroup) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AAGGroup) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetResourceInfo

`func (o *AAGGroup) GetResourceInfo() AppResource`

GetResourceInfo returns the ResourceInfo field if non-nil, zero value otherwise.

### GetResourceInfoOk

`func (o *AAGGroup) GetResourceInfoOk() (*AppResource, bool)`

GetResourceInfoOk returns a tuple with the ResourceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceInfo

`func (o *AAGGroup) SetResourceInfo(v AppResource)`

SetResourceInfo sets ResourceInfo field to given value.

### HasResourceInfo

`func (o *AAGGroup) HasResourceInfo() bool`

HasResourceInfo returns a boolean if a field has been set.

### GetServers

`func (o *AAGGroup) GetServers() []SQLServer`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *AAGGroup) GetServersOk() (*[]SQLServer, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *AAGGroup) SetServers(v []SQLServer)`

SetServers sets Servers field to given value.

### HasServers

`func (o *AAGGroup) HasServers() bool`

HasServers returns a boolean if a field has been set.

### SetServersNil

`func (o *AAGGroup) SetServersNil(b bool)`

 SetServersNil sets the value for Servers to be an explicit nil

### UnsetServers
`func (o *AAGGroup) UnsetServers()`

UnsetServers ensures that no value is present for Servers, not even an explicit nil
### GetFciClusters

`func (o *AAGGroup) GetFciClusters() []FCICluster`

GetFciClusters returns the FciClusters field if non-nil, zero value otherwise.

### GetFciClustersOk

`func (o *AAGGroup) GetFciClustersOk() (*[]FCICluster, bool)`

GetFciClustersOk returns a tuple with the FciClusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFciClusters

`func (o *AAGGroup) SetFciClusters(v []FCICluster)`

SetFciClusters sets FciClusters field to given value.

### HasFciClusters

`func (o *AAGGroup) HasFciClusters() bool`

HasFciClusters returns a boolean if a field has been set.

### SetFciClustersNil

`func (o *AAGGroup) SetFciClustersNil(b bool)`

 SetFciClustersNil sets the value for FciClusters to be an explicit nil

### UnsetFciClusters
`func (o *AAGGroup) UnsetFciClusters()`

UnsetFciClusters ensures that no value is present for FciClusters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


