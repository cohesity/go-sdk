# HyperVProtectionGroupParamsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProtectionType** | Pointer to **NullableString** | Specifies the Protection Group type. If not specified, then backup method is auto determined. Specifying RCT will forcibly use RCT backup for all VMs in this Protection Group. Available only for VMs with hardware version 8.0 and above, but is more efficient. Specifying VSS will forcibly use VSS backup for all VMs in this Protection Group. Available for VMs with hardware version 5.0 and above, but is slower than RCT backup. | [optional] 
**Objects** | Pointer to [**[]HyperVProtectionGroupObjectParams**](HyperVProtectionGroupObjectParams.md) | Specifies the objects to be included in the Protection Group. | [optional] 
**ExcludeObjectIds** | Pointer to **[]int64** | Specifies the list of IDs of the objects to not be protected by this Protection Group. This can be used to ignore specific objects under a parent object which has been included for protection. | [optional] 
**VmTagIds** | Pointer to **[][]int64** | Array of Array of VM Tag Ids that Specify VMs to Protect. Optionally specify a list of VMs to protect by listing Protection Source ids of VM Tags in this two dimensional array. Using this two dimensional array of Tag ids, the Cluster generates a list of VMs to protect which are derived from intersections of the inner arrays and union of the outer array, as shown by the following example. To protect only &#39;Eng&#39; VMs in the East and all the VMs in the West, specify the following tag id array: [ [1101, 2221], [3031] ], where 1101 is the &#39;Eng&#39; VM Tag id, 2221 is the &#39;East&#39; VM Tag id and 3031 is the &#39;West&#39; VM Tag id. The inner array [1101, 2221] produces a list of VMs that are both tagged with &#39;Eng&#39; and &#39;East&#39; (an intersection). The outer array combines the list from the inner array with list of VMs tagged with &#39;West&#39; (a union). The list of resulting VMs are protected by this Protection Group. | [optional] 
**ExcludeVmTagIds** | Pointer to **[][]int64** | Array of Arrays of VM Tag Ids that Specify VMs to Exclude. Optionally specify a list of VMs to exclude from protecting by listing Protection Source ids of VM Tags in this two dimensional array. Using this two dimensional array of Tag ids, the Cluster generates a list of VMs to exclude from protecting, which are derived from intersections of the inner arrays and union of the outer array, as shown by the following example. For example a Datacenter is selected to be protected but you want to exclude all the &#39;Former Employees&#39; VMs in the East and West but keep all the VMs for &#39;Former Employees&#39; in the South which are also stored in this Datacenter, by specifying the following tag id array: [ [1000, 2221], [1000, 3031] ], where 1000 is the &#39;Former Employee&#39; VM Tag id, 2221 is the &#39;East&#39; VM Tag id and 3031 is the &#39;West&#39; VM Tag id. The first inner array [1000, 2221] produces a list of VMs that are both tagged with &#39;Former Employees&#39; and &#39;East&#39; (an intersection). The second inner array [1000, 3031] produces a list of VMs that are both tagged with &#39;Former Employees&#39; and &#39;West&#39; (an intersection). The outer array combines the list of VMs from the two inner arrays. The list of resulting VMs are excluded from being protected this Job. | [optional] 
**CloudMigration** | Pointer to **NullableBool** | Specifies whether or not to move the workload to the cloud. | [optional] 
**SourceId** | Pointer to **NullableInt64** | Specifies the id of the parent of the objects. | [optional] [readonly] 
**SourceName** | Pointer to **NullableString** | Specifies the name of the parent of the objects. | [optional] [readonly] 

## Methods

### NewHyperVProtectionGroupParamsAllOf

`func NewHyperVProtectionGroupParamsAllOf() *HyperVProtectionGroupParamsAllOf`

NewHyperVProtectionGroupParamsAllOf instantiates a new HyperVProtectionGroupParamsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperVProtectionGroupParamsAllOfWithDefaults

