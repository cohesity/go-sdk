# ImportConfigurations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveDirectories** | Pointer to **[]string** | Selective import of active directories. | [optional] 
**All** | Pointer to **[]string** | List of which entities to import all. | [optional] 
**Clusters** | Pointer to **[]int64** | Selective import certain cluster. | [optional] 
**File** | Pointer to **NullableString** | File is the config file. | [optional] 
**Groups** | Pointer to **[]string** | Selective import certain groups. | [optional] 
**Partitions** | Pointer to **[]int64** | Selective import of Partition. | [optional] 
**PrincipalSources** | Pointer to **[]string** | Selective import of principal sources. | [optional] 
**ProtectionJobs** | Pointer to **[]int64** | Selective import of protection jobs. | [optional] 
**ProtectionPolicies** | Pointer to **[]string** | Selective import of protection policies. | [optional] 
**ProtectionSources** | Pointer to **[]int64** | Selective import of protection sources. | [optional] 
**RemoteClusters** | Pointer to **[]int64** | Selective import certain remote clusters. | [optional] 
**Roles** | Pointer to **[]string** | Selective import certain roles (by username). | [optional] 
**Sql** | Pointer to **[]int64** | Selective import of sql. | [optional] 
**Users** | Pointer to **[]string** | Selective import certain users. | [optional] 
**Vaults** | Pointer to **[]int64** | Selective import certain vaults. | [optional] 
**ViewBoxes** | Pointer to **[]int64** | Selective import certain Storage Domains (View Boxes). | [optional] 
**Views** | Pointer to **[]int64** | Selective import of views. | [optional] 

## Methods

### NewImportConfigurations

`func NewImportConfigurations() *ImportConfigurations`

NewImportConfigurations instantiates a new ImportConfigurations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportConfigurationsWithDefaults

`func NewImportConfigurationsWithDefaults() *ImportConfigurations`

NewImportConfigurationsWithDefaults instantiates a new ImportConfigurations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveDirectories

`func (o *ImportConfigurations) GetActiveDirectories() []string`

GetActiveDirectories returns the ActiveDirectories field if non-nil, zero value otherwise.

### GetActiveDirectoriesOk

`func (o *ImportConfigurations) GetActiveDirectoriesOk() (*[]string, bool)`

GetActiveDirectoriesOk returns a tuple with the ActiveDirectories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveDirectories

`func (o *ImportConfigurations) SetActiveDirectories(v []string)`

SetActiveDirectories sets ActiveDirectories field to given value.

### HasActiveDirectories

`func (o *ImportConfigurations) HasActiveDirectories() bool`

HasActiveDirectories returns a boolean if a field has been set.

### SetActiveDirectoriesNil

`func (o *ImportConfigurations) SetActiveDirectoriesNil(b bool)`

 SetActiveDirectoriesNil sets the value for ActiveDirectories to be an explicit nil

### UnsetActiveDirectories
`func (o *ImportConfigurations) UnsetActiveDirectories()`

UnsetActiveDirectories ensures that no value is present for ActiveDirectories, not even an explicit nil
### GetAll

`func (o *ImportConfigurations) GetAll() []string`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *ImportConfigurations) GetAllOk() (*[]string, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *ImportConfigurations) SetAll(v []string)`

SetAll sets All field to given value.

### HasAll

`func (o *ImportConfigurations) HasAll() bool`

HasAll returns a boolean if a field has been set.

### SetAllNil

`func (o *ImportConfigurations) SetAllNil(b bool)`

 SetAllNil sets the value for All to be an explicit nil

### UnsetAll
`func (o *ImportConfigurations) UnsetAll()`

UnsetAll ensures that no value is present for All, not even an explicit nil
### GetClusters

`func (o *ImportConfigurations) GetClusters() []int64`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *ImportConfigurations) GetClustersOk() (*[]int64, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *ImportConfigurations) SetClusters(v []int64)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *ImportConfigurations) HasClusters() bool`

HasClusters returns a boolean if a field has been set.

### SetClustersNil

`func (o *ImportConfigurations) SetClustersNil(b bool)`

 SetClustersNil sets the value for Clusters to be an explicit nil

