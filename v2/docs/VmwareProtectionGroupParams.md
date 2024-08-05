# VmwareProtectionGroupParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppConsistentSnapshot** | Pointer to **NullableBool** | Specifies whether or not to quiesce apps and the file system in order to take app consistent snapshots. | [optional] 
**EnableNBDSSLFallback** | Pointer to **NullableBool** | If this field is set to true and SAN transport backup fails, then backup will fallback to use NBDSSL transport. This field only applies if &#39;leverageSanTransport&#39; is set to true. | [optional] 
**FallbackToCrashConsistentSnapshot** | Pointer to **NullableBool** | Specifies whether or not to fallback to a crash consistent snapshot in the event that an app consistent snapshot fails. This parameter defaults to true and only changes the behavior of the operation if &#39;appConsistentSnapshot&#39; is set to &#39;true&#39;. | [optional] 
**IndexingPolicy** | Pointer to [**IndexingPolicy**](IndexingPolicy.md) |  | [optional] 
**LeverageSanTransport** | Pointer to **NullableBool** | If this field is set to true, then the backup for the objects will be performed using dedicated storage area network (SAN) instead of LAN or managment network. | [optional] 
**PrePostScript** | Pointer to [**PrePostScriptParams**](PrePostScriptParams.md) |  | [optional] 
**SkipPhysicalRDMDisks** | Pointer to **NullableBool** | Specifies whether or not to skip backing up physical RDM disks. Physical RDM disks cannot be backed up, so if you attempt to backup a VM with physical RDM disks and this value is set to &#39;false&#39;, then those VM backups will fail. | [optional] 
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
**EnableCdpSyncReplication** | Pointer to **NullableBool** | Specifies whether synchronous replication is enabled for CDP Protection Group when replication target is specified in attached policy. | [optional] 
**GlobalExcludeDisks** | Pointer to [**[]DiskInfo**](DiskInfo.md) | Specifies a list of disks to exclude from the backup. | [optional] 
**Objects** | Pointer to [**[]VmwareProtectionGroupObjectParams**](VmwareProtectionGroupObjectParams.md) | Specifies the objects to include in the backup. | [optional] 
**StandbyResourceObjects** | Pointer to [**[]VmwareProtectionGroupStandbyResourceParams**](VmwareProtectionGroupStandbyResourceParams.md) | Specifies the standby resource objects for this backup. | [optional] 

## Methods

### NewVmwareProtectionGroupParams

`func NewVmwareProtectionGroupParams() *VmwareProtectionGroupParams`

NewVmwareProtectionGroupParams instantiates a new VmwareProtectionGroupParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmwareProtectionGroupParamsWithDefaults

`func NewVmwareProtectionGroupParamsWithDefaults() *VmwareProtectionGroupParams`

NewVmwareProtectionGroupParamsWithDefaults instantiates a new VmwareProtectionGroupParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppConsistentSnapshot

`func (o *VmwareProtectionGroupParams) GetAppConsistentSnapshot() bool`

GetAppConsistentSnapshot returns the AppConsistentSnapshot field if non-nil, zero value otherwise.

### GetAppConsistentSnapshotOk

`func (o *VmwareProtectionGroupParams) GetAppConsistentSnapshotOk() (*bool, bool)`

GetAppConsistentSnapshotOk returns a tuple with the AppConsistentSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppConsistentSnapshot

`func (o *VmwareProtectionGroupParams) SetAppConsistentSnapshot(v bool)`

SetAppConsistentSnapshot sets AppConsistentSnapshot field to given value.

### HasAppConsistentSnapshot

`func (o *VmwareProtectionGroupParams) HasAppConsistentSnapshot() bool`

HasAppConsistentSnapshot returns a boolean if a field has been set.

### SetAppConsistentSnapshotNil

`func (o *VmwareProtectionGroupParams) SetAppConsistentSnapshotNil(b bool)`

 SetAppConsistentSnapshotNil sets the value for AppConsistentSnapshot to be an explicit nil

### UnsetAppConsistentSnapshot
`func (o *VmwareProtectionGroupParams) UnsetAppConsistentSnapshot()`

