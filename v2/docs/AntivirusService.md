# AntivirusService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** | Specifies the description for the Antivirus Service. | [optional] 
**IcapUri** | **NullableString** | Specifies the ICAP Uri for the Antivirus Service. | 
**Tag** | Pointer to **NullableString** | Specifies the tag of the Antivirus Service. | [optional] [readonly] 
**TagId** | Pointer to **NullableInt64** | Specifies the tag id of the Antivirus Service. | [optional] [readonly] 

## Methods

### NewAntivirusService

`func NewAntivirusService(icapUri NullableString, ) *AntivirusService`

NewAntivirusService instantiates a new AntivirusService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAntivirusServiceWithDefaults

`func NewAntivirusServiceWithDefaults() *AntivirusService`

NewAntivirusServiceWithDefaults instantiates a new AntivirusService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *AntivirusService) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AntivirusService) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AntivirusService) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AntivirusService) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AntivirusService) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AntivirusService) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIcapUri

`func (o *AntivirusService) GetIcapUri() string`

GetIcapUri returns the IcapUri field if non-nil, zero value otherwise.

### GetIcapUriOk

`func (o *AntivirusService) GetIcapUriOk() (*string, bool)`

GetIcapUriOk returns a tuple with the IcapUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcapUri

`func (o *AntivirusService) SetIcapUri(v string)`

SetIcapUri sets IcapUri field to given value.


### SetIcapUriNil

`func (o *AntivirusService) SetIcapUriNil(b bool)`

 SetIcapUriNil sets the value for IcapUri to be an explicit nil

### UnsetIcapUri
`func (o *AntivirusService) UnsetIcapUri()`

UnsetIcapUri ensures that no value is present for IcapUri, not even an explicit nil
### GetTag

`func (o *AntivirusService) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *AntivirusService) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *AntivirusService) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *AntivirusService) HasTag() bool`

HasTag returns a boolean if a field has been set.

### SetTagNil

`func (o *AntivirusService) SetTagNil(b bool)`

 SetTagNil sets the value for Tag to be an explicit nil

### UnsetTag
`func (o *AntivirusService) UnsetTag()`

UnsetTag ensures that no value is present for Tag, not even an explicit nil
### GetTagId

`func (o *AntivirusService) GetTagId() int64`

GetTagId returns the TagId field if non-nil, zero value otherwise.

### GetTagIdOk

`func (o *AntivirusService) GetTagIdOk() (*int64, bool)`

GetTagIdOk returns a tuple with the TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagId

`func (o *AntivirusService) SetTagId(v int64)`

SetTagId sets TagId field to given value.

### HasTagId

`func (o *AntivirusService) HasTagId() bool`

HasTagId returns a boolean if a field has been set.

### SetTagIdNil

`func (o *AntivirusService) SetTagIdNil(b bool)`

 SetTagIdNil sets the value for TagId to be an explicit nil

### UnsetTagId
`func (o *AntivirusService) UnsetTagId()`

UnsetTagId ensures that no value is present for TagId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


