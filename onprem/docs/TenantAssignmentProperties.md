# TenantAssignmentProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StorageDomains** | Pointer to [**StorageDomains**](StorageDomains.md) |  | [optional] 
**Objects** | Pointer to **[]int64** | A list of Ids of properties assigned to the tenant. | [optional] 
**Vlans** | Pointer to **[]string** | A list of Ids of properties assigned to the tenant. | [optional] 
**Views** | Pointer to [**GetViewsResult**](GetViewsResult.md) |  | [optional] 
**Policies** | Pointer to [**ProtectionPolicyResponseWithPagination**](ProtectionPolicyResponseWithPagination.md) |  | [optional] 

## Methods

### NewTenantAssignmentProperties

`func NewTenantAssignmentProperties() *TenantAssignmentProperties`

NewTenantAssignmentProperties instantiates a new TenantAssignmentProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantAssignmentPropertiesWithDefaults

`func NewTenantAssignmentPropertiesWithDefaults() *TenantAssignmentProperties`

NewTenantAssignmentPropertiesWithDefaults instantiates a new TenantAssignmentProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStorageDomains

`func (o *TenantAssignmentProperties) GetStorageDomains() StorageDomains`

GetStorageDomains returns the StorageDomains field if non-nil, zero value otherwise.

### GetStorageDomainsOk

`func (o *TenantAssignmentProperties) GetStorageDomainsOk() (*StorageDomains, bool)`

GetStorageDomainsOk returns a tuple with the StorageDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageDomains

`func (o *TenantAssignmentProperties) SetStorageDomains(v StorageDomains)`

SetStorageDomains sets StorageDomains field to given value.

### HasStorageDomains

`func (o *TenantAssignmentProperties) HasStorageDomains() bool`

HasStorageDomains returns a boolean if a field has been set.

### GetObjects

`func (o *TenantAssignmentProperties) GetObjects() []int64`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *TenantAssignmentProperties) GetObjectsOk() (*[]int64, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *TenantAssignmentProperties) SetObjects(v []int64)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *TenantAssignmentProperties) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### GetVlans

`func (o *TenantAssignmentProperties) GetVlans() []string`

GetVlans returns the Vlans field if non-nil, zero value otherwise.

### GetVlansOk

`func (o *TenantAssignmentProperties) GetVlansOk() (*[]string, bool)`

GetVlansOk returns a tuple with the Vlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlans

`func (o *TenantAssignmentProperties) SetVlans(v []string)`

SetVlans sets Vlans field to given value.

### HasVlans

`func (o *TenantAssignmentProperties) HasVlans() bool`

HasVlans returns a boolean if a field has been set.

### GetViews

`func (o *TenantAssignmentProperties) GetViews() GetViewsResult`

GetViews returns the Views field if non-nil, zero value otherwise.

### GetViewsOk

`func (o *TenantAssignmentProperties) GetViewsOk() (*GetViewsResult, bool)`

GetViewsOk returns a tuple with the Views field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViews

`func (o *TenantAssignmentProperties) SetViews(v GetViewsResult)`

SetViews sets Views field to given value.

### HasViews

`func (o *TenantAssignmentProperties) HasViews() bool`

HasViews returns a boolean if a field has been set.

### GetPolicies

`func (o *TenantAssignmentProperties) GetPolicies() ProtectionPolicyResponseWithPagination`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *TenantAssignmentProperties) GetPoliciesOk() (*ProtectionPolicyResponseWithPagination, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *TenantAssignmentProperties) SetPolicies(v ProtectionPolicyResponseWithPagination)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *TenantAssignmentProperties) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


