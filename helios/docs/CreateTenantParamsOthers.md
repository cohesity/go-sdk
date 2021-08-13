# CreateTenantParamsOthers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **NullableString** | Name of the Tenant. | 
**TenantIdSuffix** | **NullableString** | This suffix is used by cohesity to generate the actual tenant Id by appending the parent&#39;s tenant id to it. | 
**Description** | Pointer to **NullableString** | Description about the tenant | [optional] 
**Network** | Pointer to [**TenantNetwork**](TenantNetwork.md) |  | [optional] 

## Methods

### NewCreateTenantParamsOthers

`func NewCreateTenantParamsOthers(name NullableString, tenantIdSuffix NullableString, ) *CreateTenantParamsOthers`

NewCreateTenantParamsOthers instantiates a new CreateTenantParamsOthers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTenantParamsOthersWithDefaults

`func NewCreateTenantParamsOthersWithDefaults() *CreateTenantParamsOthers`

NewCreateTenantParamsOthersWithDefaults instantiates a new CreateTenantParamsOthers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateTenantParamsOthers) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateTenantParamsOthers) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateTenantParamsOthers) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *CreateTenantParamsOthers) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateTenantParamsOthers) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTenantIdSuffix

`func (o *CreateTenantParamsOthers) GetTenantIdSuffix() string`

GetTenantIdSuffix returns the TenantIdSuffix field if non-nil, zero value otherwise.

### GetTenantIdSuffixOk

`func (o *CreateTenantParamsOthers) GetTenantIdSuffixOk() (*string, bool)`

GetTenantIdSuffixOk returns a tuple with the TenantIdSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantIdSuffix

`func (o *CreateTenantParamsOthers) SetTenantIdSuffix(v string)`

SetTenantIdSuffix sets TenantIdSuffix field to given value.


### SetTenantIdSuffixNil

`func (o *CreateTenantParamsOthers) SetTenantIdSuffixNil(b bool)`

 SetTenantIdSuffixNil sets the value for TenantIdSuffix to be an explicit nil

### UnsetTenantIdSuffix
`func (o *CreateTenantParamsOthers) UnsetTenantIdSuffix()`

UnsetTenantIdSuffix ensures that no value is present for TenantIdSuffix, not even an explicit nil
### GetDescription

`func (o *CreateTenantParamsOthers) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateTenantParamsOthers) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateTenantParamsOthers) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateTenantParamsOthers) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateTenantParamsOthers) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateTenantParamsOthers) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetNetwork

`func (o *CreateTenantParamsOthers) GetNetwork() TenantNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *CreateTenantParamsOthers) GetNetworkOk() (*TenantNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *CreateTenantParamsOthers) SetNetwork(v TenantNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *CreateTenantParamsOthers) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


