# BannerUpdateParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **NullableString** | Specifies the content of the banner. | [optional] 
**Description** | Pointer to **NullableString** | description field is deprecated. Specifies the description of this banner. | [optional] 

## Methods

### NewBannerUpdateParameters

`func NewBannerUpdateParameters() *BannerUpdateParameters`

NewBannerUpdateParameters instantiates a new BannerUpdateParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBannerUpdateParametersWithDefaults

`func NewBannerUpdateParametersWithDefaults() *BannerUpdateParameters`

NewBannerUpdateParametersWithDefaults instantiates a new BannerUpdateParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *BannerUpdateParameters) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *BannerUpdateParameters) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *BannerUpdateParameters) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *BannerUpdateParameters) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *BannerUpdateParameters) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *BannerUpdateParameters) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetDescription

`func (o *BannerUpdateParameters) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BannerUpdateParameters) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BannerUpdateParameters) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BannerUpdateParameters) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *BannerUpdateParameters) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BannerUpdateParameters) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


