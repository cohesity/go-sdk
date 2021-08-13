# TenantAssignmentsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StorageDomains** | Pointer to **[]int64** | A list of Ids of properties assigned to the tenant. | [optional] 
**Objects** | Pointer to **[]int64** | A list of Ids of properties assigned to the tenant. | [optional] 
**Vlans** | Pointer to **[]string** | A list of Ids of properties assigned to the tenant. | [optional] 
**Views** | Pointer to **[]int64** | A list of Ids of properties assigned to the tenant. | [optional] 
**Policies** | Pointer to **[]string** | A list of Ids of properties assigned to the tenant. | [optional] 

## Methods

### NewTenantAssignmentsResult

`func NewTenantAssignmentsResult() *TenantAssignmentsResult`

NewTenantAssignmentsResult instantiates a new TenantAssignmentsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantAssignmentsResultWithDefaults

`func NewTenantAssignmentsResultWithDefaults() *TenantAssignmentsResult`

NewTenantAssignmentsResultWithDefaults instantiates a new TenantAssignmentsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStorageDomains

`func (o *TenantAssignmentsResult) GetStorageDomains() []int64`

GetStorageDomains returns the StorageDomains field if non-nil, zero value otherwise.

### GetStorageDomainsOk

`func (o *TenantAssignmentsResult) GetStorageDomainsOk() (*[]int64, bool)`

GetStorageDomainsOk returns a tuple with the StorageDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageDomains

`func (o *TenantAssignmentsResult) SetStorageDomains(v []int64)`

SetStorageDomains sets StorageDomains field to given value.

### HasStorageDomains

`func (o *TenantAssignmentsResult) HasStorageDomains() bool`

HasStorageDomains returns a boolean if a field has been set.

### GetObjects

`func (o *TenantAssignmentsResult) GetObjects() []int64`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *TenantAssignmentsResult) GetObjectsOk() (*[]int64, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *TenantAssignmentsResult) SetObjects(v []int64)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *TenantAssignmentsResult) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### GetVlans

`func (o *TenantAssignmentsResult) GetVlans() []string`

GetVlans returns the Vlans field if non-nil, zero value otherwise.

### GetVlansOk

`func (o *TenantAssignmentsResult) GetVlansOk() (*[]string, bool)`

GetVlansOk returns a tuple with the Vlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlans

`func (o *TenantAssignmentsResult) SetVlans(v []string)`

SetVlans sets Vlans field to given value.

### HasVlans

`func (o *TenantAssignmentsResult) HasVlans() bool`

HasVlans returns a boolean if a field has been set.

### GetViews

`func (o *TenantAssignmentsResult) GetViews() []int64`

GetViews returns the Views field if non-nil, zero value otherwise.

### GetViewsOk

`func (o *TenantAssignmentsResult) GetViewsOk() (*[]int64, bool)`

GetViewsOk returns a tuple with the Views field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViews

`func (o *TenantAssignmentsResult) SetViews(v []int64)`

SetViews sets Views field to given value.

### HasViews

`func (o *TenantAssignmentsResult) HasViews() bool`

HasViews returns a boolean if a field has been set.

### GetPolicies

`func (o *TenantAssignmentsResult) GetPolicies() []string`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *TenantAssignmentsResult) GetPoliciesOk() (*[]string, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *TenantAssignmentsResult) SetPolicies(v []string)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *TenantAssignmentsResult) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


