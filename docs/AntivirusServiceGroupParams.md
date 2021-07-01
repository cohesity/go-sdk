# AntivirusServiceGroupParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AntivirusServices** | Pointer to [**[]AntivirusServiceConfigParams**](AntivirusServiceConfigParams.md) | Specifies the Antivirus services for this provider. | [optional] 
**Description** | Pointer to **NullableString** | Specifies the description of the Antivirus service group. | [optional] 
**Name** | **NullableString** | Specifies the name of the Antivirus service group. | 

## Methods

### NewAntivirusServiceGroupParams

`func NewAntivirusServiceGroupParams(name NullableString, ) *AntivirusServiceGroupParams`

NewAntivirusServiceGroupParams instantiates a new AntivirusServiceGroupParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAntivirusServiceGroupParamsWithDefaults

`func NewAntivirusServiceGroupParamsWithDefaults() *AntivirusServiceGroupParams`

NewAntivirusServiceGroupParamsWithDefaults instantiates a new AntivirusServiceGroupParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAntivirusServices

`func (o *AntivirusServiceGroupParams) GetAntivirusServices() []AntivirusServiceConfigParams`

GetAntivirusServices returns the AntivirusServices field if non-nil, zero value otherwise.

### GetAntivirusServicesOk

`func (o *AntivirusServiceGroupParams) GetAntivirusServicesOk() (*[]AntivirusServiceConfigParams, bool)`

GetAntivirusServicesOk returns a tuple with the AntivirusServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntivirusServices

`func (o *AntivirusServiceGroupParams) SetAntivirusServices(v []AntivirusServiceConfigParams)`

SetAntivirusServices sets AntivirusServices field to given value.

### HasAntivirusServices

`func (o *AntivirusServiceGroupParams) HasAntivirusServices() bool`

HasAntivirusServices returns a boolean if a field has been set.

### SetAntivirusServicesNil

`func (o *AntivirusServiceGroupParams) SetAntivirusServicesNil(b bool)`

 SetAntivirusServicesNil sets the value for AntivirusServices to be an explicit nil

### UnsetAntivirusServices
`func (o *AntivirusServiceGroupParams) UnsetAntivirusServices()`

UnsetAntivirusServices ensures that no value is present for AntivirusServices, not even an explicit nil
### GetDescription

`func (o *AntivirusServiceGroupParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AntivirusServiceGroupParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AntivirusServiceGroupParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AntivirusServiceGroupParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AntivirusServiceGroupParams) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AntivirusServiceGroupParams) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetName

`func (o *AntivirusServiceGroupParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AntivirusServiceGroupParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AntivirusServiceGroupParams) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *AntivirusServiceGroupParams) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AntivirusServiceGroupParams) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


