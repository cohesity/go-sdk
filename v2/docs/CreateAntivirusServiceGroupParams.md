# CreateAntivirusServiceGroupParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AntivirusServices** | [**[]AntivirusService**](AntivirusService.md) | Specifies a list of Antivirus Services for this group. | 
**Description** | Pointer to **NullableString** | Specifies the description for the Antivirus Service group. | [optional] 
**Enabled** | Pointer to **NullableBool** | This field is currently deprecated. Specifies whether the Antivirus Group is enabled. | [optional] 
**Name** | **NullableString** | Specifies the Antivirus Service group name. | 
**State** | Pointer to **NullableString** | Specifies the state[Enable, Disable] of the group. | [optional] 

## Methods

### NewCreateAntivirusServiceGroupParams

`func NewCreateAntivirusServiceGroupParams(antivirusServices []AntivirusService, name NullableString, ) *CreateAntivirusServiceGroupParams`

NewCreateAntivirusServiceGroupParams instantiates a new CreateAntivirusServiceGroupParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAntivirusServiceGroupParamsWithDefaults

`func NewCreateAntivirusServiceGroupParamsWithDefaults() *CreateAntivirusServiceGroupParams`

NewCreateAntivirusServiceGroupParamsWithDefaults instantiates a new CreateAntivirusServiceGroupParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAntivirusServices

`func (o *CreateAntivirusServiceGroupParams) GetAntivirusServices() []AntivirusService`

GetAntivirusServices returns the AntivirusServices field if non-nil, zero value otherwise.

### GetAntivirusServicesOk

`func (o *CreateAntivirusServiceGroupParams) GetAntivirusServicesOk() (*[]AntivirusService, bool)`

GetAntivirusServicesOk returns a tuple with the AntivirusServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntivirusServices

`func (o *CreateAntivirusServiceGroupParams) SetAntivirusServices(v []AntivirusService)`

SetAntivirusServices sets AntivirusServices field to given value.


### SetAntivirusServicesNil

`func (o *CreateAntivirusServiceGroupParams) SetAntivirusServicesNil(b bool)`

 SetAntivirusServicesNil sets the value for AntivirusServices to be an explicit nil

### UnsetAntivirusServices
`func (o *CreateAntivirusServiceGroupParams) UnsetAntivirusServices()`

UnsetAntivirusServices ensures that no value is present for AntivirusServices, not even an explicit nil
### GetDescription

`func (o *CreateAntivirusServiceGroupParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateAntivirusServiceGroupParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateAntivirusServiceGroupParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateAntivirusServiceGroupParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateAntivirusServiceGroupParams) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateAntivirusServiceGroupParams) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEnabled

`func (o *CreateAntivirusServiceGroupParams) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateAntivirusServiceGroupParams) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateAntivirusServiceGroupParams) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CreateAntivirusServiceGroupParams) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *CreateAntivirusServiceGroupParams) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *CreateAntivirusServiceGroupParams) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
### GetName

`func (o *CreateAntivirusServiceGroupParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAntivirusServiceGroupParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateAntivirusServiceGroupParams) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *CreateAntivirusServiceGroupParams) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateAntivirusServiceGroupParams) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetState

`func (o *CreateAntivirusServiceGroupParams) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CreateAntivirusServiceGroupParams) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CreateAntivirusServiceGroupParams) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CreateAntivirusServiceGroupParams) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *CreateAntivirusServiceGroupParams) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *CreateAntivirusServiceGroupParams) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


