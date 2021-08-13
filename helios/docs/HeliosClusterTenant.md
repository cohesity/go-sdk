# HeliosClusterTenant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | Pointer to **NullableInt64** | Specifies the cluster id of the cluster. | [optional] 
**ClusterIncarnationId** | Pointer to **NullableInt64** | Specifies the incarnation id of the cluster. | [optional] 
**RegionId** | Pointer to **NullableString** | Specifies the region id of the cluster. Only valid for DMaaS clusters. | [optional] 
**Id** | Pointer to **NullableString** | The tenant id. | [optional] 
**Name** | Pointer to **NullableString** | Name of the Tenant. | [optional] 
**Description** | Pointer to **NullableString** | Description about the tenant. | [optional] 
**Status** | Pointer to **NullableString** | Current Status of the Tenant. | [optional] 
**Network** | Pointer to [**TenantNetwork**](TenantNetwork.md) |  | [optional] 
**CreatedAtTimeMsecs** | Pointer to **NullableInt64** | Epoch time when tenant was created. | [optional] [readonly] 
**LastUpdatedAtTimeMsecs** | Pointer to **NullableInt64** | Epoch time when tenant was last updated. | [optional] [readonly] 
**DeletedAtTimeMsecs** | Pointer to **NullableInt64** | Epoch time when tenant was last updated. | [optional] [readonly] 

## Methods

### NewHeliosClusterTenant

`func NewHeliosClusterTenant() *HeliosClusterTenant`

NewHeliosClusterTenant instantiates a new HeliosClusterTenant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeliosClusterTenantWithDefaults

`func NewHeliosClusterTenantWithDefaults() *HeliosClusterTenant`

NewHeliosClusterTenantWithDefaults instantiates a new HeliosClusterTenant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *HeliosClusterTenant) GetClusterId() int64`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *HeliosClusterTenant) GetClusterIdOk() (*int64, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *HeliosClusterTenant) SetClusterId(v int64)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *HeliosClusterTenant) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### SetClusterIdNil

`func (o *HeliosClusterTenant) SetClusterIdNil(b bool)`

 SetClusterIdNil sets the value for ClusterId to be an explicit nil

### UnsetClusterId
`func (o *HeliosClusterTenant) UnsetClusterId()`

UnsetClusterId ensures that no value is present for ClusterId, not even an explicit nil
### GetClusterIncarnationId

`func (o *HeliosClusterTenant) GetClusterIncarnationId() int64`

GetClusterIncarnationId returns the ClusterIncarnationId field if non-nil, zero value otherwise.

### GetClusterIncarnationIdOk

`func (o *HeliosClusterTenant) GetClusterIncarnationIdOk() (*int64, bool)`

GetClusterIncarnationIdOk returns a tuple with the ClusterIncarnationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIncarnationId

`func (o *HeliosClusterTenant) SetClusterIncarnationId(v int64)`

SetClusterIncarnationId sets ClusterIncarnationId field to given value.

### HasClusterIncarnationId

`func (o *HeliosClusterTenant) HasClusterIncarnationId() bool`

HasClusterIncarnationId returns a boolean if a field has been set.

### SetClusterIncarnationIdNil

`func (o *HeliosClusterTenant) SetClusterIncarnationIdNil(b bool)`

 SetClusterIncarnationIdNil sets the value for ClusterIncarnationId to be an explicit nil

### UnsetClusterIncarnationId
`func (o *HeliosClusterTenant) UnsetClusterIncarnationId()`

UnsetClusterIncarnationId ensures that no value is present for ClusterIncarnationId, not even an explicit nil
### GetRegionId

`func (o *HeliosClusterTenant) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *HeliosClusterTenant) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *HeliosClusterTenant) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.

### HasRegionId

`func (o *HeliosClusterTenant) HasRegionId() bool`

HasRegionId returns a boolean if a field has been set.

### SetRegionIdNil

`func (o *HeliosClusterTenant) SetRegionIdNil(b bool)`

 SetRegionIdNil sets the value for RegionId to be an explicit nil

### UnsetRegionId
`func (o *HeliosClusterTenant) UnsetRegionId()`

UnsetRegionId ensures that no value is present for RegionId, not even an explicit nil
### GetId

`func (o *HeliosClusterTenant) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HeliosClusterTenant) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HeliosClusterTenant) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HeliosClusterTenant) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *HeliosClusterTenant) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *HeliosClusterTenant) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *HeliosClusterTenant) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HeliosClusterTenant) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HeliosClusterTenant) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HeliosClusterTenant) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *HeliosClusterTenant) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *HeliosClusterTenant) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *HeliosClusterTenant) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HeliosClusterTenant) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HeliosClusterTenant) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HeliosClusterTenant) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *HeliosClusterTenant) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *HeliosClusterTenant) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetStatus

`func (o *HeliosClusterTenant) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HeliosClusterTenant) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HeliosClusterTenant) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HeliosClusterTenant) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *HeliosClusterTenant) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *HeliosClusterTenant) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetNetwork

`func (o *HeliosClusterTenant) GetNetwork() TenantNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *HeliosClusterTenant) GetNetworkOk() (*TenantNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *HeliosClusterTenant) SetNetwork(v TenantNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *HeliosClusterTenant) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetCreatedAtTimeMsecs

`func (o *HeliosClusterTenant) GetCreatedAtTimeMsecs() int64`

GetCreatedAtTimeMsecs returns the CreatedAtTimeMsecs field if non-nil, zero value otherwise.

### GetCreatedAtTimeMsecsOk

`func (o *HeliosClusterTenant) GetCreatedAtTimeMsecsOk() (*int64, bool)`

GetCreatedAtTimeMsecsOk returns a tuple with the CreatedAtTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAtTimeMsecs

`func (o *HeliosClusterTenant) SetCreatedAtTimeMsecs(v int64)`

SetCreatedAtTimeMsecs sets CreatedAtTimeMsecs field to given value.

### HasCreatedAtTimeMsecs

`func (o *HeliosClusterTenant) HasCreatedAtTimeMsecs() bool`

HasCreatedAtTimeMsecs returns a boolean if a field has been set.

### SetCreatedAtTimeMsecsNil

`func (o *HeliosClusterTenant) SetCreatedAtTimeMsecsNil(b bool)`

 SetCreatedAtTimeMsecsNil sets the value for CreatedAtTimeMsecs to be an explicit nil

### UnsetCreatedAtTimeMsecs
`func (o *HeliosClusterTenant) UnsetCreatedAtTimeMsecs()`

UnsetCreatedAtTimeMsecs ensures that no value is present for CreatedAtTimeMsecs, not even an explicit nil
### GetLastUpdatedAtTimeMsecs

`func (o *HeliosClusterTenant) GetLastUpdatedAtTimeMsecs() int64`

GetLastUpdatedAtTimeMsecs returns the LastUpdatedAtTimeMsecs field if non-nil, zero value otherwise.

### GetLastUpdatedAtTimeMsecsOk

`func (o *HeliosClusterTenant) GetLastUpdatedAtTimeMsecsOk() (*int64, bool)`

GetLastUpdatedAtTimeMsecsOk returns a tuple with the LastUpdatedAtTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedAtTimeMsecs

`func (o *HeliosClusterTenant) SetLastUpdatedAtTimeMsecs(v int64)`

SetLastUpdatedAtTimeMsecs sets LastUpdatedAtTimeMsecs field to given value.

### HasLastUpdatedAtTimeMsecs

`func (o *HeliosClusterTenant) HasLastUpdatedAtTimeMsecs() bool`

HasLastUpdatedAtTimeMsecs returns a boolean if a field has been set.

### SetLastUpdatedAtTimeMsecsNil

`func (o *HeliosClusterTenant) SetLastUpdatedAtTimeMsecsNil(b bool)`

 SetLastUpdatedAtTimeMsecsNil sets the value for LastUpdatedAtTimeMsecs to be an explicit nil

### UnsetLastUpdatedAtTimeMsecs
`func (o *HeliosClusterTenant) UnsetLastUpdatedAtTimeMsecs()`

UnsetLastUpdatedAtTimeMsecs ensures that no value is present for LastUpdatedAtTimeMsecs, not even an explicit nil
### GetDeletedAtTimeMsecs

`func (o *HeliosClusterTenant) GetDeletedAtTimeMsecs() int64`

GetDeletedAtTimeMsecs returns the DeletedAtTimeMsecs field if non-nil, zero value otherwise.

### GetDeletedAtTimeMsecsOk

`func (o *HeliosClusterTenant) GetDeletedAtTimeMsecsOk() (*int64, bool)`

GetDeletedAtTimeMsecsOk returns a tuple with the DeletedAtTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAtTimeMsecs

`func (o *HeliosClusterTenant) SetDeletedAtTimeMsecs(v int64)`

SetDeletedAtTimeMsecs sets DeletedAtTimeMsecs field to given value.

### HasDeletedAtTimeMsecs

`func (o *HeliosClusterTenant) HasDeletedAtTimeMsecs() bool`

HasDeletedAtTimeMsecs returns a boolean if a field has been set.

### SetDeletedAtTimeMsecsNil

`func (o *HeliosClusterTenant) SetDeletedAtTimeMsecsNil(b bool)`

 SetDeletedAtTimeMsecsNil sets the value for DeletedAtTimeMsecs to be an explicit nil

### UnsetDeletedAtTimeMsecs
`func (o *HeliosClusterTenant) UnsetDeletedAtTimeMsecs()`

UnsetDeletedAtTimeMsecs ensures that no value is present for DeletedAtTimeMsecs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


