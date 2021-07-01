# ProtectionSourceNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationNodes** | Pointer to **[]map[string]interface{}** | Array of Child Subtrees.  Specifies the child subtree used to store additional application-level Objects. Different environments use the subtree to store application-level information. For example for SQL Server, this subtree stores the SQL Server instances running on a VM. | [optional] 
**EntityPaginationParameters** | Pointer to [**PaginationParameters**](PaginationParameters.md) |  | [optional] 
**EntityPermissionInfo** | Pointer to [**EntityPermissionInformation**](EntityPermissionInformation.md) |  | [optional] 
**LogicalSize** | Pointer to **NullableInt64** | Specifies the logical size of the data in bytes for the Object on this node. Presence of this field indicates this node is a leaf node. | [optional] 
**Nodes** | Pointer to **[]map[string]interface{}** | Array of Child Nodes.  Specifies children of the current node in the Protection Sources hierarchy. When representing Objects in memory, the entire Object subtree hierarchy is represented. You can use this subtree to navigate down the Object hierarchy. | [optional] 
**ProtectedSourcesSummary** | Pointer to [**[]AggregatedSubtreeInfo**](AggregatedSubtreeInfo.md) | Array of Protected Objects.  Specifies aggregated information about all the child Objects of this node that are currently protected by a Protection Job. There is one entry for each environment that is being backed up. The aggregated information for the Object hierarchy&#39;s environment will be available at the 0th index of the vector. | [optional] 
**ProtectionSource** | Pointer to [**NullableProtectionSource**](ProtectionSource.md) | Specifies the Protection Source for the current node. | [optional] 
**RegistrationInfo** | Pointer to [**NullableRegisteredSourceInfo**](RegisteredSourceInfo.md) | Specifies registration information for a root node in a Protection Sources tree. A root node represents a registered Source on the Cohesity Cluster, such as a vCenter Server. | [optional] 
**UnprotectedSourcesSummary** | Pointer to [**[]AggregatedSubtreeInfo**](AggregatedSubtreeInfo.md) | Array of Unprotected Sources.  Specifies aggregated information about all the child Objects of this node that are not protected by any Protection Jobs. The aggregated information for the Objects hierarchy&#39;s environment will be available at the 0th index of the vector. NOTE: This list includes Objects that were protected at some point in the past but are no longer actively protected. Snapshots containing these Objects may even exist on the Cohesity Cluster and be available to recover from. | [optional] 

## Methods

### NewProtectionSourceNode

`func NewProtectionSourceNode() *ProtectionSourceNode`

NewProtectionSourceNode instantiates a new ProtectionSourceNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectionSourceNodeWithDefaults

`func NewProtectionSourceNodeWithDefaults() *ProtectionSourceNode`

NewProtectionSourceNodeWithDefaults instantiates a new ProtectionSourceNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationNodes

`func (o *ProtectionSourceNode) GetApplicationNodes() []map[string]interface{}`

GetApplicationNodes returns the ApplicationNodes field if non-nil, zero value otherwise.

### GetApplicationNodesOk

`func (o *ProtectionSourceNode) GetApplicationNodesOk() (*[]map[string]interface{}, bool)`

GetApplicationNodesOk returns a tuple with the ApplicationNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationNodes

`func (o *ProtectionSourceNode) SetApplicationNodes(v []map[string]interface{})`

SetApplicationNodes sets ApplicationNodes field to given value.

### HasApplicationNodes

`func (o *ProtectionSourceNode) HasApplicationNodes() bool`

HasApplicationNodes returns a boolean if a field has been set.

### SetApplicationNodesNil

`func (o *ProtectionSourceNode) SetApplicationNodesNil(b bool)`

 SetApplicationNodesNil sets the value for ApplicationNodes to be an explicit nil

### UnsetApplicationNodes
`func (o *ProtectionSourceNode) UnsetApplicationNodes()`

UnsetApplicationNodes ensures that no value is present for ApplicationNodes, not even an explicit nil
### GetEntityPaginationParameters

`func (o *ProtectionSourceNode) GetEntityPaginationParameters() PaginationParameters`

GetEntityPaginationParameters returns the EntityPaginationParameters field if non-nil, zero value otherwise.

### GetEntityPaginationParametersOk

