# AntivirusServiceConfigParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** | Specifies the description of the Antivirus service. This could be any additional information admin might associate with the Antivirus service. | [optional] 
**IcapUri** | **NullableString** | Specifies the ICAP uri for this Antivirus service. It is of the form icap://&lt;ip-address&gt;[:&lt;port&gt;]/&lt;service&gt; | 

## Methods

### NewAntivirusServiceConfigParams

`func NewAntivirusServiceConfigParams(icapUri NullableString, ) *AntivirusServiceConfigParams`

NewAntivirusServiceConfigParams instantiates a new AntivirusServiceConfigParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAntivirusServiceConfigParamsWithDefaults

`func NewAntivirusServiceConfigParamsWithDefaults() *AntivirusServiceConfigParams`

NewAntivirusServiceConfigParamsWithDefaults instantiates a new AntivirusServiceConfigParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *AntivirusServiceConfigParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AntivirusServiceConfigParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AntivirusServiceConfigParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AntivirusServiceConfigParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AntivirusServiceConfigParams) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AntivirusServiceConfigParams) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIcapUri

`func (o *AntivirusServiceConfigParams) GetIcapUri() string`

GetIcapUri returns the IcapUri field if non-nil, zero value otherwise.

### GetIcapUriOk

`func (o *AntivirusServiceConfigParams) GetIcapUriOk() (*string, bool)`

GetIcapUriOk returns a tuple with the IcapUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcapUri

`func (o *AntivirusServiceConfigParams) SetIcapUri(v string)`

SetIcapUri sets IcapUri field to given value.


### SetIcapUriNil

`func (o *AntivirusServiceConfigParams) SetIcapUriNil(b bool)`

 SetIcapUriNil sets the value for IcapUri to be an explicit nil

### UnsetIcapUri
`func (o *AntivirusServiceConfigParams) UnsetIcapUri()`

UnsetIcapUri ensures that no value is present for IcapUri, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


