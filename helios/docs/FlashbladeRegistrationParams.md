# FlashbladeRegistrationParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Endpoint** | **NullableString** | Specifies the Hostname or IP Address Endpoint for the Flashblade Source. | 
**ApiToken** | Pointer to **NullableString** | Specifies the API Token of the Flashblade Source | [optional] 
**BackUpSMBVolumes** | Pointer to **NullableBool** | Specifies whether or not to back up SMB Volumes. | [optional] 
**SmbCredentials** | Pointer to [**SmbMountCredentials**](SmbMountCredentials.md) |  | [optional] 
**ThrottlingConfig** | Pointer to [**NasThrottlingConfig**](NasThrottlingConfig.md) |  | [optional] 

## Methods

### NewFlashbladeRegistrationParams

`func NewFlashbladeRegistrationParams(endpoint NullableString, ) *FlashbladeRegistrationParams`

NewFlashbladeRegistrationParams instantiates a new FlashbladeRegistrationParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlashbladeRegistrationParamsWithDefaults

`func NewFlashbladeRegistrationParamsWithDefaults() *FlashbladeRegistrationParams`

NewFlashbladeRegistrationParamsWithDefaults instantiates a new FlashbladeRegistrationParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpoint

`func (o *FlashbladeRegistrationParams) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *FlashbladeRegistrationParams) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *FlashbladeRegistrationParams) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### SetEndpointNil

`func (o *FlashbladeRegistrationParams) SetEndpointNil(b bool)`

 SetEndpointNil sets the value for Endpoint to be an explicit nil

### UnsetEndpoint
`func (o *FlashbladeRegistrationParams) UnsetEndpoint()`

UnsetEndpoint ensures that no value is present for Endpoint, not even an explicit nil
### GetApiToken

`func (o *FlashbladeRegistrationParams) GetApiToken() string`

GetApiToken returns the ApiToken field if non-nil, zero value otherwise.

### GetApiTokenOk

`func (o *FlashbladeRegistrationParams) GetApiTokenOk() (*string, bool)`

GetApiTokenOk returns a tuple with the ApiToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiToken

`func (o *FlashbladeRegistrationParams) SetApiToken(v string)`

SetApiToken sets ApiToken field to given value.

### HasApiToken

`func (o *FlashbladeRegistrationParams) HasApiToken() bool`

HasApiToken returns a boolean if a field has been set.

### SetApiTokenNil

`func (o *FlashbladeRegistrationParams) SetApiTokenNil(b bool)`

 SetApiTokenNil sets the value for ApiToken to be an explicit nil

### UnsetApiToken
`func (o *FlashbladeRegistrationParams) UnsetApiToken()`

UnsetApiToken ensures that no value is present for ApiToken, not even an explicit nil
### GetBackUpSMBVolumes

`func (o *FlashbladeRegistrationParams) GetBackUpSMBVolumes() bool`

GetBackUpSMBVolumes returns the BackUpSMBVolumes field if non-nil, zero value otherwise.

### GetBackUpSMBVolumesOk

`func (o *FlashbladeRegistrationParams) GetBackUpSMBVolumesOk() (*bool, bool)`

GetBackUpSMBVolumesOk returns a tuple with the BackUpSMBVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackUpSMBVolumes

`func (o *FlashbladeRegistrationParams) SetBackUpSMBVolumes(v bool)`

SetBackUpSMBVolumes sets BackUpSMBVolumes field to given value.

### HasBackUpSMBVolumes

`func (o *FlashbladeRegistrationParams) HasBackUpSMBVolumes() bool`

HasBackUpSMBVolumes returns a boolean if a field has been set.

### SetBackUpSMBVolumesNil

`func (o *FlashbladeRegistrationParams) SetBackUpSMBVolumesNil(b bool)`

 SetBackUpSMBVolumesNil sets the value for BackUpSMBVolumes to be an explicit nil

### UnsetBackUpSMBVolumes
`func (o *FlashbladeRegistrationParams) UnsetBackUpSMBVolumes()`

UnsetBackUpSMBVolumes ensures that no value is present for BackUpSMBVolumes, not even an explicit nil
### GetSmbCredentials

`func (o *FlashbladeRegistrationParams) GetSmbCredentials() SmbMountCredentials`

GetSmbCredentials returns the SmbCredentials field if non-nil, zero value otherwise.

### GetSmbCredentialsOk

`func (o *FlashbladeRegistrationParams) GetSmbCredentialsOk() (*SmbMountCredentials, bool)`

GetSmbCredentialsOk returns a tuple with the SmbCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmbCredentials

`func (o *FlashbladeRegistrationParams) SetSmbCredentials(v SmbMountCredentials)`

SetSmbCredentials sets SmbCredentials field to given value.

### HasSmbCredentials

`func (o *FlashbladeRegistrationParams) HasSmbCredentials() bool`

HasSmbCredentials returns a boolean if a field has been set.

### GetThrottlingConfig

`func (o *FlashbladeRegistrationParams) GetThrottlingConfig() NasThrottlingConfig`

GetThrottlingConfig returns the ThrottlingConfig field if non-nil, zero value otherwise.

### GetThrottlingConfigOk

`func (o *FlashbladeRegistrationParams) GetThrottlingConfigOk() (*NasThrottlingConfig, bool)`

GetThrottlingConfigOk returns a tuple with the ThrottlingConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottlingConfig

`func (o *FlashbladeRegistrationParams) SetThrottlingConfig(v NasThrottlingConfig)`

SetThrottlingConfig sets ThrottlingConfig field to given value.

### HasThrottlingConfig

`func (o *FlashbladeRegistrationParams) HasThrottlingConfig() bool`

HasThrottlingConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


