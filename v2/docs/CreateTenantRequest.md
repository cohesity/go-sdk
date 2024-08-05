# CreateTenantRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** | Description about the tenant | [optional] 
**IsManagedOnHelios** | Pointer to **NullableBool** | Flag to indicate if tenant is managed on helios | [optional] 
**Name** | **NullableString** | Name of the Tenant. | 
**TenantIdSuffix** | **NullableString** | This suffix is used by cohesity to generate the actual tenant Id by appending the parent&#39;s tenant id to it. | 
**Network** | Pointer to [**TenantNetwork**](TenantNetwork.md) |  | [optional] 

## Methods

### NewCreateTenantRequest

`func NewCreateTenantRequest(name NullableString, tenantIdSuffix NullableString, ) *CreateTenantRequest`

NewCreateTenantRequest instantiates a new CreateTenantRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTenantRequestWithDefaults

`func NewCreateTenantRequestWithDefaults() *CreateTenantRequest`

NewCreateTenantRequestWithDefaults instantiates a new CreateTenantRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateTenantRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateTenantRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateTenantRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateTenantRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateTenantRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateTenantRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIsManagedOnHelios

`func (o *CreateTenantRequest) GetIsManagedOnHelios() bool`

GetIsManagedOnHelios returns the IsManagedOnHelios field if non-nil, zero value otherwise.

### GetIsManagedOnHeliosOk

`func (o *CreateTenantRequest) GetIsManagedOnHeliosOk() (*bool, bool)`

GetIsManagedOnHeliosOk returns a tuple with the IsManagedOnHelios field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsManagedOnHelios

`func (o *CreateTenantRequest) SetIsManagedOnHelios(v bool)`

SetIsManagedOnHelios sets IsManagedOnHelios field to given value.

### HasIsManagedOnHelios

`func (o *CreateTenantRequest) HasIsManagedOnHelios() bool`

HasIsManagedOnHelios returns a boolean if a field has been set.

### SetIsManagedOnHeliosNil

`func (o *CreateTenantRequest) SetIsManagedOnHeliosNil(b bool)`

 SetIsManagedOnHeliosNil sets the value for IsManagedOnHelios to be an explicit nil

### UnsetIsManagedOnHelios
`func (o *CreateTenantRequest) UnsetIsManagedOnHelios()`

UnsetIsManagedOnHelios ensures that no value is present for IsManagedOnHelios, not even an explicit nil
### GetName

`func (o *CreateTenantRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateTenantRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateTenantRequest) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *CreateTenantRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateTenantRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTenantIdSuffix

`func (o *CreateTenantRequest) GetTenantIdSuffix() string`

GetTenantIdSuffix returns the TenantIdSuffix field if non-nil, zero value otherwise.

### GetTenantIdSuffixOk

`func (o *CreateTenantRequest) GetTenantIdSuffixOk() (*string, bool)`

GetTenantIdSuffixOk returns a tuple with the TenantIdSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantIdSuffix

`func (o *CreateTenantRequest) SetTenantIdSuffix(v string)`

SetTenantIdSuffix sets TenantIdSuffix field to given value.


### SetTenantIdSuffixNil

`func (o *CreateTenantRequest) SetTenantIdSuffixNil(b bool)`

 SetTenantIdSuffixNil sets the value for TenantIdSuffix to be an explicit nil

### UnsetTenantIdSuffix
`func (o *CreateTenantRequest) UnsetTenantIdSuffix()`

UnsetTenantIdSuffix ensures that no value is present for TenantIdSuffix, not even an explicit nil
### GetNetwork

`func (o *CreateTenantRequest) GetNetwork() TenantNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *CreateTenantRequest) GetNetworkOk() (*TenantNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *CreateTenantRequest) SetNetwork(v TenantNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *CreateTenantRequest) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


