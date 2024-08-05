# Tenant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAtTimeMsecs** | Pointer to **NullableInt64** | Epoch time when tenant was created. | [optional] [readonly] 
**DeletedAtTimeMsecs** | Pointer to **NullableInt64** | Epoch time when tenant was last updated. | [optional] [readonly] 
**Description** | Pointer to **NullableString** | Description about the tenant. | [optional] 
**ExternalVendorMetadata** | Pointer to [**ExternalVendorTenantMetadata**](ExternalVendorTenantMetadata.md) |  | [optional] 
**Id** | Pointer to **NullableString** | The tenant id. | [optional] 
**IsManagedOnHelios** | Pointer to **NullableBool** | Flag to indicate if tenant is managed on helios | [optional] 
**LastUpdatedAtTimeMsecs** | Pointer to **NullableInt64** | Epoch time when tenant was last updated. | [optional] [readonly] 
**Name** | Pointer to **NullableString** | Name of the Tenant. | [optional] 
**Network** | Pointer to [**TenantNetwork**](TenantNetwork.md) |  | [optional] 
**Status** | Pointer to **NullableString** | Current Status of the Tenant. | [optional] 

## Methods

### NewTenant

`func NewTenant() *Tenant`

NewTenant instantiates a new Tenant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantWithDefaults

`func NewTenantWithDefaults() *Tenant`

NewTenantWithDefaults instantiates a new Tenant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAtTimeMsecs

`func (o *Tenant) GetCreatedAtTimeMsecs() int64`

GetCreatedAtTimeMsecs returns the CreatedAtTimeMsecs field if non-nil, zero value otherwise.

### GetCreatedAtTimeMsecsOk

`func (o *Tenant) GetCreatedAtTimeMsecsOk() (*int64, bool)`

GetCreatedAtTimeMsecsOk returns a tuple with the CreatedAtTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAtTimeMsecs

`func (o *Tenant) SetCreatedAtTimeMsecs(v int64)`

SetCreatedAtTimeMsecs sets CreatedAtTimeMsecs field to given value.

### HasCreatedAtTimeMsecs

`func (o *Tenant) HasCreatedAtTimeMsecs() bool`

HasCreatedAtTimeMsecs returns a boolean if a field has been set.

### SetCreatedAtTimeMsecsNil

`func (o *Tenant) SetCreatedAtTimeMsecsNil(b bool)`

 SetCreatedAtTimeMsecsNil sets the value for CreatedAtTimeMsecs to be an explicit nil

### UnsetCreatedAtTimeMsecs
`func (o *Tenant) UnsetCreatedAtTimeMsecs()`

UnsetCreatedAtTimeMsecs ensures that no value is present for CreatedAtTimeMsecs, not even an explicit nil
### GetDeletedAtTimeMsecs

`func (o *Tenant) GetDeletedAtTimeMsecs() int64`

GetDeletedAtTimeMsecs returns the DeletedAtTimeMsecs field if non-nil, zero value otherwise.

### GetDeletedAtTimeMsecsOk

`func (o *Tenant) GetDeletedAtTimeMsecsOk() (*int64, bool)`

GetDeletedAtTimeMsecsOk returns a tuple with the DeletedAtTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAtTimeMsecs

`func (o *Tenant) SetDeletedAtTimeMsecs(v int64)`

SetDeletedAtTimeMsecs sets DeletedAtTimeMsecs field to given value.

### HasDeletedAtTimeMsecs

`func (o *Tenant) HasDeletedAtTimeMsecs() bool`

HasDeletedAtTimeMsecs returns a boolean if a field has been set.

### SetDeletedAtTimeMsecsNil

`func (o *Tenant) SetDeletedAtTimeMsecsNil(b bool)`

 SetDeletedAtTimeMsecsNil sets the value for DeletedAtTimeMsecs to be an explicit nil

### UnsetDeletedAtTimeMsecs
`func (o *Tenant) UnsetDeletedAtTimeMsecs()`

UnsetDeletedAtTimeMsecs ensures that no value is present for DeletedAtTimeMsecs, not even an explicit nil
### GetDescription

