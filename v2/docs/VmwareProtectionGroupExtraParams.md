# VmwareProtectionGroupExtraParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowParallelRuns** | Pointer to **NullableBool** | Specifies whether or not this job can have parallel runs. | [optional] 
**CloudMigration** | Pointer to **NullableBool** | Specifies whether or not to move the workload to the cloud. | [optional] 
**ExcludeFilters** | Pointer to [**[]VMFilter**](VMFilter.md) | Specifies the list of exclusion filters applied during the group creation or edit. These exclusion filters can be wildcard supported strings or regular expressions. Objects satisfying these filters will be excluded during backup and also auto protected objects will be ignored if filtered by any of the filters. | [optional] 
**ExcludeObjectIds** | Pointer to **[]int64** | Specifies the list of IDs of the objects to not be protected in this backup. This field only applies if provided object id is non leaf entity such as Tag or a folder. This can be used to ignore specific objects under a parent object which has been included for protection. | [optional] 
**ExcludeVmTagIds** | Pointer to **[][]int64** | Array of Arrays of VM Tag Ids that Specify VMs to Exclude. Optionally specify a list of VMs to exclude from protecting by listing Protection Source ids of VM Tags in this two dimensional array. Using this two dimensional array of Tag ids, the Cluster generates a list of VMs to exclude from protecting, which are derived from intersections of the inner arrays and union of the outer array, as shown by the following example. For example a Datacenter is selected to be protected but you want to exclude all the &#39;Former Employees&#39; VMs in the East and West but keep all the VMs for &#39;Former Employees&#39; in the South which are also stored in this Datacenter, by specifying the following tag id array: [ [1000, 2221], [1000, 3031] ], where 1000 is the &#39;Former Employee&#39; VM Tag id, 2221 is the &#39;East&#39; VM Tag id and 3031 is the &#39;West&#39; VM Tag id. The first inner array [1000, 2221] produces a list of VMs that are both tagged with &#39;Former Employees&#39; and &#39;East&#39; (an intersection). The second inner array [1000, 3031] produces a list of VMs that are both tagged with &#39;Former Employees&#39; and &#39;West&#39; (an intersection). The outer array combines the list of VMs from the two inner arrays. The list of resulting VMs are excluded from being protected this Job. | [optional] 
**LeverageHyperflexSnapshots** | Pointer to **NullableBool** | Whether to leverage the hyperflex based snapshots for this backup. To leverage hyperflex snapshots, it has to first be registered. If hyperflex based snapshots cannot be taken, backup will fallback to the default backup method. | [optional] 
**LeverageNutanixSnapshots** | Pointer to **NullableBool** | Whether to leverage the nutanix based snapshots for this backup. To leverage nutanix snapshots, it has to first be registered. If nutanix based snapshots cannot be taken, backup will fallback to the default backup method. | [optional] 
**LeverageStorageSnapshots** | Pointer to **NullableBool** | Whether to leverage the storage array based snapshots for this backup. To leverage storage snapshots, the storage array has to be registered as a source. If storage based snapshots can not be taken, backup will fallback to the default backup method. | [optional] 
**SourceId** | Pointer to **NullableInt64** | Specifies the id of the parent of the objects. | [optional] [readonly] 
**SourceName** | Pointer to **NullableString** | Specifies the name of the parent of the objects. | [optional] [readonly] 
**VmTagIds** | Pointer to **[][]int64** | Array of Array of VM Tag Ids that Specify VMs to Protect. Optionally specify a list of VMs to protect by listing Protection Source ids of VM Tags in this two dimensional array. Using this two dimensional array of Tag ids, the Cluster generates a list of VMs to protect which are derived from intersections of the inner arrays and union of the outer array, as shown by the following example. To protect only &#39;Eng&#39; VMs in the East and all the VMs in the West, specify the following tag id array: [ [1101, 2221], [3031] ], where 1101 is the &#39;Eng&#39; VM Tag id, 2221 is the &#39;East&#39; VM Tag id and 3031 is the &#39;West&#39; VM Tag id. The inner array [1101, 2221] produces a list of VMs that are both tagged with &#39;Eng&#39; and &#39;East&#39; (an intersection). The outer array combines the list from the inner array with list of VMs tagged with &#39;West&#39; (a union). The list of resulting VMs are protected by this Protection Group. | [optional] 

## Methods

### NewVmwareProtectionGroupExtraParams