### UnsetClusters
`func (o *ImportConfigurations) UnsetClusters()`

UnsetClusters ensures that no value is present for Clusters, not even an explicit nil
### GetFile

`func (o *ImportConfigurations) GetFile() string`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *ImportConfigurations) GetFileOk() (*string, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *ImportConfigurations) SetFile(v string)`

SetFile sets File field to given value.

### HasFile

`func (o *ImportConfigurations) HasFile() bool`

HasFile returns a boolean if a field has been set.

### SetFileNil

`func (o *ImportConfigurations) SetFileNil(b bool)`

 SetFileNil sets the value for File to be an explicit nil

### UnsetFile
`func (o *ImportConfigurations) UnsetFile()`

UnsetFile ensures that no value is present for File, not even an explicit nil
### GetGroups

`func (o *ImportConfigurations) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *ImportConfigurations) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *ImportConfigurations) SetGroups(v []string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *ImportConfigurations) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### SetGroupsNil

`func (o *ImportConfigurations) SetGroupsNil(b bool)`

 SetGroupsNil sets the value for Groups to be an explicit nil

### UnsetGroups
`func (o *ImportConfigurations) UnsetGroups()`

UnsetGroups ensures that no value is present for Groups, not even an explicit nil
### GetPartitions

`func (o *ImportConfigurations) GetPartitions() []int64`

GetPartitions returns the Partitions field if non-nil, zero value otherwise.

### GetPartitionsOk

`func (o *ImportConfigurations) GetPartitionsOk() (*[]int64, bool)`

GetPartitionsOk returns a tuple with the Partitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitions

`func (o *ImportConfigurations) SetPartitions(v []int64)`

SetPartitions sets Partitions field to given value.

### HasPartitions

`func (o *ImportConfigurations) HasPartitions() bool`

HasPartitions returns a boolean if a field has been set.

### SetPartitionsNil

`func (o *ImportConfigurations) SetPartitionsNil(b bool)`

 SetPartitionsNil sets the value for Partitions to be an explicit nil

### UnsetPartitions
`func (o *ImportConfigurations) UnsetPartitions()`

UnsetPartitions ensures that no value is present for Partitions, not even an explicit nil
### GetPrincipalSources

`func (o *ImportConfigurations) GetPrincipalSources() []string`

GetPrincipalSources returns the PrincipalSources field if non-nil, zero value otherwise.

### GetPrincipalSourcesOk

`func (o *ImportConfigurations) GetPrincipalSourcesOk() (*[]string, bool)`

GetPrincipalSourcesOk returns a tuple with the PrincipalSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalSources

`func (o *ImportConfigurations) SetPrincipalSources(v []string)`

SetPrincipalSources sets PrincipalSources field to given value.

### HasPrincipalSources

`func (o *ImportConfigurations) HasPrincipalSources() bool`

HasPrincipalSources returns a boolean if a field has been set.

### SetPrincipalSourcesNil

`func (o *ImportConfigurations) SetPrincipalSourcesNil(b bool)`

 SetPrincipalSourcesNil sets the value for PrincipalSources to be an explicit nil

### UnsetPrincipalSources
`func (o *ImportConfigurations) UnsetPrincipalSources()`

UnsetPrincipalSources ensures that no value is present for PrincipalSources, not even an explicit nil
### GetProtectionJobs

`func (o *ImportConfigurations) GetProtectionJobs() []int64`

GetProtectionJobs returns the ProtectionJobs field if non-nil, zero value otherwise.

### GetProtectionJobsOk

`func (o *ImportConfigurations) GetProtectionJobsOk() (*[]int64, bool)`

GetProtectionJobsOk returns a tuple with the ProtectionJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionJobs

`func (o *ImportConfigurations) SetProtectionJobs(v []int64)`

SetProtectionJobs sets ProtectionJobs field to given value.

### HasProtectionJobs

`func (o *ImportConfigurations) HasProtectionJobs() bool`

HasProtectionJobs returns a boolean if a field has been set.

### SetProtectionJobsNil

