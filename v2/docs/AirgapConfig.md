# AirgapConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AirgapStatus** | Pointer to **NullableString** | Specifies Airgap should be enabled or disabled. | [optional] 
**ExceptionProfiles** | Pointer to **[]string** | Specifies the firewall profiles allowed when Airgap is enabled. | [optional] 

## Methods

### NewAirgapConfig

`func NewAirgapConfig() *AirgapConfig`

NewAirgapConfig instantiates a new AirgapConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAirgapConfigWithDefaults

`func NewAirgapConfigWithDefaults() *AirgapConfig`

NewAirgapConfigWithDefaults instantiates a new AirgapConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAirgapStatus

`func (o *AirgapConfig) GetAirgapStatus() string`

GetAirgapStatus returns the AirgapStatus field if non-nil, zero value otherwise.

### GetAirgapStatusOk

`func (o *AirgapConfig) GetAirgapStatusOk() (*string, bool)`

GetAirgapStatusOk returns a tuple with the AirgapStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirgapStatus

`func (o *AirgapConfig) SetAirgapStatus(v string)`

SetAirgapStatus sets AirgapStatus field to given value.

### HasAirgapStatus

`func (o *AirgapConfig) HasAirgapStatus() bool`

HasAirgapStatus returns a boolean if a field has been set.

### SetAirgapStatusNil

`func (o *AirgapConfig) SetAirgapStatusNil(b bool)`

 SetAirgapStatusNil sets the value for AirgapStatus to be an explicit nil

### UnsetAirgapStatus
`func (o *AirgapConfig) UnsetAirgapStatus()`

UnsetAirgapStatus ensures that no value is present for AirgapStatus, not even an explicit nil
### GetExceptionProfiles

`func (o *AirgapConfig) GetExceptionProfiles() []string`

GetExceptionProfiles returns the ExceptionProfiles field if non-nil, zero value otherwise.

### GetExceptionProfilesOk

`func (o *AirgapConfig) GetExceptionProfilesOk() (*[]string, bool)`

GetExceptionProfilesOk returns a tuple with the ExceptionProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceptionProfiles

`func (o *AirgapConfig) SetExceptionProfiles(v []string)`

SetExceptionProfiles sets ExceptionProfiles field to given value.

### HasExceptionProfiles

`func (o *AirgapConfig) HasExceptionProfiles() bool`

HasExceptionProfiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