`func NewVmwareProtectionGroupExtraParams() *VmwareProtectionGroupExtraParams`

NewVmwareProtectionGroupExtraParams instantiates a new VmwareProtectionGroupExtraParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmwareProtectionGroupExtraParamsWithDefaults

`func NewVmwareProtectionGroupExtraParamsWithDefaults() *VmwareProtectionGroupExtraParams`

NewVmwareProtectionGroupExtraParamsWithDefaults instantiates a new VmwareProtectionGroupExtraParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowParallelRuns

`func (o *VmwareProtectionGroupExtraParams) GetAllowParallelRuns() bool`

GetAllowParallelRuns returns the AllowParallelRuns field if non-nil, zero value otherwise.

### GetAllowParallelRunsOk

`func (o *VmwareProtectionGroupExtraParams) GetAllowParallelRunsOk() (*bool, bool)`

GetAllowParallelRunsOk returns a tuple with the AllowParallelRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowParallelRuns

`func (o *VmwareProtectionGroupExtraParams) SetAllowParallelRuns(v bool)`

SetAllowParallelRuns sets AllowParallelRuns field to given value.

### HasAllowParallelRuns

`func (o *VmwareProtectionGroupExtraParams) HasAllowParallelRuns() bool`

HasAllowParallelRuns returns a boolean if a field has been set.

### SetAllowParallelRunsNil

`func (o *VmwareProtectionGroupExtraParams) SetAllowParallelRunsNil(b bool)`

 SetAllowParallelRunsNil sets the value for AllowParallelRuns to be an explicit nil

### UnsetAllowParallelRuns
`func (o *VmwareProtectionGroupExtraParams) UnsetAllowParallelRuns()`

UnsetAllowParallelRuns ensures that no value is present for AllowParallelRuns, not even an explicit nil
### GetCloudMigration

`func (o *VmwareProtectionGroupExtraParams) GetCloudMigration() bool`

GetCloudMigration returns the CloudMigration field if non-nil, zero value otherwise.

### GetCloudMigrationOk

`func (o *VmwareProtectionGroupExtraParams) GetCloudMigrationOk() (*bool, bool)`

GetCloudMigrationOk returns a tuple with the CloudMigration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudMigration

`func (o *VmwareProtectionGroupExtraParams) SetCloudMigration(v bool)`

SetCloudMigration sets CloudMigration field to given value.

### HasCloudMigration

`func (o *VmwareProtectionGroupExtraParams) HasCloudMigration() bool`

HasCloudMigration returns a boolean if a field has been set.

### SetCloudMigrationNil

`func (o *VmwareProtectionGroupExtraParams) SetCloudMigrationNil(b bool)`

 SetCloudMigrationNil sets the value for CloudMigration to be an explicit nil

### UnsetCloudMigration
`func (o *VmwareProtectionGroupExtraParams) UnsetCloudMigration()`

UnsetCloudMigration ensures that no value is present for CloudMigration, not even an explicit nil
### GetExcludeFilters

`func (o *VmwareProtectionGroupExtraParams) GetExcludeFilters() []VMFilter`

GetExcludeFilters returns the ExcludeFilters field if non-nil, zero value otherwise.

### GetExcludeFiltersOk

`func (o *VmwareProtectionGroupExtraParams) GetExcludeFiltersOk() (*[]VMFilter, bool)`

GetExcludeFiltersOk returns a tuple with the ExcludeFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeFilters

`func (o *VmwareProtectionGroupExtraParams) SetExcludeFilters(v []VMFilter)`

SetExcludeFilters sets ExcludeFilters field to given value.

### HasExcludeFilters

`func (o *VmwareProtectionGroupExtraParams) HasExcludeFilters() bool`

HasExcludeFilters returns a boolean if a field has been set.

### SetExcludeFiltersNil

`func (o *VmwareProtectionGroupExtraParams) SetExcludeFiltersNil(b bool)`

 SetExcludeFiltersNil sets the value for ExcludeFilters to be an explicit nil

### UnsetExcludeFilters
`func (o *VmwareProtectionGroupExtraParams) UnsetExcludeFilters()`

UnsetExcludeFilters ensures that no value is present for ExcludeFilters, not even an explicit nil
### GetExcludeObjectIds

`func (o *VmwareProtectionGroupExtraParams) GetExcludeObjectIds() []*int64`

