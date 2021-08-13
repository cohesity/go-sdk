# ConfirmTenantParamsBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **NullableString** | Name of the Tenant. | 
**Description** | Pointer to **NullableString** | Description about the Tenant. | [optional] 

## Methods

### NewConfirmTenantParamsBody

`func NewConfirmTenantParamsBody(name NullableString, ) *ConfirmTenantParamsBody`

NewConfirmTenantParamsBody instantiates a new ConfirmTenantParamsBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfirmTenantParamsBodyWithDefaults

`func NewConfirmTenantParamsBodyWithDefaults() *ConfirmTenantParamsBody`

NewConfirmTenantParamsBodyWithDefaults instantiates a new ConfirmTenantParamsBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ConfirmTenantParamsBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConfirmTenantParamsBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConfirmTenantParamsBody) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *ConfirmTenantParamsBody) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ConfirmTenantParamsBody) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *ConfirmTenantParamsBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConfirmTenantParamsBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConfirmTenantParamsBody) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConfirmTenantParamsBody) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ConfirmTenantParamsBody) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ConfirmTenantParamsBody) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