`func (o *ProtectionSourceNode) GetEntityPaginationParametersOk() (*PaginationParameters, bool)`

GetEntityPaginationParametersOk returns a tuple with the EntityPaginationParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityPaginationParameters

`func (o *ProtectionSourceNode) SetEntityPaginationParameters(v PaginationParameters)`

SetEntityPaginationParameters sets EntityPaginationParameters field to given value.

### HasEntityPaginationParameters

`func (o *ProtectionSourceNode) HasEntityPaginationParameters() bool`

HasEntityPaginationParameters returns a boolean if a field has been set.

### GetEntityPermissionInfo

`func (o *ProtectionSourceNode) GetEntityPermissionInfo() EntityPermissionInformation`

GetEntityPermissionInfo returns the EntityPermissionInfo field if non-nil, zero value otherwise.

### GetEntityPermissionInfoOk

`func (o *ProtectionSourceNode) GetEntityPermissionInfoOk() (*EntityPermissionInformation, bool)`

GetEntityPermissionInfoOk returns a tuple with the EntityPermissionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityPermissionInfo

`func (o *ProtectionSourceNode) SetEntityPermissionInfo(v EntityPermissionInformation)`

SetEntityPermissionInfo sets EntityPermissionInfo field to given value.

### HasEntityPermissionInfo

`func (o *ProtectionSourceNode) HasEntityPermissionInfo() bool`

HasEntityPermissionInfo returns a boolean if a field has been set.

### GetLogicalSize

`func (o *ProtectionSourceNode) GetLogicalSize() int64`

GetLogicalSize returns the LogicalSize field if non-nil, zero value otherwise.

### GetLogicalSizeOk

`func (o *ProtectionSourceNode) GetLogicalSizeOk() (*int64, bool)`

GetLogicalSizeOk returns a tuple with the LogicalSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalSize

`func (o *ProtectionSourceNode) SetLogicalSize(v int64)`

SetLogicalSize sets LogicalSize field to given value.

### HasLogicalSize

`func (o *ProtectionSourceNode) HasLogicalSize() bool`

HasLogicalSize returns a boolean if a field has been set.

### SetLogicalSizeNil

`func (o *ProtectionSourceNode) SetLogicalSizeNil(b bool)`

 SetLogicalSizeNil sets the value for LogicalSize to be an explicit nil

### UnsetLogicalSize
`func (o *ProtectionSourceNode) UnsetLogicalSize()`

UnsetLogicalSize ensures that no value is present for LogicalSize, not even an explicit nil
### GetNodes