`func NewHyperVProtectionGroupParamsAllOfWithDefaults() *HyperVProtectionGroupParamsAllOf`

NewHyperVProtectionGroupParamsAllOfWithDefaults instantiates a new HyperVProtectionGroupParamsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtectionType

`func (o *HyperVProtectionGroupParamsAllOf) GetProtectionType() string`

GetProtectionType returns the ProtectionType field if non-nil, zero value otherwise.

### GetProtectionTypeOk

`func (o *HyperVProtectionGroupParamsAllOf) GetProtectionTypeOk() (*string, bool)`

GetProtectionTypeOk returns a tuple with the ProtectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionType

`func (o *HyperVProtectionGroupParamsAllOf) SetProtectionType(v string)`

SetProtectionType sets ProtectionType field to given value.

### HasProtectionType

`func (o *HyperVProtectionGroupParamsAllOf) HasProtectionType() bool`

HasProtectionType returns a boolean if a field has been set.

### SetProtectionTypeNil

`func (o *HyperVProtectionGroupParamsAllOf) SetProtectionTypeNil(b bool)`

 SetProtectionTypeNil sets the value for ProtectionType to be an explicit nil

### UnsetProtectionType
`func (o *HyperVProtectionGroupParamsAllOf) UnsetProtectionType()`

UnsetProtectionType ensures that no value is present for ProtectionType, not even an explicit nil
### GetObjects

`func (o *HyperVProtectionGroupParamsAllOf) GetObjects() []HyperVProtectionGroupObjectParams`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *HyperVProtectionGroupParamsAllOf) GetObjectsOk() (*[]HyperVProtectionGroupObjectParams, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *HyperVProtectionGroupParamsAllOf) SetObjects(v []HyperVProtectionGroupObjectParams)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *HyperVProtectionGroupParamsAllOf) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### GetExcludeObjectIds

`func (o *HyperVProtectionGroupParamsAllOf) GetExcludeObjectIds() []int64`

GetExcludeObjectIds returns the ExcludeObjectIds field if non-nil, zero value otherwise.

### GetExcludeObjectIdsOk

`func (o *HyperVProtectionGroupParamsAllOf) GetExcludeObjectIdsOk() (*[]int64, bool)`

GetExcludeObjectIdsOk returns a tuple with the ExcludeObjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeObjectIds

`func (o *HyperVProtectionGroupParamsAllOf) SetExcludeObjectIds(v []int64)`

SetExcludeObjectIds sets ExcludeObjectIds field to given value.

### HasExcludeObjectIds

`func (o *HyperVProtectionGroupParamsAllOf) HasExcludeObjectIds() bool`

HasExcludeObjectIds returns a boolean if a field has been set.

### GetVmTagIds

`func (o *HyperVProtectionGroupParamsAllOf) GetVmTagIds() [][]int64`

GetVmTagIds returns the VmTagIds field if non-nil, zero value otherwise.

### GetVmTagIdsOk

`func (o *HyperVProtectionGroupParamsAllOf) GetVmTagIdsOk() (*[][]int64, bool)`

GetVmTagIdsOk returns a tuple with the VmTagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmTagIds

`func (o *HyperVProtectionGroupParamsAllOf) SetVmTagIds(v [][]int64)`

SetVmTagIds sets VmTagIds field to given value.

### HasVmTagIds

`func (o *HyperVProtectionGroupParamsAllOf) HasVmTagIds() bool`

HasVmTagIds returns a boolean if a field has been set.

### SetVmTagIdsNil

`func (o *HyperVProtectionGroupParamsAllOf) SetVmTagIdsNil(b bool)`

 SetVmTagIdsNil sets the value for VmTagIds to be an explicit nil

### UnsetVmTagIds
`func (o *HyperVProtectionGroupParamsAllOf) UnsetVmTagIds()`

UnsetVmTagIds ensures that no value is present for VmTagIds, not even an explicit nil
### GetExcludeVmTagIds

