# UpdateTenantBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Name of the Tenant. | [optional] 
**Description** | Pointer to **NullableString** | Description about the tenant | [optional] 
**Network** | Pointer to [**TenantNetwork**](TenantNetwork.md) |  | [optional] 

## Methods

### NewUpdateTenantBody

`func NewUpdateTenantBody() *UpdateTenantBody`

NewUpdateTenantBody instantiates a new UpdateTenantBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTenantBodyWithDefaults

`func NewUpdateTenantBodyWithDefaults() *UpdateTenantBody`

NewUpdateTenantBodyWithDefaults instantiates a new UpdateTenantBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateTenantBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateTenantBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateTenantBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateTenantBody) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UpdateTenantBody) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateTenantBody) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *UpdateTenantBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateTenantBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateTenantBody) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateTenantBody) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateTenantBody) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateTenantBody) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetNetwork

`func (o *UpdateTenantBody) GetNetwork() TenantNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *UpdateTenantBody) GetNetworkOk() (*TenantNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *UpdateTenantBody) SetNetwork(v TenantNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *UpdateTenantBody) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


