# Privilege

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **NullableString** | Specifies the Privilege category. | [optional] 
**Description** | Pointer to **NullableString** | Specifies the description message for the Privilege. | [optional] 
**Id** | Pointer to **NullableInt32** | Specifies the Privilege id. | [optional] 
**IsAvailableOnHelios** | Pointer to **NullableBool** | Specifies whether the Privilege is available for Helios operations. | [optional] 
**IsCustomRoleDefault** | Pointer to **NullableBool** | Specifies whether the Privilege is auto assigned to custom Roles. | [optional] 
**IsSpecial** | Pointer to **NullableBool** | Specifies whether the Privilege is a special privilege. Special Privileges are not assigned to builtin &#39;Admin&#39; Role. | [optional] 
**IsViewOnly** | Pointer to **NullableBool** | Specifies whether the Privilege is a read-only privilege. Read-only Previlege only grants read access to a Role. | [optional] 
**Label** | Pointer to **NullableString** | Specifies the Privilege label. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the Privilege name. | [optional] 

## Methods

### NewPrivilege

`func NewPrivilege() *Privilege`

NewPrivilege instantiates a new Privilege object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivilegeWithDefaults

`func NewPrivilegeWithDefaults() *Privilege`

NewPrivilegeWithDefaults instantiates a new Privilege object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *Privilege) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *Privilege) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *Privilege) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *Privilege) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *Privilege) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *Privilege) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetDescription

`func (o *Privilege) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Privilege) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Privilege) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Privilege) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Privilege) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Privilege) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetId

`func (o *Privilege) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Privilege) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Privilege) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Privilege) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *Privilege) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *Privilege) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIsAvailableOnHelios

`func (o *Privilege) GetIsAvailableOnHelios() bool`

GetIsAvailableOnHelios returns the IsAvailableOnHelios field if non-nil, zero value otherwise.

### GetIsAvailableOnHeliosOk

`func (o *Privilege) GetIsAvailableOnHeliosOk() (*bool, bool)`

GetIsAvailableOnHeliosOk returns a tuple with the IsAvailableOnHelios field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAvailableOnHelios

`func (o *Privilege) SetIsAvailableOnHelios(v bool)`

SetIsAvailableOnHelios sets IsAvailableOnHelios field to given value.

### HasIsAvailableOnHelios

`func (o *Privilege) HasIsAvailableOnHelios() bool`

HasIsAvailableOnHelios returns a boolean if a field has been set.

### SetIsAvailableOnHeliosNil

`func (o *Privilege) SetIsAvailableOnHeliosNil(b bool)`

 SetIsAvailableOnHeliosNil sets the value for IsAvailableOnHelios to be an explicit nil

### UnsetIsAvailableOnHelios
`func (o *Privilege) UnsetIsAvailableOnHelios()`

UnsetIsAvailableOnHelios ensures that no value is present for IsAvailableOnHelios, not even an explicit nil
### GetIsCustomRoleDefault

`func (o *Privilege) GetIsCustomRoleDefault() bool`

GetIsCustomRoleDefault returns the IsCustomRoleDefault field if non-nil, zero value otherwise.

### GetIsCustomRoleDefaultOk

`func (o *Privilege) GetIsCustomRoleDefaultOk() (*bool, bool)`

GetIsCustomRoleDefaultOk returns a tuple with the IsCustomRoleDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCustomRoleDefault

`func (o *Privilege) SetIsCustomRoleDefault(v bool)`

SetIsCustomRoleDefault sets IsCustomRoleDefault field to given value.

### HasIsCustomRoleDefault

`func (o *Privilege) HasIsCustomRoleDefault() bool`

HasIsCustomRoleDefault returns a boolean if a field has been set.

### SetIsCustomRoleDefaultNil

`func (o *Privilege) SetIsCustomRoleDefaultNil(b bool)`

 SetIsCustomRoleDefaultNil sets the value for IsCustomRoleDefault to be an explicit nil

### UnsetIsCustomRoleDefault
`func (o *Privilege) UnsetIsCustomRoleDefault()`

UnsetIsCustomRoleDefault ensures that no value is present for IsCustomRoleDefault, not even an explicit nil
### GetIsSpecial

`func (o *Privilege) GetIsSpecial() bool`

GetIsSpecial returns the IsSpecial field if non-nil, zero value otherwise.

### GetIsSpecialOk

`func (o *Privilege) GetIsSpecialOk() (*bool, bool)`

GetIsSpecialOk returns a tuple with the IsSpecial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSpecial

`func (o *Privilege) SetIsSpecial(v bool)`

SetIsSpecial sets IsSpecial field to given value.

### HasIsSpecial

`func (o *Privilege) HasIsSpecial() bool`

HasIsSpecial returns a boolean if a field has been set.

### SetIsSpecialNil

`func (o *Privilege) SetIsSpecialNil(b bool)`

 SetIsSpecialNil sets the value for IsSpecial to be an explicit nil

### UnsetIsSpecial
`func (o *Privilege) UnsetIsSpecial()`

UnsetIsSpecial ensures that no value is present for IsSpecial, not even an explicit nil
### GetIsViewOnly

`func (o *Privilege) GetIsViewOnly() bool`

GetIsViewOnly returns the IsViewOnly field if non-nil, zero value otherwise.

### GetIsViewOnlyOk

`func (o *Privilege) GetIsViewOnlyOk() (*bool, bool)`

GetIsViewOnlyOk returns a tuple with the IsViewOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsViewOnly

`func (o *Privilege) SetIsViewOnly(v bool)`

SetIsViewOnly sets IsViewOnly field to given value.

### HasIsViewOnly

`func (o *Privilege) HasIsViewOnly() bool`

HasIsViewOnly returns a boolean if a field has been set.

### SetIsViewOnlyNil

`func (o *Privilege) SetIsViewOnlyNil(b bool)`

 SetIsViewOnlyNil sets the value for IsViewOnly to be an explicit nil

### UnsetIsViewOnly
`func (o *Privilege) UnsetIsViewOnly()`

UnsetIsViewOnly ensures that no value is present for IsViewOnly, not even an explicit nil
### GetLabel

`func (o *Privilege) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Privilege) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Privilege) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *Privilege) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *Privilege) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *Privilege) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil
### GetName

`func (o *Privilege) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Privilege) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Privilege) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Privilege) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Privilege) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Privilege) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


