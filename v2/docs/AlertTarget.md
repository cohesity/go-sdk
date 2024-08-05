# AlertTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailAddress** | **NullableString** | Specifies an email address to receive an alert. | 
**Language** | Pointer to **NullableString** | Specifies the language of the delivery target. Default value is &#39;en-us&#39;. | [optional] 
**RecipientType** | Pointer to **NullableString** | Specifies the recipient type of email recipient. Default value is &#39;kTo&#39;. | [optional] 

## Methods

### NewAlertTarget

`func NewAlertTarget(emailAddress NullableString, ) *AlertTarget`

NewAlertTarget instantiates a new AlertTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertTargetWithDefaults

`func NewAlertTargetWithDefaults() *AlertTarget`

NewAlertTargetWithDefaults instantiates a new AlertTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailAddress

`func (o *AlertTarget) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *AlertTarget) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *AlertTarget) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.


### SetEmailAddressNil

`func (o *AlertTarget) SetEmailAddressNil(b bool)`

 SetEmailAddressNil sets the value for EmailAddress to be an explicit nil

### UnsetEmailAddress
`func (o *AlertTarget) UnsetEmailAddress()`

UnsetEmailAddress ensures that no value is present for EmailAddress, not even an explicit nil
### GetLanguage

`func (o *AlertTarget) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *AlertTarget) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *AlertTarget) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *AlertTarget) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### SetLanguageNil

`func (o *AlertTarget) SetLanguageNil(b bool)`

 SetLanguageNil sets the value for Language to be an explicit nil

### UnsetLanguage
`func (o *AlertTarget) UnsetLanguage()`

UnsetLanguage ensures that no value is present for Language, not even an explicit nil
### GetRecipientType

`func (o *AlertTarget) GetRecipientType() string`

GetRecipientType returns the RecipientType field if non-nil, zero value otherwise.

### GetRecipientTypeOk

`func (o *AlertTarget) GetRecipientTypeOk() (*string, bool)`

GetRecipientTypeOk returns a tuple with the RecipientType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientType

`func (o *AlertTarget) SetRecipientType(v string)`

SetRecipientType sets RecipientType field to given value.

### HasRecipientType

`func (o *AlertTarget) HasRecipientType() bool`

HasRecipientType returns a boolean if a field has been set.

### SetRecipientTypeNil

`func (o *AlertTarget) SetRecipientTypeNil(b bool)`

 SetRecipientTypeNil sets the value for RecipientType to be an explicit nil

### UnsetRecipientType
`func (o *AlertTarget) UnsetRecipientType()`

UnsetRecipientType ensures that no value is present for RecipientType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