`func (o *ProtectionSourceNode) GetNodes() []map[string]interface{}`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *ProtectionSourceNode) GetNodesOk() (*[]map[string]interface{}, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *ProtectionSourceNode) SetNodes(v []map[string]interface{})`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *ProtectionSourceNode) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### SetNodesNil

`func (o *ProtectionSourceNode) SetNodesNil(b bool)`

 SetNodesNil sets the value for Nodes to be an explicit nil

### UnsetNodes
`func (o *ProtectionSourceNode) UnsetNodes()`

UnsetNodes ensures that no value is present for Nodes, not even an explicit nil
### GetProtectedSourcesSummary

`func (o *ProtectionSourceNode) GetProtectedSourcesSummary() []AggregatedSubtreeInfo`

GetProtectedSourcesSummary returns the ProtectedSourcesSummary field if non-nil, zero value otherwise.

### GetProtectedSourcesSummaryOk

`func (o *ProtectionSourceNode) GetProtectedSourcesSummaryOk() (*[]AggregatedSubtreeInfo, bool)`

GetProtectedSourcesSummaryOk returns a tuple with the ProtectedSourcesSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectedSourcesSummary

`func (o *ProtectionSourceNode) SetProtectedSourcesSummary(v []AggregatedSubtreeInfo)`

SetProtectedSourcesSummary sets ProtectedSourcesSummary field to given value.

### HasProtectedSourcesSummary

`func (o *ProtectionSourceNode) HasProtectedSourcesSummary() bool`

HasProtectedSourcesSummary returns a boolean if a field has been set.

### SetProtectedSourcesSummaryNil

`func (o *ProtectionSourceNode) SetProtectedSourcesSummaryNil(b bool)`

 SetProtectedSourcesSummaryNil sets the value for ProtectedSourcesSummary to be an explicit nil

### UnsetProtectedSourcesSummary
`func (o *ProtectionSourceNode) UnsetProtectedSourcesSummary()`

UnsetProtectedSourcesSummary ensures that no value is present for ProtectedSourcesSummary, not even an explicit nil
### GetProtectionSource

`func (o *ProtectionSourceNode) GetProtectionSource() ProtectionSource`

GetProtectionSource returns the ProtectionSource field if non-nil, zero value otherwise.

### GetProtectionSourceOk

`func (o *ProtectionSourceNode) GetProtectionSourceOk() (*ProtectionSource, bool)`

GetProtectionSourceOk returns a tuple with the ProtectionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionSource

`func (o *ProtectionSourceNode) SetProtectionSource(v ProtectionSource)`

SetProtectionSource sets ProtectionSource field to given value.

### HasProtectionSource

`func (o *ProtectionSourceNode) HasProtectionSource() bool`

HasProtectionSource returns a boolean if a field has been set.

### SetProtectionSourceNil

`func (o *ProtectionSourceNode) SetProtectionSourceNil(b bool)`

 SetProtectionSourceNil sets the value for ProtectionSource to be an explicit nil

### UnsetProtectionSource
`func (o *ProtectionSourceNode) UnsetProtectionSource()`

UnsetProtectionSource ensures that no value is present for ProtectionSource, not even an explicit nil
### GetRegistrationInfo

`func (o *ProtectionSourceNode) GetRegistrationInfo() RegisteredSourceInfo`

GetRegistrationInfo returns the RegistrationInfo field if non-nil, zero value otherwise.

### GetRegistrationInfoOk

`func (o *ProtectionSourceNode) GetRegistrationInfoOk() (*RegisteredSourceInfo, bool)`

GetRegistrationInfoOk returns a tuple with the RegistrationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationInfo

`func (o *ProtectionSourceNode) SetRegistrationInfo(v RegisteredSourceInfo)`

SetRegistrationInfo sets RegistrationInfo field to given value.

### HasRegistrationInfo

`func (o *ProtectionSourceNode) HasRegistrationInfo() bool`

HasRegistrationInfo returns a boolean if a field has been set.

### SetRegistrationInfoNil

`func (o *ProtectionSourceNode) SetRegistrationInfoNil(b bool)`

 SetRegistrationInfoNil sets the value for RegistrationInfo to be an explicit nil

### UnsetRegistrationInfo
`func (o *ProtectionSourceNode) UnsetRegistrationInfo()`

UnsetRegistrationInfo ensures that no value is present for RegistrationInfo, not even an explicit nil
### GetUnprotectedSourcesSummary

`func (o *ProtectionSourceNode) GetUnprotectedSourcesSummary() []AggregatedSubtreeInfo`

GetUnprotectedSourcesSummary returns the UnprotectedSourcesSummary field if non-nil, zero value otherwise.

### GetUnprotectedSourcesSummaryOk

`func (o *ProtectionSourceNode) GetUnprotectedSourcesSummaryOk() (*[]AggregatedSubtreeInfo, bool)`

GetUnprotectedSourcesSummaryOk returns a tuple with the UnprotectedSourcesSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnprotectedSourcesSummary

`func (o *ProtectionSourceNode) SetUnprotectedSourcesSummary(v []AggregatedSubtreeInfo)`

SetUnprotectedSourcesSummary sets UnprotectedSourcesSummary field to given value.

### HasUnprotectedSourcesSummary

`func (o *ProtectionSourceNode) HasUnprotectedSourcesSummary() bool`

HasUnprotectedSourcesSummary returns a boolean if a field has been set.

### SetUnprotectedSourcesSummaryNil

`func (o *ProtectionSourceNode) SetUnprotectedSourcesSummaryNil(b bool)`

 SetUnprotectedSourcesSummaryNil sets the value for UnprotectedSourcesSummary to be an explicit nil

### UnsetUnprotectedSourcesSummary
`func (o *ProtectionSourceNode) UnsetUnprotectedSourcesSummary()`

UnsetUnprotectedSourcesSummary ensures that no value is present for UnprotectedSourcesSummary, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


