# TenantAssignmentsParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectIds** | Pointer to **[]int64** | List of objects on the cluster, to be associated to the Tenant. | [optional] 
**PolicyIds** | Pointer to **[]string** | List of policies on the cluster, to be associated to the Tenant. | [optional] 
**StorageDomainIds** | Pointer to **[]int64** | List of storage domains on the cluster, to be associated to the Tenant. | [optional] 
**ViewIds** | Pointer to **[]int64** | List of views on the cluster, to be associated to the Tenant. | [optional] 
**VlanIfaceNames** | Pointer to **[]string** | List of vlans on the cluster, to be associated to the Tenant. | [optional] 

## Methods

### NewTenantAssignmentsParams

`func NewTenantAssignmentsParams() *TenantAssignmentsParams`

NewTenantAssignmentsParams instantiates a new TenantAssignmentsParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantAssignmentsParamsWithDefaults

`func NewTenantAssignmentsParamsWithDefaults() *TenantAssignmentsParams`

NewTenantAssignmentsParamsWithDefaults instantiates a new TenantAssignmentsParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectIds

`func (o *TenantAssignmentsParams) GetObjectIds() []int64`

GetObjectIds returns the ObjectIds field if non-nil, zero value otherwise.

### GetObjectIdsOk

`func (o *TenantAssignmentsParams) GetObjectIdsOk() (*[]int64, bool)`

GetObjectIdsOk returns a tuple with the ObjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectIds

`func (o *TenantAssignmentsParams) SetObjectIds(v []int64)`

SetObjectIds sets ObjectIds field to given value.

### HasObjectIds

`func (o *TenantAssignmentsParams) HasObjectIds() bool`

HasObjectIds returns a boolean if a field has been set.

### SetObjectIdsNil

`func (o *TenantAssignmentsParams) SetObjectIdsNil(b bool)`

 SetObjectIdsNil sets the value for ObjectIds to be an explicit nil

### UnsetObjectIds
`func (o *TenantAssignmentsParams) UnsetObjectIds()`

UnsetObjectIds ensures that no value is present for ObjectIds, not even an explicit nil
### GetPolicyIds

`func (o *TenantAssignmentsParams) GetPolicyIds() []string`

GetPolicyIds returns the PolicyIds field if non-nil, zero value otherwise.

### GetPolicyIdsOk

`func (o *TenantAssignmentsParams) GetPolicyIdsOk() (*[]string, bool)`

GetPolicyIdsOk returns a tuple with the PolicyIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyIds

`func (o *TenantAssignmentsParams) SetPolicyIds(v []string)`

SetPolicyIds sets PolicyIds field to given value.

### HasPolicyIds

`func (o *TenantAssignmentsParams) HasPolicyIds() bool`

HasPolicyIds returns a boolean if a field has been set.

### SetPolicyIdsNil

`func (o *TenantAssignmentsParams) SetPolicyIdsNil(b bool)`

 SetPolicyIdsNil sets the value for PolicyIds to be an explicit nil

### UnsetPolicyIds
`func (o *TenantAssignmentsParams) UnsetPolicyIds()`

UnsetPolicyIds ensures that no value is present for PolicyIds, not even an explicit nil
### GetStorageDomainIds

`func (o *TenantAssignmentsParams) GetStorageDomainIds() []int64`

GetStorageDomainIds returns the StorageDomainIds field if non-nil, zero value otherwise.

### GetStorageDomainIdsOk

`func (o *TenantAssignmentsParams) GetStorageDomainIdsOk() (*[]int64, bool)`

GetStorageDomainIdsOk returns a tuple with the StorageDomainIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageDomainIds

`func (o *TenantAssignmentsParams) SetStorageDomainIds(v []int64)`

SetStorageDomainIds sets StorageDomainIds field to given value.

### HasStorageDomainIds

`func (o *TenantAssignmentsParams) HasStorageDomainIds() bool`

HasStorageDomainIds returns a boolean if a field has been set.

### SetStorageDomainIdsNil

`func (o *TenantAssignmentsParams) SetStorageDomainIdsNil(b bool)`

 SetStorageDomainIdsNil sets the value for StorageDomainIds to be an explicit nil

### UnsetStorageDomainIds
`func (o *TenantAssignmentsParams) UnsetStorageDomainIds()`

UnsetStorageDomainIds ensures that no value is present for StorageDomainIds, not even an explicit nil
### GetViewIds

`func (o *TenantAssignmentsParams) GetViewIds() []int64`

GetViewIds returns the ViewIds field if non-nil, zero value otherwise.

### GetViewIdsOk

`func (o *TenantAssignmentsParams) GetViewIdsOk() (*[]int64, bool)`

GetViewIdsOk returns a tuple with the ViewIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewIds

`func (o *TenantAssignmentsParams) SetViewIds(v []int64)`

SetViewIds sets ViewIds field to given value.

### HasViewIds

`func (o *TenantAssignmentsParams) HasViewIds() bool`

HasViewIds returns a boolean if a field has been set.

### SetViewIdsNil

`func (o *TenantAssignmentsParams) SetViewIdsNil(b bool)`

 SetViewIdsNil sets the value for ViewIds to be an explicit nil

### UnsetViewIds
`func (o *TenantAssignmentsParams) UnsetViewIds()`

UnsetViewIds ensures that no value is present for ViewIds, not even an explicit nil
### GetVlanIfaceNames

`func (o *TenantAssignmentsParams) GetVlanIfaceNames() []string`

GetVlanIfaceNames returns the VlanIfaceNames field if non-nil, zero value otherwise.

### GetVlanIfaceNamesOk

`func (o *TenantAssignmentsParams) GetVlanIfaceNamesOk() (*[]string, bool)`

GetVlanIfaceNamesOk returns a tuple with the VlanIfaceNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanIfaceNames

`func (o *TenantAssignmentsParams) SetVlanIfaceNames(v []string)`

SetVlanIfaceNames sets VlanIfaceNames field to given value.

### HasVlanIfaceNames

`func (o *TenantAssignmentsParams) HasVlanIfaceNames() bool`

HasVlanIfaceNames returns a boolean if a field has been set.

### SetVlanIfaceNamesNil

`func (o *TenantAssignmentsParams) SetVlanIfaceNamesNil(b bool)`

 SetVlanIfaceNamesNil sets the value for VlanIfaceNames to be an explicit nil

### UnsetVlanIfaceNames
`func (o *TenantAssignmentsParams) UnsetVlanIfaceNames()`

UnsetVlanIfaceNames ensures that no value is present for VlanIfaceNames, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