GetExcludeObjectIds returns the ExcludeObjectIds field if non-nil, zero value otherwise.

### GetExcludeObjectIdsOk

`func (o *VmwareProtectionGroupExtraParams) GetExcludeObjectIdsOk() (*[]*int64, bool)`

GetExcludeObjectIdsOk returns a tuple with the ExcludeObjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeObjectIds

`func (o *VmwareProtectionGroupExtraParams) SetExcludeObjectIds(v []*int64)`

SetExcludeObjectIds sets ExcludeObjectIds field to given value.

### HasExcludeObjectIds

`func (o *VmwareProtectionGroupExtraParams) HasExcludeObjectIds() bool`

HasExcludeObjectIds returns a boolean if a field has been set.

### GetExcludeVmTagIds

`func (o *VmwareProtectionGroupExtraParams) GetExcludeVmTagIds() [][]int64`

GetExcludeVmTagIds returns the ExcludeVmTagIds field if non-nil, zero value otherwise.

### GetExcludeVmTagIdsOk

`func (o *VmwareProtectionGroupExtraParams) GetExcludeVmTagIdsOk() (*[][]int64, bool)`

GetExcludeVmTagIdsOk returns a tuple with the ExcludeVmTagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeVmTagIds

`func (o *VmwareProtectionGroupExtraParams) SetExcludeVmTagIds(v [][]int64)`

SetExcludeVmTagIds sets ExcludeVmTagIds field to given value.

### HasExcludeVmTagIds

`func (o *VmwareProtectionGroupExtraParams) HasExcludeVmTagIds() bool`

HasExcludeVmTagIds returns a boolean if a field has been set.

### GetLeverageHyperflexSnapshots

`func (o *VmwareProtectionGroupExtraParams) GetLeverageHyperflexSnapshots() bool`

GetLeverageHyperflexSnapshots returns the LeverageHyperflexSnapshots field if non-nil, zero value otherwise.

### GetLeverageHyperflexSnapshotsOk

`func (o *VmwareProtectionGroupExtraParams) GetLeverageHyperflexSnapshotsOk() (*bool, bool)`

GetLeverageHyperflexSnapshotsOk returns a tuple with the LeverageHyperflexSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeverageHyperflexSnapshots

`func (o *VmwareProtectionGroupExtraParams) SetLeverageHyperflexSnapshots(v bool)`

SetLeverageHyperflexSnapshots sets LeverageHyperflexSnapshots field to given value.

### HasLeverageHyperflexSnapshots

`func (o *VmwareProtectionGroupExtraParams) HasLeverageHyperflexSnapshots() bool`

HasLeverageHyperflexSnapshots returns a boolean if a field has been set.

### SetLeverageHyperflexSnapshotsNil

`func (o *VmwareProtectionGroupExtraParams) SetLeverageHyperflexSnapshotsNil(b bool)`

 SetLeverageHyperflexSnapshotsNil sets the value for LeverageHyperflexSnapshots to be an explicit nil

### UnsetLeverageHyperflexSnapshots
`func (o *VmwareProtectionGroupExtraParams) UnsetLeverageHyperflexSnapshots()`

UnsetLeverageHyperflexSnapshots ensures that no value is present for LeverageHyperflexSnapshots, not even an explicit nil
### GetLeverageNutanixSnapshots

`func (o *VmwareProtectionGroupExtraParams) GetLeverageNutanixSnapshots() bool`

GetLeverageNutanixSnapshots returns the LeverageNutanixSnapshots field if non-nil, zero value otherwise.

### GetLeverageNutanixSnapshotsOk

`func (o *VmwareProtectionGroupExtraParams) GetLeverageNutanixSnapshotsOk() (*bool, bool)`

GetLeverageNutanixSnapshotsOk returns a tuple with the LeverageNutanixSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeverageNutanixSnapshots

`func (o *VmwareProtectionGroupExtraParams) SetLeverageNutanixSnapshots(v bool)`

SetLeverageNutanixSnapshots sets LeverageNutanixSnapshots field to given value.

### HasLeverageNutanixSnapshots

`func (o *VmwareProtectionGroupExtraParams) HasLeverageNutanixSnapshots() bool`

HasLeverageNutanixSnapshots returns a boolean if a field has been set.

### SetLeverageNutanixSnapshotsNil