UnsetAppConsistentSnapshot ensures that no value is present for AppConsistentSnapshot, not even an explicit nil
### GetEnableNBDSSLFallback

`func (o *VmwareProtectionGroupParams) GetEnableNBDSSLFallback() bool`

GetEnableNBDSSLFallback returns the EnableNBDSSLFallback field if non-nil, zero value otherwise.

### GetEnableNBDSSLFallbackOk

`func (o *VmwareProtectionGroupParams) GetEnableNBDSSLFallbackOk() (*bool, bool)`

GetEnableNBDSSLFallbackOk returns a tuple with the EnableNBDSSLFallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNBDSSLFallback

`func (o *VmwareProtectionGroupParams) SetEnableNBDSSLFallback(v bool)`

SetEnableNBDSSLFallback sets EnableNBDSSLFallback field to given value.

### HasEnableNBDSSLFallback

`func (o *VmwareProtectionGroupParams) HasEnableNBDSSLFallback() bool`

HasEnableNBDSSLFallback returns a boolean if a field has been set.

### SetEnableNBDSSLFallbackNil

`func (o *VmwareProtectionGroupParams) SetEnableNBDSSLFallbackNil(b bool)`

 SetEnableNBDSSLFallbackNil sets the value for EnableNBDSSLFallback to be an explicit nil

### UnsetEnableNBDSSLFallback
`func (o *VmwareProtectionGroupParams) UnsetEnableNBDSSLFallback()`

UnsetEnableNBDSSLFallback ensures that no value is present for EnableNBDSSLFallback, not even an explicit nil
### GetFallbackToCrashConsistentSnapshot

`func (o *VmwareProtectionGroupParams) GetFallbackToCrashConsistentSnapshot() bool`

GetFallbackToCrashConsistentSnapshot returns the FallbackToCrashConsistentSnapshot field if non-nil, zero value otherwise.

### GetFallbackToCrashConsistentSnapshotOk

`func (o *VmwareProtectionGroupParams) GetFallbackToCrashConsistentSnapshotOk() (*bool, bool)`

GetFallbackToCrashConsistentSnapshotOk returns a tuple with the FallbackToCrashConsistentSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackToCrashConsistentSnapshot

`func (o *VmwareProtectionGroupParams) SetFallbackToCrashConsistentSnapshot(v bool)`

SetFallbackToCrashConsistentSnapshot sets FallbackToCrashConsistentSnapshot field to given value.

### HasFallbackToCrashConsistentSnapshot

`func (o *VmwareProtectionGroupParams) HasFallbackToCrashConsistentSnapshot() bool`

HasFallbackToCrashConsistentSnapshot returns a boolean if a field has been set.

### SetFallbackToCrashConsistentSnapshotNil

`func (o *VmwareProtectionGroupParams) SetFallbackToCrashConsistentSnapshotNil(b bool)`

 SetFallbackToCrashConsistentSnapshotNil sets the value for FallbackToCrashConsistentSnapshot to be an explicit nil

### UnsetFallbackToCrashConsistentSnapshot
`func (o *VmwareProtectionGroupParams) UnsetFallbackToCrashConsistentSnapshot()`

UnsetFallbackToCrashConsistentSnapshot ensures that no value is present for FallbackToCrashConsistentSnapshot, not even an explicit nil
### GetIndexingPolicy

`func (o *VmwareProtectionGroupParams) GetIndexingPolicy() IndexingPolicy`

GetIndexingPolicy returns the IndexingPolicy field if non-nil, zero value otherwise.

### GetIndexingPolicyOk

`func (o *VmwareProtectionGroupParams) GetIndexingPolicyOk() (*IndexingPolicy, bool)`

GetIndexingPolicyOk returns a tuple with the IndexingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexingPolicy

`func (o *VmwareProtectionGroupParams) SetIndexingPolicy(v IndexingPolicy)`

SetIndexingPolicy sets IndexingPolicy field to given value.

### HasIndexingPolicy

`func (o *VmwareProtectionGroupParams) HasIndexingPolicy() bool`

HasIndexingPolicy returns a boolean if a field has been set.

