# TenantInfo

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

### NewTenantInfo

`func NewTenantInfo() *TenantInfo`

NewTenantInfo instantiates a new TenantInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantInfoWithDefaults

`func NewTenantInfoWithDefaults() *TenantInfo`

NewTenantInfoWithDefaults instantiates a new TenantInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAtTimeMsecs

`func (o *TenantInfo) GetCreatedAtTimeMsecs() int64`

GetCreatedAtTimeMsecs returns the CreatedAtTimeMsecs field if non-nil, zero value otherwise.

### GetCreatedAtTimeMsecsOk

`func (o *TenantInfo) GetCreatedAtTimeMsecsOk() (*int64, bool)`

GetCreatedAtTimeMsecsOk returns a tuple with the CreatedAtTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAtTimeMsecs

`func (o *TenantInfo) SetCreatedAtTimeMsecs(v int64)`

SetCreatedAtTimeMsecs sets CreatedAtTimeMsecs field to given value.

### HasCreatedAtTimeMsecs

`func (o *TenantInfo) HasCreatedAtTimeMsecs() bool`

HasCreatedAtTimeMsecs returns a boolean if a field has been set.

### SetCreatedAtTimeMsecsNil

`func (o *TenantInfo) SetCreatedAtTimeMsecsNil(b bool)`

 SetCreatedAtTimeMsecsNil sets the value for CreatedAtTimeMsecs to be an explicit nil

### UnsetCreatedAtTimeMsecs
`func (o *TenantInfo) UnsetCreatedAtTimeMsecs()`

UnsetCreatedAtTimeMsecs ensures that no value is present for CreatedAtTimeMsecs, not even an explicit nil
### GetDeletedAtTimeMsecs

`func (o *TenantInfo) GetDeletedAtTimeMsecs() int64`

GetDeletedAtTimeMsecs returns the DeletedAtTimeMsecs field if non-nil, zero value otherwise.

### GetDeletedAtTimeMsecsOk

`func (o *TenantInfo) GetDeletedAtTimeMsecsOk() (*int64, bool)`

GetDeletedAtTimeMsecsOk returns a tuple with the DeletedAtTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAtTimeMsecs

`func (o *TenantInfo) SetDeletedAtTimeMsecs(v int64)`

SetDeletedAtTimeMsecs sets DeletedAtTimeMsecs field to given value.

### HasDeletedAtTimeMsecs

`func (o *TenantInfo) HasDeletedAtTimeMsecs() bool`

HasDeletedAtTimeMsecs returns a boolean if a field has been set.

### SetDeletedAtTimeMsecsNil

`func (o *TenantInfo) SetDeletedAtTimeMsecsNil(b bool)`

 SetDeletedAtTimeMsecsNil sets the value for DeletedAtTimeMsecs to be an explicit nil

### UnsetDeletedAtTimeMsecs
`func (o *TenantInfo) UnsetDeletedAtTimeMsecs()`

UnsetDeletedAtTimeMsecs ensures that no value is present for DeletedAtTimeMsecs, not even an explicit nil
### GetDescription

`func (o *TenantInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TenantInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TenantInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TenantInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TenantInfo) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TenantInfo) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetExternalVendorMetadata

`func (o *TenantInfo) GetExternalVendorMetadata() ExternalVendorTenantMetadata`

GetExternalVendorMetadata returns the ExternalVendorMetadata field if non-nil, zero value otherwise.

### GetExternalVendorMetadataOk

`func (o *TenantInfo) GetExternalVendorMetadataOk() (*ExternalVendorTenantMetadata, bool)`

GetExternalVendorMetadataOk returns a tuple with the ExternalVendorMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalVendorMetadata

`func (o *TenantInfo) SetExternalVendorMetadata(v ExternalVendorTenantMetadata)`

SetExternalVendorMetadata sets ExternalVendorMetadata field to given value.

### HasExternalVendorMetadata

`func (o *TenantInfo) HasExternalVendorMetadata() bool`

HasExternalVendorMetadata returns a boolean if a field has been set.

### GetId

`func (o *TenantInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TenantInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TenantInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TenantInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *TenantInfo) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *TenantInfo) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIsManagedOnHelios

`func (o *TenantInfo) GetIsManagedOnHelios() bool`

GetIsManagedOnHelios returns the IsManagedOnHelios field if non-nil, zero value otherwise.

### GetIsManagedOnHeliosOk

`func (o *TenantInfo) GetIsManagedOnHeliosOk() (*bool, bool)`

GetIsManagedOnHeliosOk returns a tuple with the IsManagedOnHelios field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsManagedOnHelios

`func (o *TenantInfo) SetIsManagedOnHelios(v bool)`

SetIsManagedOnHelios sets IsManagedOnHelios field to given value.

### HasIsManagedOnHelios

`func (o *TenantInfo) HasIsManagedOnHelios() bool`

HasIsManagedOnHelios returns a boolean if a field has been set.

### SetIsManagedOnHeliosNil

`func (o *TenantInfo) SetIsManagedOnHeliosNil(b bool)`

 SetIsManagedOnHeliosNil sets the value for IsManagedOnHelios to be an explicit nil

### UnsetIsManagedOnHelios
`func (o *TenantInfo) UnsetIsManagedOnHelios()`

UnsetIsManagedOnHelios ensures that no value is present for IsManagedOnHelios, not even an explicit nil
### GetLastUpdatedAtTimeMsecs

`func (o *TenantInfo) GetLastUpdatedAtTimeMsecs() int64`

GetLastUpdatedAtTimeMsecs returns the LastUpdatedAtTimeMsecs field if non-nil, zero value otherwise.

### GetLastUpdatedAtTimeMsecsOk

`func (o *TenantInfo) GetLastUpdatedAtTimeMsecsOk() (*int64, bool)`

GetLastUpdatedAtTimeMsecsOk returns a tuple with the LastUpdatedAtTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedAtTimeMsecs

`func (o *TenantInfo) SetLastUpdatedAtTimeMsecs(v int64)`

SetLastUpdatedAtTimeMsecs sets LastUpdatedAtTimeMsecs field to given value.

### HasLastUpdatedAtTimeMsecs

`func (o *TenantInfo) HasLastUpdatedAtTimeMsecs() bool`

HasLastUpdatedAtTimeMsecs returns a boolean if a field has been set.

### SetLastUpdatedAtTimeMsecsNil

`func (o *TenantInfo) SetLastUpdatedAtTimeMsecsNil(b bool)`

 SetLastUpdatedAtTimeMsecsNil sets the value for LastUpdatedAtTimeMsecs to be an explicit nil

### UnsetLastUpdatedAtTimeMsecs
`func (o *TenantInfo) UnsetLastUpdatedAtTimeMsecs()`

UnsetLastUpdatedAtTimeMsecs ensures that no value is present for LastUpdatedAtTimeMsecs, not even an explicit nil
### GetName

`func (o *TenantInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TenantInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TenantInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TenantInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TenantInfo) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TenantInfo) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetNetwork

`func (o *TenantInfo) GetNetwork() TenantNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *TenantInfo) GetNetworkOk() (*TenantNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *TenantInfo) SetNetwork(v TenantNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *TenantInfo) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetStatus

`func (o *TenantInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TenantInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TenantInfo) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TenantInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *TenantInfo) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *TenantInfo) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