`func (o *VmwareProtectionGroupExtraParams) SetLeverageNutanixSnapshotsNil(b bool)`

 SetLeverageNutanixSnapshotsNil sets the value for LeverageNutanixSnapshots to be an explicit nil

### UnsetLeverageNutanixSnapshots
`func (o *VmwareProtectionGroupExtraParams) UnsetLeverageNutanixSnapshots()`

UnsetLeverageNutanixSnapshots ensures that no value is present for LeverageNutanixSnapshots, not even an explicit nil
### GetLeverageStorageSnapshots

`func (o *VmwareProtectionGroupExtraParams) GetLeverageStorageSnapshots() bool`

GetLeverageStorageSnapshots returns the LeverageStorageSnapshots field if non-nil, zero value otherwise.

### GetLeverageStorageSnapshotsOk

`func (o *VmwareProtectionGroupExtraParams) GetLeverageStorageSnapshotsOk() (*bool, bool)`

GetLeverageStorageSnapshotsOk returns a tuple with the LeverageStorageSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeverageStorageSnapshots

`func (o *VmwareProtectionGroupExtraParams) SetLeverageStorageSnapshots(v bool)`

SetLeverageStorageSnapshots sets LeverageStorageSnapshots field to given value.

### HasLeverageStorageSnapshots

`func (o *VmwareProtectionGroupExtraParams) HasLeverageStorageSnapshots() bool`

HasLeverageStorageSnapshots returns a boolean if a field has been set.

### SetLeverageStorageSnapshotsNil

`func (o *VmwareProtectionGroupExtraParams) SetLeverageStorageSnapshotsNil(b bool)`

 SetLeverageStorageSnapshotsNil sets the value for LeverageStorageSnapshots to be an explicit nil

### UnsetLeverageStorageSnapshots
`func (o *VmwareProtectionGroupExtraParams) UnsetLeverageStorageSnapshots()`

UnsetLeverageStorageSnapshots ensures that no value is present for LeverageStorageSnapshots, not even an explicit nil
### GetSourceId

`func (o *VmwareProtectionGroupExtraParams) GetSourceId() int64`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *VmwareProtectionGroupExtraParams) GetSourceIdOk() (*int64, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *VmwareProtectionGroupExtraParams) SetSourceId(v int64)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *VmwareProtectionGroupExtraParams) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### SetSourceIdNil

`func (o *VmwareProtectionGroupExtraParams) SetSourceIdNil(b bool)`

 SetSourceIdNil sets the value for SourceId to be an explicit nil

### UnsetSourceId
`func (o *VmwareProtectionGroupExtraParams) UnsetSourceId()`

UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
### GetSourceName

`func (o *VmwareProtectionGroupExtraParams) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *VmwareProtectionGroupExtraParams) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *VmwareProtectionGroupExtraParams) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.

### HasSourceName

`func (o *VmwareProtectionGroupExtraParams) HasSourceName() bool`

HasSourceName returns a boolean if a field has been set.

### SetSourceNameNil

`func (o *VmwareProtectionGroupExtraParams) SetSourceNameNil(b bool)`

 SetSourceNameNil sets the value for SourceName to be an explicit nil

### UnsetSourceName
`func (o *VmwareProtectionGroupExtraParams) UnsetSourceName()`

UnsetSourceName ensures that no value is present for SourceName, not even an explicit nil
### GetVmTagIds

`func (o *VmwareProtectionGroupExtraParams) GetVmTagIds() [][]int64`

GetVmTagIds returns the VmTagIds field if non-nil, zero value otherwise.

### GetVmTagIdsOk

`func (o *VmwareProtectionGroupExtraParams) GetVmTagIdsOk() (*[][]int64, bool)`

GetVmTagIdsOk returns a tuple with the VmTagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmTagIds

`func (o *VmwareProtectionGroupExtraParams) SetVmTagIds(v [][]int64)`

SetVmTagIds sets VmTagIds field to given value.

### HasVmTagIds

`func (o *VmwareProtectionGroupExtraParams) HasVmTagIds() bool`

HasVmTagIds returns a boolean if a field has been set.

### SetVmTagIdsNil

`func (o *VmwareProtectionGroupExtraParams) SetVmTagIdsNil(b bool)`

 SetVmTagIdsNil sets the value for VmTagIds to be an explicit nil

### UnsetVmTagIds
`func (o *VmwareProtectionGroupExtraParams) UnsetVmTagIds()`

UnsetVmTagIds ensures that no value is present for VmTagIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


