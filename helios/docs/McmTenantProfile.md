# McmTenantProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **NullableString** | Specifies the tenant id. | 
**TenantName** | Pointer to **NullableString** | Specifies the tenant id. | [optional] [readonly] 
**TenantType** | **NullableString** | Specifies the type of tenant. | 
**Clusters** | Pointer to [**[]McmClusterIdentifier**](McmClusterIdentifier.md) | Specifies the list of clusters owned by this account id. | [optional] 

## Methods

### NewMcmTenantProfile

`func NewMcmTenantProfile(tenantId NullableString, tenantType NullableString, ) *McmTenantProfile`

NewMcmTenantProfile instantiates a new McmTenantProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcmTenantProfileWithDefaults

`func NewMcmTenantProfileWithDefaults() *McmTenantProfile`

NewMcmTenantProfileWithDefaults instantiates a new McmTenantProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *McmTenantProfile) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *McmTenantProfile) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *McmTenantProfile) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### SetTenantIdNil

`func (o *McmTenantProfile) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *McmTenantProfile) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetTenantName

`func (o *McmTenantProfile) GetTenantName() string`

GetTenantName returns the TenantName field if non-nil, zero value otherwise.

### GetTenantNameOk

`func (o *McmTenantProfile) GetTenantNameOk() (*string, bool)`

GetTenantNameOk returns a tuple with the TenantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantName

`func (o *McmTenantProfile) SetTenantName(v string)`

SetTenantName sets TenantName field to given value.

### HasTenantName

`func (o *McmTenantProfile) HasTenantName() bool`

HasTenantName returns a boolean if a field has been set.

### SetTenantNameNil

`func (o *McmTenantProfile) SetTenantNameNil(b bool)`

 SetTenantNameNil sets the value for TenantName to be an explicit nil

### UnsetTenantName
`func (o *McmTenantProfile) UnsetTenantName()`

UnsetTenantName ensures that no value is present for TenantName, not even an explicit nil
### GetTenantType

`func (o *McmTenantProfile) GetTenantType() string`

GetTenantType returns the TenantType field if non-nil, zero value otherwise.

### GetTenantTypeOk

`func (o *McmTenantProfile) GetTenantTypeOk() (*string, bool)`

GetTenantTypeOk returns a tuple with the TenantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantType

`func (o *McmTenantProfile) SetTenantType(v string)`

SetTenantType sets TenantType field to given value.


### SetTenantTypeNil

`func (o *McmTenantProfile) SetTenantTypeNil(b bool)`

 SetTenantTypeNil sets the value for TenantType to be an explicit nil

### UnsetTenantType
`func (o *McmTenantProfile) UnsetTenantType()`

UnsetTenantType ensures that no value is present for TenantType, not even an explicit nil
### GetClusters

`func (o *McmTenantProfile) GetClusters() []McmClusterIdentifier`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *McmTenantProfile) GetClustersOk() (*[]McmClusterIdentifier, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *McmTenantProfile) SetClusters(v []McmClusterIdentifier)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *McmTenantProfile) HasClusters() bool`

HasClusters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


