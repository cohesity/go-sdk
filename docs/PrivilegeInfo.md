# PrivilegeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrivilegeId** | Pointer to **NullableString** | Specifies unique id for a privilege. This number must be unique when creating a new privilege. Type for unique privilege Id values. All below enum values specify a value for all uniquely defined privileges in Cohesity. | [optional] 
**Category** | Pointer to **NullableString** | Specifies a category for the privilege such as &#39;Access Management&#39;. | [optional] 
**Description** | Pointer to **NullableString** | Specifies a description defining what the privilege provides. | [optional] 
**IsAvailableOnHelios** | Pointer to **NullableBool** | Specifies that this privilege is available for Helios operations. | [optional] 
**IsCustomRoleDefault** | Pointer to **NullableBool** | Specifies if this privilege is automatically assigned to custom roles. | [optional] 
**IsSpecial** | Pointer to **NullableBool** | Specifies if this privilege is automatically assigned to the default System Admin user called &#39;admin&#39;. If true, the privilege is NOT assigned to the default System Admin user called &#39;admin&#39;. By default, privileges are automatically assigned to the default System Admin user called &#39;admin&#39;. | [optional] 
**IsViewOnly** | Pointer to **NullableBool** | Specifies if privilege is view-only privilege that cannot make changes. | [optional] 
**Label** | Pointer to **NullableString** | Specifies the label for the privilege as displayed on the Cohesity Dashboard such as &#39;Access Management View&#39;. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the Cluster name for the privilege such as PRINCIPAL_VIEW. | [optional] 

## Methods

### NewPrivilegeInfo

`func NewPrivilegeInfo() *PrivilegeInfo`

NewPrivilegeInfo instantiates a new PrivilegeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivilegeInfoWithDefaults

`func NewPrivilegeInfoWithDefaults() *PrivilegeInfo`

NewPrivilegeInfoWithDefaults instantiates a new PrivilegeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrivilegeId

`func (o *PrivilegeInfo) GetPrivilegeId() string`

GetPrivilegeId returns the PrivilegeId field if non-nil, zero value otherwise.

### GetPrivilegeIdOk

`func (o *PrivilegeInfo) GetPrivilegeIdOk() (*string, bool)`

GetPrivilegeIdOk returns a tuple with the PrivilegeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilegeId

`func (o *PrivilegeInfo) SetPrivilegeId(v string)`

SetPrivilegeId sets PrivilegeId field to given value.

### HasPrivilegeId

`func (o *PrivilegeInfo) HasPrivilegeId() bool`

HasPrivilegeId returns a boolean if a field has been set.

### SetPrivilegeIdNil

`func (o *PrivilegeInfo) SetPrivilegeIdNil(b bool)`

 SetPrivilegeIdNil sets the value for PrivilegeId to be an explicit nil

### UnsetPrivilegeId
`func (o *PrivilegeInfo) UnsetPrivilegeId()`

UnsetPrivilegeId ensures that no value is present for PrivilegeId, not even an explicit nil
### GetCategory

`func (o *PrivilegeInfo) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *PrivilegeInfo) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *PrivilegeInfo) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *PrivilegeInfo) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *PrivilegeInfo) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *PrivilegeInfo) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetDescription

`func (o *PrivilegeInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PrivilegeInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PrivilegeInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PrivilegeInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PrivilegeInfo) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PrivilegeInfo) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIsAvailableOnHelios

`func (o *PrivilegeInfo) GetIsAvailableOnHelios() bool`

GetIsAvailableOnHelios returns the IsAvailableOnHelios field if non-nil, zero value otherwise.

### GetIsAvailableOnHeliosOk

`func (o *PrivilegeInfo) GetIsAvailableOnHeliosOk() (*bool, bool)`

GetIsAvailableOnHeliosOk returns a tuple with the IsAvailableOnHelios field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAvailableOnHelios

`func (o *PrivilegeInfo) SetIsAvailableOnHelios(v bool)`

SetIsAvailableOnHelios sets IsAvailableOnHelios field to given value.

### HasIsAvailableOnHelios

`func (o *PrivilegeInfo) HasIsAvailableOnHelios() bool`

HasIsAvailableOnHelios returns a boolean if a field has been set.

### SetIsAvailableOnHeliosNil

