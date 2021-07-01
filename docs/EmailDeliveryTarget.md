# EmailDeliveryTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailAddress** | Pointer to **NullableString** |  | [optional] 
**Locale** | Pointer to **NullableString** | Specifies the language in which the emails sent to the above defined mail address should be in. | [optional] 
**RecipientType** | Pointer to **NullableInt32** | Specifies the recipient type on how the emails are to received. | [optional] 

## Methods

### NewEmailDeliveryTarget

`func NewEmailDeliveryTarget() *EmailDeliveryTarget`

NewEmailDeliveryTarget instantiates a new EmailDeliveryTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailDeliveryTargetWithDefaults

`func NewEmailDeliveryTargetWithDefaults() *EmailDeliveryTarget`

NewEmailDeliveryTargetWithDefaults instantiates a new EmailDeliveryTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailAddress

`func (o *EmailDeliveryTarget) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *EmailDeliveryTarget) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *EmailDeliveryTarget) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *EmailDeliveryTarget) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### SetEmailAddressNil

`func (o *EmailDeliveryTarget) SetEmailAddressNil(b bool)`

 SetEmailAddressNil sets the value for EmailAddress to be an explicit nil

### UnsetEmailAddress
`func (o *EmailDeliveryTarget) UnsetEmailAddress()`

UnsetEmailAddress ensures that no value is present for EmailAddress, not even an explicit nil
### GetLocale

`func (o *EmailDeliveryTarget) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *EmailDeliveryTarget) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *EmailDeliveryTarget) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *EmailDeliveryTarget) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### SetLocaleNil

`func (o *EmailDeliveryTarget) SetLocaleNil(b bool)`

 SetLocaleNil sets the value for Locale to be an explicit nil

### UnsetLocale
`func (o *EmailDeliveryTarget) UnsetLocale()`

UnsetLocale ensures that no value is present for Locale, not even an explicit nil
### GetRecipientType

`func (o *EmailDeliveryTarget) GetRecipientType() int32`

GetRecipientType returns the RecipientType field if non-nil, zero value otherwise.

### GetRecipientTypeOk

`func (o *EmailDeliveryTarget) GetRecipientTypeOk() (*int32, bool)`

GetRecipientTypeOk returns a tuple with the RecipientType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientType

`func (o *EmailDeliveryTarget) SetRecipientType(v int32)`

SetRecipientType sets RecipientType field to given value.

### HasRecipientType

`func (o *EmailDeliveryTarget) HasRecipientType() bool`

HasRecipientType returns a boolean if a field has been set.

### SetRecipientTypeNil

`func (o *EmailDeliveryTarget) SetRecipientTypeNil(b bool)`

 SetRecipientTypeNil sets the value for RecipientType to be an explicit nil

### UnsetRecipientType
`func (o *EmailDeliveryTarget) UnsetRecipientType()`

UnsetRecipientType ensures that no value is present for RecipientType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


