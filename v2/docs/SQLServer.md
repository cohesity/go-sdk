# SQLServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentInfo** | Pointer to [**AgentInformation**](AgentInformation.md) |  | [optional] 
**Error** | Pointer to [**Error**](Error.md) |  | [optional] 
**Id** | Pointer to **NullableString** | Specifies the unique identifier of the SQL server host. | [optional] 
**Instances** | Pointer to [**[]SQLServerInstance**](SQLServerInstance.md) | Specifies the list of all sql instances running inside the current SQL host. | [optional] 
**IsPrimary** | Pointer to **NullableBool** | Indicates whether this is a active node of a FCI cluster or hosts primary replica of a AAG group. | [optional] 
**IsSelectedByDefault** | Pointer to **NullableBool** | Indicates to the UI whether this server should be selected by default | [optional] 
**ResourceInfo** | Pointer to [**AppResource**](AppResource.md) |  | [optional] 

## Methods

### NewSQLServer

`func NewSQLServer() *SQLServer`

NewSQLServer instantiates a new SQLServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSQLServerWithDefaults

`func NewSQLServerWithDefaults() *SQLServer`

NewSQLServerWithDefaults instantiates a new SQLServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentInfo

`func (o *SQLServer) GetAgentInfo() AgentInformation`

GetAgentInfo returns the AgentInfo field if non-nil, zero value otherwise.

### GetAgentInfoOk

`func (o *SQLServer) GetAgentInfoOk() (*AgentInformation, bool)`

GetAgentInfoOk returns a tuple with the AgentInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentInfo

`func (o *SQLServer) SetAgentInfo(v AgentInformation)`

SetAgentInfo sets AgentInfo field to given value.

### HasAgentInfo

`func (o *SQLServer) HasAgentInfo() bool`

HasAgentInfo returns a boolean if a field has been set.

### GetError

`func (o *SQLServer) GetError() Error`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *SQLServer) GetErrorOk() (*Error, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *SQLServer) SetError(v Error)`

SetError sets Error field to given value.

### HasError

`func (o *SQLServer) HasError() bool`

HasError returns a boolean if a field has been set.

### GetId

`func (o *SQLServer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SQLServer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SQLServer) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SQLServer) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *SQLServer) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SQLServer) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetInstances

`func (o *SQLServer) GetInstances() []SQLServerInstance`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *SQLServer) GetInstancesOk() (*[]SQLServerInstance, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *SQLServer) SetInstances(v []SQLServerInstance)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *SQLServer) HasInstances() bool`

HasInstances returns a boolean if a field has been set.

### SetInstancesNil

`func (o *SQLServer) SetInstancesNil(b bool)`

 SetInstancesNil sets the value for Instances to be an explicit nil

### UnsetInstances
`func (o *SQLServer) UnsetInstances()`

UnsetInstances ensures that no value is present for Instances, not even an explicit nil
### GetIsPrimary

`func (o *SQLServer) GetIsPrimary() bool`

GetIsPrimary returns the IsPrimary field if non-nil, zero value otherwise.

### GetIsPrimaryOk

`func (o *SQLServer) GetIsPrimaryOk() (*bool, bool)`

GetIsPrimaryOk returns a tuple with the IsPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrimary

`func (o *SQLServer) SetIsPrimary(v bool)`

SetIsPrimary sets IsPrimary field to given value.

### HasIsPrimary

`func (o *SQLServer) HasIsPrimary() bool`

HasIsPrimary returns a boolean if a field has been set.

### SetIsPrimaryNil

`func (o *SQLServer) SetIsPrimaryNil(b bool)`

 SetIsPrimaryNil sets the value for IsPrimary to be an explicit nil

### UnsetIsPrimary
`func (o *SQLServer) UnsetIsPrimary()`

UnsetIsPrimary ensures that no value is present for IsPrimary, not even an explicit nil
### GetIsSelectedByDefault

`func (o *SQLServer) GetIsSelectedByDefault() bool`

GetIsSelectedByDefault returns the IsSelectedByDefault field if non-nil, zero value otherwise.

### GetIsSelectedByDefaultOk

`func (o *SQLServer) GetIsSelectedByDefaultOk() (*bool, bool)`

GetIsSelectedByDefaultOk returns a tuple with the IsSelectedByDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSelectedByDefault

`func (o *SQLServer) SetIsSelectedByDefault(v bool)`

SetIsSelectedByDefault sets IsSelectedByDefault field to given value.

### HasIsSelectedByDefault

`func (o *SQLServer) HasIsSelectedByDefault() bool`

HasIsSelectedByDefault returns a boolean if a field has been set.

### SetIsSelectedByDefaultNil

`func (o *SQLServer) SetIsSelectedByDefaultNil(b bool)`

 SetIsSelectedByDefaultNil sets the value for IsSelectedByDefault to be an explicit nil

### UnsetIsSelectedByDefault
`func (o *SQLServer) UnsetIsSelectedByDefault()`

UnsetIsSelectedByDefault ensures that no value is present for IsSelectedByDefault, not even an explicit nil
### GetResourceInfo

`func (o *SQLServer) GetResourceInfo() AppResource`

GetResourceInfo returns the ResourceInfo field if non-nil, zero value otherwise.

### GetResourceInfoOk

`func (o *SQLServer) GetResourceInfoOk() (*AppResource, bool)`

GetResourceInfoOk returns a tuple with the ResourceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceInfo

`func (o *SQLServer) SetResourceInfo(v AppResource)`

SetResourceInfo sets ResourceInfo field to given value.

### HasResourceInfo

`func (o *SQLServer) HasResourceInfo() bool`

HasResourceInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


