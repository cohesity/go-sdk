# GcpNativeProtectionGroupParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExcludeObjectIds** | Pointer to **[]int64** | Specifies the objects to be excluded in the Protection Group. | [optional] 
**ExcludeVmTagIds** | Pointer to **[][]int64** | Array of Arrays of VM Tag Ids that Specify VMs to Exclude. Optionally specify a list of VMs to exclude from protecting by listing Protection Source ids of VM Tags in this two dimensional array. Using this two dimensional array of Tag ids, the Cluster generates a list of VMs to exclude from protecting, which are derived from intersections of the inner arrays and union of the outer array, as shown by the following example. For example a Datacenter is selected to be protected but you want to exclude all the &#39;Former Employees&#39; VMs in the East and West but keep all the VMs for &#39;Former Employees&#39; in the South which are also stored in this Datacenter, by specifying the following tag id array: [ [1000, 2221], [1000, 3031] ], where 1000 is the &#39;Former Employee&#39; VM Tag id, 2221 is the &#39;East&#39; VM Tag id and 3031 is the &#39;West&#39; VM Tag id. The first inner array [1000, 2221] produces a list of VMs that are both tagged with &#39;Former Employees&#39; and &#39;East&#39; (an intersection). The second inner array [1000, 3031] produces a list of VMs that are both tagged with &#39;Former Employees&#39; and &#39;West&#39; (an intersection). The outer array combines the list of VMs from the two inner arrays. The list of resulting VMs are excluded from being protected this Job. | [optional] 
**GcpDiskExclusionParams** | Pointer to [**GcpDiskExclusionParams**](GcpDiskExclusionParams.md) |  | [optional] 
**IndexingPolicy** | Pointer to [**IndexingPolicy**](IndexingPolicy.md) |  | [optional] 
**Objects** | Pointer to [**[]GcpNativeProtectionGroupObjectParams**](GcpNativeProtectionGroupObjectParams.md) | Specifies the objects to be included in the Protection Group. | [optional] 
**SourceId** | Pointer to **NullableInt64** | Specifies the id of the parent of the objects. | [optional] [readonly] 
**SourceName** | Pointer to **NullableString** | Specifies the name of the parent of the objects. | [optional] [readonly] 
**VmTagIds** | Pointer to **[][]int64** | Array of Array of VM Tag Ids that Specify VMs to Protect. Optionally specify a list of VMs to protect by listing Protection Source ids of VM Tags in this two dimensional array. Using this two dimensional array of Tag ids, the Cluster generates a list of VMs to protect which are derived from intersections of the inner arrays and union of the outer array, as shown by the following example. To protect only &#39;Eng&#39; VMs in the East and all the VMs in the West, specify the following tag id array: [ [1101, 2221], [3031] ], where 1101 is the &#39;Eng&#39; VM Tag id, 2221 is the &#39;East&#39; VM Tag id and 3031 is the &#39;West&#39; VM Tag id. The inner array [1101, 2221] produces a list of VMs that are both tagged with &#39;Eng&#39; and &#39;East&#39; (an intersection). The outer array combines the list from the inner array with list of VMs tagged with &#39;West&#39; (a union). The list of resulting VMs are protected by this Protection Group. | [optional] 

## Methods

### NewGcpNativeProtectionGroupParams

`func NewGcpNativeProtectionGroupParams() *GcpNativeProtectionGroupParams`

NewGcpNativeProtectionGroupParams instantiates a new GcpNativeProtectionGroupParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpNativeProtectionGroupParamsWithDefaults

`func NewGcpNativeProtectionGroupParamsWithDefaults() *GcpNativeProtectionGroupParams`

NewGcpNativeProtectionGroupParamsWithDefaults instantiates a new GcpNativeProtectionGroupParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExcludeObjectIds

`func (o *GcpNativeProtectionGroupParams) GetExcludeObjectIds() []int64`

GetExcludeObjectIds returns the ExcludeObjectIds field if non-nil, zero value otherwise.

### GetExcludeObjectIdsOk

`func (o *GcpNativeProtectionGroupParams) GetExcludeObjectIdsOk() (*[]int64, bool)`

GetExcludeObjectIdsOk returns a tuple with the ExcludeObjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeObjectIds

`func (o *GcpNativeProtectionGroupParams) SetExcludeObjectIds(v []int64)`

SetExcludeObjectIds sets ExcludeObjectIds field to given value.

### HasExcludeObjectIds

`func (o *GcpNativeProtectionGroupParams) HasExcludeObjectIds() bool`

HasExcludeObjectIds returns a boolean if a field has been set.

### SetExcludeObjectIdsNil

`func (o *GcpNativeProtectionGroupParams) SetExcludeObjectIdsNil(b bool)`

 SetExcludeObjectIdsNil sets the value for ExcludeObjectIds to be an explicit nil

### UnsetExcludeObjectIds
`func (o *GcpNativeProtectionGroupParams) UnsetExcludeObjectIds()`

UnsetExcludeObjectIds ensures that no value is present for ExcludeObjectIds, not even an explicit nil
### GetExcludeVmTagIds

`func (o *GcpNativeProtectionGroupParams) GetExcludeVmTagIds() [][]int64`

GetExcludeVmTagIds returns the ExcludeVmTagIds field if non-nil, zero value otherwise.

### GetExcludeVmTagIdsOk