`func (o *HyperVProtectionGroupParamsAllOf) GetExcludeVmTagIds() [][]int64`

GetExcludeVmTagIds returns the ExcludeVmTagIds field if non-nil, zero value otherwise.

### GetExcludeVmTagIdsOk

`func (o *HyperVProtectionGroupParamsAllOf) GetExcludeVmTagIdsOk() (*[][]int64, bool)`

GetExcludeVmTagIdsOk returns a tuple with the ExcludeVmTagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeVmTagIds

`func (o *HyperVProtectionGroupParamsAllOf) SetExcludeVmTagIds(v [][]int64)`

SetExcludeVmTagIds sets ExcludeVmTagIds field to given value.

### HasExcludeVmTagIds

`func (o *HyperVProtectionGroupParamsAllOf) HasExcludeVmTagIds() bool`

HasExcludeVmTagIds returns a boolean if a field has been set.

### SetExcludeVmTagIdsNil

`func (o *HyperVProtectionGroupParamsAllOf) SetExcludeVmTagIdsNil(b bool)`

 SetExcludeVmTagIdsNil sets the value for ExcludeVmTagIds to be an explicit nil

### UnsetExcludeVmTagIds
`func (o *HyperVProtectionGroupParamsAllOf) UnsetExcludeVmTagIds()`

UnsetExcludeVmTagIds ensures that no value is present for ExcludeVmTagIds, not even an explicit nil
### GetCloudMigration

`func (o *HyperVProtectionGroupParamsAllOf) GetCloudMigration() bool`

GetCloudMigration returns the CloudMigration field if non-nil, zero value otherwise.

### GetCloudMigrationOk

`func (o *HyperVProtectionGroupParamsAllOf) GetCloudMigrationOk() (*bool, bool)`

GetCloudMigrationOk returns a tuple with the CloudMigration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudMigration

`func (o *HyperVProtectionGroupParamsAllOf) SetCloudMigration(v bool)`

SetCloudMigration sets CloudMigration field to given value.

### HasCloudMigration

`func (o *HyperVProtectionGroupParamsAllOf) HasCloudMigration() bool`

HasCloudMigration returns a boolean if a field has been set.

### SetCloudMigrationNil

`func (o *HyperVProtectionGroupParamsAllOf) SetCloudMigrationNil(b bool)`

 SetCloudMigrationNil sets the value for CloudMigration to be an explicit nil

### UnsetCloudMigration
`func (o *HyperVProtectionGroupParamsAllOf) UnsetCloudMigration()`

UnsetCloudMigration ensures that no value is present for CloudMigration, not even an explicit nil
### GetSourceId

`func (o *HyperVProtectionGroupParamsAllOf) GetSourceId() int64`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *HyperVProtectionGroupParamsAllOf) GetSourceIdOk() (*int64, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *HyperVProtectionGroupParamsAllOf) SetSourceId(v int64)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *HyperVProtectionGroupParamsAllOf) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### SetSourceIdNil

`func (o *HyperVProtectionGroupParamsAllOf) SetSourceIdNil(b bool)`

 SetSourceIdNil sets the value for SourceId to be an explicit nil

### UnsetSourceId
`func (o *HyperVProtectionGroupParamsAllOf) UnsetSourceId()`

UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
### GetSourceName

`func (o *HyperVProtectionGroupParamsAllOf) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *HyperVProtectionGroupParamsAllOf) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *HyperVProtectionGroupParamsAllOf) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.

### HasSourceName

`func (o *HyperVProtectionGroupParamsAllOf) HasSourceName() bool`

HasSourceName returns a boolean if a field has been set.

### SetSourceNameNil

`func (o *HyperVProtectionGroupParamsAllOf) SetSourceNameNil(b bool)`

 SetSourceNameNil sets the value for SourceName to be an explicit nil

### UnsetSourceName
`func (o *HyperVProtectionGroupParamsAllOf) UnsetSourceName()`

UnsetSourceName ensures that no value is present for SourceName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