`func (o *Tenant) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Tenant) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Tenant) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Tenant) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Tenant) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Tenant) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetExternalVendorMetadata

`func (o *Tenant) GetExternalVendorMetadata() ExternalVendorTenantMetadata`

GetExternalVendorMetadata returns the ExternalVendorMetadata field if non-nil, zero value otherwise.

### GetExternalVendorMetadataOk

`func (o *Tenant) GetExternalVendorMetadataOk() (*ExternalVendorTenantMetadata, bool)`

GetExternalVendorMetadataOk returns a tuple with the ExternalVendorMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalVendorMetadata

`func (o *Tenant) SetExternalVendorMetadata(v ExternalVendorTenantMetadata)`

SetExternalVendorMetadata sets ExternalVendorMetadata field to given value.

### HasExternalVendorMetadata

`func (o *Tenant) HasExternalVendorMetadata() bool`

HasExternalVendorMetadata returns a boolean if a field has been set.

### GetId

`func (o *Tenant) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Tenant) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Tenant) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Tenant) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *Tenant) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *Tenant) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIsManagedOnHelios

`func (o *Tenant) GetIsManagedOnHelios() bool`

GetIsManagedOnHelios returns the IsManagedOnHelios field if non-nil, zero value otherwise.

### GetIsManagedOnHeliosOk

`func (o *Tenant) GetIsManagedOnHeliosOk() (*bool, bool)`

GetIsManagedOnHeliosOk returns a tuple with the IsManagedOnHelios field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsManagedOnHelios

`func (o *Tenant) SetIsManagedOnHelios(v bool)`

SetIsManagedOnHelios sets IsManagedOnHelios field to given value.

### HasIsManagedOnHelios

`func (o *Tenant) HasIsManagedOnHelios() bool`

HasIsManagedOnHelios returns a boolean if a field has been set.

### SetIsManagedOnHeliosNil

`func (o *Tenant) SetIsManagedOnHeliosNil(b bool)`

 SetIsManagedOnHeliosNil sets the value for IsManagedOnHelios to be an explicit nil

### UnsetIsManagedOnHelios
`func (o *Tenant) UnsetIsManagedOnHelios()`

UnsetIsManagedOnHelios ensures that no value is present for IsManagedOnHelios, not even an explicit nil
### GetLastUpdatedAtTimeMsecs

`func (o *Tenant) GetLastUpdatedAtTimeMsecs() int64`

GetLastUpdatedAtTimeMsecs returns the LastUpdatedAtTimeMsecs field if non-nil, zero value otherwise.

### GetLastUpdatedAtTimeMsecsOk

`func (o *Tenant) GetLastUpdatedAtTimeMsecsOk() (*int64, bool)`

GetLastUpdatedAtTimeMsecsOk returns a tuple with the LastUpdatedAtTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedAtTimeMsecs

`func (o *Tenant) SetLastUpdatedAtTimeMsecs(v int64)`

SetLastUpdatedAtTimeMsecs sets LastUpdatedAtTimeMsecs field to given value.

### HasLastUpdatedAtTimeMsecs

`func (o *Tenant) HasLastUpdatedAtTimeMsecs() bool`

HasLastUpdatedAtTimeMsecs returns a boolean if a field has been set.

### SetLastUpdatedAtTimeMsecsNil

`func (o *Tenant) SetLastUpdatedAtTimeMsecsNil(b bool)`

 SetLastUpdatedAtTimeMsecsNil sets the value for LastUpdatedAtTimeMsecs to be an explicit nil

### UnsetLastUpdatedAtTimeMsecs
`func (o *Tenant) UnsetLastUpdatedAtTimeMsecs()`

UnsetLastUpdatedAtTimeMsecs ensures that no value is present for LastUpdatedAtTimeMsecs, not even an explicit nil
### GetName

`func (o *Tenant) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Tenant) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Tenant) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Tenant) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Tenant) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Tenant) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetNetwork

`func (o *Tenant) GetNetwork() TenantNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *Tenant) GetNetworkOk() (*TenantNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *Tenant) SetNetwork(v TenantNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *Tenant) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetStatus

`func (o *Tenant) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Tenant) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Tenant) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Tenant) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *Tenant) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *Tenant) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


