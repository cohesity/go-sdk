# ViewDirectoryQuotaPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertLimitBytes** | Pointer to **NullableInt64** | Specifies if an alert should be triggered when the usage of this resource exceeds this quota limit. This limit is optional and is specified in bytes. If no value is specified, there is no limit. | [optional] 
**HardLimitBytes** | Pointer to **NullableInt64** | Specifies an optional quota limit on the usage allowed for this resource. This limit is specified in bytes. If no value is specified, there is no limit. | [optional] 

## Methods

### NewViewDirectoryQuotaPolicy

`func NewViewDirectoryQuotaPolicy() *ViewDirectoryQuotaPolicy`

NewViewDirectoryQuotaPolicy instantiates a new ViewDirectoryQuotaPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewDirectoryQuotaPolicyWithDefaults

`func NewViewDirectoryQuotaPolicyWithDefaults() *ViewDirectoryQuotaPolicy`

NewViewDirectoryQuotaPolicyWithDefaults instantiates a new ViewDirectoryQuotaPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertLimitBytes

`func (o *ViewDirectoryQuotaPolicy) GetAlertLimitBytes() int64`

GetAlertLimitBytes returns the AlertLimitBytes field if non-nil, zero value otherwise.

### GetAlertLimitBytesOk

`func (o *ViewDirectoryQuotaPolicy) GetAlertLimitBytesOk() (*int64, bool)`

GetAlertLimitBytesOk returns a tuple with the AlertLimitBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertLimitBytes

`func (o *ViewDirectoryQuotaPolicy) SetAlertLimitBytes(v int64)`

SetAlertLimitBytes sets AlertLimitBytes field to given value.

### HasAlertLimitBytes

`func (o *ViewDirectoryQuotaPolicy) HasAlertLimitBytes() bool`

HasAlertLimitBytes returns a boolean if a field has been set.

### SetAlertLimitBytesNil

`func (o *ViewDirectoryQuotaPolicy) SetAlertLimitBytesNil(b bool)`

 SetAlertLimitBytesNil sets the value for AlertLimitBytes to be an explicit nil

### UnsetAlertLimitBytes
`func (o *ViewDirectoryQuotaPolicy) UnsetAlertLimitBytes()`

UnsetAlertLimitBytes ensures that no value is present for AlertLimitBytes, not even an explicit nil
### GetHardLimitBytes

`func (o *ViewDirectoryQuotaPolicy) GetHardLimitBytes() int64`

GetHardLimitBytes returns the HardLimitBytes field if non-nil, zero value otherwise.

### GetHardLimitBytesOk

`func (o *ViewDirectoryQuotaPolicy) GetHardLimitBytesOk() (*int64, bool)`

GetHardLimitBytesOk returns a tuple with the HardLimitBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardLimitBytes

`func (o *ViewDirectoryQuotaPolicy) SetHardLimitBytes(v int64)`

SetHardLimitBytes sets HardLimitBytes field to given value.

### HasHardLimitBytes

`func (o *ViewDirectoryQuotaPolicy) HasHardLimitBytes() bool`

HasHardLimitBytes returns a boolean if a field has been set.

### SetHardLimitBytesNil

`func (o *ViewDirectoryQuotaPolicy) SetHardLimitBytesNil(b bool)`

 SetHardLimitBytesNil sets the value for HardLimitBytes to be an explicit nil

### UnsetHardLimitBytes
`func (o *ViewDirectoryQuotaPolicy) UnsetHardLimitBytes()`

UnsetHardLimitBytes ensures that no value is present for HardLimitBytes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