`func (o *GcpNativeProtectionGroupParams) GetExcludeVmTagIdsOk() (*[][]int64, bool)`

GetExcludeVmTagIdsOk returns a tuple with the ExcludeVmTagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeVmTagIds

`func (o *GcpNativeProtectionGroupParams) SetExcludeVmTagIds(v [][]int64)`

SetExcludeVmTagIds sets ExcludeVmTagIds field to given value.

### HasExcludeVmTagIds

`func (o *GcpNativeProtectionGroupParams) HasExcludeVmTagIds() bool`

HasExcludeVmTagIds returns a boolean if a field has been set.

### GetGcpDiskExclusionParams

`func (o *GcpNativeProtectionGroupParams) GetGcpDiskExclusionParams() GcpDiskExclusionParams`

GetGcpDiskExclusionParams returns the GcpDiskExclusionParams field if non-nil, zero value otherwise.

### GetGcpDiskExclusionParamsOk

`func (o *GcpNativeProtectionGroupParams) GetGcpDiskExclusionParamsOk() (*GcpDiskExclusionParams, bool)`

GetGcpDiskExclusionParamsOk returns a tuple with the GcpDiskExclusionParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpDiskExclusionParams

`func (o *GcpNativeProtectionGroupParams) SetGcpDiskExclusionParams(v GcpDiskExclusionParams)`

SetGcpDiskExclusionParams sets GcpDiskExclusionParams field to given value.

### HasGcpDiskExclusionParams

`func (o *GcpNativeProtectionGroupParams) HasGcpDiskExclusionParams() bool`

HasGcpDiskExclusionParams returns a boolean if a field has been set.

### GetIndexingPolicy

`func (o *GcpNativeProtectionGroupParams) GetIndexingPolicy() IndexingPolicy`

GetIndexingPolicy returns the IndexingPolicy field if non-nil, zero value otherwise.

### GetIndexingPolicyOk

`func (o *GcpNativeProtectionGroupParams) GetIndexingPolicyOk() (*IndexingPolicy, bool)`

GetIndexingPolicyOk returns a tuple with the IndexingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexingPolicy

`func (o *GcpNativeProtectionGroupParams) SetIndexingPolicy(v IndexingPolicy)`

SetIndexingPolicy sets IndexingPolicy field to given value.

### HasIndexingPolicy

`func (o *GcpNativeProtectionGroupParams) HasIndexingPolicy() bool`

HasIndexingPolicy returns a boolean if a field has been set.

### GetObjects

`func (o *GcpNativeProtectionGroupParams) GetObjects() []GcpNativeProtectionGroupObjectParams`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *GcpNativeProtectionGroupParams) GetObjectsOk() (*[]GcpNativeProtectionGroupObjectParams, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *GcpNativeProtectionGroupParams) SetObjects(v []GcpNativeProtectionGroupObjectParams)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *GcpNativeProtectionGroupParams) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### GetSourceId

`func (o *GcpNativeProtectionGroupParams) GetSourceId() int64`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *GcpNativeProtectionGroupParams) GetSourceIdOk() (*int64, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *GcpNativeProtectionGroupParams) SetSourceId(v int64)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *GcpNativeProtectionGroupParams) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### SetSourceIdNil

`func (o *GcpNativeProtectionGroupParams) SetSourceIdNil(b bool)`

 SetSourceIdNil sets the value for SourceId to be an explicit nil

### UnsetSourceId
`func (o *GcpNativeProtectionGroupParams) UnsetSourceId()`

UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
### GetSourceName

`func (o *GcpNativeProtectionGroupParams) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *GcpNativeProtectionGroupParams) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *GcpNativeProtectionGroupParams) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.

### HasSourceName

`func (o *GcpNativeProtectionGroupParams) HasSourceName() bool`

HasSourceName returns a boolean if a field has been set.

### SetSourceNameNil

`func (o *GcpNativeProtectionGroupParams) SetSourceNameNil(b bool)`

 SetSourceNameNil sets the value for SourceName to be an explicit nil

### UnsetSourceName
`func (o *GcpNativeProtectionGroupParams) UnsetSourceName()`

UnsetSourceName ensures that no value is present for SourceName, not even an explicit nil
### GetVmTagIds

`func (o *GcpNativeProtectionGroupParams) GetVmTagIds() [][]int64`

GetVmTagIds returns the VmTagIds field if non-nil, zero value otherwise.

### GetVmTagIdsOk

`func (o *GcpNativeProtectionGroupParams) GetVmTagIdsOk() (*[][]int64, bool)`

GetVmTagIdsOk returns a tuple with the VmTagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmTagIds

`func (o *GcpNativeProtectionGroupParams) SetVmTagIds(v [][]int64)`

SetVmTagIds sets VmTagIds field to given value.

### HasVmTagIds

`func (o *GcpNativeProtectionGroupParams) HasVmTagIds() bool`

HasVmTagIds returns a boolean if a field has been set.

### SetVmTagIdsNil

`func (o *GcpNativeProtectionGroupParams) SetVmTagIdsNil(b bool)`

 SetVmTagIdsNil sets the value for VmTagIds to be an explicit nil

### UnsetVmTagIds
`func (o *GcpNativeProtectionGroupParams) UnsetVmTagIds()`

UnsetVmTagIds ensures that no value is present for VmTagIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


