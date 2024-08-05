# IbmTenantMetadataParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **NullableString** | Specifies the unique identifier of the IBM&#39;s account ID. | [optional] 
**Crn** | Pointer to **NullableString** | Specifies the unique CRN associated with the tenant. | [optional] 
**CustomProperties** | Pointer to [**[]ExternalVendorCustomProperties**](ExternalVendorCustomProperties.md) | Specifies the list of custom properties associated with the tenant. External vendors can choose to set any properties inside following list. Note that the fields set inside the following will not be available for direct filtering. API callers should make sure that no sensitive information such as passwords is sent in these fields. | [optional] 
**LivenessMode** | Pointer to **NullableString** | Specifies the current liveness mode of the tenant. This mode may change based on AZ failures when vendor chooses to failover or failback the tenants to other AZs. | [optional] 
**OwnershipMode** | Pointer to **NullableString** | Specifies the current ownership mode for the tenant. The ownership of the tenant represents the active role for functioning of the tenant. | [optional] 
**ResourceGroupId** | Pointer to **NullableString** | Specifies the Resource Group ID associated with the tenant. | [optional] 

## Methods

### NewIbmTenantMetadataParams

`func NewIbmTenantMetadataParams() *IbmTenantMetadataParams`

NewIbmTenantMetadataParams instantiates a new IbmTenantMetadataParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIbmTenantMetadataParamsWithDefaults

`func NewIbmTenantMetadataParamsWithDefaults() *IbmTenantMetadataParams`

NewIbmTenantMetadataParamsWithDefaults instantiates a new IbmTenantMetadataParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *IbmTenantMetadataParams) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *IbmTenantMetadataParams) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *IbmTenantMetadataParams) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *IbmTenantMetadataParams) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountIdNil

`func (o *IbmTenantMetadataParams) SetAccountIdNil(b bool)`

 SetAccountIdNil sets the value for AccountId to be an explicit nil

### UnsetAccountId
`func (o *IbmTenantMetadataParams) UnsetAccountId()`

UnsetAccountId ensures that no value is present for AccountId, not even an explicit nil
### GetCrn

`func (o *IbmTenantMetadataParams) GetCrn() string`

GetCrn returns the Crn field if non-nil, zero value otherwise.

### GetCrnOk

`func (o *IbmTenantMetadataParams) GetCrnOk() (*string, bool)`

GetCrnOk returns a tuple with the Crn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrn

`func (o *IbmTenantMetadataParams) SetCrn(v string)`

SetCrn sets Crn field to given value.

### HasCrn

`func (o *IbmTenantMetadataParams) HasCrn() bool`

HasCrn returns a boolean if a field has been set.

### SetCrnNil

`func (o *IbmTenantMetadataParams) SetCrnNil(b bool)`

 SetCrnNil sets the value for Crn to be an explicit nil

### UnsetCrn
`func (o *IbmTenantMetadataParams) UnsetCrn()`

UnsetCrn ensures that no value is present for Crn, not even an explicit nil
### GetCustomProperties

`func (o *IbmTenantMetadataParams) GetCustomProperties() []ExternalVendorCustomProperties`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *IbmTenantMetadataParams) GetCustomPropertiesOk() (*[]ExternalVendorCustomProperties, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *IbmTenantMetadataParams) SetCustomProperties(v []ExternalVendorCustomProperties)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *IbmTenantMetadataParams) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### SetCustomPropertiesNil

`func (o *IbmTenantMetadataParams) SetCustomPropertiesNil(b bool)`

 SetCustomPropertiesNil sets the value for CustomProperties to be an explicit nil

### UnsetCustomProperties
`func (o *IbmTenantMetadataParams) UnsetCustomProperties()`

UnsetCustomProperties ensures that no value is present for CustomProperties, not even an explicit nil
### GetLivenessMode

`func (o *IbmTenantMetadataParams) GetLivenessMode() string`

GetLivenessMode returns the LivenessMode field if non-nil, zero value otherwise.

### GetLivenessModeOk

`func (o *IbmTenantMetadataParams) GetLivenessModeOk() (*string, bool)`

GetLivenessModeOk returns a tuple with the LivenessMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivenessMode

`func (o *IbmTenantMetadataParams) SetLivenessMode(v string)`

SetLivenessMode sets LivenessMode field to given value.

### HasLivenessMode

`func (o *IbmTenantMetadataParams) HasLivenessMode() bool`

HasLivenessMode returns a boolean if a field has been set.

### SetLivenessModeNil

`func (o *IbmTenantMetadataParams) SetLivenessModeNil(b bool)`

 SetLivenessModeNil sets the value for LivenessMode to be an explicit nil

### UnsetLivenessMode
`func (o *IbmTenantMetadataParams) UnsetLivenessMode()`

UnsetLivenessMode ensures that no value is present for LivenessMode, not even an explicit nil
### GetOwnershipMode

`func (o *IbmTenantMetadataParams) GetOwnershipMode() string`

GetOwnershipMode returns the OwnershipMode field if non-nil, zero value otherwise.

### GetOwnershipModeOk

`func (o *IbmTenantMetadataParams) GetOwnershipModeOk() (*string, bool)`

GetOwnershipModeOk returns a tuple with the OwnershipMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnershipMode

`func (o *IbmTenantMetadataParams) SetOwnershipMode(v string)`

SetOwnershipMode sets OwnershipMode field to given value.

### HasOwnershipMode

`func (o *IbmTenantMetadataParams) HasOwnershipMode() bool`

HasOwnershipMode returns a boolean if a field has been set.

### SetOwnershipModeNil

`func (o *IbmTenantMetadataParams) SetOwnershipModeNil(b bool)`

 SetOwnershipModeNil sets the value for OwnershipMode to be an explicit nil

### UnsetOwnershipMode
`func (o *IbmTenantMetadataParams) UnsetOwnershipMode()`

UnsetOwnershipMode ensures that no value is present for OwnershipMode, not even an explicit nil
### GetResourceGroupId

`func (o *IbmTenantMetadataParams) GetResourceGroupId() string`

GetResourceGroupId returns the ResourceGroupId field if non-nil, zero value otherwise.

### GetResourceGroupIdOk

`func (o *IbmTenantMetadataParams) GetResourceGroupIdOk() (*string, bool)`

GetResourceGroupIdOk returns a tuple with the ResourceGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroupId

`func (o *IbmTenantMetadataParams) SetResourceGroupId(v string)`

SetResourceGroupId sets ResourceGroupId field to given value.

### HasResourceGroupId

`func (o *IbmTenantMetadataParams) HasResourceGroupId() bool`

HasResourceGroupId returns a boolean if a field has been set.

### SetResourceGroupIdNil

`func (o *IbmTenantMetadataParams) SetResourceGroupIdNil(b bool)`

 SetResourceGroupIdNil sets the value for ResourceGroupId to be an explicit nil

### UnsetResourceGroupId
`func (o *IbmTenantMetadataParams) UnsetResourceGroupId()`

UnsetResourceGroupId ensures that no value is present for ResourceGroupId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