### GetLeverageSanTransport

`func (o *VmwareProtectionGroupParams) GetLeverageSanTransport() bool`

GetLeverageSanTransport returns the LeverageSanTransport field if non-nil, zero value otherwise.

### GetLeverageSanTransportOk

`func (o *VmwareProtectionGroupParams) GetLeverageSanTransportOk() (*bool, bool)`

GetLeverageSanTransportOk returns a tuple with the LeverageSanTransport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeverageSanTransport

`func (o *VmwareProtectionGroupParams) SetLeverageSanTransport(v bool)`

SetLeverageSanTransport sets LeverageSanTransport field to given value.

### HasLeverageSanTransport

`func (o *VmwareProtectionGroupParams) HasLeverageSanTransport() bool`

HasLeverageSanTransport returns a boolean if a field has been set.

### SetLeverageSanTransportNil

`func (o *VmwareProtectionGroupParams) SetLeverageSanTransportNil(b bool)`

 SetLeverageSanTransportNil sets the value for LeverageSanTransport to be an explicit nil

### UnsetLeverageSanTransport
`func (o *VmwareProtectionGroupParams) UnsetLeverageSanTransport()`

UnsetLeverageSanTransport ensures that no value is present for LeverageSanTransport, not even an explicit nil
### GetPrePostScript

`func (o *VmwareProtectionGroupParams) GetPrePostScript() PrePostScriptParams`

GetPrePostScript returns the PrePostScript field if non-nil, zero value otherwise.

### GetPrePostScriptOk

`func (o *VmwareProtectionGroupParams) GetPrePostScriptOk() (*PrePostScriptParams, bool)`

GetPrePostScriptOk returns a tuple with the PrePostScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrePostScript

`func (o *VmwareProtectionGroupParams) SetPrePostScript(v PrePostScriptParams)`

SetPrePostScript sets PrePostScript field to given value.

### HasPrePostScript

`func (o *VmwareProtectionGroupParams) HasPrePostScript() bool`

HasPrePostScript returns a boolean if a field has been set.

### GetSkipPhysicalRDMDisks

`func (o *VmwareProtectionGroupParams) GetSkipPhysicalRDMDisks() bool`

GetSkipPhysicalRDMDisks returns the SkipPhysicalRDMDisks field if non-nil, zero value otherwise.

### GetSkipPhysicalRDMDisksOk

`func (o *VmwareProtectionGroupParams) GetSkipPhysicalRDMDisksOk() (*bool, bool)`

GetSkipPhysicalRDMDisksOk returns a tuple with the SkipPhysicalRDMDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipPhysicalRDMDisks

`func (o *VmwareProtectionGroupParams) SetSkipPhysicalRDMDisks(v bool)`

SetSkipPhysicalRDMDisks sets SkipPhysicalRDMDisks field to given value.

### HasSkipPhysicalRDMDisks

`func (o *VmwareProtectionGroupParams) HasSkipPhysicalRDMDisks() bool`

HasSkipPhysicalRDMDisks returns a boolean if a field has been set.

### SetSkipPhysicalRDMDisksNil

`func (o *VmwareProtectionGroupParams) SetSkipPhysicalRDMDisksNil(b bool)`

 SetSkipPhysicalRDMDisksNil sets the value for SkipPhysicalRDMDisks to be an explicit nil

### UnsetSkipPhysicalRDMDisks
`func (o *VmwareProtectionGroupParams) UnsetSkipPhysicalRDMDisks()`

UnsetSkipPhysicalRDMDisks ensures that no value is present for SkipPhysicalRDMDisks, not even an explicit nil
### GetAllowParallelRuns

`func (o *VmwareProtectionGroupParams) GetAllowParallelRuns() bool`

GetAllowParallelRuns returns the AllowParallelRuns field if non-nil, zero value otherwise.

### GetAllowParallelRunsOk

`func (o *VmwareProtectionGroupParams) GetAllowParallelRunsOk() (*bool, bool)`

GetAllowParallelRunsOk returns a tuple with the AllowParallelRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowParallelRuns

`func (o *VmwareProtectionGroupParams) SetAllowParallelRuns(v bool)`

SetAllowParallelRuns sets AllowParallelRuns field to given value.

### HasAllowParallelRuns

`func (o *VmwareProtectionGroupParams) HasAllowParallelRuns() bool`

HasAllowParallelRuns returns a boolean if a field has been set.

### SetAllowParallelRunsNil

`func (o *VmwareProtectionGroupParams) SetAllowParallelRunsNil(b bool)`

 SetAllowParallelRunsNil sets the value for AllowParallelRuns to be an explicit nil

### UnsetAllowParallelRuns
`func (o *VmwareProtectionGroupParams) UnsetAllowParallelRuns()`

UnsetAllowParallelRuns ensures that no value is present for AllowParallelRuns, not even an explicit nil
### GetCloudMigration

`func (o *VmwareProtectionGroupParams) GetCloudMigration() bool`

GetCloudMigration returns the CloudMigration field if non-nil, zero value otherwise.

### GetCloudMigrationOk

`func (o *VmwareProtectionGroupParams) GetCloudMigrationOk() (*bool, bool)`

GetCloudMigrationOk returns a tuple with the CloudMigration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudMigration

`func (o *VmwareProtectionGroupParams) SetCloudMigration(v bool)`

SetCloudMigration sets CloudMigration field to given value.

### HasCloudMigration

`func (o *VmwareProtectionGroupParams) HasCloudMigration() bool`

HasCloudMigration returns a boolean if a field has been set.

### SetCloudMigrationNil

`func (o *VmwareProtectionGroupParams) SetCloudMigrationNil(b bool)`

 SetCloudMigrationNil sets the value for CloudMigration to be an explicit nil

### UnsetCloudMigration
`func (o *VmwareProtectionGroupParams) UnsetCloudMigration()`

UnsetCloudMigration ensures that no value is present for CloudMigration, not even an explicit nil
### GetExcludeFilters

`func (o *VmwareProtectionGroupParams) GetExcludeFilters() []VMFilter`

GetExcludeFilters returns the ExcludeFilters field if non-nil, zero value otherwise.

### GetExcludeFiltersOk

`func (o *VmwareProtectionGroupParams) GetExcludeFiltersOk() (*[]VMFilter, bool)`

GetExcludeFiltersOk returns a tuple with the ExcludeFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeFilters

`func (o *VmwareProtectionGroupParams) SetExcludeFilters(v []VMFilter)`

SetExcludeFilters sets ExcludeFilters field to given value.

### HasExcludeFilters

`func (o *VmwareProtectionGroupParams) HasExcludeFilters() bool`

HasExcludeFilters returns a boolean if a field has been set.

### SetExcludeFiltersNil

`func (o *VmwareProtectionGroupParams) SetExcludeFiltersNil(b bool)`

 SetExcludeFiltersNil sets the value for ExcludeFilters to be an explicit nil

### UnsetExcludeFilters
`func (o *VmwareProtectionGroupParams) UnsetExcludeFilters()`

UnsetExcludeFilters ensures that no value is present for ExcludeFilters, not even an explicit nil
### GetExcludeObjectIds

`func (o *VmwareProtectionGroupParams) GetExcludeObjectIds() []*int64`

GetExcludeObjectIds returns the ExcludeObjectIds field if non-nil, zero value otherwise.

### GetExcludeObjectIdsOk

`func (o *VmwareProtectionGroupParams) GetExcludeObjectIdsOk() (*[]*int64, bool)`

GetExcludeObjectIdsOk returns a tuple with the ExcludeObjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeObjectIds

`func (o *VmwareProtectionGroupParams) SetExcludeObjectIds(v []*int64)`

SetExcludeObjectIds sets ExcludeObjectIds field to given value.

### HasExcludeObjectIds

`func (o *VmwareProtectionGroupParams) HasExcludeObjectIds() bool`

HasExcludeObjectIds returns a boolean if a field has been set.

### GetExcludeVmTagIds

`func (o *VmwareProtectionGroupParams) GetExcludeVmTagIds() [][]int64`

GetExcludeVmTagIds returns the ExcludeVmTagIds field if non-nil, zero value otherwise.

