# AntivirusServiceGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **NullableString** | Specifies the Antivirus Service group name. | 
**AntivirusServices** | [**[]AntivirusService**](AntivirusService.md) | Specifies a list of Antivirus Services for this group. | 
**Enabled** | **NullableBool** | Specifies whether the Antivirus Group is enabled. | 
**Description** | Pointer to **NullableString** | Specifies the description for the Antivirus Service group. | [optional] 
**Id** | Pointer to **NullableInt64** | Specifies the Antivirus Service group id. | [optional] 

## Methods

### NewAntivirusServiceGroup

`func NewAntivirusServiceGroup(name NullableString, antivirusServices []AntivirusService, enabled NullableBool, ) *AntivirusServiceGroup`

NewAntivirusServiceGroup instantiates a new AntivirusServiceGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAntivirusServiceGroupWithDefaults

`func NewAntivirusServiceGroupWithDefaults() *AntivirusServiceGroup`

NewAntivirusServiceGroupWithDefaults instantiates a new AntivirusServiceGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AntivirusServiceGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AntivirusServiceGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AntivirusServiceGroup) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *AntivirusServiceGroup) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AntivirusServiceGroup) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetAntivirusServices

`func (o *AntivirusServiceGroup) GetAntivirusServices() []AntivirusService`

GetAntivirusServices returns the AntivirusServices field if non-nil, zero value otherwise.

### GetAntivirusServicesOk

`func (o *AntivirusServiceGroup) GetAntivirusServicesOk() (*[]AntivirusService, bool)`

GetAntivirusServicesOk returns a tuple with the AntivirusServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntivirusServices

`func (o *AntivirusServiceGroup) SetAntivirusServices(v []AntivirusService)`

SetAntivirusServices sets AntivirusServices field to given value.


### SetAntivirusServicesNil

`func (o *AntivirusServiceGroup) SetAntivirusServicesNil(b bool)`

 SetAntivirusServicesNil sets the value for AntivirusServices to be an explicit nil

### UnsetAntivirusServices
`func (o *AntivirusServiceGroup) UnsetAntivirusServices()`

UnsetAntivirusServices ensures that no value is present for AntivirusServices, not even an explicit nil
### GetEnabled

`func (o *AntivirusServiceGroup) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AntivirusServiceGroup) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AntivirusServiceGroup) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### SetEnabledNil

`func (o *AntivirusServiceGroup) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *AntivirusServiceGroup) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
### GetDescription

`func (o *AntivirusServiceGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AntivirusServiceGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AntivirusServiceGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AntivirusServiceGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AntivirusServiceGroup) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AntivirusServiceGroup) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetId

`func (o *AntivirusServiceGroup) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AntivirusServiceGroup) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AntivirusServiceGroup) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AntivirusServiceGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *AntivirusServiceGroup) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *AntivirusServiceGroup) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


