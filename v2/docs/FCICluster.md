# FCICluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to [**Error**](Error.md) |  | [optional] 
**Id** | Pointer to **NullableString** | Specifies the unique identifier of the FCI. | [optional] 
**IsSelectedByDefault** | Pointer to **NullableBool** | Indicates to the UI whether this FCI cluster should be selected by default | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the FCI. | [optional] 
**ResourceInfo** | Pointer to [**AppResource**](AppResource.md) |  | [optional] 
**Servers** | Pointer to [**[]SQLServer**](SQLServer.md) | Specifies the list of SQL servers which belongs to the given FCI.  | [optional] 

## Methods

### NewFCICluster

`func NewFCICluster() *FCICluster`

NewFCICluster instantiates a new FCICluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFCIClusterWithDefaults

`func NewFCIClusterWithDefaults() *FCICluster`

NewFCIClusterWithDefaults instantiates a new FCICluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *FCICluster) GetError() Error`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *FCICluster) GetErrorOk() (*Error, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *FCICluster) SetError(v Error)`

SetError sets Error field to given value.

### HasError

`func (o *FCICluster) HasError() bool`

HasError returns a boolean if a field has been set.

### GetId

`func (o *FCICluster) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FCICluster) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FCICluster) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FCICluster) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *FCICluster) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *FCICluster) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIsSelectedByDefault

`func (o *FCICluster) GetIsSelectedByDefault() bool`

GetIsSelectedByDefault returns the IsSelectedByDefault field if non-nil, zero value otherwise.

### GetIsSelectedByDefaultOk

`func (o *FCICluster) GetIsSelectedByDefaultOk() (*bool, bool)`

GetIsSelectedByDefaultOk returns a tuple with the IsSelectedByDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSelectedByDefault

`func (o *FCICluster) SetIsSelectedByDefault(v bool)`

SetIsSelectedByDefault sets IsSelectedByDefault field to given value.

### HasIsSelectedByDefault

`func (o *FCICluster) HasIsSelectedByDefault() bool`

HasIsSelectedByDefault returns a boolean if a field has been set.

### SetIsSelectedByDefaultNil

`func (o *FCICluster) SetIsSelectedByDefaultNil(b bool)`

 SetIsSelectedByDefaultNil sets the value for IsSelectedByDefault to be an explicit nil

### UnsetIsSelectedByDefault
`func (o *FCICluster) UnsetIsSelectedByDefault()`

UnsetIsSelectedByDefault ensures that no value is present for IsSelectedByDefault, not even an explicit nil
### GetName

`func (o *FCICluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FCICluster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FCICluster) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FCICluster) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *FCICluster) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *FCICluster) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetResourceInfo

`func (o *FCICluster) GetResourceInfo() AppResource`

GetResourceInfo returns the ResourceInfo field if non-nil, zero value otherwise.

### GetResourceInfoOk

`func (o *FCICluster) GetResourceInfoOk() (*AppResource, bool)`

GetResourceInfoOk returns a tuple with the ResourceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceInfo

`func (o *FCICluster) SetResourceInfo(v AppResource)`

SetResourceInfo sets ResourceInfo field to given value.

### HasResourceInfo

`func (o *FCICluster) HasResourceInfo() bool`

HasResourceInfo returns a boolean if a field has been set.

### GetServers

`func (o *FCICluster) GetServers() []SQLServer`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *FCICluster) GetServersOk() (*[]SQLServer, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *FCICluster) SetServers(v []SQLServer)`

SetServers sets Servers field to given value.

### HasServers

`func (o *FCICluster) HasServers() bool`

HasServers returns a boolean if a field has been set.

### SetServersNil

`func (o *FCICluster) SetServersNil(b bool)`

 SetServersNil sets the value for Servers to be an explicit nil

### UnsetServers
`func (o *FCICluster) UnsetServers()`

UnsetServers ensures that no value is present for Servers, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


