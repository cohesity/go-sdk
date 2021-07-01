# ProtectionSourceTreeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Applications** | Pointer to [**[]ApplicationInfo**](ApplicationInfo.md) | Array of applications hierarchy registered on this node.  Specifies the application type and the list of instances of the application objects. For example for SQL Server, this list provides the SQL Server instances running on a VM or a Physical Server. | [optional] 
**EntityPermissionInfo** | Pointer to [**EntityPermissionInformation**](EntityPermissionInformation.md) |  | [optional] 
**LogicalSizeBytes** | Pointer to **NullableInt64** | Specifies the logical size of the Protection Source in bytes. | [optional] 
**RegistrationInfo** | Pointer to [**NullableRegisteredSourceInfo**](RegisteredSourceInfo.md) | Specifies registration information for a root node in a Protection Sources tree. A root node represents a registered Source on the Cohesity Cluster, such as a vCenter Server. | [optional] 
**RootNode** | Pointer to [**NullableProtectionSource**](ProtectionSource.md) | Specifies the Protection Source for the root node of the Protection Source tree. | [optional] 
**Stats** | Pointer to [**NullableProtectionSummary**](ProtectionSummary.md) | Specifies the stats of protection for a Protection Source Tree. | [optional] 
**StatsByEnv** | Pointer to [**[]ProtectionSummaryByEnv**](ProtectionSummaryByEnv.md) | Specifies the breakdown of the stats of protection by environment. overrideDescription: true | [optional] 

## Methods

### NewProtectionSourceTreeInfo

`func NewProtectionSourceTreeInfo() *ProtectionSourceTreeInfo`

NewProtectionSourceTreeInfo instantiates a new ProtectionSourceTreeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectionSourceTreeInfoWithDefaults

`func NewProtectionSourceTreeInfoWithDefaults() *ProtectionSourceTreeInfo`

NewProtectionSourceTreeInfoWithDefaults instantiates a new ProtectionSourceTreeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplications

`func (o *ProtectionSourceTreeInfo) GetApplications() []ApplicationInfo`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *ProtectionSourceTreeInfo) GetApplicationsOk() (*[]ApplicationInfo, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *ProtectionSourceTreeInfo) SetApplications(v []ApplicationInfo)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *ProtectionSourceTreeInfo) HasApplications() bool`

HasApplications returns a boolean if a field has been set.

### SetApplicationsNil

`func (o *ProtectionSourceTreeInfo) SetApplicationsNil(b bool)`

 SetApplicationsNil sets the value for Applications to be an explicit nil

### UnsetApplications
`func (o *ProtectionSourceTreeInfo) UnsetApplications()`

UnsetApplications ensures that no value is present for Applications, not even an explicit nil
### GetEntityPermissionInfo

`func (o *ProtectionSourceTreeInfo) GetEntityPermissionInfo() EntityPermissionInformation`

GetEntityPermissionInfo returns the EntityPermissionInfo field if non-nil, zero value otherwise.

### GetEntityPermissionInfoOk

`func (o *ProtectionSourceTreeInfo) GetEntityPermissionInfoOk() (*EntityPermissionInformation, bool)`

GetEntityPermissionInfoOk returns a tuple with the EntityPermissionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityPermissionInfo

`func (o *ProtectionSourceTreeInfo) SetEntityPermissionInfo(v EntityPermissionInformation)`

SetEntityPermissionInfo sets EntityPermissionInfo field to given value.

### HasEntityPermissionInfo

`func (o *ProtectionSourceTreeInfo) HasEntityPermissionInfo() bool`

HasEntityPermissionInfo returns a boolean if a field has been set.

### GetLogicalSizeBytes

`func (o *ProtectionSourceTreeInfo) GetLogicalSizeBytes() int64`

GetLogicalSizeBytes returns the LogicalSizeBytes field if non-nil, zero value otherwise.

### GetLogicalSizeBytesOk

`func (o *ProtectionSourceTreeInfo) GetLogicalSizeBytesOk() (*int64, bool)`

GetLogicalSizeBytesOk returns a tuple with the LogicalSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalSizeBytes

`func (o *ProtectionSourceTreeInfo) SetLogicalSizeBytes(v int64)`

SetLogicalSizeBytes sets LogicalSizeBytes field to given value.

### HasLogicalSizeBytes

`func (o *ProtectionSourceTreeInfo) HasLogicalSizeBytes() bool`

HasLogicalSizeBytes returns a boolean if a field has been set.

### SetLogicalSizeBytesNil

`func (o *ProtectionSourceTreeInfo) SetLogicalSizeBytesNil(b bool)`

 SetLogicalSizeBytesNil sets the value for LogicalSizeBytes to be an explicit nil

### UnsetLogicalSizeBytes
`func (o *ProtectionSourceTreeInfo) UnsetLogicalSizeBytes()`

UnsetLogicalSizeBytes ensures that no value is present for LogicalSizeBytes, not even an explicit nil
### GetRegistrationInfo

`func (o *ProtectionSourceTreeInfo) GetRegistrationInfo() RegisteredSourceInfo`

GetRegistrationInfo returns the RegistrationInfo field if non-nil, zero value otherwise.

### GetRegistrationInfoOk

`func (o *ProtectionSourceTreeInfo) GetRegistrationInfoOk() (*RegisteredSourceInfo, bool)`

GetRegistrationInfoOk returns a tuple with the RegistrationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationInfo

`func (o *ProtectionSourceTreeInfo) SetRegistrationInfo(v RegisteredSourceInfo)`

SetRegistrationInfo sets RegistrationInfo field to given value.

### HasRegistrationInfo

`func (o *ProtectionSourceTreeInfo) HasRegistrationInfo() bool`

HasRegistrationInfo returns a boolean if a field has been set.

### SetRegistrationInfoNil

`func (o *ProtectionSourceTreeInfo) SetRegistrationInfoNil(b bool)`

 SetRegistrationInfoNil sets the value for RegistrationInfo to be an explicit nil

### UnsetRegistrationInfo
`func (o *ProtectionSourceTreeInfo) UnsetRegistrationInfo()`

UnsetRegistrationInfo ensures that no value is present for RegistrationInfo, not even an explicit nil
### GetRootNode

`func (o *ProtectionSourceTreeInfo) GetRootNode() ProtectionSource`

GetRootNode returns the RootNode field if non-nil, zero value otherwise.

### GetRootNodeOk

`func (o *ProtectionSourceTreeInfo) GetRootNodeOk() (*ProtectionSource, bool)`

GetRootNodeOk returns a tuple with the RootNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootNode

`func (o *ProtectionSourceTreeInfo) SetRootNode(v ProtectionSource)`

SetRootNode sets RootNode field to given value.

### HasRootNode

`func (o *ProtectionSourceTreeInfo) HasRootNode() bool`

HasRootNode returns a boolean if a field has been set.

### SetRootNodeNil

`func (o *ProtectionSourceTreeInfo) SetRootNodeNil(b bool)`

 SetRootNodeNil sets the value for RootNode to be an explicit nil

### UnsetRootNode
`func (o *ProtectionSourceTreeInfo) UnsetRootNode()`

UnsetRootNode ensures that no value is present for RootNode, not even an explicit nil
### GetStats

`func (o *ProtectionSourceTreeInfo) GetStats() ProtectionSummary`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *ProtectionSourceTreeInfo) GetStatsOk() (*ProtectionSummary, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *ProtectionSourceTreeInfo) SetStats(v ProtectionSummary)`

SetStats sets Stats field to given value.

### HasStats

`func (o *ProtectionSourceTreeInfo) HasStats() bool`

HasStats returns a boolean if a field has been set.

### SetStatsNil

`func (o *ProtectionSourceTreeInfo) SetStatsNil(b bool)`

 SetStatsNil sets the value for Stats to be an explicit nil

### UnsetStats
`func (o *ProtectionSourceTreeInfo) UnsetStats()`

UnsetStats ensures that no value is present for Stats, not even an explicit nil
### GetStatsByEnv

`func (o *ProtectionSourceTreeInfo) GetStatsByEnv() []ProtectionSummaryByEnv`

GetStatsByEnv returns the StatsByEnv field if non-nil, zero value otherwise.

### GetStatsByEnvOk

`func (o *ProtectionSourceTreeInfo) GetStatsByEnvOk() (*[]ProtectionSummaryByEnv, bool)`

GetStatsByEnvOk returns a tuple with the StatsByEnv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatsByEnv

`func (o *ProtectionSourceTreeInfo) SetStatsByEnv(v []ProtectionSummaryByEnv)`

SetStatsByEnv sets StatsByEnv field to given value.

### HasStatsByEnv

`func (o *ProtectionSourceTreeInfo) HasStatsByEnv() bool`

HasStatsByEnv returns a boolean if a field has been set.

### SetStatsByEnvNil

`func (o *ProtectionSourceTreeInfo) SetStatsByEnvNil(b bool)`

 SetStatsByEnvNil sets the value for StatsByEnv to be an explicit nil

### UnsetStatsByEnv
`func (o *ProtectionSourceTreeInfo) UnsetStatsByEnv()`

UnsetStatsByEnv ensures that no value is present for StatsByEnv, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