`func (o *ImportConfigurations) SetProtectionJobsNil(b bool)`

 SetProtectionJobsNil sets the value for ProtectionJobs to be an explicit nil

### UnsetProtectionJobs
`func (o *ImportConfigurations) UnsetProtectionJobs()`

UnsetProtectionJobs ensures that no value is present for ProtectionJobs, not even an explicit nil
### GetProtectionPolicies

`func (o *ImportConfigurations) GetProtectionPolicies() []string`

GetProtectionPolicies returns the ProtectionPolicies field if non-nil, zero value otherwise.

### GetProtectionPoliciesOk

`func (o *ImportConfigurations) GetProtectionPoliciesOk() (*[]string, bool)`

GetProtectionPoliciesOk returns a tuple with the ProtectionPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionPolicies

`func (o *ImportConfigurations) SetProtectionPolicies(v []string)`

SetProtectionPolicies sets ProtectionPolicies field to given value.

### HasProtectionPolicies

`func (o *ImportConfigurations) HasProtectionPolicies() bool`

HasProtectionPolicies returns a boolean if a field has been set.

### SetProtectionPoliciesNil

`func (o *ImportConfigurations) SetProtectionPoliciesNil(b bool)`

 SetProtectionPoliciesNil sets the value for ProtectionPolicies to be an explicit nil

### UnsetProtectionPolicies
`func (o *ImportConfigurations) UnsetProtectionPolicies()`

UnsetProtectionPolicies ensures that no value is present for ProtectionPolicies, not even an explicit nil
### GetProtectionSources

`func (o *ImportConfigurations) GetProtectionSources() []int64`

GetProtectionSources returns the ProtectionSources field if non-nil, zero value otherwise.

### GetProtectionSourcesOk

`func (o *ImportConfigurations) GetProtectionSourcesOk() (*[]int64, bool)`

GetProtectionSourcesOk returns a tuple with the ProtectionSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionSources

`func (o *ImportConfigurations) SetProtectionSources(v []int64)`

SetProtectionSources sets ProtectionSources field to given value.

### HasProtectionSources

`func (o *ImportConfigurations) HasProtectionSources() bool`

HasProtectionSources returns a boolean if a field has been set.

### SetProtectionSourcesNil

`func (o *ImportConfigurations) SetProtectionSourcesNil(b bool)`

 SetProtectionSourcesNil sets the value for ProtectionSources to be an explicit nil

### UnsetProtectionSources
`func (o *ImportConfigurations) UnsetProtectionSources()`

UnsetProtectionSources ensures that no value is present for ProtectionSources, not even an explicit nil
### GetRemoteClusters

`func (o *ImportConfigurations) GetRemoteClusters() []int64`

GetRemoteClusters returns the RemoteClusters field if non-nil, zero value otherwise.

### GetRemoteClustersOk

`func (o *ImportConfigurations) GetRemoteClustersOk() (*[]int64, bool)`

GetRemoteClustersOk returns a tuple with the RemoteClusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteClusters

`func (o *ImportConfigurations) SetRemoteClusters(v []int64)`

SetRemoteClusters sets RemoteClusters field to given value.

### HasRemoteClusters

`func (o *ImportConfigurations) HasRemoteClusters() bool`

HasRemoteClusters returns a boolean if a field has been set.

### SetRemoteClustersNil

`func (o *ImportConfigurations) SetRemoteClustersNil(b bool)`

 SetRemoteClustersNil sets the value for RemoteClusters to be an explicit nil

### UnsetRemoteClusters
`func (o *ImportConfigurations) UnsetRemoteClusters()`

UnsetRemoteClusters ensures that no value is present for RemoteClusters, not even an explicit nil
### GetRoles