### GetExcludeVmTagIdsOk

`func (o *VmwareProtectionGroupParams) GetExcludeVmTagIdsOk() (*[][]int64, bool)`

GetExcludeVmTagIdsOk returns a tuple with the ExcludeVmTagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeVmTagIds

`func (o *VmwareProtectionGroupParams) SetExcludeVmTagIds(v [][]int64)`

SetExcludeVmTagIds sets ExcludeVmTagIds field to given value.

### HasExcludeVmTagIds

`func (o *VmwareProtectionGroupParams) HasExcludeVmTagIds() bool`

HasExcludeVmTagIds returns a boolean if a field has been set.

### GetLeverageHyperflexSnapshots

`func (o *VmwareProtectionGroupParams) GetLeverageHyperflexSnapshots() bool`

GetLeverageHyperflexSnapshots returns the LeverageHyperflexSnapshots field if non-nil, zero value otherwise.

### GetLeverageHyperflexSnapshotsOk

`func (o *VmwareProtectionGroupParams) GetLeverageHyperflexSnapshotsOk() (*bool, bool)`

GetLeverageHyperflexSnapshotsOk returns a tuple with the LeverageHyperflexSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeverageHyperflexSnapshots

`func (o *VmwareProtectionGroupParams) SetLeverageHyperflexSnapshots(v bool)`

SetLeverageHyperflexSnapshots sets LeverageHyperflexSnapshots field to given value.

### HasLeverageHyperflexSnapshots

`func (o *VmwareProtectionGroupParams) HasLeverageHyperflexSnapshots() bool`

HasLeverageHyperflexSnapshots returns a boolean if a field has been set.

### SetLeverageHyperflexSnapshotsNil

`func (o *VmwareProtectionGroupParams) SetLeverageHyperflexSnapshotsNil(b bool)`

 SetLeverageHyperflexSnapshotsNil sets the value for LeverageHyperflexSnapshots to be an explicit nil

### UnsetLeverageHyperflexSnapshots
`func (o *VmwareProtectionGroupParams) UnsetLeverageHyperflexSnapshots()`

UnsetLeverageHyperflexSnapshots ensures that no value is present for LeverageHyperflexSnapshots, not even an explicit nil
### GetLeverageNutanixSnapshots

`func (o *VmwareProtectionGroupParams) GetLeverageNutanixSnapshots() bool`

GetLeverageNutanixSnapshots returns the LeverageNutanixSnapshots field if non-nil, zero value otherwise.

### GetLeverageNutanixSnapshotsOk

`func (o *VmwareProtectionGroupParams) GetLeverageNutanixSnapshotsOk() (*bool, bool)`

GetLeverageNutanixSnapshotsOk returns a tuple with the LeverageNutanixSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeverageNutanixSnapshots

`func (o *VmwareProtectionGroupParams) SetLeverageNutanixSnapshots(v bool)`

SetLeverageNutanixSnapshots sets LeverageNutanixSnapshots field to given value.

### HasLeverageNutanixSnapshots

`func (o *VmwareProtectionGroupParams) HasLeverageNutanixSnapshots() bool`

HasLeverageNutanixSnapshots returns a boolean if a field has been set.

### SetLeverageNutanixSnapshotsNil

`func (o *VmwareProtectionGroupParams) SetLeverageNutanixSnapshotsNil(b bool)`

 SetLeverageNutanixSnapshotsNil sets the value for LeverageNutanixSnapshots to be an explicit nil

### UnsetLeverageNutanixSnapshots
`func (o *VmwareProtectionGroupParams) UnsetLeverageNutanixSnapshots()`

UnsetLeverageNutanixSnapshots ensures that no value is present for LeverageNutanixSnapshots, not even an explicit nil
### GetLeverageStorageSnapshots

`func (o *VmwareProtectionGroupParams) GetLeverageStorageSnapshots() bool`

GetLeverageStorageSnapshots returns the LeverageStorageSnapshots field if non-nil, zero value otherwise.

### GetLeverageStorageSnapshotsOk

