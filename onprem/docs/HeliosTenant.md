# HeliosTenant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | The tenant id. | [optional] 
**Name** | Pointer to **NullableString** | Name of the Tenant | [optional] 
**Description** | Pointer to **NullableString** | Description about the tenant. | [optional] 
**ManagedOnHelios** | Pointer to **NullableBool** | Wether managed on helios or not. | [optional] 
**Status** | Pointer to **NullableString** | Current Status of the Tenant. | [optional] 
**Systems** | Pointer to [**[]HeliosClusterTenant**](HeliosClusterTenant.md) | Details of tenant on each system that it is living. | [optional] 
**CreatedAtTimeMsecs** | Pointer to **NullableInt64** | Epoch time when tenant was created. | [optional] [readonly] 
**LastUpdatedAtTimeMsecs** | Pointer to **NullableInt64** | Epoch time when tenant was last updated. | [optional] [readonly] 
**DeletedAtTimeMsecs** | Pointer to **NullableInt64** | Epoch time when tenant was last updated. | [optional] [readonly] 

## Methods

### NewHeliosTenant

`func NewHeliosTenant() *HeliosTenant`

NewHeliosTenant instantiates a new HeliosTenant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeliosTenantWithDefaults

`func NewHeliosTenantWithDefaults() *HeliosTenant`

NewHeliosTenantWithDefaults instantiates a new HeliosTenant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HeliosTenant) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HeliosTenant) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HeliosTenant) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HeliosTenant) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *HeliosTenant) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *HeliosTenant) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *HeliosTenant) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HeliosTenant) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HeliosTenant) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HeliosTenant) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *HeliosTenant) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *HeliosTenant) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *HeliosTenant) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HeliosTenant) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HeliosTenant) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HeliosTenant) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *HeliosTenant) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *HeliosTenant) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetManagedOnHelios

`func (o *HeliosTenant) GetManagedOnHelios() bool`

GetManagedOnHelios returns the ManagedOnHelios field if non-nil, zero value otherwise.

### GetManagedOnHeliosOk

`func (o *HeliosTenant) GetManagedOnHeliosOk() (*bool, bool)`

GetManagedOnHeliosOk returns a tuple with the ManagedOnHelios field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedOnHelios

`func (o *HeliosTenant) SetManagedOnHelios(v bool)`

SetManagedOnHelios sets ManagedOnHelios field to given value.

### HasManagedOnHelios

`func (o *HeliosTenant) HasManagedOnHelios() bool`

HasManagedOnHelios returns a boolean if a field has been set.

### SetManagedOnHeliosNil

`func (o *HeliosTenant) SetManagedOnHeliosNil(b bool)`

 SetManagedOnHeliosNil sets the value for ManagedOnHelios to be an explicit nil

### UnsetManagedOnHelios
`func (o *HeliosTenant) UnsetManagedOnHelios()`

UnsetManagedOnHelios ensures that no value is present for ManagedOnHelios, not even an explicit nil
### GetStatus