`func (o *PrivilegeInfo) SetIsAvailableOnHeliosNil(b bool)`

 SetIsAvailableOnHeliosNil sets the value for IsAvailableOnHelios to be an explicit nil

### UnsetIsAvailableOnHelios
`func (o *PrivilegeInfo) UnsetIsAvailableOnHelios()`

UnsetIsAvailableOnHelios ensures that no value is present for IsAvailableOnHelios, not even an explicit nil
### GetIsCustomRoleDefault

`func (o *PrivilegeInfo) GetIsCustomRoleDefault() bool`

GetIsCustomRoleDefault returns the IsCustomRoleDefault field if non-nil, zero value otherwise.

### GetIsCustomRoleDefaultOk

`func (o *PrivilegeInfo) GetIsCustomRoleDefaultOk() (*bool, bool)`

GetIsCustomRoleDefaultOk returns a tuple with the IsCustomRoleDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCustomRoleDefault

`func (o *PrivilegeInfo) SetIsCustomRoleDefault(v bool)`

SetIsCustomRoleDefault sets IsCustomRoleDefault field to given value.

### HasIsCustomRoleDefault

`func (o *PrivilegeInfo) HasIsCustomRoleDefault() bool`

HasIsCustomRoleDefault returns a boolean if a field has been set.

### SetIsCustomRoleDefaultNil

`func (o *PrivilegeInfo) SetIsCustomRoleDefaultNil(b bool)`

 SetIsCustomRoleDefaultNil sets the value for IsCustomRoleDefault to be an explicit nil

### UnsetIsCustomRoleDefault
`func (o *PrivilegeInfo) UnsetIsCustomRoleDefault()`

UnsetIsCustomRoleDefault ensures that no value is present for IsCustomRoleDefault, not even an explicit nil
### GetIsSpecial

`func (o *PrivilegeInfo) GetIsSpecial() bool`

GetIsSpecial returns the IsSpecial field if non-nil, zero value otherwise.

### GetIsSpecialOk

`func (o *PrivilegeInfo) GetIsSpecialOk() (*bool, bool)`

GetIsSpecialOk returns a tuple with the IsSpecial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSpecial

`func (o *PrivilegeInfo) SetIsSpecial(v bool)`

SetIsSpecial sets IsSpecial field to given value.

### HasIsSpecial

`func (o *PrivilegeInfo) HasIsSpecial() bool`

HasIsSpecial returns a boolean if a field has been set.

### SetIsSpecialNil

`func (o *PrivilegeInfo) SetIsSpecialNil(b bool)`

 SetIsSpecialNil sets the value for IsSpecial to be an explicit nil

### UnsetIsSpecial
`func (o *PrivilegeInfo) UnsetIsSpecial()`

UnsetIsSpecial ensures that no value is present for IsSpecial, not even an explicit nil
### GetIsViewOnly

`func (o *PrivilegeInfo) GetIsViewOnly() bool`

GetIsViewOnly returns the IsViewOnly field if non-nil, zero value otherwise.

### GetIsViewOnlyOk

`func (o *PrivilegeInfo) GetIsViewOnlyOk() (*bool, bool)`

GetIsViewOnlyOk returns a tuple with the IsViewOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsViewOnly

`func (o *PrivilegeInfo) SetIsViewOnly(v bool)`

SetIsViewOnly sets IsViewOnly field to given value.

### HasIsViewOnly

`func (o *PrivilegeInfo) HasIsViewOnly() bool`

HasIsViewOnly returns a boolean if a field has been set.

### SetIsViewOnlyNil

`func (o *PrivilegeInfo) SetIsViewOnlyNil(b bool)`

 SetIsViewOnlyNil sets the value for IsViewOnly to be an explicit nil

### UnsetIsViewOnly
`func (o *PrivilegeInfo) UnsetIsViewOnly()`

UnsetIsViewOnly ensures that no value is present for IsViewOnly, not even an explicit nil
### GetLabel

`func (o *PrivilegeInfo) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PrivilegeInfo) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PrivilegeInfo) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PrivilegeInfo) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *PrivilegeInfo) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *PrivilegeInfo) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil
### GetName

`func (o *PrivilegeInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PrivilegeInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PrivilegeInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PrivilegeInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PrivilegeInfo) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PrivilegeInfo) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


