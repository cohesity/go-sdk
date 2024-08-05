# CreateTenantParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** | Description about the tenant | [optional] 
**IsManagedOnHelios** | Pointer to **NullableBool** | Flag to indicate if tenant is managed on helios | [optional] 
**Name** | **NullableString** | Name of the Tenant. | 
**TenantIdSuffix** | **NullableString** | This suffix is used by cohesity to generate the actual tenant Id by appending the parent&#39;s tenant id to it. | 

## Methods

### NewCreateTenantParams

`func NewCreateTenantParams(name NullableString, tenantIdSuffix NullableString, ) *CreateTenantParams`

NewCreateTenantParams instantiates a new CreateTenantParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTenantParamsWithDefaults

`func NewCreateTenantParamsWithDefaults() *CreateTenantParams`

NewCreateTenantParamsWithDefaults instantiates a new CreateTenantParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateTenantParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateTenantParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateTenantParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateTenantParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateTenantParams) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateTenantParams) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIsManagedOnHelios

`func (o *CreateTenantParams) GetIsManagedOnHelios() bool`

GetIsManagedOnHelios returns the IsManagedOnHelios field if non-nil, zero value otherwise.

### GetIsManagedOnHeliosOk

`func (o *CreateTenantParams) GetIsManagedOnHeliosOk() (*bool, bool)`

GetIsManagedOnHeliosOk returns a tuple with the IsManagedOnHelios field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsManagedOnHelios

`func (o *CreateTenantParams) SetIsManagedOnHelios(v bool)`

SetIsManagedOnHelios sets IsManagedOnHelios field to given value.

### HasIsManagedOnHelios

`func (o *CreateTenantParams) HasIsManagedOnHelios() bool`

HasIsManagedOnHelios returns a boolean if a field has been set.

### SetIsManagedOnHeliosNil

`func (o *CreateTenantParams) SetIsManagedOnHeliosNil(b bool)`

 SetIsManagedOnHeliosNil sets the value for IsManagedOnHelios to be an explicit nil

### UnsetIsManagedOnHelios
`func (o *CreateTenantParams) UnsetIsManagedOnHelios()`

UnsetIsManagedOnHelios ensures that no value is present for IsManagedOnHelios, not even an explicit nil
### GetName

`func (o *CreateTenantParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateTenantParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateTenantParams) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *CreateTenantParams) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateTenantParams) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTenantIdSuffix

`func (o *CreateTenantParams) GetTenantIdSuffix() string`

GetTenantIdSuffix returns the TenantIdSuffix field if non-nil, zero value otherwise.

### GetTenantIdSuffixOk

`func (o *CreateTenantParams) GetTenantIdSuffixOk() (*string, bool)`

GetTenantIdSuffixOk returns a tuple with the TenantIdSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantIdSuffix

`func (o *CreateTenantParams) SetTenantIdSuffix(v string)`

SetTenantIdSuffix sets TenantIdSuffix field to given value.


### SetTenantIdSuffixNil

`func (o *CreateTenantParams) SetTenantIdSuffixNil(b bool)`

 SetTenantIdSuffixNil sets the value for TenantIdSuffix to be an explicit nil

### UnsetTenantIdSuffix
`func (o *CreateTenantParams) UnsetTenantIdSuffix()`

UnsetTenantIdSuffix ensures that no value is present for TenantIdSuffix, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


