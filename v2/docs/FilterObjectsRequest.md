# FilterObjectsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationEnvironment** | Pointer to **NullableString** | Specifies the type of application enviornment needed for filtering to be applied on. This is needed because in case of applications like SQL, Oracle, a single source can contain multiple application enviornments. | [optional] 
**FilterType** | **NullableString** | Specifies the type of filtering user wants to perform. Currently, we only support exclude type of filter. | 
**Filters** | [**[]Filter**](Filter.md) | Specifies the list of filters that need to be applied on given list of discovered objects. | 
**IncludeTenants** | Pointer to **NullableBool** | If true, the response will include objects which belongs to all tenants which the current user has permission to see. Default value is false. | [optional] [default to false]
**ObjectIds** | **[]int64** | Specifies a list of non leaf object ids to filter the leaf level objects. Non leaf object such host (physical or vm) or database instance can be specified. | 
**TenantIds** | Pointer to **[]string** | TenantIds contains list of the tenant for which objects are to be returned. | [optional] 

## Methods

### NewFilterObjectsRequest

`func NewFilterObjectsRequest(filterType NullableString, filters []Filter, objectIds []int64, ) *FilterObjectsRequest`

NewFilterObjectsRequest instantiates a new FilterObjectsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilterObjectsRequestWithDefaults

`func NewFilterObjectsRequestWithDefaults() *FilterObjectsRequest`

NewFilterObjectsRequestWithDefaults instantiates a new FilterObjectsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationEnvironment

`func (o *FilterObjectsRequest) GetApplicationEnvironment() string`

GetApplicationEnvironment returns the ApplicationEnvironment field if non-nil, zero value otherwise.

### GetApplicationEnvironmentOk

`func (o *FilterObjectsRequest) GetApplicationEnvironmentOk() (*string, bool)`

GetApplicationEnvironmentOk returns a tuple with the ApplicationEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationEnvironment

`func (o *FilterObjectsRequest) SetApplicationEnvironment(v string)`

SetApplicationEnvironment sets ApplicationEnvironment field to given value.

### HasApplicationEnvironment

`func (o *FilterObjectsRequest) HasApplicationEnvironment() bool`

HasApplicationEnvironment returns a boolean if a field has been set.

### SetApplicationEnvironmentNil

`func (o *FilterObjectsRequest) SetApplicationEnvironmentNil(b bool)`

 SetApplicationEnvironmentNil sets the value for ApplicationEnvironment to be an explicit nil

### UnsetApplicationEnvironment
`func (o *FilterObjectsRequest) UnsetApplicationEnvironment()`

UnsetApplicationEnvironment ensures that no value is present for ApplicationEnvironment, not even an explicit nil
### GetFilterType

`func (o *FilterObjectsRequest) GetFilterType() string`

GetFilterType returns the FilterType field if non-nil, zero value otherwise.

### GetFilterTypeOk

`func (o *FilterObjectsRequest) GetFilterTypeOk() (*string, bool)`

GetFilterTypeOk returns a tuple with the FilterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterType

`func (o *FilterObjectsRequest) SetFilterType(v string)`

SetFilterType sets FilterType field to given value.


### SetFilterTypeNil

`func (o *FilterObjectsRequest) SetFilterTypeNil(b bool)`

 SetFilterTypeNil sets the value for FilterType to be an explicit nil

### UnsetFilterType
`func (o *FilterObjectsRequest) UnsetFilterType()`

UnsetFilterType ensures that no value is present for FilterType, not even an explicit nil
### GetFilters

`func (o *FilterObjectsRequest) GetFilters() []Filter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *FilterObjectsRequest) GetFiltersOk() (*[]Filter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *FilterObjectsRequest) SetFilters(v []Filter)`

SetFilters sets Filters field to given value.


### SetFiltersNil

`func (o *FilterObjectsRequest) SetFiltersNil(b bool)`

 SetFiltersNil sets the value for Filters to be an explicit nil

### UnsetFilters
`func (o *FilterObjectsRequest) UnsetFilters()`

UnsetFilters ensures that no value is present for Filters, not even an explicit nil
### GetIncludeTenants

`func (o *FilterObjectsRequest) GetIncludeTenants() bool`

GetIncludeTenants returns the IncludeTenants field if non-nil, zero value otherwise.

### GetIncludeTenantsOk

`func (o *FilterObjectsRequest) GetIncludeTenantsOk() (*bool, bool)`

GetIncludeTenantsOk returns a tuple with the IncludeTenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeTenants

`func (o *FilterObjectsRequest) SetIncludeTenants(v bool)`

SetIncludeTenants sets IncludeTenants field to given value.

### HasIncludeTenants

`func (o *FilterObjectsRequest) HasIncludeTenants() bool`

HasIncludeTenants returns a boolean if a field has been set.

### SetIncludeTenantsNil

`func (o *FilterObjectsRequest) SetIncludeTenantsNil(b bool)`

 SetIncludeTenantsNil sets the value for IncludeTenants to be an explicit nil

### UnsetIncludeTenants
`func (o *FilterObjectsRequest) UnsetIncludeTenants()`

UnsetIncludeTenants ensures that no value is present for IncludeTenants, not even an explicit nil
### GetObjectIds

`func (o *FilterObjectsRequest) GetObjectIds() []int64`

GetObjectIds returns the ObjectIds field if non-nil, zero value otherwise.

### GetObjectIdsOk

`func (o *FilterObjectsRequest) GetObjectIdsOk() (*[]int64, bool)`

GetObjectIdsOk returns a tuple with the ObjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectIds

`func (o *FilterObjectsRequest) SetObjectIds(v []int64)`

SetObjectIds sets ObjectIds field to given value.


### SetObjectIdsNil

`func (o *FilterObjectsRequest) SetObjectIdsNil(b bool)`

 SetObjectIdsNil sets the value for ObjectIds to be an explicit nil

### UnsetObjectIds
`func (o *FilterObjectsRequest) UnsetObjectIds()`

UnsetObjectIds ensures that no value is present for ObjectIds, not even an explicit nil
### GetTenantIds

`func (o *FilterObjectsRequest) GetTenantIds() []string`

GetTenantIds returns the TenantIds field if non-nil, zero value otherwise.

### GetTenantIdsOk

`func (o *FilterObjectsRequest) GetTenantIdsOk() (*[]string, bool)`

GetTenantIdsOk returns a tuple with the TenantIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantIds

`func (o *FilterObjectsRequest) SetTenantIds(v []string)`

SetTenantIds sets TenantIds field to given value.

### HasTenantIds

`func (o *FilterObjectsRequest) HasTenantIds() bool`

HasTenantIds returns a boolean if a field has been set.

### SetTenantIdsNil

`func (o *FilterObjectsRequest) SetTenantIdsNil(b bool)`

 SetTenantIdsNil sets the value for TenantIds to be an explicit nil

### UnsetTenantIds
`func (o *FilterObjectsRequest) UnsetTenantIds()`

UnsetTenantIds ensures that no value is present for TenantIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


