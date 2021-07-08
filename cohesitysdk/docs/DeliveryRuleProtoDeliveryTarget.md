# DeliveryRuleProtoDeliveryTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailAddress** | Pointer to **NullableString** | List of email addresses to send notifications. | [optional] 
**EmailRecipientType** | Pointer to **NullableInt32** |  | [optional] 
**ExternalApiCurlOptions** | Pointer to **NullableString** | Specifies the curl options used to invoke above rest external api. | [optional] 
**ExternalApiUrl** | Pointer to **NullableString** | Specifies the external api to be invoked when an alert matching this rule is raised. | [optional] 
**Locale** | Pointer to **NullableString** | Locale for the delivery target. | [optional] 
**SnmpNotification** | Pointer to **NullableBool** | Need to send snmp notification for matched alerts. | [optional] 
**SyslogNotification** | Pointer to **NullableBool** | Need to write syslog for matched alerts. | [optional] 
**TenantId** | Pointer to **NullableString** | Tenant who has been assigned this target. This field is not populated within AlertsDataProto persisted in Gandalf. This is a convenience field and is populated on the fly by the Alerts component for delivery targets in the delivery_target_list within AlertProto. This field is utilized by NotificationDeliveryHelper to group delivery targets so that we could send out a single email to all the email addresses registered with the same locale by a given tenant or by the SP admin. Another approach could have been to use an internal object, but since the AlertProto contains a list of type DeliveryTarget, this field has been added to make it convenient to pass around an AlertProto object. | [optional] 

## Methods

### NewDeliveryRuleProtoDeliveryTarget

`func NewDeliveryRuleProtoDeliveryTarget() *DeliveryRuleProtoDeliveryTarget`

NewDeliveryRuleProtoDeliveryTarget instantiates a new DeliveryRuleProtoDeliveryTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliveryRuleProtoDeliveryTargetWithDefaults

`func NewDeliveryRuleProtoDeliveryTargetWithDefaults() *DeliveryRuleProtoDeliveryTarget`

NewDeliveryRuleProtoDeliveryTargetWithDefaults instantiates a new DeliveryRuleProtoDeliveryTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailAddress

`func (o *DeliveryRuleProtoDeliveryTarget) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *DeliveryRuleProtoDeliveryTarget) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *DeliveryRuleProtoDeliveryTarget) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *DeliveryRuleProtoDeliveryTarget) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### SetEmailAddressNil

`func (o *DeliveryRuleProtoDeliveryTarget) SetEmailAddressNil(b bool)`

 SetEmailAddressNil sets the value for EmailAddress to be an explicit nil

### UnsetEmailAddress
`func (o *DeliveryRuleProtoDeliveryTarget) UnsetEmailAddress()`

UnsetEmailAddress ensures that no value is present for EmailAddress, not even an explicit nil
### GetEmailRecipientType

`func (o *DeliveryRuleProtoDeliveryTarget) GetEmailRecipientType() int32`

GetEmailRecipientType returns the EmailRecipientType field if non-nil, zero value otherwise.

### GetEmailRecipientTypeOk

`func (o *DeliveryRuleProtoDeliveryTarget) GetEmailRecipientTypeOk() (*int32, bool)`

GetEmailRecipientTypeOk returns a tuple with the EmailRecipientType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailRecipientType

`func (o *DeliveryRuleProtoDeliveryTarget) SetEmailRecipientType(v int32)`

SetEmailRecipientType sets EmailRecipientType field to given value.

### HasEmailRecipientType

`func (o *DeliveryRuleProtoDeliveryTarget) HasEmailRecipientType() bool`

HasEmailRecipientType returns a boolean if a field has been set.

### SetEmailRecipientTypeNil

`func (o *DeliveryRuleProtoDeliveryTarget) SetEmailRecipientTypeNil(b bool)`

 SetEmailRecipientTypeNil sets the value for EmailRecipientType to be an explicit nil

### UnsetEmailRecipientType
`func (o *DeliveryRuleProtoDeliveryTarget) UnsetEmailRecipientType()`

UnsetEmailRecipientType ensures that no value is present for EmailRecipientType, not even an explicit nil
### GetExternalApiCurlOptions

`func (o *DeliveryRuleProtoDeliveryTarget) GetExternalApiCurlOptions() string`

GetExternalApiCurlOptions returns the ExternalApiCurlOptions field if non-nil, zero value otherwise.

### GetExternalApiCurlOptionsOk

`func (o *DeliveryRuleProtoDeliveryTarget) GetExternalApiCurlOptionsOk() (*string, bool)`

GetExternalApiCurlOptionsOk returns a tuple with the ExternalApiCurlOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalApiCurlOptions

`func (o *DeliveryRuleProtoDeliveryTarget) SetExternalApiCurlOptions(v string)`

SetExternalApiCurlOptions sets ExternalApiCurlOptions field to given value.

### HasExternalApiCurlOptions

`func (o *DeliveryRuleProtoDeliveryTarget) HasExternalApiCurlOptions() bool`

HasExternalApiCurlOptions returns a boolean if a field has been set.

### SetExternalApiCurlOptionsNil

`func (o *DeliveryRuleProtoDeliveryTarget) SetExternalApiCurlOptionsNil(b bool)`

 SetExternalApiCurlOptionsNil sets the value for ExternalApiCurlOptions to be an explicit nil

### UnsetExternalApiCurlOptions
`func (o *DeliveryRuleProtoDeliveryTarget) UnsetExternalApiCurlOptions()`

UnsetExternalApiCurlOptions ensures that no value is present for ExternalApiCurlOptions, not even an explicit nil
### GetExternalApiUrl

`func (o *DeliveryRuleProtoDeliveryTarget) GetExternalApiUrl() string`

GetExternalApiUrl returns the ExternalApiUrl field if non-nil, zero value otherwise.

### GetExternalApiUrlOk

`func (o *DeliveryRuleProtoDeliveryTarget) GetExternalApiUrlOk() (*string, bool)`

GetExternalApiUrlOk returns a tuple with the ExternalApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalApiUrl

`func (o *DeliveryRuleProtoDeliveryTarget) SetExternalApiUrl(v string)`

SetExternalApiUrl sets ExternalApiUrl field to given value.

### HasExternalApiUrl

`func (o *DeliveryRuleProtoDeliveryTarget) HasExternalApiUrl() bool`

HasExternalApiUrl returns a boolean if a field has been set.

### SetExternalApiUrlNil

`func (o *DeliveryRuleProtoDeliveryTarget) SetExternalApiUrlNil(b bool)`

 SetExternalApiUrlNil sets the value for ExternalApiUrl to be an explicit nil

### UnsetExternalApiUrl
`func (o *DeliveryRuleProtoDeliveryTarget) UnsetExternalApiUrl()`

UnsetExternalApiUrl ensures that no value is present for ExternalApiUrl, not even an explicit nil
### GetLocale

`func (o *DeliveryRuleProtoDeliveryTarget) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *DeliveryRuleProtoDeliveryTarget) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *DeliveryRuleProtoDeliveryTarget) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *DeliveryRuleProtoDeliveryTarget) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### SetLocaleNil