`func (o *HeliosTenant) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HeliosTenant) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HeliosTenant) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HeliosTenant) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *HeliosTenant) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *HeliosTenant) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetSystems

`func (o *HeliosTenant) GetSystems() []HeliosClusterTenant`

GetSystems returns the Systems field if non-nil, zero value otherwise.

### GetSystemsOk

`func (o *HeliosTenant) GetSystemsOk() (*[]HeliosClusterTenant, bool)`

GetSystemsOk returns a tuple with the Systems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystems

`func (o *HeliosTenant) SetSystems(v []HeliosClusterTenant)`

SetSystems sets Systems field to given value.

### HasSystems

`func (o *HeliosTenant) HasSystems() bool`

HasSystems returns a boolean if a field has been set.

### GetCreatedAtTimeMsecs

`func (o *HeliosTenant) GetCreatedAtTimeMsecs() int64`

GetCreatedAtTimeMsecs returns the CreatedAtTimeMsecs field if non-nil, zero value otherwise.

### GetCreatedAtTimeMsecsOk

`func (o *HeliosTenant) GetCreatedAtTimeMsecsOk() (*int64, bool)`

GetCreatedAtTimeMsecsOk returns a tuple with the CreatedAtTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAtTimeMsecs

`func (o *HeliosTenant) SetCreatedAtTimeMsecs(v int64)`

SetCreatedAtTimeMsecs sets CreatedAtTimeMsecs field to given value.

### HasCreatedAtTimeMsecs

`func (o *HeliosTenant) HasCreatedAtTimeMsecs() bool`

HasCreatedAtTimeMsecs returns a boolean if a field has been set.

### SetCreatedAtTimeMsecsNil

`func (o *HeliosTenant) SetCreatedAtTimeMsecsNil(b bool)`

 SetCreatedAtTimeMsecsNil sets the value for CreatedAtTimeMsecs to be an explicit nil

### UnsetCreatedAtTimeMsecs
`func (o *HeliosTenant) UnsetCreatedAtTimeMsecs()`

UnsetCreatedAtTimeMsecs ensures that no value is present for CreatedAtTimeMsecs, not even an explicit nil
### GetLastUpdatedAtTimeMsecs

`func (o *HeliosTenant) GetLastUpdatedAtTimeMsecs() int64`

GetLastUpdatedAtTimeMsecs returns the LastUpdatedAtTimeMsecs field if non-nil, zero value otherwise.

### GetLastUpdatedAtTimeMsecsOk

`func (o *HeliosTenant) GetLastUpdatedAtTimeMsecsOk() (*int64, bool)`

GetLastUpdatedAtTimeMsecsOk returns a tuple with the LastUpdatedAtTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedAtTimeMsecs

`func (o *HeliosTenant) SetLastUpdatedAtTimeMsecs(v int64)`

SetLastUpdatedAtTimeMsecs sets LastUpdatedAtTimeMsecs field to given value.

### HasLastUpdatedAtTimeMsecs

`func (o *HeliosTenant) HasLastUpdatedAtTimeMsecs() bool`

HasLastUpdatedAtTimeMsecs returns a boolean if a field has been set.

### SetLastUpdatedAtTimeMsecsNil

`func (o *HeliosTenant) SetLastUpdatedAtTimeMsecsNil(b bool)`

 SetLastUpdatedAtTimeMsecsNil sets the value for LastUpdatedAtTimeMsecs to be an explicit nil

### UnsetLastUpdatedAtTimeMsecs
`func (o *HeliosTenant) UnsetLastUpdatedAtTimeMsecs()`

UnsetLastUpdatedAtTimeMsecs ensures that no value is present for LastUpdatedAtTimeMsecs, not even an explicit nil
### GetDeletedAtTimeMsecs

`func (o *HeliosTenant) GetDeletedAtTimeMsecs() int64`

GetDeletedAtTimeMsecs returns the DeletedAtTimeMsecs field if non-nil, zero value otherwise.

### GetDeletedAtTimeMsecsOk

`func (o *HeliosTenant) GetDeletedAtTimeMsecsOk() (*int64, bool)`

GetDeletedAtTimeMsecsOk returns a tuple with the DeletedAtTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAtTimeMsecs

`func (o *HeliosTenant) SetDeletedAtTimeMsecs(v int64)`

SetDeletedAtTimeMsecs sets DeletedAtTimeMsecs field to given value.

### HasDeletedAtTimeMsecs

`func (o *HeliosTenant) HasDeletedAtTimeMsecs() bool`

HasDeletedAtTimeMsecs returns a boolean if a field has been set.

### SetDeletedAtTimeMsecsNil

`func (o *HeliosTenant) SetDeletedAtTimeMsecsNil(b bool)`

 SetDeletedAtTimeMsecsNil sets the value for DeletedAtTimeMsecs to be an explicit nil

### UnsetDeletedAtTimeMsecs
`func (o *HeliosTenant) UnsetDeletedAtTimeMsecs()`

UnsetDeletedAtTimeMsecs ensures that no value is present for DeletedAtTimeMsecs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