`func (o *ImportConfigurations) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *ImportConfigurations) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *ImportConfigurations) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *ImportConfigurations) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### SetRolesNil

`func (o *ImportConfigurations) SetRolesNil(b bool)`

 SetRolesNil sets the value for Roles to be an explicit nil

### UnsetRoles
`func (o *ImportConfigurations) UnsetRoles()`

UnsetRoles ensures that no value is present for Roles, not even an explicit nil
### GetSql

`func (o *ImportConfigurations) GetSql() []int64`

GetSql returns the Sql field if non-nil, zero value otherwise.

### GetSqlOk

`func (o *ImportConfigurations) GetSqlOk() (*[]int64, bool)`

GetSqlOk returns a tuple with the Sql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSql

`func (o *ImportConfigurations) SetSql(v []int64)`

SetSql sets Sql field to given value.

### HasSql

`func (o *ImportConfigurations) HasSql() bool`

HasSql returns a boolean if a field has been set.

### SetSqlNil

`func (o *ImportConfigurations) SetSqlNil(b bool)`

 SetSqlNil sets the value for Sql to be an explicit nil

### UnsetSql
`func (o *ImportConfigurations) UnsetSql()`

UnsetSql ensures that no value is present for Sql, not even an explicit nil
### GetUsers

`func (o *ImportConfigurations) GetUsers() []string`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *ImportConfigurations) GetUsersOk() (*[]string, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *ImportConfigurations) SetUsers(v []string)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *ImportConfigurations) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### SetUsersNil

`func (o *ImportConfigurations) SetUsersNil(b bool)`

 SetUsersNil sets the value for Users to be an explicit nil

### UnsetUsers
`func (o *ImportConfigurations) UnsetUsers()`

UnsetUsers ensures that no value is present for Users, not even an explicit nil
### GetVaults

`func (o *ImportConfigurations) GetVaults() []int64`

GetVaults returns the Vaults field if non-nil, zero value otherwise.

### GetVaultsOk

`func (o *ImportConfigurations) GetVaultsOk() (*[]int64, bool)`

GetVaultsOk returns a tuple with the Vaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaults

`func (o *ImportConfigurations) SetVaults(v []int64)`

SetVaults sets Vaults field to given value.

### HasVaults

`func (o *ImportConfigurations) HasVaults() bool`

HasVaults returns a boolean if a field has been set.

### SetVaultsNil

`func (o *ImportConfigurations) SetVaultsNil(b bool)`

 SetVaultsNil sets the value for Vaults to be an explicit nil

### UnsetVaults
`func (o *ImportConfigurations) UnsetVaults()`

UnsetVaults ensures that no value is present for Vaults, not even an explicit nil
### GetViewBoxes

`func (o *ImportConfigurations) GetViewBoxes() []int64`

GetViewBoxes returns the ViewBoxes field if non-nil, zero value otherwise.

### GetViewBoxesOk

`func (o *ImportConfigurations) GetViewBoxesOk() (*[]int64, bool)`

GetViewBoxesOk returns a tuple with the ViewBoxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewBoxes

`func (o *ImportConfigurations) SetViewBoxes(v []int64)`

SetViewBoxes sets ViewBoxes field to given value.

### HasViewBoxes

`func (o *ImportConfigurations) HasViewBoxes() bool`

HasViewBoxes returns a boolean if a field has been set.

### SetViewBoxesNil

`func (o *ImportConfigurations) SetViewBoxesNil(b bool)`

 SetViewBoxesNil sets the value for ViewBoxes to be an explicit nil

### UnsetViewBoxes
`func (o *ImportConfigurations) UnsetViewBoxes()`

UnsetViewBoxes ensures that no value is present for ViewBoxes, not even an explicit nil
### GetViews

`func (o *ImportConfigurations) GetViews() []int64`

GetViews returns the Views field if non-nil, zero value otherwise.

### GetViewsOk

`func (o *ImportConfigurations) GetViewsOk() (*[]int64, bool)`

GetViewsOk returns a tuple with the Views field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViews

`func (o *ImportConfigurations) SetViews(v []int64)`

SetViews sets Views field to given value.

### HasViews

`func (o *ImportConfigurations) HasViews() bool`

HasViews returns a boolean if a field has been set.

### SetViewsNil

`func (o *ImportConfigurations) SetViewsNil(b bool)`

 SetViewsNil sets the value for Views to be an explicit nil

### UnsetViews
`func (o *ImportConfigurations) UnsetViews()`

UnsetViews ensures that no value is present for Views, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