`func (o *DeliveryRuleProtoDeliveryTarget) SetLocaleNil(b bool)`

 SetLocaleNil sets the value for Locale to be an explicit nil

### UnsetLocale
`func (o *DeliveryRuleProtoDeliveryTarget) UnsetLocale()`

UnsetLocale ensures that no value is present for Locale, not even an explicit nil
### GetSnmpNotification

`func (o *DeliveryRuleProtoDeliveryTarget) GetSnmpNotification() bool`

GetSnmpNotification returns the SnmpNotification field if non-nil, zero value otherwise.

### GetSnmpNotificationOk

`func (o *DeliveryRuleProtoDeliveryTarget) GetSnmpNotificationOk() (*bool, bool)`

GetSnmpNotificationOk returns a tuple with the SnmpNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpNotification

`func (o *DeliveryRuleProtoDeliveryTarget) SetSnmpNotification(v bool)`

SetSnmpNotification sets SnmpNotification field to given value.

### HasSnmpNotification

`func (o *DeliveryRuleProtoDeliveryTarget) HasSnmpNotification() bool`

HasSnmpNotification returns a boolean if a field has been set.

### SetSnmpNotificationNil

`func (o *DeliveryRuleProtoDeliveryTarget) SetSnmpNotificationNil(b bool)`

 SetSnmpNotificationNil sets the value for SnmpNotification to be an explicit nil

### UnsetSnmpNotification
`func (o *DeliveryRuleProtoDeliveryTarget) UnsetSnmpNotification()`

UnsetSnmpNotification ensures that no value is present for SnmpNotification, not even an explicit nil
### GetSyslogNotification

`func (o *DeliveryRuleProtoDeliveryTarget) GetSyslogNotification() bool`

GetSyslogNotification returns the SyslogNotification field if non-nil, zero value otherwise.

### GetSyslogNotificationOk

`func (o *DeliveryRuleProtoDeliveryTarget) GetSyslogNotificationOk() (*bool, bool)`

GetSyslogNotificationOk returns a tuple with the SyslogNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogNotification

`func (o *DeliveryRuleProtoDeliveryTarget) SetSyslogNotification(v bool)`

SetSyslogNotification sets SyslogNotification field to given value.

### HasSyslogNotification

`func (o *DeliveryRuleProtoDeliveryTarget) HasSyslogNotification() bool`

HasSyslogNotification returns a boolean if a field has been set.

### SetSyslogNotificationNil

`func (o *DeliveryRuleProtoDeliveryTarget) SetSyslogNotificationNil(b bool)`

 SetSyslogNotificationNil sets the value for SyslogNotification to be an explicit nil

### UnsetSyslogNotification
`func (o *DeliveryRuleProtoDeliveryTarget) UnsetSyslogNotification()`

UnsetSyslogNotification ensures that no value is present for SyslogNotification, not even an explicit nil
### GetTenantId

`func (o *DeliveryRuleProtoDeliveryTarget) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *DeliveryRuleProtoDeliveryTarget) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *DeliveryRuleProtoDeliveryTarget) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *DeliveryRuleProtoDeliveryTarget) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *DeliveryRuleProtoDeliveryTarget) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *DeliveryRuleProtoDeliveryTarget) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


