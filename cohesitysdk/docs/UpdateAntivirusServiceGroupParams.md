# UpdateAntivirusServiceGroupParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AntivirusServices** | Pointer to [**[]AntivirusServiceConfigParams**](AntivirusServiceConfigParams.md) | Specifies the Antivirus services for this provider. | [optional] 
**Description** | Pointer to **NullableString** | Specifies the description of the Antivirus service group. | [optional] 
**Id** | **NullableInt64** | Specifies the Id of the Antivirus service group. | 
**IsEnabled** | Pointer to **NullableBool** | Specifies whether the antivirus service group is enabled or not. | [optional] 
**Name** | **NullableString** | Specifies the name of the Antivirus service group. | 

## Methods

### NewUpdateAntivirusServiceGroupParams

`func NewUpdateAntivirusServiceGroupParams(id NullableInt64, name NullableString, ) *UpdateAntivirusServiceGroupParams`

NewUpdateAntivirusServiceGroupParams instantiates a new UpdateAntivirusServiceGroupParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAntivirusServiceGroupParamsWithDefaults

`func NewUpdateAntivirusServiceGroupParamsWithDefaults() *UpdateAntivirusServiceGroupParams`

NewUpdateAntivirusServiceGroupParamsWithDefaults instantiates a new UpdateAntivirusServiceGroupParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAntivirusServices

`func (o *UpdateAntivirusServiceGroupParams) GetAntivirusServices() []AntivirusServiceConfigParams`

GetAntivirusServices returns the AntivirusServices field if non-nil, zero value otherwise.

### GetAntivirusServicesOk

`func (o *UpdateAntivirusServiceGroupParams) GetAntivirusServicesOk() (*[]AntivirusServiceConfigParams, bool)`

GetAntivirusServicesOk returns a tuple with the AntivirusServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntivirusServices

`func (o *UpdateAntivirusServiceGroupParams) SetAntivirusServices(v []AntivirusServiceConfigParams)`

SetAntivirusServices sets AntivirusServices field to given value.

### HasAntivirusServices

`func (o *UpdateAntivirusServiceGroupParams) HasAntivirusServices() bool`

HasAntivirusServices returns a boolean if a field has been set.

### SetAntivirusServicesNil

`func (o *UpdateAntivirusServiceGroupParams) SetAntivirusServicesNil(b bool)`

 SetAntivirusServicesNil sets the value for AntivirusServices to be an explicit nil

### UnsetAntivirusServices
`func (o *UpdateAntivirusServiceGroupParams) UnsetAntivirusServices()`

UnsetAntivirusServices ensures that no value is present for AntivirusServices, not even an explicit nil
### GetDescription

`func (o *UpdateAntivirusServiceGroupParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateAntivirusServiceGroupParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateAntivirusServiceGroupParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateAntivirusServiceGroupParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateAntivirusServiceGroupParams) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateAntivirusServiceGroupParams) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetId

`func (o *UpdateAntivirusServiceGroupParams) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateAntivirusServiceGroupParams) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateAntivirusServiceGroupParams) SetId(v int64)`

SetId sets Id field to given value.


### SetIdNil

`func (o *UpdateAntivirusServiceGroupParams) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *UpdateAntivirusServiceGroupParams) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIsEnabled

`func (o *UpdateAntivirusServiceGroupParams) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *UpdateAntivirusServiceGroupParams) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *UpdateAntivirusServiceGroupParams) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *UpdateAntivirusServiceGroupParams) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### SetIsEnabledNil

`func (o *UpdateAntivirusServiceGroupParams) SetIsEnabledNil(b bool)`

 SetIsEnabledNil sets the value for IsEnabled to be an explicit nil

### UnsetIsEnabled
`func (o *UpdateAntivirusServiceGroupParams) UnsetIsEnabled()`

UnsetIsEnabled ensures that no value is present for IsEnabled, not even an explicit nil
### GetName

`func (o *UpdateAntivirusServiceGroupParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateAntivirusServiceGroupParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateAntivirusServiceGroupParams) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *UpdateAntivirusServiceGroupParams) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateAntivirusServiceGroupParams) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


