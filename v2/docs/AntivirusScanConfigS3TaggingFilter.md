# AntivirusScanConfigS3TaggingFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | Pointer to **NullableBool** | If set, it enables tagging filter for S3 views. | [optional] 
**Mode** | Pointer to **NullableString** | Whitelist or blacklist the objects to be scanned. | [optional] 
**TagSet** | Pointer to [**[]TagSet**](TagSet.md) | List of key, value pair. If any of the tags on the object matches any tags defined in tagSet array, it&#39;s regarded as a match. | [optional] 

## Methods

### NewAntivirusScanConfigS3TaggingFilter

`func NewAntivirusScanConfigS3TaggingFilter() *AntivirusScanConfigS3TaggingFilter`

NewAntivirusScanConfigS3TaggingFilter instantiates a new AntivirusScanConfigS3TaggingFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAntivirusScanConfigS3TaggingFilterWithDefaults

`func NewAntivirusScanConfigS3TaggingFilterWithDefaults() *AntivirusScanConfigS3TaggingFilter`

NewAntivirusScanConfigS3TaggingFilterWithDefaults instantiates a new AntivirusScanConfigS3TaggingFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *AntivirusScanConfigS3TaggingFilter) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *AntivirusScanConfigS3TaggingFilter) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *AntivirusScanConfigS3TaggingFilter) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *AntivirusScanConfigS3TaggingFilter) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### SetIsEnabledNil

`func (o *AntivirusScanConfigS3TaggingFilter) SetIsEnabledNil(b bool)`

 SetIsEnabledNil sets the value for IsEnabled to be an explicit nil

### UnsetIsEnabled
`func (o *AntivirusScanConfigS3TaggingFilter) UnsetIsEnabled()`

UnsetIsEnabled ensures that no value is present for IsEnabled, not even an explicit nil
### GetMode

`func (o *AntivirusScanConfigS3TaggingFilter) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *AntivirusScanConfigS3TaggingFilter) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *AntivirusScanConfigS3TaggingFilter) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *AntivirusScanConfigS3TaggingFilter) HasMode() bool`

HasMode returns a boolean if a field has been set.

### SetModeNil

`func (o *AntivirusScanConfigS3TaggingFilter) SetModeNil(b bool)`

 SetModeNil sets the value for Mode to be an explicit nil

### UnsetMode
`func (o *AntivirusScanConfigS3TaggingFilter) UnsetMode()`

UnsetMode ensures that no value is present for Mode, not even an explicit nil
### GetTagSet

`func (o *AntivirusScanConfigS3TaggingFilter) GetTagSet() []TagSet`

GetTagSet returns the TagSet field if non-nil, zero value otherwise.

### GetTagSetOk

`func (o *AntivirusScanConfigS3TaggingFilter) GetTagSetOk() (*[]TagSet, bool)`

GetTagSetOk returns a tuple with the TagSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagSet

`func (o *AntivirusScanConfigS3TaggingFilter) SetTagSet(v []TagSet)`

SetTagSet sets TagSet field to given value.

### HasTagSet

`func (o *AntivirusScanConfigS3TaggingFilter) HasTagSet() bool`

HasTagSet returns a boolean if a field has been set.

### SetTagSetNil

`func (o *AntivirusScanConfigS3TaggingFilter) SetTagSetNil(b bool)`

 SetTagSetNil sets the value for TagSet to be an explicit nil

### UnsetTagSet
`func (o *AntivirusScanConfigS3TaggingFilter) UnsetTagSet()`

UnsetTagSet ensures that no value is present for TagSet, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