`func (o *VmwareProtectionGroupParams) GetLeverageStorageSnapshotsOk() (*bool, bool)`

GetLeverageStorageSnapshotsOk returns a tuple with the LeverageStorageSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeverageStorageSnapshots

`func (o *VmwareProtectionGroupParams) SetLeverageStorageSnapshots(v bool)`

SetLeverageStorageSnapshots sets LeverageStorageSnapshots field to given value.

### HasLeverageStorageSnapshots

`func (o *VmwareProtectionGroupParams) HasLeverageStorageSnapshots() bool`

HasLeverageStorageSnapshots returns a boolean if a field has been set.

### SetLeverageStorageSnapshotsNil

`func (o *VmwareProtectionGroupParams) SetLeverageStorageSnapshotsNil(b bool)`

 SetLeverageStorageSnapshotsNil sets the value for LeverageStorageSnapshots to be an explicit nil

### UnsetLeverageStorageSnapshots
`func (o *VmwareProtectionGroupParams) UnsetLeverageStorageSnapshots()`

UnsetLeverageStorageSnapshots ensures that no value is present for LeverageStorageSnapshots, not even an explicit nil
### GetSourceId

`func (o *VmwareProtectionGroupParams) GetSourceId() int64`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *VmwareProtectionGroupParams) GetSourceIdOk() (*int64, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *VmwareProtectionGroupParams) SetSourceId(v int64)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *VmwareProtectionGroupParams) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### SetSourceIdNil

`func (o *VmwareProtectionGroupParams) SetSourceIdNil(b bool)`

 SetSourceIdNil sets the value for SourceId to be an explicit nil

### UnsetSourceId
`func (o *VmwareProtectionGroupParams) UnsetSourceId()`

UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
### GetSourceName

`func (o *VmwareProtectionGroupParams) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *VmwareProtectionGroupParams) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *VmwareProtectionGroupParams) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.

### HasSourceName

`func (o *VmwareProtectionGroupParams) HasSourceName() bool`

HasSourceName returns a boolean if a field has been set.

### SetSourceNameNil

`func (o *VmwareProtectionGroupParams) SetSourceNameNil(b bool)`

 SetSourceNameNil sets the value for SourceName to be an explicit nil

### UnsetSourceName
`func (o *VmwareProtectionGroupParams) UnsetSourceName()`

UnsetSourceName ensures that no value is present for SourceName, not even an explicit nil
### GetVmTagIds

`func (o *VmwareProtectionGroupParams) GetVmTagIds() [][]int64`

GetVmTagIds returns the VmTagIds field if non-nil, zero value otherwise.

### GetVmTagIdsOk

`func (o *VmwareProtectionGroupParams) GetVmTagIdsOk() (*[][]int64, bool)`

GetVmTagIdsOk returns a tuple with the VmTagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmTagIds

`func (o *VmwareProtectionGroupParams) SetVmTagIds(v [][]int64)`

SetVmTagIds sets VmTagIds field to given value.

### HasVmTagIds

`func (o *VmwareProtectionGroupParams) HasVmTagIds() bool`

HasVmTagIds returns a boolean if a field has been set.

### SetVmTagIdsNil

`func (o *VmwareProtectionGroupParams) SetVmTagIdsNil(b bool)`

 SetVmTagIdsNil sets the value for VmTagIds to be an explicit nil

### UnsetVmTagIds
`func (o *VmwareProtectionGroupParams) UnsetVmTagIds()`

UnsetVmTagIds ensures that no value is present for VmTagIds, not even an explicit nil
### GetEnableCdpSyncReplication

`func (o *VmwareProtectionGroupParams) GetEnableCdpSyncReplication() bool`

GetEnableCdpSyncReplication returns the EnableCdpSyncReplication field if non-nil, zero value otherwise.

### GetEnableCdpSyncReplicationOk

`func (o *VmwareProtectionGroupParams) GetEnableCdpSyncReplicationOk() (*bool, bool)`

GetEnableCdpSyncReplicationOk returns a tuple with the EnableCdpSyncReplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCdpSyncReplication

`func (o *VmwareProtectionGroupParams) SetEnableCdpSyncReplication(v bool)`

SetEnableCdpSyncReplication sets EnableCdpSyncReplication field to given value.

### HasEnableCdpSyncReplication

`func (o *VmwareProtectionGroupParams) HasEnableCdpSyncReplication() bool`

HasEnableCdpSyncReplication returns a boolean if a field has been set.

### SetEnableCdpSyncReplicationNil

`func (o *VmwareProtectionGroupParams) SetEnableCdpSyncReplicationNil(b bool)`

 SetEnableCdpSyncReplicationNil sets the value for EnableCdpSyncReplication to be an explicit nil

### UnsetEnableCdpSyncReplication
`func (o *VmwareProtectionGroupParams) UnsetEnableCdpSyncReplication()`

UnsetEnableCdpSyncReplication ensures that no value is present for EnableCdpSyncReplication, not even an explicit nil
### GetGlobalExcludeDisks

`func (o *VmwareProtectionGroupParams) GetGlobalExcludeDisks() []DiskInfo`

GetGlobalExcludeDisks returns the GlobalExcludeDisks field if non-nil, zero value otherwise.

### GetGlobalExcludeDisksOk

`func (o *VmwareProtectionGroupParams) GetGlobalExcludeDisksOk() (*[]DiskInfo, bool)`

GetGlobalExcludeDisksOk returns a tuple with the GlobalExcludeDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalExcludeDisks

`func (o *VmwareProtectionGroupParams) SetGlobalExcludeDisks(v []DiskInfo)`

SetGlobalExcludeDisks sets GlobalExcludeDisks field to given value.

### HasGlobalExcludeDisks

`func (o *VmwareProtectionGroupParams) HasGlobalExcludeDisks() bool`

HasGlobalExcludeDisks returns a boolean if a field has been set.

### SetGlobalExcludeDisksNil

`func (o *VmwareProtectionGroupParams) SetGlobalExcludeDisksNil(b bool)`

 SetGlobalExcludeDisksNil sets the value for GlobalExcludeDisks to be an explicit nil

### UnsetGlobalExcludeDisks
`func (o *VmwareProtectionGroupParams) UnsetGlobalExcludeDisks()`

UnsetGlobalExcludeDisks ensures that no value is present for GlobalExcludeDisks, not even an explicit nil
### GetObjects

`func (o *VmwareProtectionGroupParams) GetObjects() []VmwareProtectionGroupObjectParams`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *VmwareProtectionGroupParams) GetObjectsOk() (*[]VmwareProtectionGroupObjectParams, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *VmwareProtectionGroupParams) SetObjects(v []VmwareProtectionGroupObjectParams)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *VmwareProtectionGroupParams) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### GetStandbyResourceObjects

`func (o *VmwareProtectionGroupParams) GetStandbyResourceObjects() []VmwareProtectionGroupStandbyResourceParams`

GetStandbyResourceObjects returns the StandbyResourceObjects field if non-nil, zero value otherwise.

### GetStandbyResourceObjectsOk

`func (o *VmwareProtectionGroupParams) GetStandbyResourceObjectsOk() (*[]VmwareProtectionGroupStandbyResourceParams, bool)`

GetStandbyResourceObjectsOk returns a tuple with the StandbyResourceObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandbyResourceObjects

`func (o *VmwareProtectionGroupParams) SetStandbyResourceObjects(v []VmwareProtectionGroupStandbyResourceParams)`

SetStandbyResourceObjects sets StandbyResourceObjects field to given value.

### HasStandbyResourceObjects

`func (o *VmwareProtectionGroupParams) HasStandbyResourceObjects() bool`

HasStandbyResourceObjects returns a boolean if a field has been set.

### SetStandbyResourceObjectsNil

`func (o *VmwareProtectionGroupParams) SetStandbyResourceObjectsNil(b bool)`

 SetStandbyResourceObjectsNil sets the value for StandbyResourceObjects to be an explicit nil

### UnsetStandbyResourceObjects
`func (o *VmwareProtectionGroupParams) UnsetStandbyResourceObjects()`

UnsetStandbyResourceObjects ensures that no value is present for StandbyResourceObjects, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


